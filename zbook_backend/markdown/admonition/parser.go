package admonitions

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type admonitionParser struct {
	triggerChar byte // The character that triggers the admonition, e.g., '!' or '?'
}

// NewAdmonitionParser creates a new admonition parser with a specified trigger character.
func NewAdmonitionParser(char byte) parser.BlockParser {
	return &admonitionParser{
		triggerChar: char,
	}
}

type admonitionData struct {
	ID                string   // The ID of the admonition. This enables nested admonitions with indentation
	char              byte     // Trigger character, e.g., '!' or '?'
	indent            int      // The indentation of the opening (and closing) tags (!!!{})
	length            int      // The length of the admonition, e.g. is it !!! or ???
	node              ast.Node // The node of the admonition
	contentIndent     int      // The indentation of the content relative to the previous admonition block. The first line of the content is taken as its indentation. If you want an admonition with just a code block you need to use backticks
	contentHasStarted bool     // Only used as an indicator if contentIndent has been set already
}

var admonitionInfoKey = parser.NewContextKey()

func (b *admonitionParser) Trigger() []byte {
	return []byte{b.triggerChar}
}

func (b *admonitionParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	line, _ := reader.PeekLine()
	pos := pc.BlockOffset()
	if pos < 0 || line[pos] != b.triggerChar {
		return nil, parser.NoChildren
	}
	findent := pos

	// currently useless
	admonitionChar := line[pos]
	i := pos
	for ; i < len(line) && line[i] == admonitionChar; i++ {
	}
	oAdmonitionLength := i - pos
	if oAdmonitionLength < 3 {
		return nil, parser.NoChildren
	}

	// ========================================================================== //
	// 	Without attributes we return

	if i >= len(line)-1 {
		// If there are no attributes we can't create a div because we won't know
		// if a "!!!" ends the last admonition or opens a new one
		return nil, parser.NoChildren
	}

	rest := line[i:]
	left := i + util.TrimLeftSpaceLength(rest)
	right := len(line) - 1 - util.TrimRightSpaceLength(rest)

	if left >= right {
		// As above:
		// If there are no attributes we can't create a div because we won't know
		// if a "!!!" ends the last admonition or opens a new one
		return nil, parser.NoChildren
	}

	// ========================================================================== //
	// 	With attributes we construct the node
	node := parseOpeningLine(reader, left)
	admonitionID := genRandomString(24)
	node.SetAttributeString("data-admonition", []byte(admonitionID))

	fdata := &admonitionData{
		ID:                admonitionID,
		char:              admonitionChar,
		indent:            findent,
		length:            oAdmonitionLength,
		node:              node,
		contentIndent:     0,
		contentHasStarted: false,
	}
	var fdataMap []*admonitionData

	if oldData := pc.Get(admonitionInfoKey); oldData != nil {
		fdataMap = oldData.([]*admonitionData)
		fdataMap = append(fdataMap, fdata)
	} else {
		fdataMap = []*admonitionData{fdata}
	}
	pc.Set(admonitionInfoKey, fdataMap)

	// ========================================================================== //
	// 	 check if it's an empty block

	line, _ = reader.PeekLine()
	w, pos := util.IndentWidth(line, reader.LineOffset())
	if close, _ := hasClosingTag(line, w, pos, fdata); w < fdata.indent || close {
		return node, parser.NoChildren
	}
	indentClose :=
		!util.IsBlank(line) &&
			(w < fdata.indent || (w == fdata.indent && w <= fdata.contentIndent)) // mydebug:warning,这里w < fdata.contentIndent 修改添加=
	if indentClose {
		return node, parser.NoChildren
	}
	return node, parser.HasChildren
}

