open Lib
open Core

exception InvalidNumberInputFiles

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
    validate contract;
    if print_json.contents then
      print_endline (pretty_contract contract)
