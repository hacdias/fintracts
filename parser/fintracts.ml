open Types
open Core

exception ImcompleteContract
exception InvalidNumberInputFiles

module JSON = Yojson.Safe

let parse_contract lexbuf =
  try
    Parser.main Lexer.token lexbuf
  with Lexer.Eof ->
    raise ImcompleteContract

let parse_from_channel ch =
  let lexbuf = Lexing.from_channel ch in
    parse_contract lexbuf

let parse_from_file path =
  let file = In_channel.read_all path in
    let lexbuf = Lexing.from_string file in
      parse_contract lexbuf

let pretty_contract contract =
  let json = contract_to_yojson contract in
    JSON.pretty_to_string json

let usage_msg = "fintracts [-json] <input file>"

let print_json = ref false
let input_file = ref ""

let arg_parser filename =
  match input_file.contents with
    | "" -> input_file := filename
    | _ -> raise InvalidNumberInputFiles

let speclist =
  [("-json", Arg.Set print_json, "Print parsed JSON.")]

let () =
  Arg.parse speclist arg_parser usage_msg;
  let contract = match input_file.contents with
    | "" -> parse_from_channel In_channel.stdin
    | filename -> parse_from_file filename in
    if print_json.contents then
      print_endline (pretty_contract contract)
