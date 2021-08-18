package clanglex

import (
	"fmt"
	"strings"
)

type Lexer struct {
	input string
	pos   int
}

type Token struct {
	TokenType int
	Literal   string
}

const (
	eof = iota
	word
	integer
	float
	assign
	plus
	minus
	bang
	asterisk
	slash
	percent
	lt
	gt
	eq
	ne
	gteq
	lteq
	semicolon
	lparen
	rparen
	comma
	lbrace
	rbrace
	lbracket
	rbracket
	ampersand
	tilde
	caret
	vertical
	colon
	question
	period
	backslash
	str
	letter
	arrow
	leftShift
	rightShift
	increment
	decrement
	and
	or
	plusAssigne
	minusAssigne
	asteriskAssigne
	slashAssigne
	verticalAssigne
	ampersandAssigne
	leftShiftAssigne
	rightShiftAssigne
	tildeAssigne
	caretAssigne
	percentAssigne
	keyReturn
	keyIf
	keyElse
	keyWhile
	keyDo
	keyGoto
	keyFor
	keyBreak
	keyContinue
	keySwitch
	keyCase
	keyDefault
	keyExtern
	keyVolatile
	keyConst
	keyTypedef
	keyUnion
	keyStruct
	keyEnum
	keyAttribute
	keyVoid
	keyAsm
	keySizeof
	comment
	illegal
)

func (t *Token) String() string {
	return fmt.Sprintf("tokenType:%v, literal:%s", t.TokenType, t.Literal)
}

func NewLexer(src string) *Lexer {
	return &Lexer{input: src, pos: 0}
}

func (l *Lexer) lexicalize() []*Token {
	ts := []*Token{}
	for {
		t := l.nextToken()
		ts = append(ts, t)
		if t.TokenType == eof {
			break
		}
	}
	return ts
}

