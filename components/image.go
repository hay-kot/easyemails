package components

type Image struct {
	styles *styles
	url    string
}

func NewImage(url string) Image {
	return Image{
		url:    url,
		styles: &styles{},
	}
}

func (i Image) Style(property string, value string) Image {
	i.styles.Style(property, value)
	return i
}

func (i Image) Centered() Image {
	i.styles.Style("display", "block")
	i.styles.Style("margin-left", "auto")
	i.styles.Style("margin-right", "auto")
	return i
}

func (i Image) RenderPlain() string {
	return i.url
}

func (i Image) Render() string {
	styles := i.styles.string()

	return `
<tr>
  <td style="width: 213px">
    <img
      height="auto"
      src="` + i.url + `"
      style="
        border: 0;
        display: block;
        outline: none;
        text-decoration: none;
        height: auto;
        width: 100%;
        font-size: 13px; ` + styles + `
      "
    />
  </td>
</tr>
  `
}
