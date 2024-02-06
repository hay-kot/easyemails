package components

type LineBreak struct{}

func (l LineBreak) ParagraphPlain() string {
	return "\n"
}

func (l LineBreak) Paragraph() string {
	return "<div><br></div>"
}
