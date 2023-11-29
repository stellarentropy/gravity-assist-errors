package utils

import "os"

// IsFile determines whether the provided path refers to an existing regular
// file. It distinguishes between files and directories, returning true only for
// files. Non-existent paths or errors during file information retrieval result
// in a false outcome.
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func IsDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
