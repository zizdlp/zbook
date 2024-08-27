package render

import (
	// admonitions "github.com/stefanfritsch/goldmark-admonitions"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	admonitions "github.com/zizdlp/zbook/markdown/admonition"
	"github.com/zizdlp/zbook/markdown/katex"
	"go.abhg.dev/goldmark/toc"
)

func GetMarkdownConfig() goldmark.Markdown {
	markdown := goldmark.New(
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
		goldmark.WithRendererOptions(html.WithUnsafe()),
		goldmark.WithExtensions(
			&toc.Extender{
				Title:  "content_table",
				ListID: "content_title",
			},
		),
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithExtensions(
			&admonitions.Extender{},
		),
		goldmark.WithExtensions(extension.NewCJK(extension.WithEastAsianLineBreaks(), extension.WithEscapedSpace())),
		goldmark.WithExtensions(
			extension.NewFootnote(
				extension.WithFootnoteIDPrefix([]byte("footnote-")),
				extension.WithFootnoteBacklinkHTML([]byte("^"))),
		),
		goldmark.WithExtensions(&katex.Extender{}),
	)
	return markdown
}
