# JSON Contract Specification

This is the specification of the common JSON contract format. All the tools that use the JSON format must adhere to the rules specified on this document.

- [Contract](#contract)
- [Primitives](#primitives)
  - [Date](#date)
  - [Party](#party)
  - [Signature](#signature)
  - [Money](#money)
  - [Exchange Rate](#exchange-rate)
- [Agreements](#agreements)
  - [Bond Purchase](#bond-purchase)
  - [Interest Rate Swap](#interest-rate-swap)
  - [Currency Swap](#currency-swap)
  - [Other Types to Consider](#other-types-to-consider)

## Contract

The idea is that the contract is composed by the definition of parties, the agreements and the date in which it was signed on. This is the top-level field of the JSON format.

**Properties**

| Name         | Type                                | Description                                                             |
|--------------|-------------------------------------|-------------------------------------------------------------------------|
| `parties`    | Array of [`Party`](#party)          | The parties mentioned in the contract. Must have two or more elements.  |
| `agreements` | Array of [`Agreement`](#agreements) | The list of agreements in the contract. Must have one or more elements. |
| `signatures` | Array of [`Signature`](#signature)  | The signatures of the parties of the contract.                          |

**Format Example**

```
{
  "parties": [ Party... ],
  "agreements": [ Agreement... ],
  "signatures": [ Signature... ]
}
```

## Primitives

These are the primitives of the contracts. Primitives are all object types that are used in more than one different [agreement](#agreements).

### Date

A Date is a `string` in the [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339) specification. E.g.: `2025-10-01T00:00:00Z`.

### Party

**Properties**

| Name         | Type     | Description                                                                                      |
|--------------|----------|--------------------------------------------------------------------------------------------------|
| `name`       | `string` | The full name of the party.                                                                      |
| `identifier` | `string` | An alphanumeric identifier of the party that will be used everywhere else to identify the party. |

**Format Example**

```
{
  "name": "The Party Name",
  "identifier": "TPN"
}
```

### Signature

**Properties**

| Name      | Type              | Description                                                        |
|-----------|-------------------|--------------------------------------------------------------------|
| `date`    | [`Date`](#date)   | The date of the signature.                                         |
| `parties` | Array of `string` | An array with the identifiers of parties that signed on this date. |

**Format Example**

```
{
  "date": "2020-05-16T00:00:00Z",
  "parties": ["TPN", "WS"]
}
```

### Money

Represents a monetary amount of a certain currency.

**Properties**

| Name       | Type              | Description                                     |
|------------|-------------------|-------------------------------------------------|
| `currency` | `string`          | A string representing the currency. E.g. `EUR`. |
| `amount`   | `float`  \| `int` | The amount of `currency`.                       |

**Format Example**

```
{
  "currency": "USD",
  "amount": 1234.5
}
```

### Exchange Rate

Represents an [exchange rate](https://en.wikipedia.org/wiki/Exchange_rate) between two different currencies.

**Properties**

| Name              | Type     | Description                                     |
|-------------------|----------|-------------------------------------------------|
| `baseCurrency`    | `string` | A string representing the currency. E.g. `EUR`. |
| `counterCurrency` | `string` | A string representing the currency. E.g. `USD`. |
| `rate`            | `float`  | The exchange rate.                              |

**Format Example**

```
{
  "baseCurrency": "EUR",
  "counterCurrency": "USD",
  "rate": 1.45
}
```

## Agreements

An **agreement** has a key indicating the type of the agreement and its object as value. Only one type of agreement is permitted per agreement object. Other keys can be omitted or have `null` values.

**Properties**

| Name               | Type                                                  | Description                      |
|--------------------|-------------------------------------------------------|----------------------------------|
| `bondPurchase`     | [`Bond Purchase`](#bond-purchase) \| `null`           | A Bond Purchase agreement.       |
| `currencySwap`     | [`Currency Swap`](#currency-swap) \| `null`           | A Currency Swap agreement.       |
| `interestRateSwap` | [`Interest Rate Swap`](#interest-rate-swap) \| `null` | An Interest Rate Swap agreement. |

**Format Example**

```
{
  "bondPurchase": Bond Purchase | null,
  "currencySwap": Currency Swap | null,
  "interestRateSwap": Interest Rate Swap | null
}
```

### Bond Purchase

Represents a Bond Purchase Agreement. More information on the links below:

- [https://www.investopedia.com/terms/b/bond.asp](https://www.investopedia.com/terms/b/bond.asp)
- [https://www.investopedia.com/terms/b/bond-purchase-agreement.asp](https://www.investopedia.com/terms/b/bond-purchase-agreement.asp)

**Properties**

| Name           | Type                            | Description                                                                            |
|----------------|---------------------------------|----------------------------------------------------------------------------------------|
| `issuer`       | `string`                        | The identifier of the party that issues the bond.                                      |
| `underwriter`  | `string`                        | The identifier of the party that receives the bond.                                    |
| `faceValue`    | [`Money`](#money)               | The value of the bond.                                                                 |
| `issuePrice`   | [`Money`](#money)               | The price at which the bond issuer is selling the bond.                                |
| `maturityDate` | [`Date`](#date)                 | The date at which the bond issuer will pay the bond holder the face value of the bond. |
| `coupons`      | [`Coupons`](#coupons) \| `null` | The coupons of the bond, if there are any.                                             |

**Format Example**

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

#### Coupons

The bond may pay coupons periodically.

**Properties**

| Name    | Type                     | Description                                                                      |
|---------|--------------------------|----------------------------------------------------------------------------------|
| `rate`  | `float`                  | The interest rate the bond issuer pays on the face value of the bond on `dates`. |
| `dates` | Array of [`Date`](#date) | The dates at which the bond issuer pays the coupon interest rates.               |

**Format Example**

```
{
  "rate": 1.5,
  "dates": [ Date... ]
}
```

### Interest Rate Swap

Represents an Interest Rate Swap agreement. More information on the links below:

- [https://www.investopedia.com/terms/i/interestrateswap.asp](https://www.investopedia.com/terms/i/interestrateswap.asp)

**Properties**

| Name               | Type                                             | Description                                                                                                      |
|--------------------|--------------------------------------------------|------------------------------------------------------------------------------------------------------------------|
| `notationalAmount` | [`Money`](#money)                                | The principal over which the interest amount is calculated. The principal is never exchanged, only the interest. |
| `effectiveDate`    | [`Date`](#date)                                  | The date the agreement starts.                                                                                   |
| `maturityDate`     | [`Date`](#date)                                  | The date the agreement ends.                                                                                     |
| `interest`         | Array of [`Interest Payment`](#interest-payment) | The details of the interest agreement payouts.                                                                   |

**Format Example**

```
{
  "notationalAmount": Money,
  "effectiveDate": Date,
  "maturityDate": Date,
  "interest": [ Interest Payment... ]
}
```

#### Interest Payment

Represents an interest payment. Can be either a floating rate interest, or a fixed rate interest.

**Common Properties**

| Name       | Type                     | Description                                                                        |
|------------|--------------------------|------------------------------------------------------------------------------------|
| `payer`    | `string`                 | The identifier of the party that will pay the interest rate. E.g. `TPN`.           |
| `receiver` | `string`                 | The identifier of the party that will receive the interest payments. E.g. `TPN`.   |
| `dates`    | Array of [`Date`](#date) | The dates in which the payer pays out the interest rate over the principal amount. |


**Floating Rate Properties**

| Name          | Type     | Description                                                                   |
|---------------|----------|-------------------------------------------------------------------------------|
| `initialRate` | `float`  | The interest rate used during the initial calculation period.                 |
| `rateOption`  | `string` | The derivative that tracks the exchange rate over time. E.g. `usd-libor-bba`. |


**Fixed Rate Properties**

| Name        | Type    | Description            |
|-------------|---------|------------------------|
| `fixedRate` | `float` | The agreed fixed rate. |


**Format Example**

For floating rate interest:

```
{
  "payer": "ID1",
  "receiver": "ID2",
  "dates": [ Date... ],
  "initialRate": 1.2,
  "rateOption": "usd-libor"
}
```

For fixed rate interest:

```
{
  "payer": "ID1",
  "receiver": "ID2",
  "dates": [ Date... ],
  "fixedRate": 1.2
}
```

### Currency Swap

Represents a Currency Swap agreement. More information on the links below:

- [https://www.investopedia.com/terms/c/currencyswap.asp](https://www.investopedia.com/terms/c/currencyswap.asp)
- [https://www.investopedia.com/ask/answers/051215/what-difference-between-currency-and-interest-rate-swap.asp](https://www.investopedia.com/ask/answers/051215/what-difference-between-currency-and-interest-rate-swap.asp) (Currency Swap vs. Interest Rate Swap)

**Properties**

| Name              | Type                                                       | Description                                                                                                          |
|-------------------|------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------|
| `payerA`          | `string`                                                   | The identifier of the party A.                                                                                       |
| `principalA`      | [`Money`](#money)                                          | The principal provided by `payerA`.                                                                                  |
| `payerB`          | `string`                                                   | The identifier of the party B.                                                                                       |
| `principalB`      | [`Money`](#money)                                          | The principal provided by `payerB`.                                                                                  |
| `endExchangeRate` | [`Exchange Rate`](#exchange-rate) \| `null`                | The pre-agreed rate to swap at maturity date. If non-specified, it is the implied exchange rate of the initial swap. |
| `maturityDate`    | [`Date`](#date)                                            | The date at which both parties have to swap the principals again at the `endExchangeRate`.                           |
| `effectiveDate`   | [`Date`](#date)                                            | The date at which both parties swap the principals.                                                                  |

**Format Example**

```
{
  "payerA": "IDA",
  "principalA": Money,
  "payerB": "IDB",
  "principalB": Money,
  "endExchangeRate": ExchangeRate | null,
  "effectiveDate": Date,
  "maturityDate": Date
}
```

### Other Types to Consider

- Spot Foreign Exchange Transaction Agreement
- Foreign Currency Exchange Agreement
- Equity (Stocks) Purchase Agreement
- Certificate of Deposit Purchase Agreement
