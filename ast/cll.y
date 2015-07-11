%{
package ast

var SynTree ASTNode

%}

%union {
    num int
	str  string
	node ASTNode
}

%token END_STMT ASSIGN ASM LBRACE RBRACE COLON IF ELSE
%token <str> ID NUM ASM_BODY OP
%type <node> program statement_list statement expression literal identifier
%type <node> buildin block_statement condition optional_else

%%

program
	: statement_list { SynTree = $1 }
	;

statement_list
	: statement_list statement
	{ 
		if list, ok := $1.(*BlockStmt); ok {
			list.Add($2)
			$$ = $1
		} else {
			node := BlockStatement().(*BlockStmt)
			node.Add($1, $2)
			$1.SetParent(node)
			$$ = node
		}
		$2.SetParent($$)
    	}
	| { $$ = Empty() }
	;

statement
	: ID COLON ASSIGN expression
	{
		$$ = Decleration(Decl($1), $4)
		$4.SetParent($$)
	}
	| expression { $$ = $1 }
	| END_STMT { $$ = Empty(); }
	;

expression
	: identifier ASSIGN expression
	{
		$$ = Assign($1, $3)
		$1.SetParent($$)
		$3.SetParent($$)
	}
	| ASSIGN expression
	{
		$$ = $2    
	}
	| IF condition block_statement optional_else
	{
		$$ = If($2, $3, $4)
		$2.SetParent($$)
		$3.SetParent($3)
		if $4 != nil {
			$4.SetParent($$)
		}
	}
	| buildin { $$ = $1 }
	| literal { $$ = $1 }
	| identifier { $$ = $1 }
	;

optional_else
	: ELSE block_statement
	{
		$$ = $2
	}
	| ELSE expression
	{ 
		$$ = $2
	}
	| { $$ = nil }
	;

condition
	: identifier OP identifier
	{
		$$ = Binary($1, $2, $3)
		$1.SetParent($$)
		$3.SetParent($$)
	}
	;

block_statement
	: LBRACE statement_list RBRACE
	{
		$$ = $2
	}
	;

literal
	: NUM { $$ = Literal($1, numTy) }
	;

identifier
	: ID { $$ = Id($1) }
	;

buildin
	: ASM LBRACE ASM_BODY RBRACE
	{
		$$ = Asm($3)
	}
	;

%%
