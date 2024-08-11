package admonitions

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

// A Config struct has configurations for the HTML based renderers.
type Config struct {
	Writer    html.Writer
	HardWraps bool
	XHTML     bool
	Unsafe    bool
}

// HeadingAttributeFilter defines attribute names which heading elements can have
var AdmonitionAttributeFilter = html.GlobalAttributeFilter

// A Renderer struct is an implementation of renderer.NodeRenderer that renders
// nodes as (X)HTML.
type Renderer struct {
	Config
}

// RegisterFuncs implements NodeRenderer.RegisterFuncs .
func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindAdmonition, r.renderAdmonition)
}

func (r *Renderer) renderAdmonition(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*Admonition)
	if entering {
		if n.Attributes() != nil {
			_, _ = w.WriteString("<div")
			html.RenderAttributes(w, n, AdmonitionAttributeFilter)
			_, _ = w.WriteString(">\n")
		} else {
			_, _ = w.WriteString("<div>\n")
		}
		_, _ = w.WriteString(`  <div class="adm-title">` + string(util.EscapeHTML(n.Title)) + "</div>\n  <div class=\"adm-body\">\n")
	} else {
		_, _ = w.WriteString("  </div>\n</div>\n")
	}
	return ast.WalkContinue, nil
}
