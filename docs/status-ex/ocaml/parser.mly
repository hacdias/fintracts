%{
  open Types
%}

%token <string> WORD
%token <int> INT
%token SIGNED_BY
%token AND
%token ON_THE
%token DOT
%token COMMA
%token DATE_SEP
%token OF

%start main
%type <signature> main
%type <date> date
%type <string list> parties

%%

main
  : SIGNED_BY parties ON_THE date DOT { { parties = $2; date = $4 } }
;

parties
  : WORD parties                    { $1 :: $2  }
  | COMMA WORD parties              { $2 :: $3 }
  | AND WORD                        { [$2] }
;

date
  : INT DATE_SEP OF WORD INT { { day = $1; month = $4; year = $5 }}
;