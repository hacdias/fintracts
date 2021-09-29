%{
  open Types
%}

%token <string> WORD
%token THE_PARTIES
%token COMMA SEMI_COL DOT
%token UNDERMENTIONED
%token AND

%token <string> TPARTY

%start main
%type <contract> main

%%

main:
  THE_PARTIES parties { { parties = $2 } }
;

parties
  : WORD COMMA UNDERMENTIONED WORD SEMI_COL AND parties { { name = $1; identifier = $4 } :: $7 }
  | WORD COMMA UNDERMENTIONED WORD DOT  { [{ name = $1; identifier = $4 }] }
;