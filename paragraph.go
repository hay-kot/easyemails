package easyemails

import (
	"strconv"
	"strings"
)

type RenderableParagraph interface {
	Paragraph() string
}

func WithParagraph(blocks ...RenderableParagraph) *Paragraph {
	return &Paragraph{blocks: blocks}
}

func WithParagraphSize(size int64, blocks ...RenderableParagraph) *Paragraph {
	return &Paragraph{size: size, blocks: blocks}
}

type Paragraph struct {
	size   int64
	blocks []RenderableParagraph
}

func (p *Paragraph) Render() string {
	var bldr strings.Builder

	if p.size == 0 {
		p.size = 16
	}

	if len(p.blocks) == 0 {
		return ""
	}

	bldr.WriteString(`
	<tr>
	<td
	  align="left"
	  style="
		font-size: 0px;
		padding: 10px 25px;
		word-break: break-word;
	  "
	>
	  <div
		style="
		  font-family: Roboto, Helvetica Neue, Helvetica,
			Arial, sans-serif;
		font-size: `+ strconv.FormatInt(p.size, 10) + `px;
		  line-height: 1;
		  text-align: left;
		  color: #000000;
		"
	  >
	`)

	for _, block := range p.blocks {
		bldr.WriteString(inlineMarkup(block.Paragraph()))
	}

	bldr.WriteString(`</div></td></tr>`)
	return bldr.String()
}

func (p *Paragraph) Add(block RenderableParagraph) *Paragraph {
	p.blocks = append(p.blocks, block)
	return p
}

func WithText(text string) Text {
	return Text{text: text}
}

func WithTextCentered(text string) Text {
	return Text{text: text, center: true}
}

type Text struct {
	center bool
	text   string
}

func (t Text) Paragraph() string {
	if t.center {
		return "<div style=\"text-align: center;\">" + t.text + "</div>"
	}

	return "<div>" + t.text + "</div>"
}

func WithLineBreak() LineBreak {
	return LineBreak{}
}

type LineBreak struct{}

func (l LineBreak) Paragraph() string {
	return "<div><br></div>"
}

func WithList(items ...string) *List {
	return &List{items: items}
}

type List struct {
	items []string
}

func (l *List) Paragraph() string {
	var bldr strings.Builder

	bldr.WriteString("<div><ul style=\"list-style-type: disc; line-height: 1.3;\">")

	for _, item := range l.items {
		bldr.WriteString("<li>" + item + "</li>")
	}

	bldr.WriteString("</ul></div>")

	return bldr.String()
}