// Parse the opening line for
// * admonition class
// * admonition title
// * attributes
func parseOpeningLine(reader text.Reader, left int) *Admonition {
	node := NewAdmonition()
	reader.Advance(left)

	remainingLine, _ := reader.PeekLine()
	remainingLength := len(remainingLine) - 1

	// ========================================================================== //
	// 	find class
	endClass := 0
	for ; endClass < remainingLength && remainingLine[endClass] != ' ' && remainingLine[endClass] != '{'; endClass++ {
	}
	if endClass > 0 {
		node.AdmonitionClass = remainingLine[0:endClass]
	}

	// ========================================================================== //
	// 	find title
	startTitle := endClass + util.TrimLeftSpaceLength(remainingLine[endClass:])
	endTitle := startTitle
	for ; endTitle < remainingLength && remainingLine[endTitle] != '{'; endTitle++ {
	}
	if endTitle > startTitle {
		endTitle = endTitle - util.TrimRightSpaceLength(remainingLine[startTitle:endTitle])
		if endTitle > startTitle {
			node.Title = remainingLine[startTitle:endTitle]
		}
	}

	if endTitle < remainingLength {
		reader.Advance(endTitle)
	} else {
		reader.Advance(remainingLength)
	}

	// ========================================================================== //
	// 	find attributes
	hasClass := false
	admClass := bytes.Join([][]byte{[]byte("admonition adm-"), node.AdmonitionClass}, []byte(""))

	attrs, ok := parser.ParseAttributes(reader)

	if ok {
		for _, attr := range attrs {
			oldVal := attr.Value.([]byte)
			var val []byte

			if bytes.Equal(attr.Name, []byte("class")) {
				hasClass = true
				val = bytes.Join([][]byte{admClass, oldVal}, []byte(" "))
			} else {
				val = oldVal
			}

			node.SetAttribute(attr.Name, val)
		}
	}

	if !hasClass {
		node.SetAttribute([]byte("class"), admClass)
	}

	return node
}

func (b *admonitionParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	// ========================================================================== //
	// Get admonitionID from node

	rawAdmonitionID, ok := node.AttributeString("data-admonition")
	if !ok {
		fmt.Println("Admonition ID is missing")
	}
	admonitionID := string(rawAdmonitionID.([]byte))

	// ========================================================================== //
	// 	Get admonition for current admonition
	rawdata := pc.Get(admonitionInfoKey)
	fdataMap := rawdata.([]*admonitionData)

	// This should not happen
	if len(fdataMap) == 0 {
		fmt.Printf("we're in an admonition block but have no state data. This should not happen")
		return parser.Close
	}

	var fdata *admonitionData
	var flevel int
	for flevel = 0; flevel < len(fdataMap); flevel++ {
		fdata = fdataMap[flevel]
		if fdata.ID == admonitionID {
			break
		}
	}

	// ========================================================================== //
	// 	Set indentation level if it hasn't been set yet

	line, segment := reader.PeekLine()
	w, pos := util.IndentWidth(line, reader.LineOffset())

	if !fdata.contentHasStarted && !util.IsBlank(line[pos:]) {
		fdata.contentHasStarted = true
		fdata.contentIndent = w

		fdataMap[flevel] = fdata
		pc.Set(admonitionInfoKey, fdataMap)
	}

	// ========================================================================== //
	// Are we closing the node?
	// * Either the indentation is below the indentation of the opening tags
	// * or it is at the level of the opening tags but the content was indented
	// * or there is a closing tag and we're in the deepest admonition block
	close, newline := hasClosingTag(line, w, pos, fdata)
	if close && flevel == len(fdataMap)-1 {
		reader.Advance(segment.Stop - segment.Start - newline + segment.Padding)

		node.SetAttributeString("data-admonition", []byte(fmt.Sprint(flevel)))

		fdataMap = fdataMap[:flevel]
		pc.Set(admonitionInfoKey, fdataMap)

		return parser.Close
	}
	//// w: 是新起一行文本的最左边位置
	//// fdata.contentIndent：是admonition 应该的最左位置：！

	indentClose :=
		!util.IsBlank(line) &&
			(w < fdata.indent || (w == fdata.indent && w <= fdata.contentIndent)) // mydebug:warning,这里w < fdata.contentIndent 修改添加=
	if indentClose {
		node.SetAttributeString("data-admonition", []byte(fmt.Sprint(flevel)))

		fdataMap = fdataMap[:flevel]
		pc.Set(admonitionInfoKey, fdataMap)

		return parser.Close
	}

	if fdata.contentIndent > 0 {
		dontJumpLineEnd := segment.Stop - segment.Start - 1
		if fdata.contentIndent < dontJumpLineEnd {
			dontJumpLineEnd = fdata.contentIndent
		}

		reader.Advance(dontJumpLineEnd)
	}

	return parser.Continue | parser.HasChildren
}

func (b *admonitionParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
}

func (b *admonitionParser) CanInterruptParagraph() bool {
	return true
}

func (b *admonitionParser) CanAcceptIndentedLine() bool {
	return false
}

func hasClosingTag(line []byte, w int, pos int, fdata *admonitionData) (bool, int) {
	// else, check for the correct number of closing chars and provide the info
	// necessary to advance the reader
	if w == fdata.indent {
		i := pos
		for ; i < len(line) && line[i] == fdata.char; i++ {
		}
		length := i - pos

		if length >= fdata.length && util.IsBlank(line[i:]) {
			newline := 1
			if line[len(line)-1] != '\n' {
				newline = 0
			}

			return true, newline
		}
	}

	return false, 0
}
