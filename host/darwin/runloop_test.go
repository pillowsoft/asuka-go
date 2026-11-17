//go:build darwin && !ios

package darwin

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/egoist/asuka/internal/darwin"
)

// commonModes is the constant, not a string that reads the same.
func TestCommonModesIsTheConstant(t *testing.T) {
	onMain(func() {
		Load()
		if commonModes() == 0 {
			t.Fatal("commonModes is nil")
		}
		if commonModes() == String("kCFRunLoopCommonModes") {
			t.Fatal("commonModes is a new string, which a run loop takes for an ordinary mode")
		}
	})
}

// No file passes a run-loop mode as a new string: see commonModes.
func TestNoRunLoopModeStrings(t *testing.T) {
	files, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		if strings.HasSuffix(f, "_test.go") {
			continue
		}
		b, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		for i, line := range strings.Split(string(b), "\n") {
			if strings.Contains(line, `String("kCFRunLoop`) || strings.Contains(line, `String("NSRunLoop`) || strings.Contains(line, `String("NSDefaultRunLoopMode`) {
				t.Errorf("%s:%d passes a run-loop mode as a new string; use the constant (commonModes)", f, i+1)
			}
		}
	}
}
