package markdown

import (
	"fmt"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
	"io/ioutil"
)

func ExtractFile(path string) ([]CommandBlock, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return Extract(content)
}

func Extract(content []byte) ([]CommandBlock, error) {
	var ret []CommandBlock

	reader := text.NewReader(content)
	root := goldmark.DefaultParser().Parse(reader)
	cb := &CommandBlock{}
	err := ast.Walk(root, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		s := ast.WalkStatus(ast.WalkContinue)
		if !entering {
			return s, nil
		}
		switch n.Kind() {
		case ast.KindHeading:
			cb.Title = nodeRawString(n, content)
		case ast.KindFencedCodeBlock:
			cb.Command = nodeRawString(n, content)
			ret = append(ret, *cb)
			cb = &CommandBlock{}
		}
		return s, nil
	})
	return ret, err
}

func nodeRawString(node ast.Node, content []byte) string {
	ret := ""
	for i := 0; i < node.Lines().Len(); i++ {
		line := node.Lines().At(i)
		ret += fmt.Sprintf("%s", line.Value(content))
	}
	return ret
}
