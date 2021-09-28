(* type party = {
  name: string;
  identifier: string
} [@@deriving yojson] *)

type contract = {
  parties: string
} [@@deriving yojson]
