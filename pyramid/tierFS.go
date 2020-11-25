package pyramid

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/treeverse/lakefs/block"
)

type storageProps struct {
	t                  StorageType
	localDir           string
	blockStoragePrefix string
}

// ImmutableTierFS is a filesystem where written files are never edited.
// All files are stored in the block storage. Local paths are treated as a
// cache layer that will be evicted according to the given eviction algorithm.
type ImmutableTierFS struct {
	adaptor block.Adapter
	config  map[StorageType]storageProps

	// TODO: use refs anc last-access for the eviction algorithm
	refCount   map[StorageType]map[string]int
	lastAccess map[StorageType]map[string]time.Time
}

// mapping between supported storage types and their prefix
var types = map[StorageType]string{
	StorageTypeSSTable:      "sstables",
	StorageTypeTreeManifest: "trees",
}

func NewTierFS(adaptor block.Adapter, localdir, blockStoragePrefix string) *ImmutableTierFS {
	refCount := map[StorageType]map[string]int{}
	lastAccess := map[StorageType]map[string]time.Time{}
	config := map[StorageType]storageProps{}
	for t, prefix := range types {
		refCount[t] = map[string]int{}
		lastAccess[t] = map[string]time.Time{}
		config[t] = storageProps{
			t:                  t,
			localDir:           path.Join(localdir, prefix),
			blockStoragePrefix: path.Join(blockStoragePrefix, prefix),
		}
	}

	return &ImmutableTierFS{
		adaptor:    adaptor,
		refCount:   refCount,
		lastAccess: lastAccess,
		config:     config,
	}
}

// Store uploads the local file to the block storage for persistence.
func (tfs *ImmutableTierFS) Store(t StorageType, ns, filename string) error {
	localpath := tfs.localpath(t, filename)
	f, err := os.Open(localpath)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return fmt.Errorf("file stat: %w", err)
	}

	return tfs.adaptor.Put(tfs.objPointer(t, ns, filename), stat.Size(), f, block.PutOpts{})
}

// Load returns the a file descriptor to the local file.
// If the file is missing from the local disk, it will try to fetch it from the block storage.
func (tfs *ImmutableTierFS) Load(t StorageType, ns, filename string) (*File, error) {
	localPath := tfs.localpath(t, filename)
	fh, err := os.Open(localPath)
	if err != nil {
		if os.IsNotExist(err) {
			fh, err = tfs.readFromBlockStorage(t, ns, filename)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("open file: %w", err)
		}
	}

	// TODO: refs thread-safe
	tfs.refCount[t][filename] = tfs.refCount[t][filename] + 1
	return &File{
		fh: fh,
		access: func() {
			tfs.lastAccess[t][filename] = time.Now()
		},
		release: func() {
			tfs.refCount[t][filename] = tfs.refCount[t][filename] - 1
		},
	}, nil
}

func (tfs *ImmutableTierFS) readFromBlockStorage(t StorageType, ns string, filename string) (*os.File, error) {
	reader, err := tfs.adaptor.Get(tfs.objPointer(t, ns, filename), 0)
	if err != nil {
		return nil, fmt.Errorf("read from block storage: %w", err)
	}
	defer reader.Close()

	localPath := tfs.localpath(t, filename)
	writer, err := os.Create(localPath)
	if err != nil {
		return nil, fmt.Errorf("creating file: %w", err)
	}
	defer writer.Close()

	if _, err := io.Copy(writer, reader); err != nil {
		return nil, fmt.Errorf("copying date to file: %w", err)
	}

	fh, err := os.Open(localPath)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	return fh, nil
}

func (tfs *ImmutableTierFS) localpath(t StorageType, filename string) string {
	return path.Join(tfs.config[t].localDir, filename)
}

func (tfs *ImmutableTierFS) blockStoragePath(t StorageType, filename string) string {
	return path.Join(tfs.config[t].blockStoragePrefix, filename)
}

func (tfs *ImmutableTierFS) objPointer(t StorageType, ns string, filename string) block.ObjectPointer {
	return block.ObjectPointer{
		StorageNamespace: ns,
		Identifier:       tfs.blockStoragePath(t, filename),
	}
}
