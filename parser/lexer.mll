{
  open Parser
  exception Eof
}

let ws = [' ' '\t' '\n' '\r']*
let word = ['a'-'z''A'-'Z']+
let int = ['0'-'9']+

rule token = parse
  | ws                          { token lexbuf }
  | "The" ws "parties" ws ":"   { THE_PARTIES }
  | "Signed" ws "by"            { SIGNED_BY }
  | "on" ws "the"               { ON_THE }
  | "undermentioned" ws "as"    { UNDERMENTIONED }
  | "and"                       { AND }
  | "st"                        { DATE_SEP }
  | "th"                        { DATE_SEP }
  | "rd"                        { DATE_SEP }
  | "of"                        { OF }
  | ','                         { COMMA }
  | ';'                         { SEMICOLON }
  | '.'                         { DOT }
  | word as w                   { WORD(w) }
  | int as i                    { INT(int_of_string i) }
  | eof                         { raise Eof }
