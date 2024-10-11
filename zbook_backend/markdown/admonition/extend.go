package admonitions

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// Extender allows you to use admonitions in markdown with different triggers.
//
// Admonitions are a markdown extension that styles markdown as
// boxes with titles, triggered by characters like '!' or '?'
//
// !!!note This is the title {.some-additional-class}
// This is the admonition content
//
// Example with a different trigger:
// ???note Another type of admonition
//
// You can now pass multiple trigger characters to support different
// types of admonitions.
type Extender struct {
	priority     int    // optional int != 0. the priority value for parser and renderer. Defaults to 100.
	triggerChars []byte // list of trigger characters (e.g., '!', '?')
}

// NewExtender creates an Extender with specified trigger characters and priority.
func NewExtender(priority int, triggerChars []byte) *Extender {
	return &Extender{
		priority:     priority,
		triggerChars: triggerChars,
	}
}

// Extend adds block parsers and node renderers to the markdown instance.
func (e *Extender) Extend(md goldmark.Markdown) {
	priority := 100
	if e.priority != 0 {
		priority = e.priority
	}

	// Add a parser for each trigger character
	for _, char := range e.triggerChars {
		md.Parser().AddOptions(
			parser.WithBlockParsers(
				util.Prioritized(NewAdmonitionParser(char), priority),
			),
		)
	}

	// Add the renderer
	md.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(&Renderer{}, priority),
		),
	)
}
