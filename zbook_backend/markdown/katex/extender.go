package katex

import (
	"github.com/bluele/gcache"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type Extender struct {
}

func (e *Extender) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(&Parser{}, 0),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(&HTMLRenderer{
			cacheInline:  gcache.New(5000).ARC().Build(),
			cacheDisplay: gcache.New(5000).ARC().Build(),
		}, 0),
	))
}
