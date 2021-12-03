package generator

import (
	"embed"
)

var (
	//go:embed templates
	templatesFS embed.FS
)

type IndexesInfo struct {
	Comment   string
	ConstName string
	Label     string
	Method    string
	Use       bool
}

type FieldInfo struct {
	FsTag      string
	Field      string
	FieldType  string
	IsUnique   bool
	IndexerTag string
	Indexes    []*IndexesInfo
}

type templateParameter struct {
	AppVersion        string
	PackageName       string
	ImportName        string
	GeneratedFileName string
	FileName          string
	StructName        string
	StructNameRef     string
	ModelImportPath   string
	MockGenPath       string
	MockOutputPath    string

	RepositoryStructName    string
	RepositoryInterfaceName string

	KeyFieldName string
	KeyFieldType string
	KeyValueName string // lower camel case

	FieldInfos []*FieldInfo

	EnableIndexes       bool
	FieldInfoForIndexes *FieldInfo
	BoolCriteriaCnt     int
	SliceExist          bool

	AutomaticGeneration bool
	IsSubCollection     bool

	MetaFieldsEnabled bool
}
