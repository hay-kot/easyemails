package components

// LineBreak is an empty line inside a Paragraph. Plain text output omits it,
// because every block in a plain text paragraph already ends with a blank
// line.
type LineBreak struct{}

func (l LineBreak) ParagraphPlain() string {
	return "\n"
}

func (l LineBreak) Paragraph() string {
	return "<div><br></div>"
}
