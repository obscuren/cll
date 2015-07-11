package ast

import (
	"bytes"
	"fmt"
	"strings"
)

var _n = 0

const identSize = ".  "

func indent(n int) string {
	if n > 0 {
		return string(bytes.Repeat([]byte(identSize), n))
	}
	return ""
}

func sprintf(n int, format string, v ...interface{}) string {
	return fmt.Sprintf("%s%s", indent(n), fmt.Sprintf(format, v...))
}

func stringTree(n ASTNode, no int) (ret string) {
	ret += sprintf(0, "%T{\n", n)
	switch n := n.(type) {
	case *GenDecl:
		ret += sprintf(no+1, "Decl: %s", stringTree(n.Decl, no+1))
		ret += sprintf(no+1, "Value: %s", stringTree(n.Value, no+1))
	case *BlockStmt:
		ret += sprintf(no+1, "List: []Stmt (len = %d) {\n", len(n.List()))
		for i, node := range n.List() {
			ret += sprintf(no+2, "%d: %s", i, stringTree(node, no+2))
		}
		ret += sprintf(no+1, "}\n")
	case *LiteralNode:
		ret += sprintf(no+1, "Type: %v\n", n.Type)
		ret += sprintf(no+1, "Value: \"%v\"\n", n.Value)
	case *AssignExpr:
		ret += sprintf(no+1, "Lhs: %s", stringTree(n.Lhs, no+1))
		ret += sprintf(no+1, "Rhs: %s", stringTree(n.Rhs, no+1))

	case *DeclObj:
		ret += sprintf(no+1, "Id: %s\n", n.Id)
	case *AsmExpr:
		ret += sprintf(no+1, "Asm: \"%s\"\n", strings.Replace(n.Asm, "\n", " ", -1))
	case *IfExpr:
		ret += sprintf(no+1, "Cond: %s", stringTree(n.Cond, no+1))
		ret += sprintf(no+1, "Body: %s", stringTree(n.Body, no+1))
		if n.Else != nil {
			ret += sprintf(no+1, "Else: %s", stringTree(n.Else, no+1))
		}
	case *BinaryExpr:
		ret += sprintf(no+1, "X: %s", stringTree(n.X, no+1))
		ret += sprintf(no+1, "Op: %s\n", n.Op)
		ret += sprintf(no+1, "Y: %s", stringTree(n.Y, no+1))
	case *Ident:
		ret += sprintf(no+1, "Name: %s\n", n.Name)
	case nil, *EmptyNode:
		ret += sprintf(no+1, "%s<nil>\n", identSize)
	}
	ret += sprintf(no, "}\n")
	return
}

// Print pretty prints the Abstract Syntax Tree
func Print(n ASTNode) {
	_n = 0
	items := strings.Split(stringTree(n, 0), "\n")
	for i, item := range items {
		fmt.Printf("%6d  %s\n", i, item)
	}
}
