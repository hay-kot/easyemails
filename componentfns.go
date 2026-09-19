package easyemails

import "github.com/hay-kot/easyemails/components"

// WithParagraph returns a block that groups Text, List, and LineBreak blocks.
func WithParagraph(blocks ...components.RenderableParagraph) *components.Paragraph {
	return components.NewParagraph(blocks...)
}

// WithText returns a line of text for a Paragraph. The text supports inline
// markup: **bold**, *italic*, and [links](https://example.com).
func WithText(text string) *components.Text {
	return components.NewText(text)
}

// WithLineBreak returns an empty line for a Paragraph.
func WithLineBreak() components.LineBreak {
	return components.LineBreak{}
}

// WithList returns a bulleted list for a Paragraph. Call Ordered on the result
// for a numbered list.
func WithList(items ...string) *components.List {
	return components.NewList(items...)
}

// WithButton returns a block with a button that links to url.
func WithButton(text, url string) *components.Button {
	return components.NewButton(text, url)
}

// WithImage returns a block with an image.
func WithImage(url string) *components.Image {
	return components.NewImage(url)
}
