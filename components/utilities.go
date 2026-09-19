package components

import "strings"

// or returns the first non-zero value
func or[T comparable](a, b T) T {
	var zero T

	if a == zero {
		return b
	}
	return a
}

// styles holds inline CSS declarations in the order they were added. It does
// no validation or deduplication, so for a repeated property the last
// declaration wins, as CSS specifies.
type styles []string

func (s *styles) add(property, value string) {
	*s = append(*s, property+": "+value+";")
}

func (s styles) String() string {
	return escape(strings.Join(s, " "))
}

// suffix returns the declarations with a leading space, for appending to a
// style attribute that already has declarations.
func (s styles) suffix() string {
	if len(s) == 0 {
		return ""
	}
	return " " + s.String()
}

// attr returns a style attribute with a leading space, or an empty string
// when there are no declarations.
func (s styles) attr() string {
	if len(s) == 0 {
		return ""
	}
	return ` style="` + s.String() + `"`
}

// htmlEscaper escapes text for HTML text nodes and double-quoted attribute
// values. It leaves apostrophes alone, unlike html.EscapeString, because
// every attribute in the output uses double quotes.
var htmlEscaper = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
	`"`, "&quot;",
)

func escape(s string) string {
	return htmlEscaper.Replace(s)
}

// fontFamily is a system font stack. Most email clients ignore web fonts, so
// the template loads none and every component uses this stack.
const fontFamily = `-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif`

// blockCellStyle is the style of the table cell around each top-level block.
// All blocks share it so that their edges line up.
const blockCellStyle = "font-size: 0px; padding: 10px 25px; word-break: break-word;"
