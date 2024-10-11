package admonitions

import (
	"github.com/yuin/goldmark/ast"
)

// Admonition struct represents a fenced code block of Markdown text.
type Admonition struct {
	ast.BaseBlock
	AdmonitionClass []byte
	Title           []byte
	TriggerChar     byte // New field for trigger type, e.g., '!', '?', etc.
}

// Dump implements Node.Dump.
func (n *Admonition) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

// KindAdmonition is a NodeKind of the Admonition node.
var KindAdmonition = ast.NewNodeKind("Admonition")

// Kind implements Node.Kind.
func (n *Admonition) Kind() ast.NodeKind {
	return KindAdmonition
}

// NewAdmonition returns a new Admonition node.
func NewAdmonition(triggerChar byte) *Admonition {
	return &Admonition{
		BaseBlock:   ast.BaseBlock{},
		TriggerChar: triggerChar, // Set the trigger type on creation
	}
}
