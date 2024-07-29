package osi

import "os"

type FileInfo = os.FileInfo

func Lstat(name string) (FileInfo, error) {
	return os.Lstat(name)
}

func Stat(name string) (FileInfo, error) {
	return os.Stat(name)
}
