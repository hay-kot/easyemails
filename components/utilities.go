package components

func orDefault[T comparable](a, b T) T {
	var zero T

	if a == zero {
		return b
	}
	return a
}
