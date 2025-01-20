package main

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/pb33f/libopenapi/index"
	"gopkg.in/yaml.v3"
)

type DummyFS struct {
}

func (d *DummyFS) Open(name string) (fs.File, error) {
	fmt.Printf("SEARCH %v\n", name)

	f, err := os.Open(name)
	if err != nil {
		fmt.Println("ERROROROROROR")
		panic(err)
	}

	return f, err
}

func (d *DummyFS) GetFiles() map[string]index.RolodexFile {
	fmt.Println("GET FILES!!!!!")
	m := make(map[string]index.RolodexFile)
	return m
}

func main() {
	fmt.Println("START")

	petstore, _ := os.ReadFile("openapi.yaml")

	var rootNode yaml.Node
	_ = yaml.Unmarshal(petstore, &rootNode)

	indexConfig := &index.SpecIndexConfig{
		AllowRemoteLookup:           false,
		AllowFileLookup:             true,
		AvoidCircularReferenceCheck: true,
	}

	rolodex := index.NewRolodex(indexConfig)

	// fsCfg := &index.LocalFSConfig{
	// 	BaseDirectory: "/",
	// 	IndexConfig:   indexConfig,
	// 	DirFS:         &DummyFS{},
	// }

	rolodex.AddLocalFS("/", &DummyFS{})

	// fileFS, err := index.NewLocalFSWithConfig(fsCfg)
	// if err != nil {
	// 	panic(err)
	// }

	// set the root node of the rolodex (this is the root of the spec)
	rolodex.SetRootNode(&rootNode)

	rolodex.AddLocalFS("/", &DummyFS{})

	indexingError := rolodex.IndexTheRolodex()
	if indexingError != nil {
		panic(indexingError)
	}

	// // resolve the petstore
	// rolodex.Resolve()

	// // extract the resolver from the root index.
	// resolver := rolodex.GetRootIndex().GetResolver()

	// // print out some interesting information discovered when visiting all the references.
	// fmt.Printf("%d errors repored\n", len(rolodex.GetCaughtErrors()))
	// fmt.Printf("%d references visited\n", resolver.GetReferenceVisited())
	// fmt.Printf("%d journeys taken\n", resolver.GetJourneysTaken())
	// fmt.Printf("%d index visits\n", resolver.GetIndexesVisited())
	// fmt.Printf("%d relatives seen\n", resolver.GetRelativesSeen())

	// // create a new document from specification bytes
	// document, err := libopenapi.NewDocument(petstore)

	// fmt.Printf("%v %v\n", err, document)

	// indexConfig := index.CreateClosedAPIIndexConfig()
	// rolodex := index.NewRolodex(indexConfig)
	// rolodex.AddLocalFS("/", &DummyFS{})

	// rolodex.Resolve()

	// resolver := rolodex.GetRootIndex().GetResolver()

	// fmt.Printf("%d errors repored\n", len(rolodex.GetCaughtErrors()))
	// fmt.Printf("%d references visited\n", resolver.GetReferenceVisited())
	// fmt.Printf("%d journeys taken\n", resolver.GetJourneysTaken())
	// fmt.Printf("%d index visits\n", resolver.GetIndexesVisited())
	// fmt.Printf("%d relatives seen\n", resolver.GetRelativesSeen())

	// rolodex.AddExternalIndex()
	// rolodex.AddRemoteFS()
	// f, err := os.ReadFile("openapi.yaml")

	// l := openapi3.NewLoader()
	// l.IsExternalRefsAllowed = true
	// doc, err := l.LoadFromFile("openapi.yaml")

	// fmt.Printf("%v %v\n", doc, err)
	// l.LoadFromData()
	// openapi3.T
}
