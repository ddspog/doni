package cmd

import (
	"fmt"

	"github.com/ddspog/doni/project"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of Hugo",
	Long:  `All softwares has versions. This is doni's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dino experiment tool v" + project.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
