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
	// DefaultLogo is the URL of the image at the top of the email. An empty
	// value means no logo.
	DefaultLogo = ""
	// DefaultPrimaryColor is the color of buttons and links.
	DefaultPrimaryColor = "#0c4a6e"
	// DefaultPrimaryTextColor is the color of text on a DefaultPrimaryColor
	// background, such as a button label.
	DefaultPrimaryTextColor = "#ffffff"
	// DefaultBorderColor is the color of the borders and dividers.
	DefaultBorderColor = "#d1d5db"
)

// template is a single-column layout. Outlook for Windows ignores max-width and
// padding on a div, so an Outlook-only conditional table gives it the same
// width and padding.
//
//go:embed templates/basetemplate.html
var template string

//go:embed templates/header.html
var headerTemplate string

// templateHead, templateMiddle, and templateTail are the parts of the template
// around the {{ .Header }} and {{ .Content }} tokens. Render writes the parts
// and the blocks between them into one buffer.
var templateHead, templateMiddle, templateTail = splitTemplate(template)

func splitTemplate(t string) (head, middle, tail string) {
	head, rest, ok := strings.Cut(t, "{{ .Header }}")
	if !ok {
		panic("easyemails: template has no {{ .Header }} token")
	}
	middle, tail, ok = strings.Cut(rest, "{{ .Content }}")
	if !ok {
		panic("easyemails: template has no {{ .Content }} token")
	}
	return head, middle, tail
}

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

// Logo sets the URL of the image at the top of the email. An empty URL
// removes the logo and the divider below it.
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
	var (
		content = make([]string, len(b.blocks))
		size    = len(template) + len(headerTemplate) + len(b.logo)
	)
	for i, block := range b.blocks {
		content[i] = block.Render()
		size += len(content[i])
	}

	// The tokens are longer than typical values, so size is an upper bound
	// unless the colors or the logo URL are unusually long.
	var w strings.Builder
	w.Grow(size)

	b.writeExpanded(&w, templateHead)
	if b.logo != "" {
		b.writeExpanded(&w, headerTemplate)
	}
	b.writeExpanded(&w, templateMiddle)
	for _, c := range content {
		b.writeExpanded(&w, c)
	}
	b.writeExpanded(&w, templateTail)

	return w.String()
}

// writeExpanded writes s to w with each token replaced by its value. It writes
// unknown tokens unchanged.
func (b *Builder) writeExpanded(w *strings.Builder, s string) {
	for {
		start := strings.Index(s, "{{ .")
		if start < 0 {
			break
		}
		end := strings.Index(s[start:], " }}")
		if end < 0 {
			break
		}
		end += start + len(" }}")

		w.WriteString(s[:start])
		w.WriteString(b.tokenValue(s[start:end]))
		s = s[end:]
	}
	w.WriteString(s)
}

func (b *Builder) tokenValue(token string) string {
	switch token {
	case "{{ .Logo }}":
		return html.EscapeString(b.logo)
	case "{{ .PrimaryColor }}":
		return b.primaryColor
	case "{{ .PrimaryTextColor }}":
		return b.primaryTextColor
	case "{{ .BorderColor }}":
		return b.borderColor
	default:
		return token
	}
}
