%{
  open Types
%}

%token <int> INT
%token <float> FLOAT MONEY
%token COMMA DOT COLON SEMICOLON PERCENT

%token <string> WORD
%token AND AS A AN AGREE AGREES AGREEMENT AFOREMENTIONED AMOUNT
%token BY BOND
%token COUPONS
%token DEFINED DATES DATE DATE_SEP
%token ENTER EFFECTIVE
%token FOR FOLLOWS FOLLOWING FIXED FLOATING
%token HEREBY HAS
%token IN IS ISSUING INTEREST INITIALLY
%token MATURITY
%token NOTATIONAL
%token OF ON OVER OPTION
%token PARTIES PAID PURCHASE PRINCIPAL PAY
%token REACHES RATE
%token SIGNED SELLING SWAP
%token TO THE TRANSACTION TERMINATION
%token UNDERMENTIONED
%token WITH WILL

%start main
%type <contract> main

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
  : HEREBY ENTER IN A BOND PURCHASE AGREEMENT
    DEFINED AS FOLLOWS COLON bond_purchase_agreement                    { {
                                                                            bondPurchase = $12;
                                                                            interestRateSwap = None;
                                                                            currencySwap = None
                                                                        } }
  | HEREBY ENTER IN AN INTEREST RATE SWAP TRANSACTION AGREEMENT
    DEFINED AS FOLLOWS COLON interest_rate_swap_agreement               { {
                                                                            bondPurchase = None;
                                                                            interestRateSwap = $14;
                                                                            currencySwap = None
                                                                        } }
;

bond_purchase_agreement
  : WORD AGREES ON ISSUING AND SELLING A BOND OF money
    TO WORD FOR money DOT THE AFOREMENTIONED BOND REACHES MATURITY
    ON THE date DOT bond_coupons                                        { Some {
                                                                            seller = $1;
                                                                            payer = $12;
                                                                            issuePrice = $14;
                                                                            faceValue = $10;
                                                                            maturityDate = $23;
                                                                            coupons = $25
                                                                        } }
;

bond_coupons
  :                                                                     { None }
  | THE BOND HAS COUPONS WITH AN INTEREST RATE OF
    FLOAT PERCENT PAID ON THE FOLLOWING DATES COLON dates DOT           { Some { rate = $10; dates = $18 } }
;

interest_rate_swap_agreement
  : THE PARTIES AGREE ON AN INTEREST RATE SWAP TRANSACTION OVER
    THE NOTATIONAL PRINCIPAL OF money COMMA WITH AN EFFECTIVE DATE
    AS OF THE date AND TERMINATION ON THE date DOT interest_payments    { Some {
                                                                            notationalAmount = $15;
                                                                            effectiveDate =  $24;
                                                                            maturityDate = $29;
                                                                            interest = $31
                                                                        } }
;

interest_payments
  : interest_payment                                                    { [$1] }
  | interest_payment interest_payments                                  { $1 :: $2 }
;

interest_payment
  : WORD WILL PAY A FIXED RATE INTEREST OF FLOAT PERCENT OVER THE
    NOTATIONAL AMOUNT ON THE FOLLOWING DATES COLON dates DOT            { {
                                                                            payer = $1;
                                                                            dates = $20;
                                                                            fixedRate = $9;
                                                                            initialRate = 0.0;
                                                                            interestRateOption = ""
                                                                        } }
  | WORD WILL PAY A FLOATING RATE INTEREST COMMA INITIALLY DEFINED AS
    FLOAT PERCENT COMMA OVER THE NOTATIONAL AMOUNT ON THE FOLLOWING
    DATES COLON dates DOT THE FLOATING RATE OPTION IS WORD DOT          { {
                                                                            payer = $1;
                                                                            dates = $24;
                                                                            fixedRate = 0.0;
                                                                            initialRate = $12;
                                                                            interestRateOption = $31
                                                                        } }
;