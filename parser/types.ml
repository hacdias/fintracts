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

type money = {
  currency: string;
  amount: float
} [@@deriving yojson]

type coupons = {
  rate: float;
  dates: date list;
} [@@deriving yojson]

type bondPurchase = {
  seller: string;
  payer: string;
  issuePrice: money;
  faceValue: money;
  maturityDate: date;
  coupons: coupons option;
} [@@deriving yojson]

type interestRateSwap = {
  notationalAmount: money;
  effectiveDate: date;
  maturityDate: date
  (* TODO: add interest *)
} [@@deriving yojson]

type exchangeRate = {
  currencyTop: string;
  currencyBottom: string;
  rate: float
} [@@deriving yojson]

type currencySwap = {
  principalA: money;
  principalB: money;
  payerA: string;
  payerB: string;
  maturityDate: date;
  impliedExchangeRate: exchangeRate;
  endExhcangeRate: exchangeRate
  (* TODO: add optional interest *)
} [@@deriving yojson]

type agreement = {
  bondPurchase: bondPurchase option;
  interestRateSwap: interestRateSwap option;
  currencySwap: currencySwap option
} [@@deriving yojson]

type contract = {
  parties: party list;
  agreements: agreement list;
  signature: signature
} [@@deriving yojson]
