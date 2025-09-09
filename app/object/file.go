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
	fileContent, err := os.ReadFile(objectFilePath)
	if err != nil {
		return nil, err
	}

	zlibReader, err := zlib.NewReader(bytes.NewReader(fileContent))
	if err != nil {
		return nil, err
	}
	defer zlibReader.Close()

	content, err := io.ReadAll(zlibReader)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func WriteObjectFile(objectFilePath string, obj Object) error {
	content := obj.String()
	var buf bytes.Buffer
	zlibWriter := zlib.NewWriter(&buf)
	_, err := zlibWriter.Write([]byte(content))
	if err != nil {
		return err
	}
	err = zlibWriter.Close()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(path.Dir(objectFilePath), 0755); err != nil {
		return err
	}

	return os.WriteFile(objectFilePath, buf.Bytes(), 0644)
}
