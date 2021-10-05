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
        "at", AT;
        "agree", AGREE;
        "agrees", AGREES;
        "agreement", AGREEMENT;
        "amount", AMOUNT;
        "amounts", AMOUNTS;
        "aforementioned", AFOREMENTIONED;
        "by", BY;
        "be", BE;
        "bond", BOND;
        "back", BACK;
        "coupons", COUPONS;
        "currency", CURRENCY;
        "defined", DEFINED;
        "date", DATE;
        "dates", DATES;
        "enter", ENTER;
        "effective", EFFECTIVE;
        "exchanged", EXCHANGED;
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
        "purchase", PURCHASE;
        "principal", PRINCIPAL;
        "pay", PAY;
        "rd", DATE_SEP;
        "rate", RATE;
        "reaches", REACHES;
        "signed", SIGNED;
        "selling", SELLING;
        "swap", SWAP;
        "shall", SHALL;
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
  | word as w                                         { try
                                                          Hashtbl.find keyword_table (String.lowercase_ascii w)
                                                        with Not_found ->
                                                          WORD(w) }
  | int as i                                          { INT(int_of_string i) }
  | float as f                                        { FLOAT(float_of_string f) }
  | money as m                                        { MONEY(float_of_money m) }
  | ','                                               { COMMA }
  | '.'                                               { DOT }
  | ';'                                               { SEMICOLON }
  | ':'                                               { COLON }
  | '%'                                               { PERCENT }
  | '/'                                               { SLASH }
  | eof                                               { raise Eof }
