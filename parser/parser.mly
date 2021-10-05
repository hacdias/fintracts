%{
  open Types
%}

%token <string> WORD
%token <int> INT
%token <float> FLOAT MONEY
%token SIGNED_BY THE_PARTIES
%token AND OF TO FOR ON THE
%token COMMA SEMICOLON DOT PERCENT
%token DATE_SEP
%token UNDERMENTIONED_AS HEREBY_ENTER DEFINED_AS
%token BOND_PURCHASE_AGREEMENT AGREE_BOND_OF MATURITY_ON

%start main
%type <contract> main
%type <signature> signature
%type <date> date
%type <date list> dates
%type <string list> signature_parties parties_name
%type <party list> parties 
%type <bondPurchase option> bond_purchase_agreement 
%type <money> money 
%type <agreement> agreement 
%type <agreement list> agreements 

%%

main
  : parties agreements signature                                        { { parties = $1; agreements = $2; signature = $3; } }
;

parties
  : THE_PARTIES parties                                                 { $2 }
  | parties_name UNDERMENTIONED_AS WORD SEMICOLON AND parties           { { name = (String.concat " " $1); identifier = $3 } :: $6 }
  | parties_name UNDERMENTIONED_AS WORD DOT                             { [{ name = (String.concat " " $1); identifier = $3 }] }
;

parties_name
  : WORD parties_name                                                   { $1 :: $2 }
  | WORD COMMA                                                          { [$1]}
;

signature
  : SIGNED_BY signature_parties ON THE date DOT                         { { parties = $2; date = $5 } }
;

signature_parties
  : WORD signature_parties                                              { $1 :: $2  }
  | COMMA WORD signature_parties                                        { $2 :: $3 }
  | AND WORD                                                            { [$2] }
;

dates
  : date dates                                                          { $1 :: $2 }
  | COMMA date dates                                                    { $2 :: $3 }
  | AND date                                                            { [$2] }
;

date
  : INT DATE_SEP OF WORD INT                                            { { day = $1; month = $4; year = $5 }}
;

money
  : WORD MONEY                                                          { { currency = $1; amount = $2 }}
;

agreements
  : agreement                                                           { [$1] }
  | agreement agreements                                                { $1 :: $2 }
;

agreement
  : HEREBY_ENTER BOND_PURCHASE_AGREEMENT
    DEFINED_AS bond_purchase_agreement                                  { {
                                                                            bondPurchase = $4;
                                                                            interestRateSwap = None;
                                                                            currencySwap = None
                                                                        } }
;

bond_purchase_agreement
  : WORD AGREE_BOND_OF money TO WORD FOR money DOT MATURITY_ON date DOT { Some {
                                                                            seller = $1;
                                                                            payer = $5;
                                                                            issuePrice = $7;
                                                                            faceValue = $3;
                                                                            maturityDate = $10;
                                                                            coupons = None
                                                                        } }
;
