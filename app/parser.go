package app

import (
	"go/ast"
	"go/parser"
	"go/token"
)

//fileParserRes 格式化结构
type fileParserRes struct {
	fset *token.FileSet
	f    *ast.File
	cmap ast.CommentMap
}

//getFileAST 得到抽象语法树
func getFileAST(path string) (*fileParserRes, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	cmap := ast.NewCommentMap(fset, f, f.Comments)

	res := &fileParserRes{
		fset: fset,
		f:    f,
		cmap: cmap,
	}

	return res, nil
}
