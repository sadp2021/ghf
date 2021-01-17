package app

//Pipe 流程执行
type Pipe struct {
	parser *fileParserRes
}

//Exec 执行
func (p *Pipe) Exec(path string) (err error) {
	p.parser, err = getFileAST(path)
	if err != nil {
		return
	}

	// 去掉 generate 命令与 build tag
	p.parser.f.Comments = p.parser.f.Comments[1:]

	gen := newGenerate()
	buf, err := gen.walk(p.parser)
	if err != nil {
		return
	}

	buf, err = formatIO(p.parser.fset, p.parser.f, buf)
	if err != nil {
		return
	}

	err = writeToFile(path, buf)
	if err != nil {
		return
	}

	return
}
