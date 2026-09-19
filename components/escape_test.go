package components_test

import (
	"strings"
	"testing"

	"github.com/hay-kot/easyemails/components"
)

func Test_Render_EscapesAttributesAndText(t *testing.T) {
	const input = `"><script>x</script>`
	const escaped = `&quot;&gt;&lt;script&gt;x&lt;/script&gt;`

	tests := []struct {
		name     string
		rendered string
	}{
		{"Button text", components.NewButton(input, "https://example.com").Render()},
		{"Button URL", components.NewButton("Click", input).Render()},
		{"Image URL", components.NewImage(input).Render()},
		{"Image alt", components.NewImage("https://example.com/a.png").Alt(input).Render()},
		{"Image style", components.NewImage("https://example.com/a.png").Style("font-family", input).Render()},
		{"Text", components.NewText(input).Paragraph()},
		{"Text style", components.NewText("x").Style("font-family", input).Paragraph()},
		{"List item", components.NewList(input).Paragraph()},
		{"Paragraph style", components.NewParagraph(components.NewText("x")).Style("font-family", input).Render()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if strings.Contains(tt.rendered, "<script>") {
				t.Errorf("rendered output contains unescaped input:\n%s", tt.rendered)
			}
			if !strings.Contains(tt.rendered, escaped) {
				t.Errorf("rendered output does not contain %q:\n%s", escaped, tt.rendered)
			}
		})
	}
}
