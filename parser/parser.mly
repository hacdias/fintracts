%{
  open Types
%}

%token DATE_SEP
%token BOND_PURCHASE_AGREEMENT MATURITY_ON COUPONS_RATE_OF PAID_ON
%token INTEREST_RATE_SWAP_AGREEMENT AGREE_INTEREST_RATE_SWAP_OVER WITH_EFFECTIVE_DATE AND_TERMINATION
%token WILL_PAY FIXED_INTEREST FLOATING_INTEREST INITIALLY_DEFINED OVER_ON_DATES FLOATING_OPTION_IS

%token <int> INT
%token <float> FLOAT MONEY

%token <string> WORD
%token AND AS A AN AGREES
%token BY BOND
%token DEFINED
%token ENTER
%token FOR FOLLOWS
%token HEREBY
%token IN ISSUING
%token OF ON
%token PARTIES
%token SIGNED SELLING
%token TO THE
%token UNDERMENTIONED

%token <char> PUNCTUATION
%token COMMA DOT COLON SEMICOLON PERCENT

%start main
%type <contract> main

%{
  (* Maybe this is not needed *)
%}

%type <signature> signature
%type <date> date
%type <date list> dates
%type <string list> signature_parties parties_name
%type <party list> parties
%type <bondPurchase option> bond_purchase_agreement
%type <interestRateSwap option> interest_rate_swap_agreement
%type <coupons option> bond_coupons
%type <money> money
%type <agreement> agreement
%type <agreement list> agreements

%%

main
  : parties agreements signature                                        { { parties = $1; agreements = $2; signature = $3; } }
;

parties
  : THE PARTIES COLON parties                                           { $4 }
  | parties_name UNDERMENTIONED AS WORD SEMICOLON AND parties           { { name = (String.concat " " $1); identifier = $4 } :: $7 }
  | parties_name UNDERMENTIONED AS WORD DOT                             { [{ name = (String.concat " " $1); identifier = $4 }] }
;

parties_name
  : WORD parties_name                                                   { $1 :: $2 }
  | WORD COMMA                                                          { [$1]}
;

signature
  : SIGNED BY signature_parties ON THE date DOT                         { { parties = $3; date = $6 } }
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
  : HEREBY ENTER IN A BOND_PURCHASE_AGREEMENT
    DEFINED AS FOLLOWS COLON bond_purchase_agreement                    { {
                                                                            bondPurchase = $10;
                                                                            interestRateSwap = None;
                                                                            currencySwap = None
                                                                        } }
  | HEREBY ENTER IN AN INTEREST_RATE_SWAP_AGREEMENT
    DEFINED AS FOLLOWS COLON interest_rate_swap_agreement               { {
                                                                            bondPurchase = None;
                                                                            interestRateSwap = $10;
                                                                            currencySwap = None
                                                                        } }
;

bond_purchase_agreement
  : WORD AGREES ON ISSUING AND SELLING A BOND OF money
    TO WORD FOR money DOT MATURITY_ON date DOT bond_coupons            { Some {
                                                                            seller = $1;
                                                                            payer = $12;
                                                                            issuePrice = $14;
                                                                            faceValue = $10;
                                                                            maturityDate = $17;
                                                                            coupons = $19
                                                                        } }
;

bond_coupons
  :                                                                     { None }
  | COUPONS_RATE_OF FLOAT PERCENT PAID_ON dates DOT                     { Some { rate = $2; dates = $5 } }
;

interest_rate_swap_agreement
  : AGREE_INTEREST_RATE_SWAP_OVER money WITH_EFFECTIVE_DATE date
    AND_TERMINATION date DOT interest_payments                                           { Some {
                                                                            notationalAmount = $2;
                                                                            effectiveDate =  $4;
                                                                            maturityDate = $6;
                                                                            interest = $8
                                                                        } }
;

interest_payments
  : interest_payment                                                    { [$1] }
  | interest_payment interest_payments                                  { $1 :: $2 }
;

interest_payment
  : WORD WILL_PAY FIXED_INTEREST OF FLOAT PERCENT OVER_ON_DATES
    dates DOT                                                           { {
                                                                            payer = $1;
                                                                            dates = $8;
                                                                            fixedRate = $5;
                                                                            initialRate = 0.0;
                                                                            interestRateOption = ""
                                                                        } }
  | WORD WILL_PAY FLOATING_INTEREST INITIALLY_DEFINED FLOAT PERCENT
    COMMA OVER_ON_DATES dates DOT FLOATING_OPTION_IS WORD DOT           { {
                                                                            payer = $1;
                                                                            dates = $9;
                                                                            fixedRate = 0.0;
                                                                            initialRate = $5;
                                                                            interestRateOption = $12
                                                                        } }
;