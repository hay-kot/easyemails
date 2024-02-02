package components

func orDefault(a, b string) string {
	if a == "" {
		return b
	}
	return a
}
