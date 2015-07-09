package ir

import (
	"container/list"
	"strings"

	"github.com/obscuren/cll/ast"
)

func ParseTree(tree ast.ASTNode) (*list.List, error) {
	list := list.New()

	switch node := tree.(type) {
	case *ast.StmtList:
		for _, stmt := range node.List() {
			l, err := ParseTree(stmt)
			if err != nil {
				return nil, err
			}
			list.PushBackList(l)
		}
		return list, nil

	case *ast.AsmExpr:
		asm := strings.Split(node.Asm, "\n")
		for _, asm := range asm {
			list.PushBack(strings.TrimSpace(asm))
		}
		return list, nil
	}
	return list, nil
}
