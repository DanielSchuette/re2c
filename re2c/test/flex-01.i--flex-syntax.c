/* Generated by re2c */

{
	YYCTYPE yych;
	if (YYLIMIT <= YYCURSOR) YYFILL(1);
	yych = *YYCURSOR;
	switch (yych) {
	case 'a':	goto yy3;
	case 'b':	goto yy5;
	default:	goto yy2;
	}
yy2:
yy3:
	++YYCURSOR;
	{ return "a"; }
yy5:
	++YYCURSOR;
	{ return "b"; }
}

re2c: warning: line 9: control flow is undefined for strings that match '[\x0-\x60\x63-\xFF]', use default rule '*' [-Wundefined-control-flow]