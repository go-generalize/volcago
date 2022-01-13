//go:build !emulator
// +build !emulator

package generator

import (
	"os"
	"path/filepath"
	"testing"
)

func run(t *testing.T, structName string, subCollection bool) {
	t.Helper()

	gen, err := NewGenerator(".")

	if err != nil {
		t.Fatalf("NewGenerator failed: %+v", err)
	}

	opt := NewDefaultGenerateOption()
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

		run(t, "Task", false)
		run(t, "Lock", false)
	})

	t.Run("IDSpecified", func(tr *testing.T) {
		if err := os.Chdir(filepath.Join(root, "testfiles/not_auto")); err != nil {
			tr.Fatalf("chdir failed: %+v", err)
		}

		run(t, "Task", false)

		run(t, "SubTask", true)
	})
}
