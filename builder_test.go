package easyemails_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/hay-kot/easyemails"
	"github.com/hay-kot/easyemails/internal/snapshot"
)

func exampleEmail() *easyemails.Builder {
	return easyemails.NewBuilder().
		Logo("https://example.com/logo.png").
		Add(
			easyemails.WithParagraph(
				easyemails.WithText("Hello, **world**!"),
				easyemails.WithLineBreak(),
				easyemails.WithText("This is a test email, it works with [markdown](https://example.com)."),
				easyemails.WithList(
					"[Google](https://google.com) is a search engine.",
					"Item 2",
					"Item 3",
				),
				easyemails.WithLineBreak(),
				easyemails.WithText("It supports **bold** and *italic* text.").Centered(),
			),
			easyemails.WithImage("https://example.com/image.png").Alt("An example image").Centered(),
			easyemails.WithButton("Click me", "https://example.com").Centered(),
			easyemails.WithParagraph(
				easyemails.WithText("[Website](https://example.com/website) · [Unsubscribe](https://example.com/unsubscribe)").Centered(),
			).FontSize(12),
		)
}

func largeEmail() *easyemails.Builder {
	b := easyemails.NewBuilder()
	for i := range 50 {
		b.Add(
			easyemails.WithParagraph(
				easyemails.WithText("Paragraph "+strconv.Itoa(i)+" has **bold**, *italic*, and a [link](https://example.com)."),
				easyemails.WithList("one", "two", "three"),
			),
			easyemails.WithButton("Button "+strconv.Itoa(i), "https://example.com"),
		)
	}
	return b
}

func Test_Builder_HTMLSnapshot(t *testing.T) {
	snapshot.Match(t, ".html", exampleEmail().Render())
}

func Test_Builder_RenderPlainSnapshot(t *testing.T) {
	snapshot.Match(t, ".txt", exampleEmail().RenderPlain())
}

func Test_Builder_Render_NoLogo(t *testing.T) {
	rendered := easyemails.NewBuilder().
		Add(easyemails.WithParagraph(easyemails.WithText("Hello"))).
		Render()

	if strings.Contains(rendered, "<img") {
		t.Errorf("rendered output contains a logo:\n%s", rendered)
	}
	if strings.Contains(rendered, "border-top") {
		t.Errorf("rendered output contains the header divider:\n%s", rendered)
	}
}

func Benchmark_Builder_Render(b *testing.B) {
	b.Run("example", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			exampleEmail().Render()
		}
	})
	b.Run("large", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			largeEmail().Render()
		}
	})
}

func Benchmark_Builder_RenderPlain(b *testing.B) {
	b.Run("example", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			exampleEmail().RenderPlain()
		}
	})
	b.Run("large", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			largeEmail().RenderPlain()
		}
	})
}
