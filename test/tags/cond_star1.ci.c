/* Generated by re2c */

{
	YYCTYPE yych;
	switch (YYGETCONDITION()) {
	case yycc1:
		goto yyc_c1;
	case yycc2:
		goto yyc_c2;
	}
/* *********************************** */
yyc_c1:
	if ((YYLIMIT - YYCURSOR) < 3) YYFILL(3);
	yych = *YYCURSOR;
	switch (yych) {
	case 'a':	goto yy5;
	default:	goto yy3;
	}
yy3:
	++YYCURSOR;
	{}
yy5:
	yych = *++YYCURSOR;
	YYCTXMARKER = YYCURSOR;
	switch (yych) {
	case 'b':	goto yy7;
	default:	goto yy6;
	}
yy6:
	YYCURSOR = YYCTXMARKER;
	{}
yy7:
	yych = *++YYCURSOR;
	switch (yych) {
	case 'b':	goto yy9;
	case 'c':
		YYCTXMARKER = YYCURSOR;
		goto yy11;
	default:
		YYCTXMARKER = YYCURSOR;
		goto yy8;
	}
yy8:
	YYCURSOR = YYCTXMARKER;
	{}
yy9:
	++YYCURSOR;
	if (YYLIMIT <= YYCURSOR) YYFILL(1);
	yych = *YYCURSOR;
	switch (yych) {
	case 'b':	goto yy9;
	default:	goto yy6;
	}
yy11:
	++YYCURSOR;
	if (YYLIMIT <= YYCURSOR) YYFILL(1);
	yych = *YYCURSOR;
	switch (yych) {
	case 'c':	goto yy11;
	default:	goto yy8;
	}
/* *********************************** */
yyc_c2:
	if ((YYLIMIT - YYCURSOR) < 2) YYFILL(2);
	yych = *YYCURSOR;
	switch (yych) {
	case 'a':	goto yy17;
	default:	goto yy15;
	}
yy15:
	++YYCURSOR;
	{}
yy17:
	yych = *++YYCURSOR;
	YYCTXMARKER = YYCURSOR;
	switch (yych) {
	case 'b':	goto yy19;
	default:	goto yy23;
	}
yy18:
	YYCURSOR = YYCTXMARKER;
	{}
yy19:
	++YYCURSOR;
	if (YYLIMIT <= YYCURSOR) YYFILL(1);
	yych = *YYCURSOR;
	switch (yych) {
	case 'b':	goto yy19;
	default:	goto yy21;
	}
yy21:
	YYCURSOR = YYCTXMARKER;
	{}
yy22:
	++YYCURSOR;
	if (YYLIMIT <= YYCURSOR) YYFILL(1);
	yych = *YYCURSOR;
yy23:
	switch (yych) {
	case 'c':	goto yy22;
	default:	goto yy18;
	}
}

