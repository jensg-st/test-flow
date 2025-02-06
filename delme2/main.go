package main

import (
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"

	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/bundler"
	"github.com/pb33f/libopenapi/datamodel"
	"github.com/pb33f/libopenapi/index"
)

type direktivOpenAPIFS struct {
	// fileStore filestore.FileStore
	// ns        string
}

func (d *direktivOpenAPIFS) Open(name string) (fs.File, error) {

	if name == "." {

	}
	_, file, no, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("called from %s#%d\n", file, no)
	}

	fmt.Printf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!RESOLVE %v\n", name)

	// ff := filepath.Join("/home/jens/go/src/github/jensg-st/test-flow", name)
	ff := name
	f, err := os.Open(ff)
	fmt.Println(err)

	fmt.Printf("DONE %v %v\n", f, err)
	return f, err
}

func (d *direktivOpenAPIFS) GetFiles() map[string]index.RolodexFile {
	fmt.Println("GET FILES!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	_, file, no, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("get files called from %s#%d\n", file, no)
	}
	return make(map[string]index.RolodexFile)
}

func main() {

	// check out the mother of all exploded specifications
	// tmp, _ := os.MkdirTemp("", "openapi")
	// cmd := exec.Command("git", "clone", "https://github.com/digitalocean/openapi", tmp)
	// defer os.RemoveAll(filepath.Join(tmp, "openapi"))

	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }

	tmp := "/home/jens/go/src/github/jensg-st/test-flow"

	spec, _ := filepath.Abs(filepath.Join(tmp+"/deep1/deep21", "gateway.yaml"))
	specBytes, _ := os.ReadFile(spec)

	// localFSConf := index.LocalFSConfig{
	// 	BaseDirectory: "/home/jens/go/src/github/jensg-st/test-flow",
	// 	// IndexConfig:   idxConfig,
	// 	// FileFilters:   config.FileFilter,
	// 	// BasePath: "/home/jens/go/src/github/jensg-st/test-flow",
	// 	DirFS: &direktivOpenAPIFS{},
	// }
	// fileFS, _ := index.NewLocalFSWithConfig(&localFSConf)
	// fmt.Println(fileFS)
	doc, err := libopenapi.NewDocumentWithConfiguration([]byte(specBytes), &datamodel.DocumentConfiguration{
		BasePath:                tmp + "/deep1/deep21",
		ExtractRefsSequentially: true,
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})),
		AllowFileReferences:   true,
		AllowRemoteReferences: true,
		// LocalFS:               fileFS,
		LocalFS: &direktivOpenAPIFS{},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("-----------------------1")

	v3Doc, errs := doc.BuildV3Model()
	if len(errs) > 0 {
		panic(errs)
	}

	fmt.Println("-----------------------2")

	bytes, e := bundler.BundleDocument(&v3Doc.Model)
	fmt.Println(e)
	fmt.Println(string(bytes))
	//... do something with the bytes
}
