package components

// Image is a top-level block that shows an image at the full width of the
// email. Use Style to set a different width.
type Image struct {
	url       string
	alt       string
	alignment Alignment
	styles    styles
}

func NewImage(url string) *Image {
	return &Image{url: url}
}

// Alt sets the alternative text. Plain text output uses the alternative text
// in place of the image.
func (i *Image) Alt(text string) *Image {
	i.alt = text
	return i
}

// Style adds an inline CSS declaration to the img element.
func (i *Image) Style(property string, value string) *Image {
	i.styles.add(property, value)
	return i
}

func (i *Image) Align(alignment Alignment) *Image {
	i.alignment = alignment
	return i
}

func (i *Image) Centered() *Image {
	return i.Align(AlignCenter)
}

func (i *Image) RenderPlain() string {
	return i.alt
}

// margins aligns the img. The img is display: block, so the td align
// attribute has no effect on it.
func (i *Image) margins() string {
	switch i.alignment {
	case AlignCenter:
		return " margin-left: auto; margin-right: auto;"
	case AlignRight:
		return " margin-left: auto; margin-right: 0;"
	default:
		return ""
	}
}

func (i *Image) Render() string {
	return `
<tr>
  <td style="width: 213px">
    <img
      height="auto"
      src="` + i.url + `"
      alt="` + i.alt + `"
      style="
        border: 0;
        display: block;
        outline: none;
        text-decoration: none;
        height: auto;
        width: 100%;
        font-size: 13px;` + i.margins() + ` ` + i.styles.String() + `
      "
    />
  </td>
</tr>
  `
}
