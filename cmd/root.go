package cmd

import (
	"io/fs"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/YibangHeng/same/file"
	"github.com/YibangHeng/same/group"
	"github.com/YibangHeng/same/render"
)

var rootCmd = &cobra.Command{
	Use:     "same",
	Short:   "Scan the same files in the folder",
	Long:    `Scan the same files in the folder.`,
	Version: "v0.1.0-dev",

	Run: func(cmd *cobra.Command, args []string) {
		var fileList []fs.DirEntry
		file.List(viper.GetString("file.directory"), &fileList)
		m := new(group.SizeGrouper).Group(fileList)
		fileList = fileList[0:0]
		for _, s := range m {
			if len(s) > 1 {
				fileList = append(fileList, s...)
			}
		}
		md5Grouper := new(group.MD5Grouper)
		render.Table(md5Grouper.Group(fileList), "MD5")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// File.
	rootCmd.Flags().StringP("directory", "d", ".", "Directory to scan")
	rootCmd.Flags().BoolP("recursive", "r", false, "Scan files recursively")

	// Format.
	rootCmd.Flags().Bool("no-trunc", false, "Do not truncate output")

	// Help
	rootCmd.Flags().BoolP("version", "v", false, "Display version")
	rootCmd.Flags().BoolP("help", "h", false, "Display help")

	_ = viper.BindPFlag("file.directory", rootCmd.Flag("directory"))
	_ = viper.BindPFlag("file.recursive", rootCmd.Flag("recursive"))
	_ = viper.BindPFlag("format.no-trunc", rootCmd.Flag("no-trunc"))
}
