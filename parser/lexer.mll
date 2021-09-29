{
  (* open Parser *)
  exception Eof
}

let ws = [' ' '\t' '\n' '\r']
let w = ['a'-'z''A'-'Z']+

rule token = parse
  | ws                          { token lexbuf }
  | "The" ws* "parties" ws* ":" { print_endline "Enter parties"; parties lexbuf }
  | eof                         { raise Eof }

and parties = parse
  | ws                          { parties lexbuf }
  | (w as name) ws* ',' ws*
    "undermentioned" ws* "as" ws*
    (w as nick)                 { print_endline name; print_endline nick; parties lexbuf }
  | ';' ws* "and" ws*           { parties lexbuf }
  | '.'                         { print_endline "Exit parties"; token lexbuf }

and name_until_comma = parse 
  | char                        { back lexbuf }
  | w                           { print_endline w; parties lexbuf }
