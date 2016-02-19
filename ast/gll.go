//line ast/gll.y:2
package ast

import __yyfmt__ "fmt"

//line ast/gll.y:2
var SynTree ASTNode

//line ast/gll.y:8
type yySymType struct {
	yys  int
	num  int
	str  string
	node ASTNode
}

const END_STMT = 57346
const ASSIGN = 57347
const ASM = 57348
const LBRACE = 57349
const RBRACE = 57350
const COLON = 57351
const IF = 57352
const ELSE = 57353
const ID = 57354
const NUM = 57355
const ASM_BODY = 57356
const OP = 57357

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"END_STMT",
	"ASSIGN",
	"ASM",
	"LBRACE",
	"RBRACE",
	"COLON",
	"IF",
	"ELSE",
	"ID",
	"NUM",
	"ASM_BODY",
	"OP",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line ast/gll.y:120

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 21
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 52

var yyAct = [...]int{

	5, 23, 7, 25, 2, 26, 8, 12, 24, 16,
	17, 9, 19, 17, 13, 29, 22, 6, 8, 12,
	14, 35, 27, 9, 32, 4, 13, 24, 31, 30,
	34, 33, 6, 8, 12, 21, 20, 15, 9, 28,
	4, 13, 8, 12, 18, 10, 11, 9, 3, 17,
	13, 1,
}
var yyPact = [...]int{

	-1000, -1000, 28, -1000, 11, -1000, -1000, 32, 37, -2,
	-1000, -1000, 29, -1000, 30, 37, -1000, -1000, 20, -12,
	-9, 37, -1000, 4, -1000, -2, 16, -1000, -1000, 1,
	13, -1000, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 51, 4, 48, 0, 46, 2, 45, 1, 44,
	39,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 3, 3, 3, 4, 4, 4,
	4, 4, 4, 10, 10, 10, 9, 8, 5, 6,
	7,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 4, 1, 1, 3, 2, 4,
	1, 1, 1, 2, 2, 0, 3, 3, 1, 1,
	4,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, 12, -4, 4, -6, 5, 10,
	-7, -5, 6, 13, 9, 5, -4, 12, -9, -6,
	7, 5, -4, -8, 7, 15, 14, -4, -10, 11,
	-2, -6, 8, -8, -4, 8,
}
var yyDef = [...]int{

	3, -2, 1, 2, 19, 5, 6, 12, 0, 0,
	10, 11, 0, 18, 0, 0, 8, 19, 0, 0,
	0, 0, 7, 15, 3, 0, 0, 4, 9, 0,
	0, 16, 20, 13, 14, 17,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lookahead func() int
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
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

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yytoken := -1 // yychar translated into internal numbering
	yyrcvr.lookahead = func() int { return yychar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yychar = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
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
		yychar, yytoken = yylex1(yylex, &yylval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yychar = -1
		yytoken = -1
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
			yychar, yytoken = yylex1(yylex, &yylval)
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
			if yyn < 0 || yyn == yytoken {
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
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yychar = -1
			yytoken = -1
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
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
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
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ast/gll.y:22
		{
			SynTree = yyDollar[1].node
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ast/gll.y:27
		{
			if list, ok := yyDollar[1].node.(*BlockStmt); ok {
				list.Add(yyDollar[2].node)
				yyVAL.node = yyDollar[1].node
			} else {
				node := BlockStatement().(*BlockStmt)
				node.Add(yyDollar[1].node, yyDollar[2].node)
				yyDollar[1].node.SetParent(node)
				yyVAL.node = node
			}
			yyDollar[2].node.SetParent(yyVAL.node)
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ast/gll.y:39
		{
			yyVAL.node = Empty()
		}
	case 4:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ast/gll.y:44
		{
			yyVAL.node = Decleration(Decl(yyDollar[1].str), yyDollar[4].node)
			yyDollar[4].node.SetParent(yyVAL.node)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ast/gll.y:48
		{
			yyVAL.node = yyDollar[1].node
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ast/gll.y:49
		{
			yyVAL.node = Empty()
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ast/gll.y:54
		{
			yyVAL.node = Assign(yyDollar[1].node, yyDollar[3].node)
			yyDollar[1].node.SetParent(yyVAL.node)
			yyDollar[3].node.SetParent(yyVAL.node)
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ast/gll.y:60
		{
			yyVAL.node = yyDollar[2].node
		}
	case 9:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ast/gll.y:64
		{
			yyVAL.node = If(yyDollar[2].node, yyDollar[3].node, yyDollar[4].node)
			yyDollar[2].node.SetParent(yyVAL.node)
			yyDollar[3].node.SetParent(yyDollar[3].node)
			if yyDollar[4].node != nil {
				yyDollar[4].node.SetParent(yyVAL.node)
			}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ast/gll.y:72
		{
			yyVAL.node = yyDollar[1].node
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ast/gll.y:73
		{
			yyVAL.node = yyDollar[1].node
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ast/gll.y:74
		{
			yyVAL.node = yyDollar[1].node
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ast/gll.y:79
		{
			yyVAL.node = yyDollar[2].node
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ast/gll.y:83
		{
			yyVAL.node = yyDollar[2].node
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ast/gll.y:86
		{
			yyVAL.node = nil
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ast/gll.y:91
		{
			yyVAL.node = Binary(yyDollar[1].node, yyDollar[2].str, yyDollar[3].node)
			yyDollar[1].node.SetParent(yyVAL.node)
			yyDollar[3].node.SetParent(yyVAL.node)
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ast/gll.y:100
		{
			yyVAL.node = yyDollar[2].node
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ast/gll.y:106
		{
			yyVAL.node = Literal(yyDollar[1].str, numTy)
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ast/gll.y:110
		{
			yyVAL.node = Id(yyDollar[1].str)
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ast/gll.y:115
		{
			yyVAL.node = Asm(yyDollar[3].str)
		}
	}
	goto yystack /* stack new state and value */
}
