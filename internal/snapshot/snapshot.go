// Package snapshot compares test output with snapshot files.
package snapshot

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// dir is the directory, relative to the package under test, that holds the
// snapshot files.
const dir = ".snapshots"

// updateEnv is the environment variable that makes Match write snapshots
// instead of comparing them.
const updateEnv = "UPDATE_SNAPSHOTS"

// Match compares got with the snapshot file for the current test,
// .snapshots/<test name><ext>, in the directory of the package under test. A
// "/" in a subtest name becomes "-" in the file name.
// Match fails the test when the file does not exist or does not match.
//
// When the UPDATE_SNAPSHOTS environment variable is set, Match writes got to
// the file and does not fail.
func Match(t testing.TB, ext, got string) {
	t.Helper()

	path := filepath.Join(dir, strings.ReplaceAll(t.Name(), "/", "-")+ext)
	// The trailing newline keeps the format of the files that cupaloy wrote.
	content := got + "\n"

	if _, update := os.LookupEnv(updateEnv); update {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			t.Fatalf("create snapshot directory: %v", err)
			return
		}
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatalf("write snapshot: %v", err)
			return
		}
		t.Logf("updated snapshot %s", path)
		return
	}

	want, err := os.ReadFile(path)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		t.Fatalf("snapshot %s does not exist, run the tests with %s=true to create it", path, updateEnv)
	case err != nil:
		t.Fatalf("read snapshot: %v", err)
	case string(want) != content:
		t.Errorf("snapshot %s does not match, run the tests with %s=true to update it:\n%s",
			path, updateEnv, diff(string(want), content))
	}
}
