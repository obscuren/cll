package ir

import (
	"container/list"
	"fmt"
	"strings"

	"github.com/obscuren/gll/ast"
)

type variable struct{}

type scope struct {
	vtable map[string]variable
}

func (s *scope) declVar(n string) error {
	if _, ok := s.vtable[n]; ok {
		return fmt.Errorf("variable '%s' redeclared")
	}
	return nil
}

func newScope() *scope {
	return &scope{vtable: make(map[string]variable)}
}

var (
	scopes  []*scope
	current *scope
)

func ParseTree(tree ast.ASTNode) (*list.List, error) {
	list := list.New()

	switch node := tree.(type) {
	case *ast.BlockStmt:
		previous := current
		current = newScope()
		scopes = append(scopes, current)
		for _, stmt := range node.List() {
			l, err := ParseTree(stmt)
			if err != nil {
				return nil, err
			}
			list.PushBackList(l)
		}
		current = previous
		return list, nil
	case *ast.GenDecl:
		// declarations never push back lists
		_, err := ParseTree(node.Decl)
		if err != nil {
			return nil, err
		}

		value, err := ParseTree(node.Value)
		if err != nil {
			return nil, err
		}
		list.PushBackList(value)
	case *ast.DeclObj:
		return nil, current.declVar(node.Id)
	case *ast.AsmExpr:
		asm := strings.Split(node.Asm, "\n")
		for _, asm := range asm {
			cmt := strings.Split(asm, ";")
			if len(cmt) > 0 && len(cmt[0]) > 0 && cmt[0] != ";" {
				list.PushBack(strings.TrimSpace(cmt[0]))
			}
		}
		return list, nil
	}
	return list, nil
}
