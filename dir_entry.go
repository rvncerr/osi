package osi

import "os"

type DirEntry = os.DirEntry

func ReadDir(name string) ([]DirEntry, error) {
	return os.ReadDir(name)
}
