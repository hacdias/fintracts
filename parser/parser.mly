%token <int> INT
%token <string> WORD
%token EOL
%token THE_PARTIES
%token COMMA SEMI_COL DOT
%token UNDERMENTIONED
%token AND

%start main
%type <Types.contract> main

%%

main:
  THE_PARTIES parties { { parties = $2 } }
;

parties
  : WORD COMMA UNDERMENTIONED WORD SEMI_COL AND parties { $1 :: $7 }
  | WORD COMMA UNDERMENTIONED WORD DOT  { [$1] }
;