package components

import (
	"regexp"
	"strings"
)

var (
	inlineLinkRe   = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	inlineBoldRe   = regexp.MustCompile(`\*\*(.*?)\*\*`)
	inlineItalicRe = regexp.MustCompile(`\*(.*?)\*`)
)

// stripMarkup converts inline markup to plain text. A link becomes its text
// followed by its URL.
func stripMarkup(markup string) string {
	return replaceMarkup(markup, "$1 $2", "$1", "$1")
}

// inlineMarkup escapes markup for HTML, then converts the inline markup to HTML
// elements. The markup characters (*, [, ], (, and )) need no escape, so the
// conversion still finds them.
func inlineMarkup(markup string) string {
	return replaceMarkup(escape(markup), `<a href="$2">$1</a>`, "<strong>$1</strong>", "<em>$1</em>")
}

// replaceMarkup replaces links, then bold, then italic text with the regexp
// templates. The order matters: bold must go before italic, because the
// italic pattern also matches "**".
//
// Most text has little or no markup. The Contains checks skip each regexp
// when its pattern cannot match, and cost much less than a regexp pass.
func replaceMarkup(s, link, bold, italic string) string {
	if strings.Contains(s, "](") {
		s = inlineLinkRe.ReplaceAllString(s, link)
	}
	if strings.Contains(s, "**") {
		s = inlineBoldRe.ReplaceAllString(s, bold)
	}
	if strings.Contains(s, "*") {
		s = inlineItalicRe.ReplaceAllString(s, italic)
	}
	return s
}
