// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
)

func (p *Parser) parseIf() (ast.Node, error) {
	node := ast.If{}.Init(p.peek.Line, p.peek.Type)

	p.next() // 'if' or 'elif'

	e, err := p.parseExpression(constants.PRECEDENCE_LOWEST)
	if err != nil {
		return nil, err
	}
	node.SetCondition(e)

	e, err = p.parseBlock()
	if err != nil {
		return nil, err
	}
	if e, ok := e.(ast.Block); ok {
		node.SetConsequence(e)
	} else {
		return nil, p.Throw(e.GetLine(), constants.SYNTAX_EXPECTED, "block statement")
	}

	if p.peek.Type == constants.TOKEN_ELIF {
		e, err := p.parseIf()
		if err != nil {
			return nil, err
		}
		node.SetAlternative(ast.Block{}.Set(p.peek.Line, []ast.Node{e}))
	} else if p.peek.Type == constants.TOKEN_ELSE {
		p.next() // 'else'
		e, err = p.parseBlock()
		if err != nil {
			return nil, err
		}
		if e, ok := e.(ast.Block); ok {
			node.SetAlternative(e)
		} else {
			return nil, p.Throw(e.GetLine(), constants.SYNTAX_EXPECTED, "block statement")
		}
	}

	return node, nil
}

func (p *Parser) parseTernary(left ast.Node) (ast.Node, error) {
	node := ast.Ternary{}.Init(p.peek.Line, p.peek.Type)
	p.next()
	node.SetCondition(left)
	e, err := p.parseExpression(constants.PRECEDENCE_LOWEST)
	if err != nil {
		return nil, err
	}
	node.SetConsequence(e)

	if p.peek.Type == constants.TOKEN_COLONCOLON {
		p.next()
		e, err = p.parseExpression(constants.PRECEDENCE_LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetAlternative(e)
	}
	return node, nil
}
