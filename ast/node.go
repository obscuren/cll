package ast

// ASTNode is the basic buliding block of the Abstract Syntax Tree. All nodes
// must implement its interface
type ASTNode interface {
	SetParent(ASTNode)
	Parent() ASTNode
}

// Empty node represents void
type EmptyNode struct {
	parent ASTNode
}

func Empty() ASTNode {
	return &EmptyNode{nil}
}

func (en *EmptyNode) SetParent(n ASTNode) { en.parent = n }
func (en *EmptyNode) Parent() ASTNode     { return en.parent }

// BlockStmt is a list of statements and expressions
type BlockStmt struct {
	parent ASTNode
	list   []ASTNode
}

func BlockStatement() ASTNode {
	return &BlockStmt{nil, nil}
}

func (sl *BlockStmt) SetParent(n ASTNode) { sl.parent = n }
func (sl *BlockStmt) Parent() ASTNode     { return sl.parent }
func (sl *BlockStmt) List() []ASTNode     { return sl.list }
func (sl *BlockStmt) Add(n ...ASTNode) {
	for _, node := range n {
		if _, ok := node.(*EmptyNode); !ok {
			sl.list = append(sl.list, node)
		}
	}
}

// GenDecl is a generic variable decleration
type GenDecl struct {
	parent ASTNode
	Decl   ASTNode
	Value  ASTNode
}

func Decleration(decl ASTNode, n1 ASTNode) ASTNode {
	return &GenDecl{nil, decl, n1}
}

func (vd *GenDecl) SetParent(n ASTNode) { vd.parent = n }
func (vd *GenDecl) Parent() ASTNode     { return vd.parent }

// AssignExpr represents an assignment operation in the tree
type AssignExpr struct {
	parent ASTNode
	Lhs    ASTNode
	Rhs    ASTNode
}

func Assign(l, r ASTNode) ASTNode {
	return &AssignExpr{nil, l, r}
}

func (ae *AssignExpr) SetParent(n ASTNode) { ae.parent = n }
func (ae *AssignExpr) Parent() ASTNode     { return ae.parent }

type litType byte

const (
	numTy litType = iota
	strTy
)

func (l litType) String() string {
	switch l {
	case numTy:
		return "NUM"
	case strTy:
		return "STR"
	default:
		return "UNKNOWN"
	}
}

// DeclObj is a decleration of an object (variable)
type DeclObj struct {
	parent ASTNode
	Id     string
}

func Decl(id string) ASTNode {
	return &DeclObj{nil, id}
}

func (ln *DeclObj) Parent() ASTNode     { return ln.parent }
func (ln *DeclObj) SetParent(n ASTNode) { ln.parent = n }

// LeteralNode is a language literal (number, string, etc)
type LiteralNode struct {
	parent ASTNode
	Value  string
	Type   litType
}

func Literal(v string, t litType) *LiteralNode {
	return &LiteralNode{nil, v, t}
}

func (ln *LiteralNode) Parent() ASTNode     { return ln.parent }
func (ln *LiteralNode) SetParent(n ASTNode) { ln.parent = n }

// AsmExpr represents a block of assembly
type AsmExpr struct {
	parent ASTNode
	Asm    string
}

func Asm(asm string) ASTNode {
	return &AsmExpr{nil, asm}
}

func (en *AsmExpr) SetParent(n ASTNode) { en.parent = n }
func (en *AsmExpr) Parent() ASTNode     { return en.parent }

// IfExp is an if expression with an optional else expression
type IfExpr struct {
	parent ASTNode
	Cond   ASTNode
	Body   ASTNode
	Else   ASTNode
}

func If(cond, body, els ASTNode) ASTNode {
	return &IfExpr{nil, cond, body, els}
}

func (en *IfExpr) SetParent(n ASTNode) { en.parent = n }
func (en *IfExpr) Parent() ASTNode     { return en.parent }

// BinaryExpr is a binary condition
type BinaryExpr struct {
	parent ASTNode
	X, Y   ASTNode
	Op     string
}

func Binary(x ASTNode, op string, y ASTNode) ASTNode {
	return &BinaryExpr{nil, x, y, op}
}

func (en *BinaryExpr) SetParent(n ASTNode) { en.parent = n }
func (en *BinaryExpr) Parent() ASTNode     { return en.parent }

// Ident reprents a variable (identifier)
type Ident struct {
	parent ASTNode
	Name   string
}

func Id(name string) ASTNode {
	return &Ident{nil, name}
}

func (en *Ident) SetParent(n ASTNode) { en.parent = n }
func (en *Ident) Parent() ASTNode     { return en.parent }
