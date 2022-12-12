package parser

import (
	"os"

	"golang.org/x/mod/modfile"
)

func GetModules(gomodPath string) ([]string, error) {
	var modules []string
	content, err := os.ReadFile(gomodPath)
	if err != nil {
		return nil, err
	}
	mod, err := modfile.Parse("go.mod", content, nil)
	if err != nil {
		return nil, err
	}

	for _, req := range mod.Require {
		if subpackageRe.Match([]byte(req.Mod.Path)) && !packagesToSkip[req.Mod.Path] {
			modules = append(modules, req.Mod.String())
		}
	}

	return modules, nil
}
