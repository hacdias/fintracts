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
        "agree", AGREE;
        "agrees", AGREES;
        "amount", AMOUNT;
        "aforementioned", AFOREMENTIONED;
        "by", BY;
        "bond", BOND;
        "coupons", COUPONS;
        "defined", DEFINED;
        "date", DATE;
        "dates", DATES;
        "enter", ENTER;
        "effective", EFFECTIVE;
        "for", FOR;
        "follows", FOLLOWS;
        "following", FOLLOWING;
        "fixed", FIXED;
        "floating", FLOATING;
        "hereby", HEREBY;
        "has", HAS;
        "in", IN;
        "is", IS;
        "issuing", ISSUING;
        "interest", INTEREST;
        "initially", INITIALLY;
        "maturity", MATURITY;
        "notational", NOTATIONAL;
        "of", OF;
        "on", ON;
        "over", OVER;
        "option", OPTION;
        "parties", PARTIES;
        "paid", PAID;
        "principal", PRINCIPAL;
        "pay", PAY;
        "rd", DATE_SEP;
        "rate", RATE;
        "reaches", REACHES;
        "signed", SIGNED;
        "selling", SELLING;
        "swap", SWAP;
        "st", DATE_SEP;
        "th", DATE_SEP;
        "the", THE;
        "to", TO;
        "transaction", TRANSACTION;
        "termination", TERMINATION;
        "undermentioned", UNDERMENTIONED;
        "with", WITH;
        "will", WILL ]
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
  | word as w                                         { try
                                                          Hashtbl.find keyword_table (String.lowercase_ascii w)
                                                        with Not_found ->
                                                          WORD(w) }
  | ','                                               { COMMA }
  | '.'                                               { DOT }
  | ';'                                               { SEMICOLON }
  | ':'                                               { COLON }
  | '%'                                               { PERCENT }
  | money as m                                        { MONEY(float_of_money m) }
  | int as i                                          { INT(int_of_string i) }
  | float as f                                        { FLOAT(float_of_string f) }
  | eof                                               { raise Eof }
