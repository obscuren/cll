package ast

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// This lexer is heavily inspired on Rob Pike's lexer

func isAlphaNumeric(t rune) bool {
	return unicode.IsLetter(t)
}

func isSpace(t rune) bool {
	return unicode.IsSpace(t)
}

func isNumber(t rune) bool {
	return unicode.IsNumber(t)
}

func isOperator(t rune) bool {
	return strings.IndexRune("=+-/*><^!%&|", t) >= 0
}

const eof = -2

type stateFn func(*Lexer) stateFn
type itemType int

const (
	itemEof          itemType = 0
	itemIdentifier            = ID
	itemNumber                = NUM
	itemEndStatement          = END_STMT
	itemAssign                = ASSIGN
	itemAsm                   = ASM
	itemAsmBody               = ASM_BODY

	itemColon  = COLON
	itemLbrace = LBRACE
	itemRbrace = RBRACE
)

type item struct {
	typ itemType
	val string
}

/*
func lexLambda(l *Lexer) stateFn {
	count := 0
out:
	for {
		switch l.next() {
		case '{':
			if count == 0 {
				l.emit(itemLeftBracket)

				l.ignore()
			}
			count++
		case '}':
			count--
			if count == 0 {
				l.backup()

				break out
			}
		}
	}

	l.emit(itemCode)

	l.next()
	l.emit(itemRightBracket)

	return lexText(l)
}
*/

func lexStatement(l *Lexer) stateFn {
	acceptance := Alpha
	l.acceptRun(acceptance)

	if l.accept("_1234567890") {
		acceptance += "_1234567890"
	}
	l.acceptRun(acceptance)

	switch l.blob() {
	case "asm":
		l.emit(itemAsm)

		return lexAsm
	default:
		l.emit(itemIdentifier)
	}

	return lexText
}

const Numbers = "1234567890"
const Alpha = "abcdefghijklmnopqrstuwvxyzABCDEFGHIJKLMNOPQRSTUWVXYZ"

/*
func lexArray(l *Lexer) stateFn {
	if !l.accept("(") {
		l.err = fmt.Errorf("Exepcted '('")
		return nil
	}

	l.emit(itemLeftPar)

	if !l.accept(Numbers) {
		l.err = fmt.Errorf("Expected number")
		return nil
	}

	l.acceptRun(Numbers)

	l.emit(itemNumber)

	if !l.accept(")") {
		l.err = fmt.Errorf("Expected ')'")
		return nil
	}

	l.emit(itemRightPar)

	return lexText
}
*/

func lexNumber(l *Lexer) stateFn {
	digits := "0123456789"
	if l.accept("0") && l.accept("xX") {
		digits = "0123456789abcdefABCDEF"
	}

	l.acceptRun(digits)

	l.emit(itemNumber)

	return lexText
}

func lexAsm(l *Lexer) stateFn {
out:
	for {
		switch r := l.next(); {
		case isSpace(r):
			l.ignore()
		case r == '{':
			l.emit(itemLbrace)
		case r == '}':
			l.backup()

			break out
		default:
			if !l.acceptRunUntill('}') {
				return nil
			}

			l.emit(itemAsmBody)
		}
	}

	return lexText
}

/* TODO
func lexInsidePar(l *Lexer) stateFn {
	for {
		switch r := l.next(); {
		case isComma(r):
			l.ignore()
		case isAlphaNumeric(r):
			l.backup()

			return lexStatement
		}
	}
}
*/

func lexOperator(l *Lexer) stateFn {
	// The only special case there is, assignment

	acceptance := "="
	if !l.accept("=") {
		if l.accept("!") {
			acceptance += "!"
		} else {
			acceptance += "-/*+><^%&|"
		}
	}

	l.acceptRun(acceptance)

	switch l.blob() {
	case "=":
		l.emit(itemAssign)
		/*
			case "&":
				l.emit(itemAnd)
			case "*":
				l.emit(itemMul)
			case "++", "--":
				l.emit(itemDop)
			default:
				l.emit(itemOp)
		*/
	}

	return lexText
}

/*
func lexInsideString(l *Lexer) stateFn {
	if !l.acceptRunUntill('"') {
		l.err = fmt.Errorf("Expected '\"'")
		return nil
	}

	l.emit(itemStr)

	l.next()
	l.emit(itemQuote)

	return lexText
}
*/

