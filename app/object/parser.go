package object

import (
	"crypto/sha1"
	"fmt"
	"strconv"
)

const (
	ModeFile = "100644"
	ModeDir  = "40000"
	ModeExec = "100755"
	ModeLink = "120000"
)

type Object struct {
	Type    string
	Size    int
	Content string
	Entries []TreeEntry
}

type TreeEntry struct {
	Type string
	Mode string
	Name string
	Hash [20]byte
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

	idx := 0
	for rawContent[idx] != ' ' {
		idx++
	}

	objType := string(rawContent[:idx])
	idx++ // Skip the space
	switch objType {
	case "blob":
		return parseBlob(rawContent, idx)
	case "tree":
		return parseTree(rawContent, idx)
	default:
		return Object{}, fmt.Errorf("unknown object type: [%s]", objType)
	}
}

func parseBlob(rawContent []byte, idx int) (Object, error) {
	obj := Object{
		Type: "blob",
	}

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
	obj.Type = "blob"

	return obj, nil
}

func parseTree(rawContent []byte, idx int) (Object, error) {
	obj := Object{
		Type: "tree",
	}

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

	obj.Entries = []TreeEntry{}

	for idx < len(rawContent) {
		modeStart := idx
		for rawContent[idx] != ' ' {
			idx++
		}
		mode := string(rawContent[modeStart:idx])
		idx++ // Skip the space

		nameStart := idx
		for rawContent[idx] != 0 {
			idx++
		}
		name := string(rawContent[nameStart:idx])
		idx++ // Skip the null byte

		var hashBytes [20]byte
		copy(hashBytes[:], rawContent[idx:idx+20])
		idx += 20

		obj.Entries = append(obj.Entries, TreeEntry{
			Mode: mode,
			Name: name,
			Hash: hashBytes,
			Type: getTypeFromMode(mode),
		})

	}

	return obj, nil
}

func getTypeFromMode(mode string) string {
	switch mode {
	case ModeFile, ModeExec, ModeLink:
		return "blob"
	case ModeDir:
		return "tree"
	default:
		return ""
	}
}
