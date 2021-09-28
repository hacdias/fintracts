{
  open Parser
  exception Eof
}

rule token = parse
    [' ' '\t']          { token lexbuf }     (* skip blanks *)
  | ['\n' ]             { EOL }
  | ['a'-'z''A'-'Z']+  as lxm  { IDENT(lxm) }
  | eof                 { raise Eof }
