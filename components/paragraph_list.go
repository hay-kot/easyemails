package components

import (
	"strconv"
	"strings"
)

// List is a bulleted or numbered list inside a Paragraph. Items support the
// same inline markup as Text.
type List struct {
	ordered bool
	items   []string
}

func NewList(items ...string) *List {
	return &List{items: items}
}

// Ordered makes the list numbered instead of bulleted.
func (l *List) Ordered() *List {
	l.ordered = true
	return l
}

func (l *List) ParagraphPlain() string {
	var bldr strings.Builder

	for i, item := range l.items {
		if l.ordered {
			bldr.WriteString(strconv.Itoa(i+1) + ". ")
		} else {
			bldr.WriteString("- ")
		}
		bldr.WriteString(stripMarkup(item) + "\n")
	}

	return bldr.String()
}

func (l *List) Paragraph() string {
	tag, styleType := "ul", "disc"
	if l.ordered {
		tag, styleType = "ol", "decimal"
	}

	var bldr strings.Builder

	bldr.WriteString(`<div><` + tag + ` style="list-style-type: ` + styleType + `; line-height: 1.3;">`)

	for _, item := range l.items {
		bldr.WriteString("<li>" + inlineMarkup(item) + "</li>")
	}

	bldr.WriteString(`</` + tag + `></div>`)

	return bldr.String()
}
