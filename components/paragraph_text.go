package components

type Text struct {
	text      string
	styles    *styles
}

func NewText(text string) Text {
	return Text{
		text:   text,
		styles: &styles{},
	}
}

func (t Text) Style(property string, value string) Text {
	t.styles.Style(property, value)
	return t
}

func (t Text) Centered() Text {
	t.styles.Style("text-align", "center")
	return t
}

func (t Text) Align(alignment string) Text {
	t.styles.Style("text-align", alignment)
	return t
}

func (t Text) Size(size string) Text {
	t.styles.Style("font-size", size)
	return t
}

func (t Text) ParagraphPlain() string {
	return t.text + "\n"
}

func (t Text) Paragraph() string {
	styles := t.styles.string()
	return `<div style="` + styles + `">` + t.text + `</div>`
}
