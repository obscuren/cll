package ast

type ASTNode interface {
	SetParent(ASTNode)
	Parent() ASTNode
}

type StmtList struct {
	parent ASTNode
	list   []ASTNode
}

func StatementList() ASTNode {
	return &StmtList{nil, nil}
}

func (sl *StmtList) SetParent(n ASTNode) { sl.parent = n }
func (sl *StmtList) Parent() ASTNode     { return sl.parent }
func (sl *StmtList) List() []ASTNode     { return sl.list }
func (sl *StmtList) Add(n ...ASTNode) {
	for _, node := range n {
		if _, ok := node.(*EmptyNode); !ok {
			sl.list = append(sl.list, node)
		}
	}
}

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

type DeclNode struct {
	parent ASTNode
	Id     string
}

func Decl(id string) ASTNode {
	return &DeclNode{nil, id}
}

func (ln *DeclNode) Parent() ASTNode     { return ln.parent }
func (ln *DeclNode) SetParent(n ASTNode) { ln.parent = n }

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

type AsmExpr struct {
	parent ASTNode
	Asm    string
}

func Asm(asm string) ASTNode {
	return &AsmExpr{nil, asm}
}

func (en *AsmExpr) SetParent(n ASTNode) { en.parent = n }
func (en *AsmExpr) Parent() ASTNode     { return en.parent }

type EmptyNode struct {
	parent ASTNode
}

func Empty() ASTNode {
	return &EmptyNode{nil}
}

func (en *EmptyNode) SetParent(n ASTNode) { en.parent = n }
func (en *EmptyNode) Parent() ASTNode     { return en.parent }
