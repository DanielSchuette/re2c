// re2c:eof = 0;
// re2c:define:YYFILL = "fill(@@)";
// re2c:yyfill:check = 1;
/*!re2c
  re2c:eof = 0;
  re2c:define:YYFILL = "fill(@@)";
  re2c:yyfill:check = 1;
  $ { eof }
  [a] { a }
  *   { x }
*/
