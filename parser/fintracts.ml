open Types

module JSON = Yojson.Safe

let main =
  try
    let lexbuf = Lexing.from_channel stdin in
      let result = Parser.main Lexer.token lexbuf in
        let yo = contract_to_yojson result in
          print_endline (JSON.to_string yo);
  with Lexer.Eof ->
    exit 0
