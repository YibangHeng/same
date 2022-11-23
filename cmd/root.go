package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/YibangHeng/same/file"
	"github.com/YibangHeng/same/group"
	"github.com/YibangHeng/same/render"
)

func root(cmd *cobra.Command, args []string) {
	// If no path specified, use ".".
	if len(args) == 0 {
		args = append(args, ".")
	}

	// Get all files.
	var fileList []file.FileInfo
	file.List(&fileList, args...)

	// Find files with same size first.
	m := new(group.SizeGrouper).Group(fileList)
	fileList = fileList[0:0]
	for _, s := range m {
		fileList = append(fileList, s...)
	}

	// Find files with same md5 value and render.
	render.Table(new(group.MD5Grouper).Group(fileList), "MD5")
}

var rootCmd = &cobra.Command{
	Use:     "same",
	Short:   "Scan the same files in the folder",
	Long:    `Scan the same files in the folder.`,
	Version: "v0.1.0-dev",
	Run:     root,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// File.
	rootCmd.Flags().BoolP("recursive", "r", false, "Scan files recursively")

	// Format.
	rootCmd.Flags().Bool("no-trunc", false, "Do not truncate output")

	// Help
	rootCmd.Flags().BoolP("version", "v", false, "Display version")
	rootCmd.Flags().BoolP("help", "h", false, "Display help")

	_ = viper.BindPFlag("file.recursive", rootCmd.Flag("recursive"))
	_ = viper.BindPFlag("format.no-trunc", rootCmd.Flag("no-trunc"))
}
