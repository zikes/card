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
	var output strings.Builder

	if s.Header != "" {
		output.WriteString(s.Header + "\n")
		output.WriteString(strings.Repeat("-", len(s.Header)) + "\n")
		if s.HeaderMargin > 0 {
			output.WriteString(strings.Repeat("\n", s.HeaderMargin))
		}
	}

	longest := 0

	for _, v := range s.Items {
		if len(v.Label) > longest {
			longest = len(v.Label)
		}
	}

	for _, v := range s.Items {
		if v.Label == "" && v.Data == "" {
			output.WriteString("\n")
			continue
		}
		output.WriteString(strings.Repeat(" ", s.ListIndent))
		output.WriteString(color.New(color.Bold).Sprint(align.Right(longest, v.Label)))
		output.WriteString(": ")
		output.WriteString(v.Data)
		output.WriteString("\n")
	}

	return output.String()
}
