//line ast/cll.y:2
package ast

import __yyfmt__ "fmt"

//line ast/cll.y:2
var SynTree ASTNode

//line ast/cll.y:8
type yySymType struct {
	yys  int
	num  int
	str  string
	node ASTNode
}

const VAR = 57346
const END_STMT = 57347
const ASSIGN = 57348
const ASM = 57349
const LBRACE = 57350
const RBRACE = 57351
const ID = 57352
const NUM = 57353
const ASM_BODY = 57354

var yyToknames = []string{
	"VAR",
	"END_STMT",
	"ASSIGN",
	"ASM",
	"LBRACE",
	"RBRACE",
	"ID",
	"NUM",
	"ASM_BODY",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line ast/cll.y:83

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 15
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 26

var yyAct = []int{

	5, 4, 6, 8, 12, 20, 11, 11, 13, 16,
	21, 17, 15, 8, 12, 18, 19, 11, 13, 7,
	9, 10, 3, 2, 14, 1,
}
var yyPact = []int{

	-1000, -1000, -3, -1000, -4, -1000, -1000, 6, 7, -1000,
	-1000, -1000, 3, -1000, 7, 7, -1000, -7, -1000, -1000,
	1, -1000,
}
var yyPgo = []int{

	0, 25, 23, 22, 0, 21, 19, 20,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 4, 4, 4,
	4, 4, 5, 6, 7,
}
var yyR2 = []int{

	0, 1, 2, 0, 3, 1, 1, 3, 2, 1,
	1, 1, 1, 1, 4,
}
var yyChk = []int{

	-1000, -1, -2, -3, 4, -4, 5, -6, 6, -7,
	-5, 10, 7, 11, -6, 6, -4, 8, -4, -4,
	12, 9,
}
var yyDef = []int{

	3, -2, 1, 2, 0, 5, 6, 11, 0, 9,
	10, 13, 0, 12, 0, 0, 8, 0, 4, 7,
	0, 14,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line ast/cll.y:22
		{
			SynTree = yyS[yypt-0].node
		}
	case 2:
		//line ast/cll.y:27
		{
			if list, ok := yyS[yypt-1].node.(*StmtList); ok {
				list.Add(yyS[yypt-0].node)
				yyVAL.node = yyS[yypt-1].node
			} else {
				node := StatementList().(*StmtList)
				node.Add(yyS[yypt-1].node, yyS[yypt-0].node)
				yyS[yypt-1].node.SetParent(node)
				yyVAL.node = node
			}
			yyS[yypt-0].node.SetParent(yyVAL.node)
		}
	case 3:
		//line ast/cll.y:39
		{
			yyVAL.node = Empty()
		}
	case 4:
		//line ast/cll.y:44
		{
			yyVAL.node = Decleration(yyS[yypt-1].node, yyS[yypt-0].node)
			yyS[yypt-0].node.SetParent(yyVAL.node)
		}
	case 5:
		//line ast/cll.y:48
		{
			yyVAL.node = yyS[yypt-0].node
		}
	case 6:
		//line ast/cll.y:49
		{
			yyVAL.node = Empty()
		}
	case 7:
		//line ast/cll.y:54
		{
			yyVAL.node = Assign(yyS[yypt-2].node, yyS[yypt-0].node)
			yyS[yypt-2].node.SetParent(yyVAL.node)
			yyS[yypt-0].node.SetParent(yyVAL.node)
		}
	case 8:
		//line ast/cll.y:60
		{
			yyVAL.node = yyS[yypt-0].node
		}
	case 9:
		//line ast/cll.y:63
		{
			yyVAL.node = yyS[yypt-0].node
		}
	case 10:
		//line ast/cll.y:64
		{
			yyVAL.node = yyS[yypt-0].node
		}
	case 11:
		//line ast/cll.y:65
		{
			yyVAL.node = yyS[yypt-0].node
		}
	case 12:
		//line ast/cll.y:69
		{
			yyVAL.node = Literal(yyS[yypt-0].str, numTy)
		}
	case 13:
		//line ast/cll.y:73
		{
			yyVAL.node = Decl(yyS[yypt-0].str)
		}
	case 14:
		//line ast/cll.y:78
		{
			yyVAL.node = Asm(yyS[yypt-1].str)
		}
	}
	goto yystack /* stack new state and value */
}
