package internal

import (
	"go/ast"
)

const (
	contextIn  = "In"
	contextOut = "Out"
	get        = "Get"
	set        = "Set"
	goaCtx     = "github.com/wesovilabs/goa/api/context"
)

var (
	SelectorOutGet = selectorContextOperation(contextOut, get)
	SelectorOutSet = selectorContextOperation(contextOut, set)
	SelectorInGet  = selectorContextOperation(contextIn, get)
	SelectorInSet  = selectorContextOperation(contextIn, set)
)

func selectorContextOperation(name string, op string) *ast.SelectorExpr {
	return &ast.SelectorExpr{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   NewIdentObjVar(varGoaContext),
				Sel: NewIdent(name),
			},
		},
		Sel: NewIdent(op),
	}
}
