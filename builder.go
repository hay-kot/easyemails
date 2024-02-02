package easyemails

import (
	_ "embed"
	"strings"
)

var (
	ImageLogoHeader = "https://placehold.co/800x200"
	// ColorPrimary is the default primary color for the email template. This
	// can be overridden with the WithPrimaryColor method on each builder. When
	// all you're emails have the same color, you can set this variable
	// to avoid repeating the same color in every builder.
	ColorPrimary = "#0c4a6e"
	// ColorPrimaryText is the default primary color for text for the email template.
	// This can be overridden with the WithPrimaryColor method on each builder. When
	// all you're emails have the same color, you can set this variable
	// to avoid repeating the same color in every builder.
	ColorPrimaryText = "#ffffff"

	// ColorBorder is the default border color for the email template. This
	// can be overridden with the WithBorderColor method on each builder. When
	// all you're emails have the same color, you can set this variable
	// to avoid repeating the same color in every builder.
	ColorBorder = "#d1d5db"
)

//go:embed templates/basetemplate.html
var template string

// Renderable is an interface that all blocks must implement to be rendered
// in the email template. Any string returned by Renderable is assumed to be
// valid HTML for use within an email.
//
// Note that the result of Render() will _still_ be processed by the markup function
// to format links, bold, and italic texts.
type Renderable interface {
	Render() string
}

func or(a, b string) string {
	if a == "" {
		return b
	}
	return a
}

// Builder is a struct that holders the blocks to be rendered in the email.
type Builder struct {
	blocks       []Renderable
	logo         string
	primaryColor string
	primaryText  string
	borderColor  string
}

// NewBuilder creates a new Builder with the default values.
func NewBuilder() *Builder {
	return &Builder{
		logo:         ImageLogoHeader,
		primaryColor: ColorPrimary,
		primaryText:  ColorPrimaryText,
		borderColor:  ColorBorder,
	}
}

// WithLogo sets the logo for the email template.
func (b *Builder) WithLogo(logo string) *Builder {
	b.logo = logo
	return b
}

// WithPrimaryColor sets the primary color and text color for the buttons or
// other accent elements.
func (b *Builder) WithPrimaryColor(color string, text string) *Builder {
	b.primaryColor = color
	b.primaryText = text
	return b
}

func (b *Builder) WithBorderColor(color string) *Builder {
	b.borderColor = color
	return b
}

func (b *Builder) Render() string {
	var (
		pc = or(b.primaryColor, ColorPrimary)
		bc = or(b.borderColor, ColorBorder)
		pt = or(b.primaryText, ColorPrimaryText)
		l  = or(b.logo, ImageLogoHeader)
	)

	var bldr strings.Builder

	for _, block := range b.blocks {
		bldr.WriteString(block.Render())
	}

	rendered := strings.Replace(template, "{{ .Content }}", bldr.String(), 1)
	rendered = strings.Replace(rendered, "{{ .ImageHeader }}", l, 1)
	rendered = strings.ReplaceAll(rendered, "{{ .PrimaryColor }}", pc)
	rendered = strings.ReplaceAll(rendered, "{{ .PrimaryText }}", pt)
	rendered = strings.ReplaceAll(rendered, "{{ .BorderColor }}", bc)

	return rendered
}

func (b *Builder) Add(block ...Renderable) *Builder {
	b.blocks = append(b.blocks, block...)
	return b
}
