// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
)

func (p *Parser) parseIndex(left ast.Node) (ast.Node, error) {
	node := ast.Index{}.Init(p.peek.Line, left, p.peek.Value)
	p.next()
	e, err := p.parseExpression(constants.PRECEDENCE_DOT)
	if err != nil {
		return nil, err
	}
	node.SetRight(e)
	return node, nil
}