func (l *Lexer) nextToken() *Token {
	// スペースをとばす
	for {
		i := l.pos
		if i >= len(l.input) {
			break
		}
		c := l.input[i]
		if c != ' ' && c != '\t' && c != '\n' && c != '\r' {
			break
		}
		l.pos++
	}

	// ソースの終端
	if l.pos >= len(l.input) {
		return &Token{TokenType: eof, Literal: "eof"}
	}

	var tk *Token
	c := l.input[l.pos]
	switch c {
	case '=':
		tk = &Token{TokenType: assign, Literal: "="}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: eq, Literal: "=="}
			l.pos++
		}
	case '+':
		tk = &Token{TokenType: plus, Literal: "+"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '+' {
			tk = &Token{TokenType: increment, Literal: "++"}
			l.pos++
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: plusAssigne, Literal: "+="}
			l.pos++
		}
	case '-':
		tk = &Token{TokenType: minus, Literal: "-"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '>' {
			// ->
			tk = &Token{TokenType: arrow, Literal: "->"}
			l.pos++
		} else if l.input[l.pos] == '-' {
			tk = &Token{TokenType: decrement, Literal: "--"}
			l.pos++
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: minusAssigne, Literal: "-="}
			l.pos++
		}
	case '!':
		tk = &Token{TokenType: bang, Literal: "!"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: ne, Literal: "!="}
			l.pos++
		}
	case '*':
		tk = &Token{TokenType: asterisk, Literal: "*"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: asteriskAssigne, Literal: "*="}
			l.pos++
		}
	case '/':
		tk = &Token{TokenType: slash, Literal: "/"}
		l.pos++
		if l.pos >= len(l.input) {
			// 何もしない
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: slashAssigne, Literal: "/="}
			l.pos++
		} else if l.input[l.pos] == '*' {
			// comment
			l.pos++
			com := l.readComment()
			tk = &Token{TokenType: comment, Literal: com}
		}
	case '<':
		tk = &Token{TokenType: lt, Literal: "<"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '<' {
			tk = &Token{TokenType: leftShift, Literal: "<<"}
			l.pos++
			if l.pos >= len(l.input) {
			} else if l.input[l.pos] == '=' {
				tk = &Token{TokenType: leftShiftAssigne, Literal: "<<="}
				l.pos++
			}
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: lteq, Literal: "<="}
			l.pos++
		}
	case '>':
		tk = &Token{TokenType: gt, Literal: ">"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '>' {
			tk = &Token{TokenType: rightShift, Literal: ">>"}
			l.pos++
			if l.pos >= len(l.input) {
			} else if l.input[l.pos] == '=' {
				tk = &Token{TokenType: rightShiftAssigne, Literal: ">>="}
				l.pos++
			}
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: gteq, Literal: ">="}
			l.pos++
		}
	case ';':
		tk = &Token{TokenType: semicolon, Literal: ";"}
		l.pos++
	case '(':
		tk = &Token{TokenType: lparen, Literal: "("}
		l.pos++
	case ')':
		tk = &Token{TokenType: rparen, Literal: ")"}
		l.pos++
	case ',':
		tk = &Token{TokenType: comma, Literal: ","}
		l.pos++
	case '{':
		tk = &Token{TokenType: lbrace, Literal: "{"}
		l.pos++
	case '}':
		tk = &Token{TokenType: rbrace, Literal: "}"}
		l.pos++
	case '[':
		tk = &Token{TokenType: lbracket, Literal: "["}
		l.pos++
	case ']':
		tk = &Token{TokenType: rbracket, Literal: "]"}
		l.pos++
	case '&':
		tk = &Token{TokenType: ampersand, Literal: "&"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '&' {
			tk = &Token{TokenType: and, Literal: "&&"}
			l.pos++
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: ampersandAssigne, Literal: "&="}
			l.pos++
		}
	case '~':
		tk = &Token{TokenType: tilde, Literal: "~"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: tildeAssigne, Literal: "~="}
			l.pos++
		}
	case '^':
		tk = &Token{TokenType: caret, Literal: "^"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: caretAssigne, Literal: "^="}
			l.pos++
		}
	case '|':
		tk = &Token{TokenType: vertical, Literal: "|"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '|' {
			tk = &Token{TokenType: or, Literal: "||"}
			l.pos++
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: verticalAssigne, Literal: "|="}
			l.pos++
		}
	case '%':
		tk = &Token{TokenType: percent, Literal: "%"}
		l.pos++
		if l.pos >= len(l.input) {
		} else if l.input[l.pos] == '=' {
			tk = &Token{TokenType: percentAssigne, Literal: "%="}
			l.pos++
		}
	case ':':
		tk = &Token{TokenType: colon, Literal: ":"}
		l.pos++
	case '?':
		tk = &Token{TokenType: question, Literal: "?"}
		l.pos++
	case '.':
		tk = &Token{TokenType: period, Literal: "."}
		l.pos++
	case '\\':
		tk = &Token{TokenType: backslash, Literal: "\\"}
		l.pos++
	case '\'':
		tk = l.readLetter()
	case '"':
		tk = l.readString()
	case '#':
		tk = l.readHashComment()
		l.pos++
	default:
		if isLetter(c) {
			tk = l.readWord()
		} else if isDec(c) {
			tk = l.readNumber()
		}
	}
	return tk
}

func (l *Lexer) readWord() *Token {
	// ワードの終わりの次まで pos を進める
	var next int
	for next = l.pos; next < len(l.input); next++ {
		c := l.input[next]
		if !isLetter(c) && !isDec(c) {
			break
		}
	}
	w := l.input[l.pos:next]
	tk := l.determineKeyword(w)
	l.pos = next
	return tk
}

