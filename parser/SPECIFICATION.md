# English Contract Specification

This specification implies knowledge from the [JSON contract specification](../../SPECIFICATION.md) and the available agreement types.

- [Format](#format)
- [Parties](#parties)
- [Signature](#signature)
- [Agreement(s)](#agreements)
  - [Bond Purchase](#bond-purchase)
  - [Interest Rate Swap Transaction](#interest-rate-swap-transaction)
  - [Currency Swap Transaction](#currency-swap-transaction)

## Format

The idea is to be able to write a contract as naturally as possible. However, we have to take into account the limitations of text processing and parsing. Since we are not working with natural language processing, the best idea is to specify a format that, albeit natural looking, can be easily parsed and is rigid.

All contracts must be defined as:

```
<Parties>
<Agreement>+
<Signature>+
```

## Parties

Each party is composed by its name and abbreviation. There can be 2 or more parties. The abbreviation is used as a reference to the party in the rest of the contract.

```
The parties:

  [<Name>, undermentioned as <Abbreviation>; and]+
  <Name>, undermentioned as <Abbreviation>.
```

## Signature

The contract must be signed by all parties and contain the date.

```
Signed by <Party Abbreviation> [, <Party Abbreviation>]+ and <Party Abbreviation> on <Date>.
```

## Agreement(s)

A contract may have one or more agreements. An agreement is just a transaction type. All agreements are specified as below:

```
Hereby enter in (a|an) <Agreement Type> defined as follows:

  <Agreement Text>
```

### Bond Purchase

#### Format

```
<Abbreviation> agrees on issuing and selling a bond of <Money> to
<Abbreviation> for <Money>. The aforementioned bond reaches maturity
on the <Date>.

[The bond pays coupons with an interest rate of <Float>% paid on the following
dates: <Date> [, <Date>]* and <Date>.]
```

#### Examples

Selling a discounted bond without coupons:

```
The parties:

  Will Smith, undermentioned as WS; and
  The Bank, undermentioned as TB.

Hereby enter in a Bond Purchase Agreement defined as follows:

  TB agrees on issuing and selling a bond of EUR 10,000.00 to WS
  for EUR 9,800.00. The aforementioned bond reaches maturity on the
  1st of October 2025.

Signed by WS and TB on the 24th of September 2021.
```

Selling a bond for its face value with coupons:

```
The parties:

  Will Smith, undermentioned as WS; and
  The Bank, undermentioned as TB.

Hereby enter in a Bond Purchase Agreement defined as follows:

  TB agrees on issuing and selling a bond of EUR 10,000.00 to WS
  for EUR 9,800.00. The aforementioned bond reaches maturity on the
  1st of October 2025.

  The bond pays coupons with an interest rate of 1.2% paid on the following
  dates: 1st of October 2021, 1st of October 2022, 1st of October 2023, 1st
  of October 2024 and 1st of October 2025.

Signed by WS and TB on the 24th of September 2021.
```

### Interest Rate Swap Transaction

#### Format

```
The parties agree on an interest rate swap transaction over the notational
principal of <Money>, with an effective date as of the <Date> and termination
on the <Date>.

[<Abbreviation> will pay <Abbreviation> a fixed rate interest of <Float>% over the notational amount
on the following dates: <Date> [, <Date>]* and <Date>.]+

[<Abbreviation> will pay <Abbreviation> a floating rate interest, initially defined as <Float>%, over
the notational amount on the following dates: <Date> [, <Date>]* and <Date>. The floating
rate option is <Option>.]+
```

#### Examples

Fixed to Floating Example

```
The parties:

  The Bank 1, undermentioned as TB1; and
  The Bank 2, undermentioned as TB2.

Hereby enter in an Interest Rate Swap Transaction Agreement defined as follows:

  The parties agree on an interest rate swap transaction over the notational
  principal of USD 10,000,000.00, with an effective date as of the 1st of
  October 2021 and termination on the 15th of October 2025.

  TB1 will pay TB2 a fixed rate interest of 3.5% over the notational amount
  on the following dates: 15th of October 2021, 15th of October 2022, 15th of
  October 2023, 15th of October 2024 and 15th of October 2025.

  TB2 will pay TB1 a floating rate interest, initially defined as 2.4%, over
  the notational amount on the following dates: 23rd of September 2021, 23rd
  of September 2022, 23rd of September 2023, 23rd of September 2024 and 23rd of
  September 2025. The floating rate option is USD-LIBOR.

Signed by TB1 and TB2 on the 15th of September 2021.
```

### Currency Swap Transaction

#### Format

```
The parties agree on a currency swap transaction effective as of the
<Date> and termination on the <Date>.

<Abbreviation> will pay a principal amount of <Money>, and the
<Abbreviation> will pay a principal amount of <Money>. [At maturity, the
principal amounts shall be exchanged back with an interest rate of <Rate>.]

[<Abbreviation> will pay <Abbreviation> a fixed rate interest of <Float>% over the notational amount
on the following dates: <Date> [, <Date>]* and <Date>.]*

[<Abbreviation> will pay <Abbreviation> a floating rate interest, initially defined as <Float>%, over
the notational amount on the following dates: <Date> [, <Date>]* and <Date>. The floating
rate option is <Option>.]*
```

#### Examples

An example where the implied exchange rate USD/EUR is 1.25.

```
The parties:

  The Bank 1, undermentioned as TB1; and
  The Bank 2, undermentioned as TB2.

Hereby enter in a Currency Swap Transaction Agreement defined as follows:

  The parties agree on a currency swap transaction effective as of the
  1st of October 2021 and termination on the 1st of October 2025.

  TB1 will pay TB2 a principal amount of USD 10,000,000.00, and the
  TB2 will pay TB1 a principal amount of EUR 12,500,000.00.

Signed by TB1 and TB2 on the 15th of September 2021.
```

With interest over the loans **and** an end exchange rate of USD/EUR 1.30:

```
The parties:

  The Bank 1, undermentioned as TB1; and
  The Bank 2, undermentioned as TB2.

Hereby enter in a Currency Swap Transaction Agreement defined as follows:

  The parties agree on a currency swap transaction effective as of the
  1st of October 2021 and termination on the 1st of October 2025.

  TB1 will pay a principal amount of USD 10,000,000.00, and the
  TB2 will pay a principal amount of EUR 12,500,000.00. At maturity,
  the principal amounts shall be exchanged back with an interest
  rate of USD/EUR 1.45.

  TB2 will pay TB1 a fixed rate interest of 3.5% over the notational amount
  on the following dates: 15th of October 2021, 15th of October 2022, 15th of
  October 2023, 15th of October 2024 and 15th of October 2025.

  TB1 will pay TB2 a floating rate interest, initially defined as 2.4%, over
  the notational amount on the following dates: 15th of October 2021, 15th
  of October 2022, 15th of October 2023, 15th of October 2024 and 15th of
  October 2025. The floating rate option is USD-LIBOR.

Signed by TB1 and TB2 on the 15th of September 2021.
```
