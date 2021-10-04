type date = {
  day: int;
  month: string;
  year: int
} [@@deriving yojson]

type signature = {
  parties: string list;
  date: date
} [@@deriving yojson]

type party = {
  name: string;
  identifier: string
} [@@deriving yojson]

type contract = {
  parties: party list;
  signature: signature
} [@@deriving yojson]
