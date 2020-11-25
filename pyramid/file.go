package pyramid

import "os"

type File struct {
	fh      *os.File
	access  func()
	release func()
}

func (f *File) Read(p []byte) (n int, err error) {
	f.access()
	return f.fh.Read(p)
}

func (f *File) ReadAt(p []byte, off int64) (n int, err error) {
	f.access()
	return f.fh.ReadAt(p, off)
}

func (f *File) Write(p []byte) (n int, err error) {
	f.access()
	return f.fh.Write(p)
}

func (f *File) Stat() (os.FileInfo, error) {
	return f.fh.Stat()
}

func (f *File) Sync() error {
	return f.fh.Sync()
}

func (f *File) Close() error {
	f.release()
	return f.Close()
}
