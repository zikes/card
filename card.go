package card

import (
	"strings"

	"github.com/fatih/color"
	"github.com/zikes/align"
)

type Section struct {
	Header       string
	HeaderMargin int
	ListIndent   int
	Items        []Item
}

type Item struct {
	Label string
	Data  string
}

func (s Section) String() string {
	output := ""

	if s.Header != "" {
		output += s.Header + "\n"
		output += strings.Repeat("-", len(s.Header)) + "\n"
		if s.HeaderMargin > 0 {
			output += strings.Repeat("\n", s.HeaderMargin)
		}
	}

	longest := 0

	for _, v := range s.Items {
		if len(v.Label) > longest {
			longest = len(v.Label)
		}
	}

	for _, v := range s.Items {
		output += strings.Repeat(" ", s.ListIndent)
		output += color.New(color.Bold).Sprint(align.Right(longest, v.Label))
		output += ": "
		output += v.Data
		output += "\n"
	}

	return output
}
