package file

import (
	"io"
	"os"
)

func IsValidFilePath(path string) bool {
    _, err := os.Stat(path)
    if os.IsNotExist(err) {
        return false
    }
    return err == nil
}

func ReadFileAsBytes(path string) ([]byte, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }

    defer file.Close()

    fileContents, err := io.ReadAll(file)
    if err != nil {
        return nil, err
    }
	return fileContents, nil
}