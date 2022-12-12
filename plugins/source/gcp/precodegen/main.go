package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	"github.com/cloudquery/plugins/source/gcp/precodegen/internal/parser"
)

////go:embed templates/*.go.tpl
// var templateFS embed.FS

var (
	currentFilename string
	currentDir      string
)

func main() {
	var ok bool
	_, currentFilename, _, ok = runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	currentDir = path.Dir(currentFilename)
	var updateGoMod bool
	if len(os.Args) > 1 && os.Args[1] == "--update-go-mod" {
		updateGoMod = true
	}

	if updateGoMod {
		packagesToGoGet, err := parser.DiscoverSubpackages()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("go getting %d packages\n", len(packagesToGoGet))
		args := []string{"get", "-u"}
		args = append(args, packagesToGoGet...)
		cmd := exec.Command("go", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}

	modules, err := parser.GetModules(path.Join(currentDir, "../go.mod"))
	if err != nil {
		log.Fatal(err)
	}
	for _, module := range modules { 
		if !strings.HasPrefix(module, "cloud.google.com/go/run") {
			continue
		}
		fmt.Println(module)
		tables, err := parser.CreateTablesFromPackage(module)
		if err != nil {
			log.Fatal(err)
		}
		for _, table := range tables {
			fmt.Println(table.NewFuncName)
		}
	}
}