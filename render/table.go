package render

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/ozgio/strutil"
	"github.com/spf13/viper"
	"github.com/yibangheng/same/file"
)

var (
	keyLength = 5
	emptyTail = ""
)

// trunc truncates a string to keyLength length,
// but keep others (e.g. number) as it is.
//
// If --no-trunc specified, string will also be
// as it is.
func trunc(a file.Any) string {
	if str, ok := a.(string); ok && !viper.GetBool("format.no-trunc") {
		return strutil.Summary(str, keyLength, emptyTail)
	} else {
		return str
	}
}

func Table(m map[file.Any][]file.EntryInfoType, keys string) {
	if len(m) == 0 {
		if viper.GetBool("file.recursive") {
			fmt.Fprint(os.Stdout, "No same file found.\n")
		} else {
			fmt.Fprint(os.Stdout, "No same file found. Did you forget to use -r to scan sub folders?\n")
		}
		return
	}

	tb := table.NewWriter()
	tb.SetStyle(table.Style{
		Box:     table.BoxStyle{MiddleVertical: "  "},
		Options: table.Options{SeparateColumns: true},
	})
	tb.AppendHeader(table.Row{keys, "FILES"})

	tb.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AutoMerge: true},
	})
	tb.SetOutputMirror(os.Stdout)

	for k, v := range m {
		for _, vv := range v {
			tb.AppendRow(table.Row{trunc(k), vv.GetRelativeName()})
		}
	}
	tb.Render()
}
