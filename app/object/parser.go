package object

import "strconv"

type Object struct {
	Type    string
	Size    int
	Content string
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
