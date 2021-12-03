package generator

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/go-utils/cont"
	"github.com/go-utils/plural"
)

func (g *structGenerator) getFuncMap() template.FuncMap {
	return template.FuncMap{
		"Parse": func(fieldType string) string {
			fieldType = strings.TrimPrefix(fieldType, "[]")
			fn := "Int"
			switch fieldType {
			case typeInt:
			case typeInt64:
				fn = "Int64"
			case typeFloat64:
				fn = "Float64"
			case typeString:
				fn = "String"
			case typeBool:
				fn = "Bool"
			default:
				panic("invalid types")
			}
			return fn
		},
		"HasSuffix": func(s, suffix string) bool {
			return strings.HasSuffix(s, suffix)
		},
		"HasSlice": func(types string) bool {
			return strings.HasPrefix(types, "[]")
		},
		"HasMap": func(types string) bool {
			return strings.HasPrefix(types, typeMap)
		},
		"PluralForm": func(word string) string {
			return plural.Convert(word)
		},
		"IndexerInfo": func(fieldInfo *FieldInfo) (comment string) {
			if fieldInfo.IndexerTag == "" {
				return
			}
			comment += fmt.Sprintf(`// The value of the "indexer" tag = "%s"`, fieldInfo.IndexerTag)
			items := make([]string, 0)
			for _, index := range fieldInfo.Indexes {
				if !index.Use {
					continue
				}
				if !cont.Contains(items, index.Method) {
					items = append(items, index.Method)
				}
			}
			if len(items) > 3 {
				comment += "\n\t\t\t// "
				comment += strings.Join(items, "/")
				comment += " is valid."
			}
			return
		},
		"GetFunc": func() string {
			raw := fmt.Sprintf(
				"Get(ctx context.Context, %s %s, opts ...GetOption) (*%s, error)",
				g.param.KeyValueName, g.param.KeyFieldType, g.param.StructNameRef,
			)
			return raw
		},
		"GenerateUpdateParam": func(fis []*FieldInfo) string {
			buf := bytes.Buffer{}

			layers := make([]string, 0)
			for _, f := range fis {
				if f.IsUnique {
					continue
				}

				split := strings.Split(f.Field, ".")

				common := 0
				for common < len(split)-1 &&
					common < len(layers) &&
					split[common] == layers[common] {
					common++
				}

				for i := len(layers) - 1; i >= common; i-- {
					buf.WriteString("}\n")
				}
				for i := common; i < len(split)-1; i++ {
					buf.WriteString(fmt.Sprintf("%s struct {\n", split[i]))
				}
				layers = split[:len(split)-1]

				buf.WriteString(fmt.Sprintf("%s interface{}\n", split[len(split)-1]))
			}

			for i := len(layers) - 1; i >= 0; i-- {
				buf.WriteString("}\n")
			}

			return buf.String()
		},
		"GenerateSearchParam": func(fis []*FieldInfo) string {
			buf := bytes.Buffer{}

			layers := make([]string, 0)
			for _, f := range fis {
				split := strings.Split(f.Field, ".")

				common := 0
				for common < len(split)-1 &&
					common < len(layers) &&
					split[common] == layers[common] {
					common++
				}

				for i := len(layers) - 1; i >= common; i-- {
					buf.WriteString("}\n")
				}
				for i := common; i < len(split)-1; i++ {
					buf.WriteString(fmt.Sprintf("%s struct {\n", split[i]))
				}
				layers = split[:len(split)-1]

				buf.WriteString(fmt.Sprintf("%s *QueryChainer\n", split[len(split)-1]))
			}

			for i := len(layers) - 1; i >= 0; i-- {
				buf.WriteString("}\n")
			}

			return buf.String()
		},
		"GetWithDocFunc": func() string {
			raw := fmt.Sprintf(
				"GetWithDoc(ctx context.Context, doc *firestore.DocumentRef, opts ...GetOption) (*%s, error)",
				g.param.StructNameRef,
			)
			return raw
		},
		"InsertFunc": func() string {
			return fmt.Sprintf("Insert(ctx context.Context, subject *%s) (_ %s, err error)", g.param.StructNameRef, g.param.KeyFieldType)
		},
		"UpdateFunc": func() string {
			return fmt.Sprintf("Update(ctx context.Context, subject *%s) (err error)", g.param.StructNameRef)
		},
		"StrictUpdateFunc": func() string {
			return fmt.Sprintf(
				"StrictUpdate(ctx context.Context, id string, param *%sUpdateParam, opts ...firestore.Precondition) error",
				g.param.StructName,
			)
		},
		"DeleteFunc": func() string {
			return fmt.Sprintf("Delete(ctx context.Context, subject *%s, opts ...DeleteOption) (err error)", g.param.StructNameRef)
		},
		"DeleteByFunc": func() string {
			raw := fmt.Sprintf(
				"DeleteBy%s(ctx context.Context, %s %s, opts ...DeleteOption) (err error)",
				g.param.KeyFieldName, g.param.KeyValueName, g.param.KeyFieldType,
			)
			return raw
		},
		"GetMultiFunc": func() string {
			raw := fmt.Sprintf(
				"GetMulti(ctx context.Context, %s []%s, opts ...GetOption) ([]*%s, error)",
				plural.Convert(g.param.KeyValueName), g.param.KeyFieldType, g.param.StructNameRef,
			)
			return raw
		},
		"InsertMultiFunc": func() string {
			return fmt.Sprintf("InsertMulti(ctx context.Context, subjects []*%s) (_ []%s, er error)", g.param.StructNameRef, g.param.KeyFieldType)
		},
		"UpdateMultiFunc": func() string {
			return fmt.Sprintf("UpdateMulti(ctx context.Context, subjects []*%s) (er error)", g.param.StructNameRef)
		},
		"DeleteMultiFunc": func() string {
			return fmt.Sprintf("DeleteMulti(ctx context.Context, subjects []*%s, opts ...DeleteOption) (er error)", g.param.StructNameRef)
		},
		"DeleteMultiByFunc": func() string {
			raw := fmt.Sprintf(
				"DeleteMultiBy%s(ctx context.Context, %s []%s, opts ...DeleteOption) (er error)",
				plural.Convert(g.param.KeyFieldName), plural.Convert(g.param.KeyValueName), g.param.KeyFieldType,
			)
			return raw
		},
		"SearchFunc": func() string {
			return fmt.Sprintf(
				"Search(ctx context.Context, param *%sSearchParam, q *firestore.Query) ([]*%s, error)",
				g.param.StructName, g.param.StructNameRef)
		},
		"GetWithTxFunc": func() string {
			raw := fmt.Sprintf(
				"GetWithTx(tx *firestore.Transaction, %s %s, opts ...GetOption) (*%s, error)",
				g.param.KeyValueName, g.param.KeyFieldType, g.param.StructNameRef,
			)
			return raw
		},
		"GetWithDocWithTxFunc": func() string {
			raw := fmt.Sprintf(
				"GetWithDocWithTx(tx *firestore.Transaction, doc *firestore.DocumentRef, opts ...GetOption) (*%s, error)",
				g.param.StructNameRef,
			)
			return raw
		},
		"InsertWithTxFunc": func() string {
			return fmt.Sprintf(
				"InsertWithTx(ctx context.Context, tx *firestore.Transaction, subject *%s) (_ %s, err error)",
				g.param.StructNameRef, g.param.KeyFieldType,
			)
		},
		"UpdateWithTxFunc": func() string {
			return fmt.Sprintf("UpdateWithTx(ctx context.Context, tx *firestore.Transaction, subject *%s) (err error)", g.param.StructNameRef)
		},
		"StrictUpdateWithTxFunc": func() string {
			return fmt.Sprintf(
				"StrictUpdateWithTx(tx *firestore.Transaction, id string, param *%sUpdateParam, opts ...firestore.Precondition) error",
				g.param.StructName,
			)
		},
		"DeleteWithTxFunc": func() string {
			return fmt.Sprintf(
				"DeleteWithTx(ctx context.Context, tx *firestore.Transaction, subject *%s, opts ...DeleteOption) (err error)",
				g.param.StructNameRef,
			)
		},
		"DeleteByWithTxFunc": func() string {
			return fmt.Sprintf(
				"DeleteBy%sWithTx(ctx context.Context, tx *firestore.Transaction, %s %s, opts ...DeleteOption) (err error)",
				g.param.KeyFieldName, g.param.KeyValueName, g.param.KeyFieldType,
			)
		},
		"SearchWithTxFunc": func() string {
			return fmt.Sprintf(
				"SearchWithTx(tx *firestore.Transaction, param *%sSearchParam, q *firestore.Query) ([]*%s, error)",
				g.param.StructName, g.param.StructNameRef)
		},
		"GetMultiWithTxFunc": func() string {
			return fmt.Sprintf(
				"GetMultiWithTx(tx *firestore.Transaction, %s []%s, opts ...GetOption) ([]*%s, error)",
				plural.Convert(g.param.KeyValueName), g.param.KeyFieldType, g.param.StructNameRef,
			)
		},
		"InsertMultiWithTxFunc": func() string {
			return fmt.Sprintf(
				"InsertMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*%s) (_ []string, er error)",
				g.param.StructNameRef,
			)
		},
		"UpdateMultiWithTxFunc": func() string {
			return fmt.Sprintf("UpdateMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*%s) (er error)", g.param.StructNameRef)
		},
		"DeleteMultiWithTxFunc": func() string {
			return fmt.Sprintf("DeleteMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*%s, opts ...DeleteOption) (er error)", g.param.StructNameRef)
		},
		"DeleteMultiByWithTxFunc": func() string {
			raw := fmt.Sprintf(
				"DeleteMultiBy%sWithTx(ctx context.Context, tx *firestore.Transaction, %s []%s, opts ...DeleteOption) (er error)",
				plural.Convert(g.param.KeyFieldName), plural.Convert(g.param.KeyValueName), g.param.KeyFieldType,
			)
			return raw
		},
		"LookUpFieldByName": func(fieldInfos []*FieldInfo, name string) *FieldInfo {
			for _, fi := range fieldInfos {
				if fi.Field == name {
					return fi
				}
			}

			return nil
		},
	}
}
