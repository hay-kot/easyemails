package components

import "strings"

type Text struct {
	alignment string
	text      string
	styles    []string
}

func NewText(text string) Text {
	return Text{text: text}
}

func (t Text) Style(property string, value string) Text {
	style := property + ": " + value + ";"
	t.styles = append(t.styles, style)
	return t
}

func (t Text) joinStyles() string {
	return strings.Join(t.styles, " ")
}

func (t Text) Centered() Text {
	t.alignment = "center"
	return t
}

func (t Text) Align(alignment string) Text {
	return t.Style("text-align", alignment)
}

func (t Text) Size(size string) Text {
	return t.Style("font-size", size)
}

func (t Text) ParagraphPlain() string {
	return t.text + "\n"
}

func (t Text) Paragraph() string {
	styles := t.joinStyles()
	return `<div style="` + styles + `">` + t.text + `</div>`
}
