{
  open Parser
  open Types
  exception Eof

  let keyword_table = Hashtbl.create 53
  let _ =
    List.iter (fun (kwd, tok) -> Hashtbl.add keyword_table kwd tok)
      [ "a", A;
        "an", AN;
        "and", AND;
        "as", AS;
        "agrees", AGREES;
        "aforementioned", AFOREMENTIONED;
        "by", BY;
        "bond", BOND;
        "coupons", COUPONS;
        "defined", DEFINED;
        "dates", DATES;
        "enter", ENTER;
        "for", FOR;
        "follows", FOLLOWS;
        "following", FOLLOWING;
        "hereby", HEREBY;
        "has", HAS;
        "in", IN;
        "issuing", ISSUING;
        "interest", INTEREST;
        "maturity", MATURITY;
        "of", OF;
        "on", ON;
        "parties", PARTIES;
        "paid", PAID;
        "rd", DATE_SEP;
        "rate", RATE;
        "reaches", REACHES;
        "signed", SIGNED;
        "selling", SELLING;
        "st", DATE_SEP;
        "th", DATE_SEP;
        "the", THE;
        "to", TO;
        "undermentioned", UNDERMENTIONED;
        "with", WITH ]

  let punctuation_table = Hashtbl.create 53
  let _ =
    List.iter (fun (kwd, tok) -> Hashtbl.add punctuation_table kwd tok)
      [ ',', COMMA;
        '.', DOT;
        ':', COLON;
        ';', SEMICOLON;
        '%', PERCENT ]
}

let ws = [' ' '\t' '\n' '\r']*
let word = ['a'-'z''A'-'Z']+
let int = ['0'-'9']+
let float = (['0'-'9']+['.'])['0'-'9']+
let money = ['0'-'9']?['0'-'9']?['0'-'9']?([',']['0'-'9']['0'-'9']['0'-'9'])*['.']['0'-'9']['0'-'9']
let punctuation = [',''.'';'':''%']

rule token = parse
  | ws                                                { token lexbuf }
  | "Bond" ws "Purchase" ws "Agreement"               { BOND_PURCHASE_AGREEMENT }
  | "Interest" ws "Rate" ws "Swap" ws "Transaction"
    ws "Agreement"                                    { INTEREST_RATE_SWAP_AGREEMENT }
  | "The" ws "parties" ws "agree" ws "on" ws "an"
    ws "interest" ws "rate" ws "swap" ws
    "transaction" ws "over" ws "the" ws "notational"
    ws "principal" ws "of"                            { AGREE_INTEREST_RATE_SWAP_OVER }
  | "," ws "with" ws "an" ws "effective" ws "date"
    ws "as" ws "of" ws "the"                          { WITH_EFFECTIVE_DATE }
  | "and" ws "termination"  ws "on" ws "the"          { AND_TERMINATION }
  | "will" ws "pay" ws "a"                            { WILL_PAY }
  | "fixed" ws "rate" ws "interest"                   { FIXED_INTEREST }
  | "floating" ws "rate" ws "interest"                { FLOATING_INTEREST }
  | "," ws "initially" ws "defined" ws  "as"          { INITIALLY_DEFINED }
  | "over" ws "the" ws "notational" ws "amount" ws
    "on" ws "the" ws "following" ws "dates" ws ":"    { OVER_ON_DATES }
  | "The" ws "floating" ws "rate" ws "option" ws "is" { FLOATING_OPTION_IS }
  | word as w                                         { try
                                                          Hashtbl.find keyword_table (String.lowercase_ascii w)
                                                        with Not_found ->
                                                          WORD(w) }
  | punctuation as p                                  { try
                                                          Hashtbl.find punctuation_table p
                                                        with Not_found ->
                                                          PUNCTUATION(p) }
  | money as m                                        { MONEY(float_of_money m) }
  | int as i                                          { INT(int_of_string i) }
  | float as f                                        { FLOAT(float_of_string f) }
  | eof                                               { raise Eof }
