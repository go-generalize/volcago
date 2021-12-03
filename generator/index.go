package generator

import (
	"fmt"
	"sort"
	"strings"

	"github.com/fatih/structtag"
	"github.com/iancoleman/strcase"
)

const (
	equal      = "equal"
	like       = "like"
	prefix     = "prefix"
	suffix     = "suffix"
	biunigrams = "Biunigrams"
)

func isUseIndexer(filters []string, p1, p2 string) bool {
	for _, filter := range filters {
		switch filter {
		case p1, p2:
			return true
		}
	}
	return false
}

func (g *structGenerator) appendIndexer(tags *structtag.Tags, fsTagBase string, fieldInfo *FieldInfo) (*FieldInfo, error) {
	filters := make([]string, 0)
	if tags != nil {
		if tag, err := validateFirestoreTag(tags); err != nil {
			return nil, err
		} else if tag != "" {
			fieldInfo.FsTag = tag
			if fsTagBase != "" {
				fieldInfo.FsTag = fsTagBase + "." + tag
			}
		}

		idr, err := tags.Get("indexer")
		if err == nil {
			fieldInfo.IndexerTag = idr.Value()
			filters = strings.Split(idr.Value(), ",")
		}
	}

	patterns := [4]string{
		prefix, suffix, like, equal,
	}
	fieldLabel := g.structName + "IndexLabel"

	for i := range patterns {
		idx := &IndexesInfo{
			ConstName: strings.ReplaceAll(fieldLabel+fieldInfo.Field+strcase.ToCamel(patterns[i]), ".", "_"),
			Label:     uppercaseExtraction(fieldInfo.Field, g.dupMap),
			Method:    "Add",
		}

		switch patterns[i] {
		case prefix:
			idx.Use = isUseIndexer(filters, "p", prefix)
			idx.Method += strcase.ToCamel(prefix)
			idx.Comment = fmt.Sprintf("prefix-match of %s", fieldInfo.Field)
		case suffix:
			idx.Use = isUseIndexer(filters, "s", suffix)
			idx.Method += strcase.ToCamel(suffix)
			idx.Comment = fmt.Sprintf("suffix-match of %s", fieldInfo.Field)
		case like:
			idx.Use = isUseIndexer(filters, "l", like)
			idx.Method += biunigrams
			idx.Comment = fmt.Sprintf("like-match of %s", fieldInfo.Field)
		case equal:
			idx.Use = isUseIndexer(filters, "e", equal)
			idx.Comment = fmt.Sprintf("perfect-match of %s", fieldInfo.Field)
		}

		if fieldInfo.FieldType != typeString {
			idx.Method = "AddSomething"
		}

		fieldInfo.Indexes = append(fieldInfo.Indexes, idx)
	}

	sort.Slice(fieldInfo.Indexes, func(i, j int) bool {
		return fieldInfo.Indexes[i].Method < fieldInfo.Indexes[j].Method
	})

	return fieldInfo, nil
}

func uppercaseExtraction(name string, dupMap map[string]int) (lower string) {
	defer func() {
		if _, ok := dupMap[lower]; ok {
			lower = fmt.Sprintf("%s%d", lower, dupMap[lower])
		}
	}()
	for i, x := range name {
		switch {
		case 65 <= x && x <= 90:
			x += 32
			fallthrough
		case 97 <= x && x <= 122:
			if i == 0 {
				lower += string(x)
			}
			if _, ok := dupMap[lower]; !ok {
				dupMap[lower] = 1
				return
			}

			if dupMap[lower] >= 9 && len(name) > i+1 {
				lower += string(name[i+1])
				continue
			}
			dupMap[lower]++
			return
		}
	}
	return
}
