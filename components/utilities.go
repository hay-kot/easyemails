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
	return strings.Join(s, " ")
}
