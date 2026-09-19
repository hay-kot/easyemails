package components

import "strconv"

// Text is a line of text inside a Paragraph. The text supports inline markup:
// **bold**, *italic*, and [links](https://example.com).
type Text struct {
	text   string
	styles styles
}

func NewText(text string) *Text {
	return &Text{text: text}
}

// Style adds an inline CSS declaration to the text.
func (t *Text) Style(property string, value string) *Text {
	t.styles.add(property, value)
	return t
}

func (t *Text) Align(alignment Alignment) *Text {
	return t.Style("text-align", string(alignment))
}

func (t *Text) Centered() *Text {
	return t.Align(AlignCenter)
}

// FontSize sets the font size in pixels.
func (t *Text) FontSize(px int) *Text {
	return t.Style("font-size", strconv.Itoa(px)+"px")
}

func (t *Text) ParagraphPlain() string {
	return stripMarkup(t.text) + "\n"
}

func (t *Text) Paragraph() string {
	return `<div` + t.styles.attr() + `>` + inlineMarkup(t.text) + `</div>`
}
