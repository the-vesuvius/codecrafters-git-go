package object

import "fmt"

func LsTree(treeSha string, nameOnly bool) error {
	treeFilePath := GetFilePathFromHash(treeSha)
	treeContent, err := GetDecompressedFileContent(treeFilePath)
	if err != nil {
		return err
	}

	treeObj, err := Parse(treeContent)
	if err != nil {
		return err
	}

	for _, entry := range treeObj.Entries {
		if nameOnly {
			fmt.Println(entry.Name)
		} else {
			fmt.Printf("%6s %4s %x\t%s\n", entry.Mode, entry.Type, entry.Hash, entry.Name)
		}
	}
	return nil
}