func lexComment(l *Lexer) stateFn {
	l.acceptRunUntill('\n')

	return lexText
}

var Lineno int

// Lex text attempts to identify the current state that *might*
// be and calls the appropriate lexing method. The lexing method
// should then take care of anything that is current (even validating)
func lexText(l *Lexer) stateFn {
	for {
		switch r := l.next(); {
		case r == '\n':
			Lineno++

			l.emit(itemEndStatement)
			//l.ignore()
		case r == '\t':
			l.ignore()
		case isSpace(r): // Check whether this is a space (which we ignore)
			l.ignore()
		case isAlphaNumeric(r) || r == '_': // Check if it's alpha numeric (var, if, else etc)
			l.backup()

			return lexStatement
		case isNumber(r): // Check if it's a number (constant)
			l.backup()

			return lexNumber
		case r == '{':
			l.emit(itemLbrace)
		case r == '}':
			l.emit(itemRbrace)
			/*
				case r == '[':
					l.emit(itemLeftBracket)
				case r == ']':
					l.emit(itemRightBracket)
				case r == '(':
					l.emit(itemLeftPar)
				case r == ')':
					l.emit(itemRightPar)
				case r == '.':
					l.emit(itemDot)
				case r == ',':
					l.emit(itemComma)
				case r == '"':
					l.emit(itemQuote)

					return lexInsideString
			*/
		case r == '/', r == '#':
			return lexComment
		case r == ';':
			l.emit(itemEndStatement)
		case r == ':':
			l.emit(itemColon)
		case isOperator(r):
			l.backup()

			return lexOperator
		default:
			return nil
		}
	}

	return nil
}

type Lexer struct {
	name   string
	input  string
	start  int
	pos    int
	width  int
	state  stateFn
	items  chan item
	err    error
	lineno int
}

func lexer(name, input string) *Lexer {
	l := &Lexer{
		name:  name,
		input: input,
		state: lexText,
		items: make(chan item, 20),
	}
	Lineno = 0

	return l
}

// Grabs the current blob of text
func (l *Lexer) blob() string {
	return l.input[l.start:l.pos]
}

// Emits a new item on to item channel for processing
func (l *Lexer) emit(t itemType) {
	l.items <- item{t, l.blob()}
	l.start = l.pos
}

// Accepts checks whether the given input matches the next rune
func (l *Lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}

	l.backup()

	return false
}

// Continues *eating* the next rune until no longer valid
func (l *Lexer) acceptRegexp(valid string) bool {
	if MatchRegexp(valid, []byte(string(l.next()))) {
		return true
	}
	l.backup()

	return false
}

func MatchRegexp(reg string, str []byte) bool {
	ok, _ := regexp.Match(reg, str)
	return ok
}

func (l *Lexer) acceptRunRegexp(valid string) {
	for r := l.next(); MatchRegexp(valid, []byte(string(r))); {
	}
	l.backup()
}

func (l *Lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

func (l *Lexer) acceptRunUntill(until rune) bool {
	// Continues running until a rune is found
	var i rune
	for i = l.next(); strings.IndexRune(string(until), i) == -1; i = l.next() {
		if i == eof {
			break
		}
	}

	l.backup()

	if i == eof {
		return false
	}

	return true
}

// Grabs the next item of the channel and returns it, or nil if we're done
func (l *Lexer) nextItem() item {
	for {
		select {
		case item := <-l.items:
			return item
		default:
			if l.state == nil {
				return item{}
			}

			l.state = l.state(l)
		}
	}

	panic("not reached")
}

// Takes the next rune and returns it or returns EOF
func (l *Lexer) next() (rune rune) {
	if l.pos >= len(l.input) {
		l.width = 0

		return eof
	}
	rune, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width

	return rune
}

// Look ahead
func (l *Lexer) peek() rune {
	rune := l.next()
	l.backup()
	return rune
}

// Backup a previous *next*
func (l *Lexer) backup() {
	l.pos -= l.width
}

// Ignore the current rune
func (l *Lexer) ignore() {
	l.start = l.pos
}

// yacc's lexing method
func (l *Lexer) Lex(lval *yySymType) int {
	item := l.nextItem()
	lval.str = item.val

	return int(item.typ)
}

func (l *Lexer) Error(s string) {
	l.err = fmt.Errorf("line %d: %s: %s", Lineno, s, l.blob())
}
