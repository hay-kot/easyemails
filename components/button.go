package components

// Button is a top-level block that renders a link as a button in the
// builder's primary colors.
type Button struct {
	text      string
	url       string
	alignment Alignment
}

func NewButton(text, url string) *Button {
	return &Button{text: text, url: url}
}

func (b *Button) Align(alignment Alignment) *Button {
	b.alignment = alignment
	return b
}

func (b *Button) Centered() *Button {
	return b.Align(AlignCenter)
}

func (b *Button) RenderPlain() string {
	return b.text + " " + b.url
}

func (b *Button) Render() string {
	alignment := string(or(b.alignment, AlignLeft))

	return `<tr>
	<td align="` + alignment + `" 
		vertical-align="middle"
		style="
		font-size: 0px;
		padding: 10px 25px;
		word-break: break-word;
		"
	>
		<table border="0" cellpadding="0" cellspacing="0" role="presentation"
           style="border-collapse: separate; width: 40%; line-height: 100%;"
		>
			<tbody>
				<tr>
					<td
						align="center" 
		        bgcolor="{{ .PrimaryColor }}" 
	        	role="presentation" valign="middle"
						style="
						border: none;
						border-radius: 3px;
						cursor: auto;
						mso-padding-alt: 10px 25px;
						background: {{ .PrimaryColor }};
						"
					>
						<a
						href="` + escape(b.url) + `"
						style="
							display: inline-block;
							background: {{ .PrimaryColor }};
							color: {{ .PrimaryTextColor }};
							font-family: Ubuntu, Helvetica, Arial,
							sans-serif;
							font-size: 13px;
							font-weight: normal;
							line-height: 120%;
							margin: 0;
							text-decoration: none;
							text-transform: none;
							padding: 10px 25px;
							mso-padding-alt: 0px;
							border-radius: 3px;
						"
						>
						` + escape(b.text) + `
						</a>
					</td>
				</tr>
			</tbody>
		</table>
	</td>
</tr>
`
}
