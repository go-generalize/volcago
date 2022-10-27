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
	FsTag          string
	Field          string
	FieldType      string
	IsUnique       bool
	IsDocumentID   bool
	IndexerTag     string
	Indexes        []*IndexesInfo
	NullableFields []string
}

type UniqueInfo struct {
	Field string
	FsTag string
}

type templateParameter struct {
	AppVersion        string
	PackageName       string
	CollectionName    string
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

	FieldInfos  []*FieldInfo
	UniqueInfos []*UniqueInfo

	EnableIndexes       bool
	FieldInfoForIndexes *FieldInfo
	BoolCriteriaCnt     int
	SliceExist          bool

	AutomaticGeneration bool
	IsSubCollection     bool

	MetaFieldsEnabled bool
}
