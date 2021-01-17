package app

import (
	"bufio"
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

//formatIO 格式化文件
func formatIO(fset *token.FileSet, f *ast.File, buf *bytes.Buffer) (*bytes.Buffer, error) {
	var res bytes.Buffer
	if err := format.Node(&res, fset, f); err != nil {
		return nil, err
	}

	res.Write(buf.Bytes())

	temp, _ := format.Source(res.Bytes())
	res.Reset()
	res.Write(temp)

	return &res, nil
}

func writeToFile(path string, buf *bytes.Buffer) error {
	genPath, err := getGenFilePath(path)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(genPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	writer := bufio.NewWriter(f)
	_, err = writer.WriteString(buf.String())
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

func getGenFilePath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	dirPath := filepath.Dir(absPath)
	fileName := filepath.Base(absPath)
	dotIndex := strings.Index(fileName, ".")
	genFileName := fileName[:dotIndex] + "_gen" + fileName[dotIndex:]

	genPath := filepath.Join(dirPath, genFileName)

	return genPath, nil
}
