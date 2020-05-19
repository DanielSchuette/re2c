// Code generated by re2c, DO NOT EDIT.
package main

import "fmt"

const (
	yycinit = iota
	yycbin
	yycdec
	yychex
	yycoct
)


func Lex(str string) int {
	var cursor, ctxmarker int
	cond := yycinit
	n := 0

	
{
	var yych byte
	switch (cond) {
	case yycinit:
		goto yyc_init
	case yycbin:
		goto yyc_bin
	case yycdec:
		goto yyc_dec
	case yychex:
		goto yyc_hex
	case yycoct:
		goto yyc_oct
	}
/* *********************************** */
yyc_init:
	yych = str[cursor]
	switch (yych) {
	case '0':
		goto yy5
	case '1':
		fallthrough
	case '2':
		fallthrough
	case '3':
		fallthrough
	case '4':
		fallthrough
	case '5':
		fallthrough
	case '6':
		fallthrough
	case '7':
		fallthrough
	case '8':
		fallthrough
	case '9':
		ctxmarker = cursor
		goto yy7
	default:
		goto yy3
	}
yy3:
	cursor += 1
	{ return -1 }
yy5:
	cursor += 1
	yych = str[cursor]
	switch (yych) {
	case 'B':
		fallthrough
	case 'b':
		goto yy9
	case 'X':
		fallthrough
	case 'x':
		goto yy11
	default:
		goto yy6
	}
yy6:
	cond = yycoct;
	goto yyc_oct;
yy7:
	cursor += 1
	cursor = ctxmarker
	cond = yycdec;
	goto yyc_dec;
yy9:
	cursor += 1
	cond = yycbin;
	goto yyc_bin;
yy11:
	cursor += 1
	cond = yychex;
	goto yyc_hex;
/* *********************************** */
yyc_bin:
	yych = str[cursor]
	switch (yych) {
	case 0x00:
		goto yy15
	case '0':
		fallthrough
	case '1':
		goto yy19
	default:
		goto yy17
	}
yy15:
	cursor += 1
	{ return n }
yy17:
	cursor += 1
	{ return -1 }
yy19:
	cursor += 1
	{ n = n*2 + int(str[cursor-1] - '0'); goto yyc_bin; }
/* *********************************** */
yyc_dec:
	yych = str[cursor]
	switch (yych) {
	case 0x00:
		goto yy23
	case '0':
		fallthrough
	case '1':
		fallthrough
	case '2':
		fallthrough
	case '3':
		fallthrough
	case '4':
		fallthrough
	case '5':
		fallthrough
	case '6':
		fallthrough
	case '7':
		fallthrough
	case '8':
		fallthrough
	case '9':
		goto yy27
	default:
		goto yy25
	}
yy23:
	cursor += 1
	{ return n }
yy25:
	cursor += 1
	{ return -1 }
yy27:
	cursor += 1
	{ n = n*10 + int(str[cursor-1] - '0'); goto yyc_dec; }
/* *********************************** */
yyc_hex:
	yych = str[cursor]
	switch (yych) {
	case 0x00:
		goto yy31
	case '0':
		fallthrough
	case '1':
		fallthrough
	case '2':
		fallthrough
	case '3':
		fallthrough
	case '4':
		fallthrough
	case '5':
		fallthrough
	case '6':
		fallthrough
	case '7':
		fallthrough
	case '8':
		fallthrough
	case '9':
		goto yy35
	case 'A':
		fallthrough
	case 'B':
		fallthrough
	case 'C':
		fallthrough
	case 'D':
		fallthrough
	case 'E':
		fallthrough
	case 'F':
		goto yy37
	case 'a':
		fallthrough
	case 'b':
		fallthrough
	case 'c':
		fallthrough
	case 'd':
		fallthrough
	case 'e':
		fallthrough
	case 'f':
		goto yy39
	default:
		goto yy33
	}
yy31:
	cursor += 1
	{ return n }
yy33:
	cursor += 1
	{ return -1 }
yy35:
	cursor += 1
	{ n = n*16 + int(str[cursor-1] - '0'); goto yyc_hex; }
yy37:
	cursor += 1
	{ n = n*16 + int(str[cursor-1] - 'A') + 10; goto yyc_hex; }
yy39:
	cursor += 1
	{ n = n*16 + int(str[cursor-1] - 'a') + 10; goto yyc_hex; }
/* *********************************** */
yyc_oct:
	yych = str[cursor]
	switch (yych) {
	case 0x00:
		goto yy43
	case '0':
		fallthrough
	case '1':
		fallthrough
	case '2':
		fallthrough
	case '3':
		fallthrough
	case '4':
		fallthrough
	case '5':
		fallthrough
	case '6':
		fallthrough
	case '7':
		goto yy47
	default:
		goto yy45
	}
yy43:
	cursor += 1
	{ return n }
yy45:
	cursor += 1
	{ return -1 }
yy47:
	cursor += 1
	{ n = n*8 + int(str[cursor-1] - '0'); goto yyc_oct; }
}

}

func test(str string, num int) {
	if res := Lex(str); res != num {
		panic(fmt.Sprintf("failed %s: expected %d, got %d", str, num, res))
	}
}

func main() {
	test("123\000", 123)
	test("0b101\000", 5)
	test("0x10Ff\000", 4096 + 255)
	test("0112\000", 74)
	test("0\000", 0)
	test("\000", -1)
}
golang/009_conditions.ci--lang(go).re:28:19: warning: rule in condition 'init' matches empty string [-Wmatch-empty-string]