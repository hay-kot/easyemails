package components

import "regexp"

var (
	inlineLinkRe   = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	inlineBoldRe   = regexp.MustCompile(`\*\*(.*?)\*\*`)
	inlineItalicRe = regexp.MustCompile(`\*(.*?)\*`)
)

// runs all the inline markup functions
func inlineMarkup(markup string) string {
	markup = inlineLinks(markup)
	markup = inlineBold(markup)
	markup = inlineItalic(markup)

	return markup
}

func inlineBold(markup string) string {
	// Replace markdown-style bold with HTML bold
	result := inlineBoldRe.ReplaceAllString(markup, "<strong>$1</strong>")

	return result
}

func inlineItalic(markup string) string {
	// Replace markdown-style italics with HTML italics
	result := inlineItalicRe.ReplaceAllString(markup, "<em>$1</em>")

	return result
}

// Takes in a string of markup text and returns the string with all markdown
// style links converted to inline links
func inlineLinks(markup string) string {
	// Looking for
	//
	// [link text](http://www.example.com)
	//
	// and replacing with
	//
	// <a href="http://www.example.com">link text</a>

	// Replace markdown-style links with inline links
	result := inlineLinkRe.ReplaceAllStringFunc(markup, func(match string) string {
		parts := inlineLinkRe.FindStringSubmatch(match)
		linkText := parts[1]
		linkURL := parts[2]
		inlineLink := "<a href=\"" + linkURL + "\">" + linkText + "</a>"
		return inlineLink
	})

	return result
}
