package admonitions

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// This extender allows you to use admonitions in markdown
//
// Admonitions are a markdown extension that allows you to style markdown as
// nice boxes with a title. This is done by way of wrapping other elements in
// divs with classes starting with "adm-"
//
// !!!note This is the title {.some-additional-class}
// This is the admonition
//
// ## with a header
//
// !!!danger Nesting is possible!
// ```R
// X <- as.data.table(iris)
// X[Species != "virginica", mean(Sepal.Length), Species]
// ```
// !!!
// !!!
type Extender struct {
	priority int // optional int != 0. the priority value for parser and renderer. Defaults to 100.
}

// This implements the Extend method for goldmark-admonitions.Extender
func (e *Extender) Extend(md goldmark.Markdown) {
	priority := 100

	if e.priority != 0 {
		priority = e.priority
	}
	md.Parser().AddOptions(
		parser.WithBlockParsers(
			util.Prioritized(&admonitionParser{}, priority),
		),
	)
	md.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(&Renderer{}, priority),
		),
	)
}
