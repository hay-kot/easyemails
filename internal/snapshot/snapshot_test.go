package snapshot

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_diff(t *testing.T) {
	tests := []struct {
		name string
		a, b string
		want string
	}{
		{
			name: "Changed line",
			a:    "a\nb\nc",
			b:    "a\nB\nc",
			want: "  a\n- b\n+ B\n  c\n",
		},
		{
			name: "Added line",
			a:    "a\nb",
			b:    "a\nb\nc",
			want: "  a\n  b\n+ c\n",
		},
		{
			name: "Removed line",
			a:    "a\nb\nc",
			b:    "a\nc",
			want: "  a\n- b\n  c\n",
		},
		{
			name: "Unchanged lines far from a change collapse",
			a:    "1\n2\n3\n4\n5\n6\n7\n8\n9\n10",
			b:    "1\n2\n3\n4\n5\nfive\n7\n8\n9\n10",
			want: "...\n  3\n  4\n  5\n- 6\n+ five\n  7\n  8\n  9\n...\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diff(tt.a, tt.b); got != tt.want {
				t.Errorf("diff() =\n%s\nwant\n%s", got, tt.want)
			}
		})
	}
}

// recorder records a failure instead of failing the test, so that the tests
// can check the failure paths of Match.
type recorder struct {
	testing.TB
	failure string
}

func (r *recorder) Errorf(format string, args ...any) {
	r.failure = fmt.Sprintf(format, args...)
}

func (r *recorder) Fatalf(format string, args ...any) {
	r.failure = fmt.Sprintf(format, args...)
}

func Test_Match(t *testing.T) {
	t.Chdir(t.TempDir())
	// t.Setenv restores the value from the caller after the test, so the test
	// also works under UPDATE_SNAPSHOTS=true.
	t.Setenv(updateEnv, "")
	unsetUpdateEnv(t)

	r := &recorder{TB: t}
	Match(r, ".txt", "a\nb")
	if !strings.Contains(r.failure, "does not exist") {
		t.Fatalf("Match() with no snapshot: failure = %q, want a missing snapshot failure", r.failure)
	}

	t.Setenv(updateEnv, "true")
	Match(t, ".txt", "a\nb")
	written, err := os.ReadFile(filepath.Join(dir, "Test_Match.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if string(written) != "a\nb\n" {
		t.Fatalf("snapshot file = %q, want %q", written, "a\nb\n")
	}
	unsetUpdateEnv(t)

	Match(t, ".txt", "a\nb")

	r = &recorder{TB: t}
	Match(r, ".txt", "a\nc")
	if !strings.Contains(r.failure, "- b\n+ c\n") {
		t.Fatalf("Match() with a changed value: failure = %q, want a diff", r.failure)
	}
}

func unsetUpdateEnv(t *testing.T) {
	t.Helper()
	if err := os.Unsetenv(updateEnv); err != nil {
		t.Fatal(err)
	}
}
