package easyemails

import "fmt"

func WithButton(text, url string) Button {
	return Button{text: text, url: url}
}

type Button struct {
	text string
	url  string
}

func (b Button) Render() string {
	return fmt.Sprintf(`<tr>
	<td align="center" vertical-align="middle"
		style="
		font-size: 0px;
		padding: 10px 25px;
		word-break: break-word;
		"
	>
		<table border="0" cellpadding="0" cellspacing="0" role="presentation"
		style="
			border-collapse: separate;
			width: 40%%;
			line-height: 100%%;
		"
		>
			<tbody>
				<tr>
					<td
						align="center" bgcolor="{{ .PrimaryColor }}" role="presentation" valign="middle"
						style="
						border: none;
						border-radius: 3px;
						cursor: auto;
						mso-padding-alt: 10px 25px;
						background: {{ .PrimaryColor }};
						"
					>
						<a
						href="%s"
						style="
							display: inline-block;
							background: {{ .PrimaryColor }};
							color: {{ .PrimaryText }};
							font-family: Ubuntu, Helvetica, Arial,
							sans-serif;
							font-size: 13px;
							font-weight: normal;
							line-height: 120%%;
							margin: 0;
							text-decoration: none;
							text-transform: none;
							padding: 10px 25px;
							mso-padding-alt: 0px;
							border-radius: 3px;
						"
						>
						%s
						</a>
					</td>
				</tr>
			</tbody>
		</table>
	</td>
</tr>
`, b.url, b.text)
}
