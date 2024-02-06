package components_test

import (
	"testing"

	"github.com/hay-kot/easyemails/components"
)

func Test_Button_PlainSnapshot(t *testing.T) {
	btn := components.Button{}.
		Text("Click me!").
		URL("https://example.com")

	ss := textSnapshot()

	ss.SnapshotT(t, btn.RenderPlain())
}

func Test_Button_HTMLSnapshot(t *testing.T) {
	btn := components.Button{}.
		Text("Click me!").
		URL("https://example.com")

	ss := htmlSnapshot()

	ss.SnapshotT(t, btn.Render())
}
