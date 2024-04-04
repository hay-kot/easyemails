package components

type Button struct {
	text      string
	url       string
	alignment string
}

func (b Button) Text(text string) Button {
	b.text = text
	return b
}

func (b Button) URL(url string) Button {
	b.url = url
	return b
}

func (b Button) Align(alignment string) Button {
	b.alignment = alignment
	return b
}

func (b Button) RenderPlain() string {
	return b.text + " " + b.url
}

func (b Button) Render() string {
	alignment := or(b.alignment, "left")

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
						href="` + b.url + `"
						style="
							display: inline-block;
							background: {{ .PrimaryColor }};
							color: {{ .PrimaryText }};
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
						` + b.text + `
						</a>
					</td>
				</tr>
			</tbody>
		</table>
	</td>
</tr>
`
}
