package main

import (
	"context"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"runtime"

	"github.com/getkin/kin-openapi/openapi3"
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

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		fmt.Println(url)
		// return []byte(""), nil
		return os.ReadFile(url.String())
	}

	// loader.

	doc, err := loader.LoadFromFile("../deep1/deep21/gateway.yaml")
	if err != nil {
		panic(err)
	}

	doc.InternalizeRefs(context.Background(), nil)

	b, _ := doc.MarshalJSON()
	fmt.Println(string(b))

}
