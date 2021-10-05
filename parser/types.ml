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

type interestPayments = {
  payer: string;
  dates: date list;
  fixedRate: float;
  initialRate: float;
  interestRateOption: string;
} [@@deriving yojson]

type interestRateSwap = {
  notationalAmount: money;
  effectiveDate: date;
  maturityDate: date;
  interest: interestPayments list
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
  endExhcangeRate: exchangeRate;
  interest: interestPayments list option
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

let float_of_money money =
  let rep = Str.global_replace (Str.regexp "\\,") "" (money) in
    float_of_string rep
