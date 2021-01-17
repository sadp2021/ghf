package app

import (
	"bytes"
	"errors"
	"go/ast"
	"regexp"
	"strings"
	"text/template"
)

const (
	buildComment = iota
	noneComment
)

const (
	buildName = "build"
)

type buildInfo struct {
	Name  string
	Field []struct {
		Name     string
		Title    string
		TypeName string
	}
}

func newGenerate() *generate {
	return &generate{}
}

type generate struct {
}

func (g *generate) walk(parser *fileParserRes) (*bytes.Buffer, error) {
	cmap := parser.cmap

	for name, value := range cmap {
		commentType := g.extractComment(value[0].Text())
		switch commentType {
		case buildComment:
			genBuf, err := g.genBuild(name)
			if err != nil {
				return nil, err
			}
			return genBuf, nil
		default:
			continue
		}
	}

	return nil, nil
}

func (g *generate) extractComment(text string) int {
	match := regexp.MustCompile(".*@([a-z]*)\\s(.*)")
	res := match.FindStringSubmatch(text)

	if len(res) >= 2 {
		switch res[1] {
		case buildName:
			return buildComment
		default:
			return noneComment
		}
	}

	return noneComment
}

func (g *generate) genBuild(node ast.Node) (*bytes.Buffer, error) {
	genDecl, ok := node.(*ast.GenDecl)
	if !ok {
		return nil, errors.New("type error")
	}

	typeSpec, ok := genDecl.Specs[0].(*ast.TypeSpec)
	if !ok {
		return nil, errors.New("type error")
	}

	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		return nil, errors.New("type error")
	}

	info := buildInfo{
		Name: typeSpec.Name.Name,
		Field: make([]struct {
			Name     string
			Title    string
			TypeName string
		}, 0),
	}

	for _, item := range structType.Fields.List {
		name := item.Names[0].Name
		name = strings.ToLower(name[0:1]) + name[1:]
		title := strings.ToUpper(name[0:1]) + name[1:]

		temp := struct {
			Name     string
			Title    string
			TypeName string
		}{
			Name:     name,
			Title:    title,
			TypeName: item.Type.(*ast.Ident).Name,
		}

		info.Field = append(info.Field, temp)
	}

	t := template.Must(template.New("buildTemplate").Parse(buildTemplate))

	var buf bytes.Buffer
	err := t.Execute(&buf, info)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}
