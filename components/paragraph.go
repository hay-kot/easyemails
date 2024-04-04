package components

import (
	"strconv"
	"strings"
)

type RenderableParagraph interface {
	ParagraphPlain() string
	Paragraph() string
}

func NewParagraph(blocks ...RenderableParagraph) *Paragraph {
	return &Paragraph{blocks: blocks}
}

type Paragraph struct {
	size   int64
	blocks []RenderableParagraph
}

func (p *Paragraph) FontSize(size int64) *Paragraph {
	p.size = size
	return p
}

func (p *Paragraph) RenderPlain() string {
	var bldr strings.Builder
	for _, block := range p.blocks {
		_, ok := block.(LineBreak)
		if ok {
			continue
		}

		bldr.WriteString(block.ParagraphPlain() + "\n")
	}
	return strings.TrimSpace(stripMarkup(bldr.String()))
}

func (p *Paragraph) Render() string {
	if len(p.blocks) == 0 {
		return ""
	}

	var bldr strings.Builder

	p.size = or(p.size, 16)

	bldr.WriteString(`<tr>
	<td align="left" style="font-size: 0px; padding: 10px 25px; word-break: break-word;" >
	  <div style="
		  font-family: Roboto, Helvetica Neue, Helvetica,
			Arial, sans-serif;
	  	font-size: ` + strconv.FormatInt(p.size, 10) + `px;
		  line-height: 1;
		  text-align: left;
		  color: #000000;
		" >`)

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
