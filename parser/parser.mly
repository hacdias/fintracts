%{
  open Types
%}

%token <string> WORD
%token <int> INT
%token SIGNED_BY THE_PARTIES
%token AND OF
%token COMMA SEMICOLON DOT
%token DATE_SEP
%token ON_THE
%token UNDERMENTIONED

%start main
%type <contract> main
%type <signature> signature
%type <date> date
%type <date list> dates
%type <string list> signature_parties parties_name
%type <party list> parties 

%%

main
  : parties signature                                         { { parties = $1; signature = $2; } }
;

parties
  : THE_PARTIES parties                                       { $2 }
  | parties_name UNDERMENTIONED WORD SEMICOLON AND parties    { { name = (String.concat " " $1); identifier = $3 } :: $6 }
  | parties_name UNDERMENTIONED WORD DOT                      { [{ name = (String.concat " " $1); identifier = $3 }] }
;

parties_name
  : WORD parties_name { $1 :: $2 }
  | WORD COMMA        { [$1]}
;

signature
  : SIGNED_BY signature_parties ON_THE date DOT               { { parties = $2; date = $4 } }
;

signature_parties
  : WORD signature_parties                                    { $1 :: $2  }
  | COMMA WORD signature_parties                              { $2 :: $3 }
  | AND WORD                                                  { [$2] }
;

dates
  : date dates                                                { $1 :: $2 }
  | COMMA date dates                                          { $2 :: $3 }
  | AND date                                                  { [$2] }
;

date
  : INT DATE_SEP OF WORD INT                                  { { day = $1; month = $4; year = $5 }}
;
