package fileio

import "os"

func ReadFile(filePath string) []byte {
	contents, err := os.ReadFile(filePath)

	if err != nil {
		panic(err)
	}
	return contents
}
