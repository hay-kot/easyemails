package components

type LineBreak struct{}

func (l LineBreak) Paragraph() string {
	return "<div><br></div>"
}
