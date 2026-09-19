package components

import (
	"strconv"
	"strings"
)

// RenderableParagraph is a block inside a Paragraph. Paragraph returns HTML
// and ParagraphPlain returns plain text. The Paragraph inserts both as-is, so
// Paragraph must return valid HTML with all text escaped.
type RenderableParagraph interface {
	ParagraphPlain() string
	Paragraph() string
}

// Paragraph is a top-level block that groups Text, List, and LineBreak
// blocks.
type Paragraph struct {
	fontSize  int
	alignment Alignment
	styles    styles
	blocks    []RenderableParagraph
}

func NewParagraph(blocks ...RenderableParagraph) *Paragraph {
	return &Paragraph{blocks: blocks}
}

func (p *Paragraph) Add(blocks ...RenderableParagraph) *Paragraph {
	p.blocks = append(p.blocks, blocks...)
	return p
}

// FontSize sets the font size in pixels. The default is 16.
func (p *Paragraph) FontSize(px int) *Paragraph {
	p.fontSize = px
	return p
}

func (p *Paragraph) Align(alignment Alignment) *Paragraph {
	p.alignment = alignment
	return p
}

func (p *Paragraph) Centered() *Paragraph {
	return p.Align(AlignCenter)
}

// Style adds an inline CSS declaration to the element that wraps the
// paragraph's blocks.
func (p *Paragraph) Style(property string, value string) *Paragraph {
	p.styles.add(property, value)
	return p
}

func (p *Paragraph) RenderPlain() string {
	var bldr strings.Builder
	for _, block := range p.blocks {
		if _, ok := block.(LineBreak); ok {
			continue
		}

		bldr.WriteString(block.ParagraphPlain() + "\n")
	}
	return strings.TrimSpace(bldr.String())
}

func (p *Paragraph) Render() string {
	if len(p.blocks) == 0 {
		return ""
	}

	var (
		alignment = string(or(p.alignment, AlignLeft))
		fontSize  = strconv.Itoa(or(p.fontSize, 16))
		bldr      strings.Builder
	)

	bldr.WriteString(`<tr>
	<td align="` + alignment + `" style="font-size: 0px; padding: 10px 25px; word-break: break-word;" >
	  <div style="
		  font-family: Roboto, Helvetica Neue, Helvetica,
			Arial, sans-serif;
	  	font-size: ` + fontSize + `px;
		  line-height: 1;
		  text-align: ` + alignment + `;
		  color: #000000;` + p.styles.String() + `
		" >`)

	for _, block := range p.blocks {
		bldr.WriteString(block.Paragraph())
	}

	bldr.WriteString(`</div></td></tr>`)
	return bldr.String()
}
