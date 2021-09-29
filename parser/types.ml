type party = {
  name: string;
  identifier: string
} [@@deriving yojson]

type contract = {
  parties: party list
} [@@deriving yojson]
