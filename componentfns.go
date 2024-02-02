package easyemails

import "github.com/hay-kot/easyemails/components"

func WithButton(text, url string) components.Button {
	return components.Button{}.
		Text(text).
		URL(url)
}

func WithParagraph(blocks ...components.RenderableParagraph) *components.Paragraph {
	return components.NewParagraph(blocks...)
}

func WithText(text string) components.Text {
	return components.NewText(text)
}

func WithLineBreak() components.LineBreak {
	return components.LineBreak{}
}

func WithList(items ...string) *components.List {
	return components.NewList(items...)
}