func (l *Lexer) readNumber() *Token {
	var next int
	isFloat := false

	next = l.pos
	c := l.input[next]
	if c == '0' {
		next++
		c = l.input[next]
		switch c {
		case 'x':
			// 16進数
			next++
		case 'b':
			// 2進数
			next++
		case '.':
			// 小数
			next++
			isFloat = true
		default:
			if isDec(c) {
				// 8進数
				next++
			} else {
				// ゼロ

			}
		}
	}

	// ワードの終わりの次まで pos を進める
	for ; next < len(l.input); next++ {
		c = l.input[next]
		if c == '.' {
			isFloat = true
		} else if c == 'u' || c == 'U' || c == 'l' || c == 'L' {
			continue
		} else if !isHex(c) {
			break
		}
	}
	w := l.input[l.pos:next]
	l.pos = next

	var tk *Token
	if isFloat {
		tk = &Token{TokenType: float, Literal: w}
	} else {
		tk = &Token{TokenType: integer, Literal: w}
	}
	return tk
}

func (l *Lexer) readString() *Token {
	var next int

	// 次の " を探す
	for next = l.pos + 1; next < len(l.input); next++ {
		// エスケープシーケンス考慮
		if l.input[next] == '\\' && l.input[next+1] == '\\' {
			next++
		} else if l.input[next] == '\\' && l.input[next+1] == '"' {
			next++
		} else if l.input[next] == '"' {
			break
		}
	}
	// 次の pos に進める
	next++
	w := l.input[l.pos:next]
	l.pos = next
	return &Token{TokenType: str, Literal: w}
}

func (l *Lexer) readHashComment() *Token {
	// # の次の文字に移動
	l.pos++
	var next int
	for i := l.pos; i <= len(l.input); i++ {
		next = i
		if next >= len(l.input) {
			break
		}
		c := l.input[next]
		if c == '\n' || c == '\r' {
			break
		}
	}
	tk := &Token{TokenType: comment, Literal: l.input[l.pos:next]}
	l.pos = next
	return tk
}

func (l *Lexer) readLetter() *Token {

	l.pos++
	var s []byte
	c := l.input[l.pos]
	if c == '\\' {
		s = l.getEscC()
		l.pos++
	} else {
		s = append(s, c)
		l.pos++
		l.pos++
	}
	return &Token{TokenType: letter, Literal: string(s)}
}

func (l *Lexer) getEscC() []byte {
	res := []byte{}
	res = append(res, l.input[l.pos])
	l.pos++
	if l.input[l.pos] >= '0' && l.input[l.pos] <= '9' {
		for l.input[l.pos] >= '0' && l.input[l.pos] <= '9' {
			res = append(res, l.input[l.pos])
			l.pos++
		}
	} else {
		res = append(res, l.input[l.pos])
		l.pos++
	}

	return res
}

func (l *Lexer) newIllegal() *Token {
	tk := &Token{TokenType: illegal, Literal: l.input[l.pos:]}
	l.pos = len(l.input)
	return tk
}

func (l *Lexer) determineKeyword(w string) *Token {
	if strings.Compare("return", w) == 0 {
		return &Token{TokenType: keyReturn, Literal: w}
	} else if strings.Compare("if", w) == 0 {
		return &Token{TokenType: keyIf, Literal: w}
	} else if strings.Compare("else", w) == 0 {
		return &Token{TokenType: keyElse, Literal: w}
	} else if strings.Compare("while", w) == 0 {
		return &Token{TokenType: keyWhile, Literal: w}
	} else if strings.Compare("do", w) == 0 {
		return &Token{TokenType: keyDo, Literal: w}
	} else if strings.Compare("goto", w) == 0 {
		return &Token{TokenType: keyGoto, Literal: w}
	} else if strings.Compare("for", w) == 0 {
		return &Token{TokenType: keyFor, Literal: w}
	} else if strings.Compare("break", w) == 0 {
		return &Token{TokenType: keyBreak, Literal: w}
	} else if strings.Compare("continue", w) == 0 {
		return &Token{TokenType: keyContinue, Literal: w}
	} else if strings.Compare("switch", w) == 0 {
		return &Token{TokenType: keySwitch, Literal: w}
	} else if strings.Compare("case", w) == 0 {
		return &Token{TokenType: keyCase, Literal: w}
	} else if strings.Compare("default", w) == 0 {
		return &Token{TokenType: keyDefault, Literal: w}
	} else if strings.Compare("extern", w) == 0 {
		return &Token{TokenType: keyExtern, Literal: w}
	} else if strings.Compare("volatile", w) == 0 {
		return &Token{TokenType: keyVolatile, Literal: w}
	} else if strings.Compare("const", w) == 0 {
		return &Token{TokenType: keyConst, Literal: w}
	} else if strings.Compare("typedef", w) == 0 {
		return &Token{TokenType: keyTypedef, Literal: w}
	} else if strings.Compare("union", w) == 0 {
		return &Token{TokenType: keyUnion, Literal: w}
	} else if strings.Compare("struct", w) == 0 {
		return &Token{TokenType: keyStruct, Literal: w}
	} else if strings.Compare("enum", w) == 0 {
		return &Token{TokenType: keyEnum, Literal: w}
	} else if strings.Compare("__attribute__", w) == 0 {
		return &Token{TokenType: keyAttribute, Literal: w}
	} else if strings.Compare("void", w) == 0 {
		return &Token{TokenType: keyVoid, Literal: w}
	} else if strings.Compare("__asm", w) == 0 {
		return &Token{TokenType: keyAsm, Literal: w}
	} else if strings.Compare("sizeof", w) == 0 {
		return &Token{TokenType: keySizeof, Literal: w}
	} else {
		return &Token{TokenType: word, Literal: w}
	}
}

