package weather

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

func TestComments(t *testing.T) {
	filename := "weather.go"

	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, filename, nil, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	wantedComments := 4
	got := len(f.Comments)
	if got != wantedComments {
		t.Errorf("Incorrect number of comments: got %d, want %d", got, wantedComments)
	}

	testPackageComment(t, f)

	ast.Inspect(f, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.GenDecl:
			for _, v := range n.Specs {
				testIdentifierComment(t, v.(*ast.ValueSpec))
			}
		case *ast.FuncDecl:
			testFunctionComment(t, n)
		}
		return true
	})
}

func testPackageComment(t *testing.T, node *ast.File) {
	t.Helper()
	if node.Doc == nil {
		t.Errorf("Package weather doesn't have a comment")
	}

	want := "Package"
	packageComment := node.Doc.Text()
	if !strings.HasPrefix(packageComment, want) {
		t.Errorf("Package comment should start with '// %s ...': got %s", want, packageComment)
	}
}

func testIdentifierComment(t *testing.T, node *ast.ValueSpec) {
	t.Helper()
	if node.Doc == nil {
		t.Errorf("Exported identifier %s should have a comment", node.Names[0])
	}
}

func testFunctionComment(t *testing.T, node *ast.FuncDecl) {
	t.Helper()
	funcName := node.Name.Name
	if node.Doc == nil {
		t.Errorf("Exported function %s() has no comment", funcName)
	}
	funcComment := node.Doc.Text()
	if !strings.HasPrefix(node.Doc.Text(), funcName) {
		t.Errorf("Function comment for function %s should start with `// %s ...`: got %s", funcName, funcName, funcComment)
	}
}
