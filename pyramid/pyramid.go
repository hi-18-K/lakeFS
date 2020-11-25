package pyramid

// StorageType is enumeration of the different tiered storage types
// supported by pyramid.
type StorageType int

const (
	StorageTypeSSTable StorageType = iota
	StorageTypeTreeManifest
)

// FS is pyramid abstraction of filesystem where the persistent storage-layer is the block storage.
// Files on the Local disk are transient and might be cleaned up by the eviction policy.
type FS interface {
	// Store uploads the file to the block storage.
	// It may remove the file from the local storage.
	Store(t StorageType, filename string) error

	// Open finds the referenced file and returns the file descriptor.
	// If file isn't in the local disk, it is fetched from the block storage.
	Open(t StorageType, filename string) (*File,error)
}