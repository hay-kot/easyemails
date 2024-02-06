package components

type Text struct {
	alignment string
	text      string
}

func NewText(text string) Text {
	return Text{text: text}
}

func (t Text) Centered() Text {
	t.alignment = "center"
	return t
}

func (t Text) Align(alignment string) Text {
	t.alignment = alignment
	return t
}

func (t Text) ParagraphPlain() string {
	return t.text + "\n"
}

func (t Text) Paragraph() string {
	alignment := orDefault(t.alignment, "left")
	return `<div style="text-align: ` + alignment + `;">` + t.text + `</div>`
}
