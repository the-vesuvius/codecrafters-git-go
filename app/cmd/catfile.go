package cmd

import (
	"fmt"

	"github.com/codecrafters-io/git-starter-go/app/object"
	"github.com/spf13/cobra"
)

var (
	catFilePrettyPrint bool
)

var catFile = &cobra.Command{
	Use:  "cat-file",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if catFilePrettyPrint {
			objectHash := args[0]
			objectFilePath := object.GetFilePathFromHash(objectHash)
			content, err := object.GetDecompressedFileContent(objectFilePath)
			if err != nil {
				fmt.Printf("Error reading object file: %v\n", err)
				return
			}

			obj, err := object.Parse(content)
			if err != nil {
				fmt.Printf("Error parsing object: %v\n", err)
				return
			}

			fmt.Printf("%s", obj.Content)
		}

	},
}

func init() {
	catFile.Flags().BoolVarP(&catFilePrettyPrint, "pretty-print", "p", false, "pretty-print <object> content")
	catFile.MarkFlagsOneRequired("pretty-print")
}
