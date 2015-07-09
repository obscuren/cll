%{
package ast

var SynTree ASTNode

%}

%union {
    num int
	str  string
	node ASTNode
}

%token VAR END_STMT ASSIGN ASM LBRACE RBRACE
%token <str> ID NUM ASM_BODY
%type <node> program statement_list statement expression literal variable
%type <node> buildin

%%

program
	: statement_list { SynTree = $1 }
	;

statement_list
	: statement_list statement
	{ 
		if list, ok := $1.(*StmtList); ok {
			list.Add($2)
			$$ = $1
		} else {
			node := StatementList().(*StmtList)
			node.Add($1, $2)
			$1.SetParent(node)
			$$ = node
		}
		$2.SetParent($$)
    	}
	| { $$ = Empty() }
	;

statement
	: VAR variable expression
	{
		$$ = Decleration($2, $3)
		$3.SetParent($$)
	}
	| expression { $$ = $1 }
	| END_STMT { $$ = Empty(); }
	;

expression
	: variable ASSIGN expression
	{
		$$ = Assign($1, $3)
		$1.SetParent($$)
		$3.SetParent($$)
	}
	| ASSIGN expression
	{
		$$ = $2    
	}
	| buildin { $$ = $1 }
	| literal { $$ = $1 }
	| variable { $$ = $1 }
	;

literal
	: NUM { $$ = Literal($1, numTy) }
	;

variable
	: ID { $$ = Decl($1) }
	;

buildin
	: ASM LBRACE ASM_BODY RBRACE
	{
		$$ = Asm($3)
	}
	;

%%
