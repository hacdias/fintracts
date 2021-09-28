%token <int> INT
%token <string> IDENT
%token EOL

%start main
%type <Types.contract> main

%%

main:
  IDENT EOL { { parties = $1 } }
;