func (l *Lexer) readComment() string {
	res := []byte{}

	for !l.isCommentEnd() {
		res = append(res, l.input[l.pos])
		l.pos++
	}
	l.pos++
	l.pos++
	// next

	return string(res)
}

func (l *Lexer) isCommentEnd() bool {
	// */ か確認
	return l.input[l.pos] == '*' && l.input[l.pos+1] == '/'
}

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}

func isHex(c byte) bool {
	return '0' <= c && c <= '9' || 'a' <= c && c <= 'f' || 'A' <= c && c <= 'F'
}

func isDec(c byte) bool {
	return '0' <= c && c <= '9'
}

// isTypeToken
func (t *Token) isTypeToken() bool {

	switch t.TokenType {
	case word:
	case asterisk:
	case keyConst:
	case keyVoid:
	case keyStruct:
	case keyUnion:
	case keyEnum:
	case keyVolatile:
	case caret:
		// clang でコンパイルした場合型の種類に^が含まれる？
	default:
		return false
	}

	return true
}

func (t *Token) isOperator() bool {
	switch t.TokenType {
	case assign:
	case plus:
	case minus:
	case asterisk:
	case slash:
	case lt:
	case gt:
	case eq:
	case gteq:
	case lteq:
	case ne:
	case ampersand:
	case tilde:
	case caret:
	case vertical:
	case question:
	case leftShift:
	case rightShift:
	case increment:
	case decrement:
	case or:
	case and:
	case percent:
	case colon:
	case plusAssigne:
	case minusAssigne:
	case asteriskAssigne:
	case slashAssigne:
	case verticalAssigne:
	case ampersandAssigne:
	case leftShiftAssigne:
	case rightShiftAssigne:
	case tildeAssigne:
	case caretAssigne:
	case percentAssigne:
	default:
		return false
	}
	return true
}

func (t *Token) isPrefixExpression() bool {
	switch t.TokenType {
	case minus:
	case plus:
	case increment:
	case decrement:
	case tilde:
	case bang:
	case asterisk:
	case ampersand:
	default:
		return false
	}
	return true
}

func (t *Token) isPostExpression() bool {
	switch t.TokenType {
	case lparen:
	case increment:
	case decrement:
	default:
		return false
	}
	return true
}

func (t *Token) isCompoundOp() bool {
	switch t.TokenType {
	case plusAssigne:
	case minusAssigne:
	case asteriskAssigne:
	case slashAssigne:
	case verticalAssigne:
	case ampersandAssigne:
	case leftShiftAssigne:
	case rightShiftAssigne:
	case tildeAssigne:
	case caretAssigne:
	case percentAssigne:
	default:
		return false
	}
	return true
}

func (t *Token) isToken(t2 int) bool {
	return t.TokenType == t2
}
