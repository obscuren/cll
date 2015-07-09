package ast

func ParseFile(str string, src string) ASTNode {
	tokens := lexer(str, src)
	yyParse(tokens)

	return SynTree
}
