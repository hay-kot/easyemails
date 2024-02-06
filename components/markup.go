package components

import "regexp"

var (
	inlineLinkRe   = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	inlineBoldRe   = regexp.MustCompile(`\*\*(.*?)\*\*`)
	inlineItalicRe = regexp.MustCompile(`\*(.*?)\*`)
)

// stripMarkup takes in a string of markup text and returns the string with all
// markup removed an transformed into plain text
func stripMarkup(markup string) string {
	markup = stripLinks(markup)
	markup = stripBold(markup)
	markup = stripItalic(markup)

	return markup
}

func stripBold(markup string) string {
	// Replace markdown-style bold with plain text
	result := inlineBoldRe.ReplaceAllString(markup, "$1")
	return result
}

func stripItalic(markup string) string {
	result := inlineItalicRe.ReplaceAllString(markup, "$1")
	return result
}

// transforms [link text](http://www.example.com) into link text
// link text http://www.example.com
func stripLinks(markup string) string {
	result := inlineLinkRe.ReplaceAllStringFunc(markup, func(match string) string {
		parts := inlineLinkRe.FindStringSubmatch(match)
		linkText := parts[1]
		linkURL := parts[2]
		return linkText + " " + linkURL
	})

	return result
}

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
