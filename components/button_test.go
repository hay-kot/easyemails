package components_test

import (
	"testing"

	"github.com/hay-kot/easyemails/components"
	"github.com/hay-kot/easyemails/internal/snapshot"
)

func Test_Button_PlainSnapshot(t *testing.T) {
	btn := components.NewButton("Click me!", "https://example.com")

	snapshot.Match(t, ".txt", btn.RenderPlain())
}

func Test_Button_HTMLSnapshot(t *testing.T) {
	btn := components.NewButton("Click me!", "https://example.com")

	snapshot.Match(t, ".html", btn.Render())
}
