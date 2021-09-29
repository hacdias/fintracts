{
  open Parser
  exception Eof
}

let ws = [' ' '\t' '\n' '\r']
let w = ['a'-'z''A'-'Z']+

rule token = parse
  | ws                          { token lexbuf }
  | "The" ws* "parties" ws* ":" { parties lexbuf }
  | w as w { print_endline w; token lexbuf }
  | eof                         { raise Eof }

and parties = parse
  | ws                                        { parties lexbuf }
  | "undermentioned" ws* "as" ws* (w as w) { PARTY_NICK(w) }
  | w                                         {
    let buf = Buffer.create 17 in
      Buffer.add_string buf (Lexing.lexeme lexbuf);
      name_until_comma (buf) lexbuf
  }
  | ';' ws* "and" ws*                         { parties lexbuf }
  | '.'                                       { token lexbuf }

and name_until_comma buf = parse 
  | ','                                 { PARTY_NAME (Buffer.contents buf) }
  | ws                                  {
    Buffer.add_string buf (Lexing.lexeme lexbuf);
      name_until_comma buf lexbuf
    }
  | w                            {
    Buffer.add_string buf (Lexing.lexeme lexbuf);
      name_until_comma buf lexbuf
    }
