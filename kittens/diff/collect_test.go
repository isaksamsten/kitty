// License: GPLv3 Copyright: 2023, Kovid Goyal, <kovid at kovidgoyal.net>

package diff

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"kitty/tools/utils"

	"github.com/google/go-cmp/cmp"
)

var _ = fmt.Print

func TestDiffCollectWalk(t *testing.T) {
	tdir := t.TempDir()
	j := func(x ...string) string { return filepath.Join(append([]string{tdir}, x...)...) }
	os.MkdirAll(j("a", "b"), 0o700)
	os.WriteFile(j("a/b/c"), nil, 0o600)
	os.WriteFile(j("b"), nil, 0o600)
	os.WriteFile(j("d"), nil, 0o600)
	os.WriteFile(j("e"), nil, 0o600)
	os.WriteFile(j("#d#"), nil, 0o600)
	os.WriteFile(j("e~"), nil, 0o600)
	os.MkdirAll(j("f"), 0o700)
	os.WriteFile(j("f/g"), nil, 0o600)
	os.WriteFile(j("h space"), nil, 0o600)

	expected_names := utils.NewSetWithItems("d", "e", "f/g", "h space")
	expected_pmap := map[string]string{
		"d":       j("d"),
		"e":       j("e"),
		"f/g":     j("f/g"),
		"h space": j("h space"),
	}
	names := utils.NewSet[string](16)
	pmap := make(map[string]string, 16)
	if err := walk(tdir, []string{"*~", "#*#", "b"}, names, pmap, map[string]string{}); err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(
		utils.Sort(expected_names.AsSlice(), func(a, b string) bool { return a < b }),
		utils.Sort(names.AsSlice(), func(a, b string) bool { return a < b }),
	); diff != "" {
		t.Fatal(diff)
	}
	if diff := cmp.Diff(expected_pmap, pmap); diff != "" {
		t.Fatal(diff)
	}
}
