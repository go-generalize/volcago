package fsutil

import (
	"path/filepath"

	"golang.org/x/xerrors"
)

// IsSamePath returns whether the two paths are to the same
func IsSamePath(a, b string) (bool, error) {
	a, err := filepath.Abs(a)

	if err != nil {
		return false, xerrors.Errorf("failed to get absolute path for %s: %w", a, err)
	}

	b, err = filepath.Abs(b)

	if err != nil {
		return false, xerrors.Errorf("failed to get absolute path for %s: %w", b, err)
	}

	return filepath.Clean(a) == filepath.Clean(b), nil
}
