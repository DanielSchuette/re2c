// re2c:define:YYGETCONDITION:naked = 0;
// re2c:decorate = 0;
/*!re2c
  re2c:define:YYGETCONDITION:naked = 0;
  re2c:decorate = 0;
  <c1> [a] => c2 { a }
  <c2> [b] => c1 { b }
  <*> *          { x }
*/
