{
  open Parser
  exception Eof
}

let ws = [' ' '\t' '\n' '\r']*
let word = ['a'-'z''A'-'Z']+
let int = ['0'-'9']+

rule token = parse
  | ws                          { token lexbuf }
  | "Signed" ws "by"            { SIGNED_BY }
  | "on" ws "the"               { ON_THE }
  | "and"                       { AND }
  | "st"                        { DATE_SEP }
  | "th"                        { DATE_SEP }
  | "rd"                        { DATE_SEP }
  | "of"                        { OF }
  | ','                         { COMMA }
  | '.'                         { DOT }
  | word as w                   { WORD(w) }
  | int as i                    { INT(int_of_string i) }
  | eof                         { raise Eof }
