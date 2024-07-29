package osi

import (
	"io"
	"os"
	"time"
)

type File struct {
	*os.File
}

func Crate(name string) (*File, error) {
	f, err := os.Create(name)
	return &File{f}, err
}

func CreateTemp(dir, pattern string) (*File, error) {
	f, err := os.CreateTemp(dir, pattern)
	return &File{f}, err
}

func NewFile(fd uintptr, name string) *File {
	return &File{os.NewFile(fd, name)}
}

func Open(name string) (*File, error) {
	f, err := os.Open(name)
	return &File{f}, err
}

func OpenFile(name string, flag int, perm os.FileMode) (*File, error) {
	f, err := os.OpenFile(name, flag, perm)
	return &File{f}, err
}

func (f *File) Chdir() error {
	return os.Chdir(f.Name())
}

func (f *File) Chmod(mode os.FileMode) error {
	return os.Chmod(f.Name(), mode)
}

func (f *File) Chown(uid, gid int) error {
	return os.Chown(f.Name(), uid, gid)
}

func (f *File) Close() error {
	return f.File.Close()
}

func (f *File) Fd() uintptr {
	return f.File.Fd()
}

func (f *File) Name() string {
	return f.File.Name()
}

func (f *File) Read(b []byte) (int, error) {
	return f.File.Read(b)
}

func (f *File) ReadAt(b []byte, off int64) (int, error) {
	return f.File.ReadAt(b, off)
}

func (f *File) ReadDir(n int) ([]os.DirEntry, error) {
	return f.File.ReadDir(n)
}

func (f *File) ReadFrom(r io.Reader) (int64, error) {
	return f.File.ReadFrom(r)
}

func (f *File) Readdir(n int) ([]os.FileInfo, error) {
	return f.File.Readdir(n)
}

func (f *File) Readdirnames(n int) ([]string, error) {
	return f.File.Readdirnames(n)
}

func (f *File) Seek(offset int64, whence int) (int64, error) {
	return f.File.Seek(offset, whence)
}

func (f *File) SetDeadline(t time.Time) error {
	return f.File.SetDeadline(t)
}

func (f *File) SetReadDeadline(t time.Time) error {
	return f.File.SetReadDeadline(t)
}

func (f *File) SetWriteDeadline(t time.Time) error {
	return f.File.SetWriteDeadline(t)
}

func (f *File) Stat() (os.FileInfo, error) {
	return f.File.Stat()
}

func (f *File) Sync() error {
	return f.File.Sync()
}

func (f *File) SyscallConn() (interface{}, error) {
	return f.File.SyscallConn()
}

func (f *File) Truncate(size int64) error {
	return f.File.Truncate(size)
}

func (f *File) Write(b []byte) (int, error) {
	return f.File.Write(b)
}

func (f *File) WriteAt(b []byte, off int64) (int, error) {
	return f.File.WriteAt(b, off)
}

func (f *File) WriteString(s string) (int, error) {
	return f.File.WriteString(s)
}

func (f *File) WriteTo(w io.Writer) (int64, error) {
	return f.File.WriteTo(w)
}
