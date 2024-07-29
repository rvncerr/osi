package osi

import (
	"io/fs"
	"os"
	"syscall"
	"time"
)

const (
	O_RDONLY int = syscall.O_RDONLY
	O_WRONLY int = syscall.O_WRONLY
	O_RDWR   int = syscall.O_RDWR
	O_APPEND int = syscall.O_APPEND
	O_CREATE int = syscall.O_CREAT
	O_EXCL   int = syscall.O_EXCL
	O_SYNC   int = syscall.O_SYNC
	O_TRUNC  int = syscall.O_TRUNC
)

const (
	SEEK_SET int = 0
	SEEK_CUR int = 1
	SEEK_END int = 2
)

const (
	PathSeparator     = os.PathSeparator
	PathListSeparator = os.PathListSeparator
)

const (
	ModeDir        = fs.ModeDir
	ModeAppend     = fs.ModeAppend
	ModeExclusive  = fs.ModeExclusive
	ModeTemporary  = fs.ModeTemporary
	ModeSymlink    = fs.ModeSymlink
	ModeDevice     = fs.ModeDevice
	ModeNamedPipe  = fs.ModeNamedPipe
	ModeSocket     = fs.ModeSocket
	ModeSetuid     = fs.ModeSetuid
	ModeSetgid     = fs.ModeSetgid
	ModeCharDevice = fs.ModeCharDevice
	ModeSticky     = fs.ModeSticky
	ModeIrregular  = fs.ModeIrregular
	ModeType       = fs.ModeType
	ModePerm       = fs.ModePerm
)

const DevNull = os.DevNull

var (
	ErrInvalid = fs.ErrInvalid

	ErrPermission = fs.ErrPermission
	ErrExist      = fs.ErrExist
	ErrNotExist   = fs.ErrNotExist
	ErrClosed     = fs.ErrClosed

	ErrNoDeadline       = os.ErrNoDeadline
	ErrDeadlineExceeded = os.ErrDeadlineExceeded
)

var (
	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)

var Args []string = os.Args

var ErrProcessDone = os.ErrProcessDone

func Chdir(path string) error {
	return os.Chdir(path)
}

func Chmod(name string, mode FileMode) error {
	return os.Chmod(name, mode)
}

func Chown(name string, uid, gid int) error {
	return os.Chown(name, uid, gid)
}

func Chtimes(name string, atime, mtime time.Time) error {
	return os.Chtimes(name, atime, mtime)
}

func Clearenv() {
	os.Clearenv()
}

func DirFS(name string) fs.FS {
	return os.DirFS(name)
}

func Environ() []string {
	return os.Environ()
}

func Executable() (string, error) {
	return os.Executable()
}

func Exit(code int) {
	os.Exit(code)
}

func Expand(s string, mapping func(string) string) string {
	return os.Expand(s, mapping)
}

func ExpandEnv(s string) string {
	return os.ExpandEnv(s)
}

func Getegid() int {
	return os.Getegid()
}

func Getenv(key string) string {
	return os.Getenv(key)
}

func Geteuid() int {
	return os.Geteuid()
}

func Getgid() int {
	return os.Getgid()
}

func Getgroups() ([]int, error) {
	return os.Getgroups()
}

func Getpagesize() int {
	return os.Getpagesize()
}

func Getpid() int {
	return os.Getpid()
}

func Getppid() int {
	return os.Getppid()
}

func Getuid() int {
	return os.Getuid()
}

func Getwd() (string, error) {
	return os.Getwd()
}

func Hostname() (name string, err error) {
	return os.Hostname()
}

func IsExist(err error) bool {
	return os.IsExist(err)
}

func IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

func IsPathSeparator(c uint8) bool {
	return os.IsPathSeparator(c)
}

func IsPermission(err error) bool {
	return os.IsPermission(err)
}

func IsTimeout(err error) bool {
	return os.IsTimeout(err)
}

func Lchown(name string, uid, gid int) error {
	return os.Lchown(name, uid, gid)
}

func Link(oldname, newname string) error {
	return os.Link(oldname, newname)
}

func LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}

func Mkdir(name string, perm FileMode) error {
	return os.Mkdir(name, perm)
}

func MkdirAll(path string, perm FileMode) error {
	return os.MkdirAll(path, perm)
}

func MkdirTemp(dir, pattern string) (name string, err error) {
	return os.MkdirTemp(dir, pattern)
}

func NewSyscallError(syscall string, err error) error {
	return os.NewSyscallError(syscall, err)
}

func Pipe() (r *File, w *File, err error) {
	ro, wo, err := os.Pipe()
	return &File{ro}, &File{wo}, err
}

func Readfile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func Readlink(name string) (string, error) {
	return os.Readlink(name)
}

func Remove(name string) error {
	return os.Remove(name)
}

func RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func SameFile(fi1, fi2 FileInfo) bool {
	return os.SameFile(fi1, fi2)
}

func Setenv(key, value string) error {
	return os.Setenv(key, value)
}

func Symlink(oldname, newname string) error {
	return os.Symlink(oldname, newname)
}

func TempDir() string {
	return os.TempDir()
}

func Truncate(name string, size int64) error {
	return os.Truncate(name, size)
}

func Unsetenv(key string) error {
	return os.Unsetenv(key)
}

func UserCacheDir() (string, error) {
	return os.UserCacheDir()
}

func UserConfigDir() (string, error) {
	return os.UserConfigDir()
}

func UserHomeDir() (string, error) {
	return os.UserHomeDir()
}

func WriteFile(filename string, data []byte, perm FileMode) error {
	return os.WriteFile(filename, data, perm)
}
