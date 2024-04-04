package components_test

import (
	"testing"

	"github.com/hay-kot/easyemails/components"
)

func Test_Image_PlainSnapshot(t *testing.T) {
	img := components.NewImage("https://example.com/image.png").
		Style("border", "1px solid #000").
		Centered()

	ss := htmlSnapshot()

	ss.SnapshotT(t, img.RenderPlain())
}

func Test_Image_HTMLSnapshot(t *testing.T) {
	img := components.NewImage("https://example.com/image.png").
		Style("border", "1px solid #000").
		Centered()

	ss := htmlSnapshot()

	ss.SnapshotT(t, img.Render())
}
