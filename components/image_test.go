package components_test

import (
	"testing"

	"github.com/hay-kot/easyemails/components"
	"github.com/hay-kot/easyemails/internal/snapshot"
)

func Test_Image_PlainSnapshot(t *testing.T) {
	img := components.NewImage("https://example.com/image.png").
		Alt("An example image").
		Style("border", "1px solid #000").
		Centered()

	snapshot.Match(t, ".txt", img.RenderPlain())
}

func Test_Image_HTMLSnapshot(t *testing.T) {
	img := components.NewImage("https://example.com/image.png").
		Alt("An example image").
		Style("border", "1px solid #000").
		Centered()

	snapshot.Match(t, ".html", img.Render())
}
