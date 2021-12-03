//go:build emulator
// +build emulator

package generator

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func execTest(t *testing.T) {
	t.Helper()

	b, err := exec.Command("go", "test", "./tests", "-v", "-tags", "internal").CombinedOutput()

	if err != nil {
		t.Fatalf("go test failed: %+v(%s)", err, string(b))
	}
}

func run(t *testing.T, structName string, useMeta, subCollection bool) {
	t.Helper()

	gen, err := NewGenerator(".")

	if err != nil {
		t.Fatalf("NewGenerator failed: %+v", err)
	}

	opt := NewDefaultGenerateOption()
	opt.UseMetaField = useMeta
	opt.Subcollection = subCollection

	if err := gen.Generate(structName, opt); err != nil {
		t.Fatalf("Generate failed: %+v", err)
	}
}

func TestGenerator(t *testing.T) {
	root, err := os.Getwd()

	if err != nil {
		t.Fatalf("failed to getwd: %+v", err)
	}

	t.Run("AutomaticIDGeneration", func(tr *testing.T) {
		if err := os.Chdir(filepath.Join(root, "testfiles/auto")); err != nil {
			tr.Fatalf("chdir failed: %+v", err)
		}

		run(t, "Task", false, false)
		run(t, "Lock", true, false)

		execTest(tr)
	})

	t.Run("IDSpecified", func(tr *testing.T) {
		if err := os.Chdir(filepath.Join(root, "testfiles/not_auto")); err != nil {
			tr.Fatalf("chdir failed: %+v", err)
		}

		run(t, "Task", false, false)

		run(t, "SubTask", false, true)

		execTest(tr)
	})
}
