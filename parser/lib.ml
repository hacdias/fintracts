open Core
open Types

module JSON = Yojson.Safe

exception ImcompleteContract
exception ValidationError of string

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

let ensure_same_length l1 l2 =
  let ll1 = List.length l1 in
  let ll2 = List.length l2 in
  if ll1 <> ll2 then raise (ValidationError "lists must be the same length")

let rec ids_from_parties_rec parties ids =
  match parties with
    | [] -> ids
    | p::tail -> ids_from_parties_rec tail (p.identifier :: ids)

let ids_from_parties parties = ids_from_parties_rec parties []

let rec ensure_lists_equal l1 l2 =
  match (l1, l2) with
    | [], [] -> ()
    | [], _ -> raise (ValidationError "lists must be the same length")
    | _, [] -> raise (ValidationError "lists must be the same length")
    | (h::t), (hh::tt) -> if (String.compare h hh) = 0
      then ensure_lists_equal t tt
      else raise (ValidationError "lists have different elements")

let validate contract =
  ensure_same_length contract.parties contract.signature.parties;
  let ids1 = ids_from_parties contract.parties in
    let sids1 = List.sort ~compare:String.compare ids1 in
      let sids2 = List.sort ~compare:String.compare contract.signature.parties in
        ensure_lists_equal sids1 sids2;
  (* TODO: validate agreements *)
  ();
