package admonitions

import (
	"github.com/yuin/goldmark/ast"
)

// A Admonition struct represents a fenced code block of Markdown text.
type Admonition struct {
	ast.BaseBlock
	AdmonitionClass []byte
	Title           []byte
}

// Dump implements Node.Dump .
func (n *Admonition) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

// KindAdmonition is a NodeKind of the Admonition node.
var KindAdmonition = ast.NewNodeKind("Admonition")

// Kind implements Node.Kind.
func (n *Admonition) Kind() ast.NodeKind {
	return KindAdmonition
}

// NewAdmonition return a new Admonition node.
func NewAdmonition() *Admonition {
	return &Admonition{
		BaseBlock: ast.BaseBlock{},
	}
}
