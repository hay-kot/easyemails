package components

import "strings"

type List struct {
	listType string
	items    []string
}

func NewList(items ...string) *List {
	return &List{items: items}
}

func (l *List) Ordered() *List {
	l.listType = "ol"
	return l
}

func (l *List) ParagraphPlain() string {
	var bldr strings.Builder

	for _, item := range l.items {
		bldr.WriteString("- " + item + "\n")
	}

	return bldr.String()
}

func (l *List) Paragraph() string {
	var bldr strings.Builder

	listType := orDefault(l.listType, "ul")

	styleType := "disc"
	if listType == "ol" {
		styleType = "decimal"
	}

	bldr.WriteString(`<div><` + listType + ` style="list-style-type: ` + styleType + `; line-height: 1.3;">`)

	for _, item := range l.items {
		bldr.WriteString("<li>" + item + "</li>")
	}

	bldr.WriteString(`</` + listType + `></div>`)

	return bldr.String()
}
