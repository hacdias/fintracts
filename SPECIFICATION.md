# JSON Specification

This is the JSON specification of the contracts format. It goes well along with the [English Contract Specification](./parser/SPECIFICATION.md), as the latter explains the meaning of each field.

- [Primitives](#primitives)
  - [Date](#date)
  - [Party](#party)
  - [Money](#money)
  - [Coupons](#coupons)
  - [Exchange Rate](#exchange-rate)
  - [Interest Payment](#interest-payment)
  - [Agreement](#agreement)
  - [Contract](#contract)
- [Agreements](#agreements)
  - [Bond Purchase](#bond-purchase)
  - [Interest Rate Swap](#interest-rate-swap)
  - [Currency Swap](#currency-swap)

## Primitives

### Date

A Date is represented by a string on the [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339) format.

### Party

```json
{
  "name": "The Party Name",
  "identifier": "TPN"
}
```

### Money

```json
{
  "currency": "SYMBOL",
  "amount": 1234.5
}
```

### Coupons

```json
{
  "rate": 1.5,
  "dates": [ Date... ]
}
```

### Exchange Rate

```json
{
  "baseCurrency": "EUR",
  "counterCurrency": "USD",
  "rate": 1.45
}
```

### Interest Payment

For floating rate interest:

```json
{
  "payer": "ID1",
  "dates": [ Date... ],
  "initialRate": 1.2,
  "rateOption": "usd-libor"
}
```

For fixed rate interest:

```json
{
  "payer": "ID1",
  "dates": [ Date... ],
  "fixedRate": 1.2
}
```

### Agreement

An agreement has a key indicating the type of the agreement and its object as value. Only one type of agreement is permitted per agreement object. Other keys can be omitted or have `null` values.

```json
{
  "bondPurchase": Bond Purchase | null,
  "currencySwap": Currency Swap | null,
  "interestRateSwap": Interest Rate Swap | null
}
```

### Contract

```json
{
  "parties": [ Party... ],
  "agreements": [ Agreement... ],
  "signedOn": Date
}
```

## Agreements

### Bond Purchase

```json
{
  "issuer": "ID1",
  "underwriter": "ID2",
  "faceValue": Money,
  "issuePrice": Money,
  "maturityDate": Date,
  "Coupons": Coupons | null
}
```
### Interest Rate Swap

```json
{
  "notationalAmount": Money,
  "effectiveDate": Date,
  "maturityDate": Date,
  "interest": [ Interest Payment... ]
}
```

### Currency Swap

```json
{
  "payerA": "IDA",
  "principalA": Money,
  "payerB": "IDB",
  "principalB": Money,
  "impliedExchangeRate": ExchangeRate,
  "endExchangeRate": ExchangeRate | null,
  "interest": [ Interest Payment... ] | null,
  "effectiveDate": Date,
  "maturityDate": Date
}
```
