package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-generalize/volcago/generator"
)

var (
	isShowVersion   = flag.Bool("v", false, "print version")
	isSubCollection = flag.Bool("sub-collection", false, "is SubCollection")
	outputDir       = flag.String("o", "./", "Specify directory to generate code in")
	packageName     = flag.String("p", "", "Specify the package name, default is the same as the original package")
	collectionName  = flag.String("c", "", "Specify the collection name, default is the same as the struct name")
	mockGenPath     = flag.String("mockgen", "mockgen", "Specify mockgen path")
	mockOutputPath  = flag.String("mock-output", "mock/mock_{{ .GeneratedFileName }}/mock_{{ .GeneratedFileName }}.go", "Specify directory to generate mock code in")
)

func main() {
	flag.Parse()

	if *isShowVersion {
		fmt.Printf("Firestore Model Generator: %s\n", AppVersion)
		return
	}

	l := flag.NArg()
	if l < 1 {
		fmt.Fprintln(os.Stderr, "You have to specify the struct name of target")
		os.Exit(1)
	}

	gen, err := generator.NewGenerator(".")

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize generator: %+v\n", err)
		os.Exit(1)
	}
	gen.AppVersion = AppVersion

	structName := flag.Arg(0)

	err = gen.Generate(structName, generator.GenerateOption{
		OutputDir:      *outputDir,
		PackageName:    *packageName,
		CollectionName: *collectionName,
		MockGenPath:    *mockGenPath,
		MockOutputPath: *mockOutputPath,
		Subcollection:  *isSubCollection,
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to generate repository for %s: %+v\n", structName, err)
		os.Exit(1)
	}
}
