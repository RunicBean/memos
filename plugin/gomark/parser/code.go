package parser

import (
	"github.com/usememos/memos/plugin/gomark/ast"
	"github.com/usememos/memos/plugin/gomark/parser/tokenizer"
)

type CodeParser struct{}

var defaultCodeParser = &CodeParser{}

func NewCodeParser() *CodeParser {
	return defaultCodeParser
}

func (*CodeParser) Match(tokens []*tokenizer.Token) (int, bool) {
	if len(tokens) < 3 {
		return 0, false
	}
	if tokens[0].Type != tokenizer.Backtick {
		return 0, false
	}

	contentTokens, matched := []*tokenizer.Token{}, false
	for _, token := range tokens[1:] {
		if token.Type == tokenizer.Newline {
			return 0, false
		}
		if token.Type == tokenizer.Backtick {
			matched = true
			break
		}
		contentTokens = append(contentTokens, token)
	}
	if !matched || len(contentTokens) == 0 {
		return 0, false
	}
	return len(contentTokens) + 2, true
}

func (p *CodeParser) Parse(tokens []*tokenizer.Token) ast.Node {
	size, ok := p.Match(tokens)
	if size == 0 || !ok {
		return nil
	}

	contentTokens := tokens[1 : size-1]
	return &ast.Code{
		Content: tokenizer.Stringify(contentTokens),
	}
}
