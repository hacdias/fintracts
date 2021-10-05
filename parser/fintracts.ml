open Types

module JSON = Yojson.Safe

let main =
  let got_contract = ref false in
  try
    let lexbuf = Lexing.from_channel stdin in
      let result = Parser.main Lexer.token lexbuf in
        let yo = contract_to_yojson result in
          got_contract := true;
          print_endline (JSON.pretty_to_string yo);
  with Lexer.Eof ->
    if got_contract.contents == true then
      exit 0
    else
      print_endline "Fatal: contract unterminated.";
      exit 1
