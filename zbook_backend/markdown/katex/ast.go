package katex

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

type Inline struct {
	ast.BaseInline

	Equation []byte
}

func (n *Inline) Inline() {}

func (n *Inline) IsBlank(source []byte) bool {
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		text := c.(*ast.Text).Segment
		if !util.IsBlank(text.Value(source)) {
			return false
		}
	}
	return true
}

func (n *Inline) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

var KindInline = ast.NewNodeKind("Inline")

func (n *Inline) Kind() ast.NodeKind {
	return KindInline
}

type Block struct {
	ast.BaseInline

	Equation []byte
}

func (n *Block) IsBlank(source []byte) bool {
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		text := c.(*ast.Text).Segment
		if !util.IsBlank(text.Value(source)) {
			return false
		}
	}
	return true
}

func (n *Block) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

var KindBlock = ast.NewNodeKind("Block")

func (n *Block) Kind() ast.NodeKind {
	return KindBlock
}
