package utils

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

func ReadFile(filename string) []byte {
	bytesRead, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error reading file: %s", err.Error())
	}

	return bytesRead
}

func GetFileNameWithoutExtension(path string) string {
	filename := filepath.Base(path)
	return filename[:len(filename)-len(filepath.Ext(filename))]
}

func WriteBytesToFile(fileName string, bytes []byte) error {
	return ioutil.WriteFile(fileName, bytes, 0777)
}
