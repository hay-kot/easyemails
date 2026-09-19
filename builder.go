package easyemails

import (
	_ "embed"
	"html"
	"strings"
)

// The Default values configure every Builder that NewBuilder creates. Set them
// once at program start when all of your emails share a logo or colors, and
// use the Builder methods to change a single email. Changes do not affect
// builders that already exist.
var (
	// DefaultLogo is the URL of the image at the top of the email.
	DefaultLogo = "https://placehold.co/800x200"
	// DefaultPrimaryColor is the color of buttons and links.
	DefaultPrimaryColor = "#0c4a6e"
	// DefaultPrimaryTextColor is the color of text on a DefaultPrimaryColor
	// background, such as a button label.
	DefaultPrimaryTextColor = "#ffffff"
	// DefaultBorderColor is the color of the borders and dividers.
	DefaultBorderColor = "#d1d5db"
)

//go:embed templates/basetemplate.html
var template string

// Renderable is a top-level block of an email. Render returns one or more
// table rows (<tr>) and RenderPlain returns the block as plain text.
//
// The Builder inserts the output of Render into the document as-is, so it
// must be valid HTML with all text escaped. The Builder replaces these tokens
// in the output with its colors: {{ .PrimaryColor }}, {{ .PrimaryTextColor }},
// and {{ .BorderColor }}.
type Renderable interface {
	RenderPlain() string
	Render() string
}

// Builder holds the blocks and settings of one email.
type Builder struct {
	blocks           []Renderable
	logo             string
	primaryColor     string
	primaryTextColor string
	borderColor      string
}

// NewBuilder returns a Builder with the Default values.
func NewBuilder() *Builder {
	return &Builder{
		logo:             DefaultLogo,
		primaryColor:     DefaultPrimaryColor,
		primaryTextColor: DefaultPrimaryTextColor,
		borderColor:      DefaultBorderColor,
	}
}

// Logo sets the URL of the image at the top of the email.
func (b *Builder) Logo(url string) *Builder {
	b.logo = url
	return b
}

// PrimaryColor sets the color of buttons and links.
func (b *Builder) PrimaryColor(color string) *Builder {
	b.primaryColor = color
	return b
}

// PrimaryTextColor sets the color of text on a primary color background.
func (b *Builder) PrimaryTextColor(color string) *Builder {
	b.primaryTextColor = color
	return b
}

// BorderColor sets the color of the borders and dividers.
func (b *Builder) BorderColor(color string) *Builder {
	b.borderColor = color
	return b
}

// Add appends blocks to the email.
func (b *Builder) Add(blocks ...Renderable) *Builder {
	b.blocks = append(b.blocks, blocks...)
	return b
}

// RenderPlain returns the email as plain text, with a blank line between
// blocks.
func (b *Builder) RenderPlain() string {
	var bldr strings.Builder

	for _, block := range b.blocks {
		plain := block.RenderPlain()
		if plain == "" {
			continue
		}
		if bldr.Len() > 0 {
			bldr.WriteString("\n\n")
		}
		bldr.WriteString(plain)
	}

	return bldr.String()
}

// Render returns the email as an HTML document.
func (b *Builder) Render() string {
	var bldr strings.Builder

	for _, block := range b.blocks {
		bldr.WriteString(block.Render())
	}

	rendered := strings.Replace(template, "{{ .Content }}", bldr.String(), 1)
	rendered = strings.Replace(rendered, "{{ .Logo }}", html.EscapeString(b.logo), 1)
	rendered = strings.ReplaceAll(rendered, "{{ .PrimaryColor }}", b.primaryColor)
	rendered = strings.ReplaceAll(rendered, "{{ .PrimaryTextColor }}", b.primaryTextColor)
	rendered = strings.ReplaceAll(rendered, "{{ .BorderColor }}", b.borderColor)

	return rendered
}
