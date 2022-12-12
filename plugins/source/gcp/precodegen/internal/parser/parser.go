package parser

import (
	"fmt"
	"go/ast"
	goparser "go/parser"
	"go/token"
	"os"
	"os/exec"
	"path"
	"regexp"
	"sort"
	"strings"
)

// github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources

var newGlobalFuncsToSkip = map[string]bool{
}

var newFuncToSkipPerPackage = map[string]map[string]bool{
}

var reNewClient = regexp.MustCompile(`New[a-zA-Z]*Client`)
var reListRequest = regexp.MustCompile(`List[a-zA-Z]*|AggregatedList`)

var supportedPagerParams = [][]string{
	{"options"},
}

var supportedNewClientParams = [][]string{
	{"credential", "options"},
	{"subscriptionID", "credential", "options"},
}

type function struct {
	receiver    string
	name        string
	ast         *ast.FuncDecl
	paramNames  []string
	returnTypes []string
}

type structAST struct {
	name string
	ast  *ast.StructType
}

func parseURLFromFunc(fn *ast.FuncDecl) string {
	for _, stmt := range fn.Body.List {
		if expr, ok := stmt.(*ast.AssignStmt); ok {
			if len(expr.Lhs) == 1 && len(expr.Rhs) == 1 {
				if lhs, ok := expr.Lhs[0].(*ast.Ident); ok {
					if lhs.Name == "urlPath" {
						if rhs, ok := expr.Rhs[0].(*ast.BasicLit); ok {
							return strings.Replace(rhs.Value, "\"", "", -1)
						}
					}
				}
			}
		}
	}
	return ""
}

func compareStrArrays(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isArrayExist(arr [][]string, item []string) bool {
	for _, a := range arr {
		if compareStrArrays(a, item) {
			return true
		}
	}
	return false
}

func getParamNames(fn *ast.FieldList) []string {
	var params []string
	for _, p := range fn.List {
		for _, name := range p.Names {
			params = append(params, name.Name)
		}
	}
	return params
}

func getReturnTypes(fn *ast.FieldList) []string {
	var params []string
	for _, p := range fn.List {
		if ident, ok := p.Type.(*ast.Ident); ok {
			params = append(params, ident.Name)
		} else if star, ok := p.Type.(*ast.StarExpr); ok {
			if index, ok := star.X.(*ast.IndexExpr); ok {
				if ident, ok := index.Index.(*ast.Ident); ok {
					params = append(params, ident.Name)
				}
			}
		}
	}
	return params
}

// returns reciever and method name that matches re
func findFunctions(pkgs map[string]*ast.Package, re *regexp.Regexp) map[string]*function {
	var funcs map[string]*function = make(map[string]*function)
	fmt.Println("starting find functinos")
	for _, pack := range pkgs {
		fmt.Println("package: " + pack.Name)
		for _, f := range pack.Files {
			fmt.Println(f.Name.Name)
			for _, d := range f.Decls {
				if fn, isFn := d.(*ast.FuncDecl); isFn {
					fmt.Println(fn.Name.Name)
					if re.MatchString(fn.Name.Name) {
						fun := function{
							name: fn.Name.Name,
							ast:  fn,
						}
						// if function is a method extract receiver name
						if fn.Recv != nil && len(fn.Recv.List) == 1 {
							receiver := fn.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name
							fun.receiver = receiver
						}
						fun.paramNames = getParamNames(fn.Type.Params)
						if fn.Type != nil && fn.Type.Results != nil {
							fun.returnTypes = getReturnTypes(fn.Type.Results)
						}
						if fun.receiver != "" {
							funcs[fun.receiver+"."+fun.name] = &fun
						} else {
							funcs[fun.name] = &fun
						}
					}
				}
			}
		}
	}
	return funcs
}

// Get a package in format of github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/xxx/yy@v0.1.1
func CreateTablesFromPackage(pkg string) ([]*Table, error) {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		output, err := exec.Command("go", "env", "GOPATH").Output()
		if err != nil {
			return nil, err
		}
		goPath = strings.TrimSpace(string(output))
	}
	pkgWithoutVersion := strings.Split(pkg, "@")[0]
	if packagesToSkip[pkgWithoutVersion] {
		return nil, nil
	}
	cacheDir := goPath + "/pkg/mod"
	// this maps client name to tables
	tables := make(map[string]*Table)
	set := token.NewFileSet()
	pkgPath := path.Join(cacheDir, pkg)
	// thats because azure had to be special with uppercase
	pkgPath = strings.Replace(pkgPath, "A", "!A", 1)
	pkgs, err := goparser.ParseDir(set, pkgPath, nil, 0)
	if err != nil {
		return nil, err
	}
	newXClientFuncstions := findFunctions(pkgs, reNewClient)
	fmt.Println("what")
	fmt.Println(len(newXClientFuncstions))
	for _, fn := range newXClientFuncstions {
		fmt.Println(fn.name)
		if newGlobalFuncsToSkip[fn.name] {
		 // || !isArrayExist(supportedNewClientParams, fn.paramNames) 
			continue
		}
		if _, ok := newFuncToSkipPerPackage[pkgWithoutVersion]; ok {
			if newFuncToSkipPerPackage[pkgWithoutVersion][fn.name] {
				continue
			}
		}
		tables[strings.TrimPrefix(fn.name, "New")] = &Table{
			NewFuncName: fn.name,
		}
	}
	var result []*Table
	listMethods := findFunctions(pkgs, reListRequest)
	for _, method := range listMethods {
		if _, ok := tables[method.receiver]; ok {
			tables[method.receiver].ListMethod = method.name
			result = append(result, tables[method.receiver])
		}
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].NewFuncName < result[j].NewFuncName
	})
	return result, nil
}
