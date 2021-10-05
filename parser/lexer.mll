{
  open Parser
  open Types
  exception Eof

  let keyword_table = Hashtbl.create 53
  let _ =
    List.iter (fun (kwd, tok) -> Hashtbl.add keyword_table kwd tok)
      [ "and", AND;
        "of", OF;
        "to", TO;
        "on", ON;
        "the", THE;
        "for", FOR;
        "st", DATE_SEP;
        "th", DATE_SEP;
        "rd", DATE_SEP ]
}

let ws = [' ' '\t' '\n' '\r']*
let word = ['a'-'z''A'-'Z']+
let int = ['0'-'9']+
let float = (['0'-'9']+['.'])['0'-'9']+
let money = ['0'-'9']?['0'-'9']?['0'-'9']?([',']['0'-'9']['0'-'9']['0'-'9'])*['.']['0'-'9']['0'-'9']

rule token = parse
  | ws                                                { token lexbuf }
  | "The" ws "parties" ws ":"                         { THE_PARTIES }
  | "Signed" ws "by"                                  { SIGNED_BY }
  | "undermentioned" ws "as"                          { UNDERMENTIONED_AS }
  | "Hereby" ws "enter" ws "in" ws ("a" | "an")       { HEREBY_ENTER }
  | "defined" ws "as" ws "follows" ws ":"             { DEFINED_AS }
  | "Bond" ws "Purchase" ws "Agreement"               { BOND_PURCHASE_AGREEMENT }
  | "agrees" ws "on" ws "issuing" ws "and"
    ws "selling" ws "a" ws "bond" ws "of"             { AGREE_BOND_OF }
  | "The" ws "aforementioned" ws "bond" ws
    "reaches" ws "maturity" ws "on" ws "the"          { MATURITY_ON }
  | ','                                               { COMMA }
  | ';'                                               { SEMICOLON }
  | '.'                                               { DOT }
  | '%'                                               { PERCENT }
  | word as w                                         { try
                                                          Hashtbl.find keyword_table w
                                                        with Not_found ->
                                                          WORD(w) }
  | money as m                                        { MONEY(float_of_money m) }
  | int as i                                          { INT(int_of_string i) }
  | float as f                                        { FLOAT(float_of_string f) }
  | eof                                               { raise Eof }
