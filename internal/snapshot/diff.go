package snapshot

import "strings"

// contextLines is the number of unchanged lines that Diff shows around each
// change.
const contextLines = 3

type edit struct {
	op   byte // ' ' for an unchanged line, '-' for a line only in a, '+' for a line only in b
	text string
}

// diff returns a line diff from a to b. Lines only in a start with "-", lines
// only in b start with "+", and unchanged lines start with " ". Unchanged
// lines more than contextLines away from a change collapse into "...".
//
// diff uses a longest common subsequence table, which takes time and memory
// proportional to the product of the line counts. That is fine for snapshots
// of a few hundred lines.
func diff(a, b string) string {
	edits := lineEdits(strings.Split(a, "\n"), strings.Split(b, "\n"))

	keep := make([]bool, len(edits))
	for i, e := range edits {
		if e.op == ' ' {
			continue
		}
		for j := max(0, i-contextLines); j <= min(len(edits)-1, i+contextLines); j++ {
			keep[j] = true
		}
	}

	var out strings.Builder
	skipped := false
	for i, e := range edits {
		if !keep[i] {
			if !skipped {
				out.WriteString("...\n")
				skipped = true
			}
			continue
		}
		skipped = false
		out.WriteByte(e.op)
		out.WriteByte(' ')
		out.WriteString(e.text)
		out.WriteByte('\n')
	}
	return out.String()
}

func lineEdits(a, b []string) []edit {
	// lcs[i][j] is the length of the longest common subsequence of a[i:] and
	// b[j:].
	lcs := make([][]int, len(a)+1)
	for i := range lcs {
		lcs[i] = make([]int, len(b)+1)
	}
	for i := len(a) - 1; i >= 0; i-- {
		for j := len(b) - 1; j >= 0; j-- {
			if a[i] == b[j] {
				lcs[i][j] = lcs[i+1][j+1] + 1
			} else {
				lcs[i][j] = max(lcs[i+1][j], lcs[i][j+1])
			}
		}
	}

	// On a tie, the walk takes the line from a first, so a removed line
	// comes before the line that replaces it.
	edits := make([]edit, 0, max(len(a), len(b)))
	i, j := 0, 0
	for i < len(a) || j < len(b) {
		switch {
		case i < len(a) && j < len(b) && a[i] == b[j]:
			edits = append(edits, edit{' ', a[i]})
			i++
			j++
		case i < len(a) && (j == len(b) || lcs[i+1][j] >= lcs[i][j+1]):
			edits = append(edits, edit{'-', a[i]})
			i++
		default:
			edits = append(edits, edit{'+', b[j]})
			j++
		}
	}
	return edits
}
