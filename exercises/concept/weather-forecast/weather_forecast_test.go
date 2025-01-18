package weather

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

func TestComments(t *testing.T) {
	filename := "weather_forecast.go"

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
			if n.Lparen.IsValid() {
				for _, v := range n.Specs {
					testBlockIdentifierComment(t, v.(*ast.ValueSpec))
				}
			} else {
				testIdentifierComment(t, n)
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
		t.Errorf("Package weather should have a comment")
	}

	packageName := node.Name.Name
	want := "Package " + packageName
	packageComment := node.Doc.Text()

	if ok, errStr := testComment("Package", packageName, packageComment, want); !ok {
		t.Error(errStr)
	}

}
func testIdentifierComment(t *testing.T, node *ast.GenDecl) {
	t.Helper()

	identifierName := node.Specs[0].(*ast.ValueSpec).Names[0].Name
	if node.Doc == nil {
		t.Errorf("Exported identifier %s should have a comment", identifierName)
	}

	identifierComment := node.Doc.Text()
	want := identifierName

	if ok, errStr := testComment("Variable", identifierName, identifierComment, want); !ok {
		t.Error(errStr)
	}
}

func testBlockIdentifierComment(t *testing.T, node *ast.ValueSpec) {
	t.Helper()

	identifierName := node.Names[0].Name
	if node.Doc == nil {
		t.Errorf("Exported identifier %s should have a comment", identifierName)
	}

	identifierComment := node.Doc.Text()
	want := identifierName

	if ok, errStr := testComment("Variable", identifierName, identifierComment, want); !ok {
		t.Error(errStr)
	}

}

func testFunctionComment(t *testing.T, node *ast.FuncDecl) {
	t.Helper()
	funcName := node.Name.Name
	if node.Doc == nil {
		t.Errorf("Exported function %s() should have a comment", funcName)
	}

	funcComment := node.Doc.Text()
	want := funcName

	if ok, errStr := testComment("Function", funcName, funcComment, want); !ok {
		t.Error(errStr)
	}
}

func testComment(entityKind, entityName, comment, wantedPrefix string) (ok bool, errString string) {

	trimmedComment := strings.TrimSpace(comment)
	lowerEntity := strings.ToLower(entityKind)

	// Check if comment has wanted prefix
	if !strings.HasPrefix(trimmedComment, wantedPrefix) {
		errorString := fmt.Sprintf("%s comment for %s '%s' should start with '// %s ...': got '// %s'",
			entityKind, lowerEntity, entityName, wantedPrefix, trimmedComment)
		return false, errorString
	}

	// Check if comment content is empty
	commentContent := strings.TrimPrefix(trimmedComment, wantedPrefix)
	commentContent = strings.TrimSpace(commentContent)
	commentContent = strings.TrimSuffix(commentContent, ".")

	if commentContent == "" {
		lowerEntity := strings.ToLower(entityKind)
		errorString := fmt.Sprintf("%s comment of '%s' should provide a description of the %s, e.g '// %s <%s_description>'",
			entityKind, entityName, lowerEntity, wantedPrefix, lowerEntity)
		return false, errorString
	}

	// Check if comment ends in a period
	if !strings.HasSuffix(trimmedComment, ".") {
		return false, fmt.Sprintf("%s comment for %s '%s' should end with a period (.)",
			entityKind, lowerEntity, entityName)
	}

	return true, ""
}
