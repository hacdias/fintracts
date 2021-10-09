# JSON Contract Specification

This is the JSON specification of the contracts format. It goes well along with the [English Contract Specification](./ENGLISH-SPECIFICATION.md), as the latter explains the meaning of each field.

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

```
{
  "name": "The Party Name",
  "identifier": "TPN"
}
```

### Money

```
{
  "currency": "SYMBOL",
  "amount": 1234.5
}
```

### Coupons

```
{
  "rate": 1.5,
  "dates": [ Date... ]
}
```

### Exchange Rate

```
{
  "baseCurrency": "EUR",
  "counterCurrency": "USD",
  "rate": 1.45
}
```

### Interest Payment

For floating rate interest:

```
{
  "payer": "ID1",
  "dates": [ Date... ],
  "initialRate": 1.2,
  "rateOption": "usd-libor"
}
```

For fixed rate interest:

```
{
  "payer": "ID1",
  "dates": [ Date... ],
  "fixedRate": 1.2
}
```

### Agreement

An agreement has a key indicating the type of the agreement and its object as value. Only one type of agreement is permitted per agreement object. Other keys can be omitted or have `null` values.

```
{
  "bondPurchase": Bond Purchase | null,
  "currencySwap": Currency Swap | null,
  "interestRateSwap": Interest Rate Swap | null
}
```

### Contract

```
{
  "parties": [ Party... ],
  "agreements": [ Agreement... ],
  "signedOn": Date
}
```

## Agreements

### Bond Purchase

```
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

```
{
  "notationalAmount": Money,
  "effectiveDate": Date,
  "maturityDate": Date,
  "interest": [ Interest Payment... ]
}
```

### Currency Swap

```
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
