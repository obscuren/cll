package ast

// ParseFile parses the given file and returns the Abstract Syntax Tree
func ParseFile(str string, src string) ASTNode {
	tokens := lexer(str, src)
	yyParse(tokens)

	return SynTree
}
