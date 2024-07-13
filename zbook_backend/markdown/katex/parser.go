package katex

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type Parser struct {
}

func (s *Parser) Trigger() []byte {
	return []byte{'$'}
}

func (s *Parser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	buf := block.Source()
	ln, pos := block.Position()

	lstart := pos.Start
	lend := pos.Stop
	line := buf[lstart:lend]

	var start, end, advance int

	trigger := line[0]

	display := len(line) > 1 && line[1] == trigger

	if display { // Display
		start = lstart + 2

		offset := 2

	L:
		for x := 0; x < 20; x++ {
			for j := offset; j < len(line); j++ {
				if len(line) > j+1 && line[j] == trigger && line[j+1] == trigger {
					end = lstart + j
					advance = 2
					break L
				}
			}
			if lend == len(buf) {
				break
			}
			if end == 0 {
				rest := buf[lend:]
				j := 1
				for j < len(rest) && rest[j] != '\n' {
					j++
				}
				lstart = lend
				lend += j
				line = buf[lstart:lend]
				ln++
				offset = 0
			}
		}

	} else { // Inline
		start = lstart + 1

		for i := 1; i < len(line); i++ {
			c := line[i]
			if c == '\\' {
				i++
				continue
			}
			if c == trigger {
				end = lstart + i
				advance = 1
				break
			}
		}
		if end >= len(buf) || buf[end] != trigger {
			return nil
		}
	}

	if start >= end {
		return nil
	}

	newpos := end + advance
	if newpos < lend {
		block.SetPosition(ln, text.NewSegment(newpos, lend))
	} else {
		block.Advance(newpos)
	}

	if display {
		return &Block{
			Equation: buf[start:end],
		}
	} else {
		return &Inline{
			Equation: buf[start:end],
		}
	}
}
