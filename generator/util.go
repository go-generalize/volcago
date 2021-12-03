package generator

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/fatih/structtag"
	"golang.org/x/xerrors"
)

var (
	valueCheck = regexp.MustCompile("^[0-9a-zA-Z_]+$")
)

func validateFirestoreTag(tags *structtag.Tags) (string, error) {
	fsTag, err := tags.Get("firestore")
	if err != nil {
		return "", nil
	}

	tag := strings.Split(fsTag.Value(), ",")[0]
	if !valueCheck.MatchString(tag) {
		return "", xerrors.New("key field for firestore should have other than blanks and symbols tag")
	}

	if unicode.IsDigit(rune(tag[0])) {
		return "", xerrors.New("key field for firestore should have indexerPrefix other than numbers required")
	}

	return tag, nil
}
