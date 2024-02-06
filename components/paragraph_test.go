package components_test

import (
	"testing"

	"github.com/hay-kot/easyemails/components"
)

func Test_Paragraph_PlainSnapshot(t *testing.T) {
	para := components.NewParagraph(
		components.NewText("**Hello**, *World*!"),
		components.NewText("This is a test paragraph."),
		components.NewList("Item 1", "Item 2", "Item 3"),
		components.NewText("Goodbye!"),
	)

	ss := textSnapshot()

	ss.SnapshotT(t, para.RenderPlain())
}

func Test_Paragraph_HTMLSnapshot(t *testing.T) {
	para := components.NewParagraph(
		components.NewText("**Hello**, *World*!"),
		components.NewText("This is a test paragraph."),
		components.NewList("Item 1", "Item 2", "Item 3"),
		components.NewText("Goodbye!"),
	)

	ss := htmlSnapshot()

	ss.SnapshotT(t, para.Render())
}
