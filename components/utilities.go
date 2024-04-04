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

// styles is a helper struct to manage styles within a component
// it provides a nieve way to add styles to a component and render them.
// It does not provide any validation or type checking. It will simply
// add the styles to the component by concatenating the string.
type styles struct {
	styles []string
}

func (s *styles) Style(property string, value string) {
	style := property + ": " + value + ";"
	s.styles = append(s.styles, style)
}

func (t *styles) string() string {
	return strings.Join(t.styles, " ")
}
