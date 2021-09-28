open Types

module JSON = Yojson.Safe

let main =
  try
    let lexbuf = Lexing.from_channel stdin in
    while true do
      let result = Parser.main Lexer.token lexbuf in
        let yo = contract_to_yojson result in
          print_endline (JSON.to_string yo);
    done
  with Lexer.Eof ->
    exit 0

(* let main =
  let c = {
    parties = [{ name = "Bank"; identifier = "B" }]
  } in 
    let yo = contract_to_yojson c in
      print_endline (JSON.to_string yo); *)