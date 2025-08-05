package object

import (
	"crypto/sha1"
	"fmt"
	"strconv"
)

type Object struct {
	Type    string
	Size    int
	Content string
}

func (o Object) String() string {
	return fmt.Sprintf("%s %d\x00%s", o.Type, o.Size, o.Content)
}

func (o Object) Hash() string {
	sha1Hasher := sha1.New()
	sha1Hasher.Write([]byte(o.String()))
	return fmt.Sprintf("%x", sha1Hasher.Sum(nil))
}

func Parse(rawContent []byte) (Object, error) {
	obj := Object{}
	idx := 0
	for rawContent[idx] != ' ' {
		idx++
	}
	obj.Type = string(rawContent[:idx])
	idx++ // Skip the space
	sizeStart := idx
	for rawContent[idx] != 0 {
		idx++
	}
	sizeStr := string(rawContent[sizeStart:idx])
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return Object{}, err
	}
	obj.Size = size
	idx++ // Skip the null byte
	obj.Content = string(rawContent[idx : idx+size])
	return obj, nil
}
