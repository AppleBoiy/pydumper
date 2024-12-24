package util

import (
	"path/filepath"
	"syscall"
)

func GetFreeSpace() (int, error) {
	root := filepath.Clean("/")
	var stat syscall.Statfs_t
	err := syscall.Statfs(root, &stat)
	if err != nil {
		return 0, err
	}

	free := stat.Bfree * uint64(stat.Bsize)
	return int(free), nil
}
