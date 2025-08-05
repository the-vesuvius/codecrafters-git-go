package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/codecrafters-io/git-starter-go/app/object"
	"github.com/spf13/cobra"
)

var writeObject bool

var hashObject = &cobra.Command{
	Use:  "hash-object",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filepath := args[0]

		file, err := os.Open(filepath)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			return
		}
		defer file.Close()
		content, err := io.ReadAll(file)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}
		obj := object.Object{
			Type:    "blob",
			Size:    len(content),
			Content: string(content),
		}
		objectHash := obj.Hash()
		fmt.Printf("%s\n", objectHash)

		if writeObject {
			objectFilePath := object.GetFilePathFromHash(objectHash)
			err = object.WriteObjectFile(objectFilePath, obj)
			if err != nil {
				fmt.Printf("Error writing object file: %v\n", err)
				return
			}
		}
	},
}

func init() {
	hashObject.Flags().BoolVarP(&writeObject, "write", "w", false, "write the object into the object database")
}
