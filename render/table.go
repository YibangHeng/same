package render

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/ozgio/strutil"
	"github.com/spf13/viper"

	"github.com/YibangHeng/same/group"
)

type Any = group.Any
type Type = group.Type

var (
	keyLength = 5
	emptyTail = ""
)

// trunc truncates a string to keyLength length,
// but keep others (e.g. number) as it is.
//
// If --no-trunc specified, string will also be
// as it is.
func trunc(a Any) string {
	if str, ok := a.(string); ok && !viper.GetBool("format.no-trunc") {
		return strutil.Summary(str, keyLength, emptyTail)
	} else {
		return str
	}
}

func Table(m map[Any][]Type, keys string) {
	tb := table.NewWriter()
	tb.AppendHeader(table.Row{keys, "FILES"})

	tb.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AutoMerge: true},
	})
	tb.SetOutputMirror(os.Stdout)

	for k, v := range m {
		if len(v) > 1 {
			for _, vv := range v {
				tb.AppendRow(table.Row{trunc(k), vv.Name()})
			}
		}
	}
	tb.Render()
}
