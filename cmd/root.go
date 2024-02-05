package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/yibangheng/same/constant"
	"github.com/yibangheng/same/file"
	"github.com/yibangheng/same/group"
	"github.com/yibangheng/same/render"
)

func root(cmd *cobra.Command, args []string) {
	// If no path specified, use ".".
	if len(args) == 0 {
		args = append(args, ".")
	}

	// Get all files.
	var fileList []file.EntryInfo
	file.List(&fileList, args...)

	// Find files with same size first.
	m := new(group.SizeGrouper).Group(fileList)
	fileList = fileList[0:0]
	for _, s := range m {
		fileList = append(fileList, s...)
	}

	// Find files with same md5 value and render.
	render.Render(new(group.MD5Grouper).Group(fileList))
}

var rootCmd = &cobra.Command{
	Use:                   "same [FLAG]... [DIRECTORY]...",
	DisableFlagsInUseLine: true,
	Short:                 "Find same files in folder(s)",
	Long:                  `Find same files in folder(s).`,
	Version:               constant.Version,
	Run:                   root,
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
	rootCmd.Flags().Bool("ignore-empty-file", false, "Do not count empty file")
	rootCmd.Flags().Bool("ignore-hidden-file", false, "Do not count hidden file")

	// Format.
	rootCmd.Flags().BoolP("json", "j", false, "Print in JSON format")
	rootCmd.Flags().Bool("no-trunc", false, "Do not truncate output")

	// Help.
	rootCmd.Flags().BoolP("version", "v", false, "Display version")
	rootCmd.Flags().BoolP("help", "h", false, "Display help")

	_ = viper.BindPFlag("file.recursive", rootCmd.Flag("recursive"))
	_ = viper.BindPFlag("file.ignore-empty-file", rootCmd.Flag("ignore-empty-file"))
	_ = viper.BindPFlag("file.ignore-hidden-file", rootCmd.Flag("ignore-hidden-file"))
	_ = viper.BindPFlag("format.json", rootCmd.Flag("json"))
	_ = viper.BindPFlag("format.no-trunc", rootCmd.Flag("no-trunc"))

	// Version.
	rootCmd.SetVersionTemplate(`{{printf "%s" .Version}}{{"\n"}}`)
}
