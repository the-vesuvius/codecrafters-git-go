package object

import (
	"bytes"
	"compress/zlib"
	"io"
	"os"
	"path"

	"github.com/codecrafters-io/git-starter-go/app/constants"
)

func GetFilePathFromHash(objectHash string) string {
	return path.Join(constants.ObjectsDirPath, objectHash[:2], objectHash[2:])
}

func GetDecompressedFileContent(objectFilePath string) ([]byte, error) {
	content, err := os.ReadFile(objectFilePath)
	if err != nil {
		return nil, err
	}

	zlibReader, err := zlib.NewReader(bytes.NewReader(content))
	if err != nil {
		return nil, err
	}
	defer zlibReader.Close()

	content, err = io.ReadAll(zlibReader)
	if err != nil {
		return nil, err
	}
	return content, nil
}
