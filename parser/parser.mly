%{
  open Types
%}

%token <string> PARTY_NAME PARTY_NICK
%token END_PARTIES

%start main
%type <contract> main
%type <party list> parties

%%

main:
  parties { { parties = $1 } }
;

parties
  : PARTY_NAME PARTY_NICK { [{ name = $1; identifier = $1 }] }
;
