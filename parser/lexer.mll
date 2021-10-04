{
  open Parser
  exception Eof

  let keyword_table = Hashtbl.create 53
  let _ =
    List.iter (fun (kwd, tok) -> Hashtbl.add keyword_table kwd tok)
      [ "and", AND;
        "of", OF;
        "st", DATE_SEP; 
        "th", DATE_SEP; 
        "rd", DATE_SEP ]
}

let ws = [' ' '\t' '\n' '\r']*
let word = ['a'-'z''A'-'Z']+
let int = ['0'-'9']+
let float = (['0'-'9']+['.'])['0'-'9']+

rule token = parse
  | ws                          { token lexbuf }
  | "The" ws "parties" ws ":"   { THE_PARTIES }
  | "Signed" ws "by"            { SIGNED_BY }
  | "on" ws "the"               { ON_THE }
  | "undermentioned" ws "as"    { UNDERMENTIONED }
  | ','                         { COMMA }
  | ';'                         { SEMICOLON }
  | '.'                         { DOT }
  | '%'                         { PERCENT }
  | word as w                   { try
                                    Hashtbl.find keyword_table w
                                  with Not_found ->
                                    WORD(w) }
  | int as i                    { INT(int_of_string i) }
  | float as f                  { FLOAT(float_of_string f) }
  | eof                         { raise Eof }
