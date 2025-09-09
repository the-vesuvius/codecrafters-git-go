package cmd

import (
	"fmt"

	"github.com/codecrafters-io/git-starter-go/app/object"
	"github.com/spf13/cobra"
)

var nameOnly bool

var lsTree = &cobra.Command{
	Use:   "ls-tree",
	Args:  cobra.MinimumNArgs(1),
	Short: "List the contents of a tree object",
	Run: func(cmd *cobra.Command, args []string) {
		treeSha := args[0]
		if err := object.LsTree(treeSha, nameOnly); err != nil {
			fmt.Printf("Error listing tree contents: %v\n", err)
			return
		}
	},
}

func init() {
	lsTree.Flags().BoolVarP(&nameOnly, "name-only", "", false, "list only filenames")
}
