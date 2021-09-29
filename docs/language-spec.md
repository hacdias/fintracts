# Language Specification

- [Contract Format](#contract-format)
- [Contract Types](#contract-types)
  - [Bond Purchase Agreement](#bond-purchase-agreement)
  - [Interest Rate Swap Transaction Agreement](#interest-rate-swap-transaction-agreement)
  - [Currency Swap Transaction Agreement](#currency-swap-transaction-agreement)
  - [Other Agreements to Consider](#other-agreements-to-consider)
- [BNF Definition](#bnf-definition)

## Contract Format

The idea is to write a contract as naturally as possible. However, we have to take into account the limitations of text processing and parsing. For this project, we suggest a language that, even though looks natural and can be read, has a rigid grammar and structure.

All contracts must define:

- The *parties* involved.
- The *transaction* and the *assets* involved.
- The *signature* and date.

Suggestion:

1. Starting by defining the parties, where `Abbreviation` is the short name of the party that will be used on the rest of the contract. For example, the "Nasdaq Stock Market" could use NSM as abbreviation.
2. The contract type and agreement details.
3. The signature and date.

```
The parties:
	
	<Party Name>, undermentioned as <Abbreviation>; and
	<Party Name>, undermentioned as <Abbreviation>.

Hereby enter in (a|an) <Contract Type> defined as follows:

	<Contract Type Dependent Text>

Signed by <Party Abbreviation> and <Party Abbreviation> on <Date>.
```

## Contract Types

Contract types are just transaction types. Each contract refers to a different transaction. And each type of transaction has its own intrinsic properties.

### Bond Purchase Agreement

#### Information

- [https://www.investopedia.com/terms/b/bond.asp](https://www.investopedia.com/terms/b/bond.asp)
- [https://www.investopedia.com/terms/b/bond-purchase-agreement.asp](https://www.investopedia.com/terms/b/bond-purchase-agreement.asp)

#### Properties

- *Issuer*: the one who issues the bond.
- *Underwriter:* the one who purchases the bond.
- *Face value*: the value of the bond.
- *Issue price*: the price at which the bond issuer originally sells the bond.
- *Maturity date*, the date the bond issuer will pay the bond holder the face value of the bond.
- If there are coupons (can be defined in the agreement section).
    - *Coupon interest rate*: the interest rate the bond issuer pays on the face value of the bond at the coupon dates.
    - *Coupon dates*: the dates at which the bond issuer pays the coupon interest rate.
- *Withdraw conditions*: conditions in which the underwriter may withdraw. Should probably be free text conditions (not on examples yet).

#### Examples

Selling a discounted bond without coupons:

```
The parties:

	Will Smith, undermentioned as WS; and
	De National Bank, undermentioned as DNB.

Hereby enter in a Bond Purchase Agreement defined as follows:

	The DNB agrees on issuing and selling a bond of EUR 10,000.00 to WS
	for EUR 9,800.00. The aforementioned bond reaches maturity on the
	1st of October 2025.

Signed by WS and DNB on the 24th of September 2021.
```

Selling a bond for its face value with coupons:

```
The parties:

	Will Smith, undermentioned as WS; and
	De national Bank, undermentioned as DNB.

Hereby enter in a Bond Purchase Agreement defined as follows:

	The DNB agrees on issuing and selling a bond of EUR 10,000.00 to WS
	for EUR 9,800.00. The aforementioned bond reaches maturity on the
	1st of October 2025.

	The bond has coupons with an interest rate of 1.2% paid on the following
	dates: 1st of October 2021, 1st of October 2022, 1st of October 2023, 1st
	of October 2024 and 1st of October 2025.

Signed by WS and DNB on the 24th of September 2021.
```

### Interest Rate Swap Transaction Agreement

#### Information

- [https://www.investopedia.com/terms/i/interestrateswap.asp](https://www.investopedia.com/terms/i/interestrateswap.asp)
- There are different types of swap: fixed to floating, floating to fixed, and floating to floating.

#### Properties

- *Notational* *Principal Amount:* the principal over which the interest amount is calculated. In these types of agreements, the principal is never exchanged, only the interest.
- *Effective Date*: the date the agreement starts.
- *Maturity Date*: the date of the end of the agreement.
- For **both** types of rates:
    - *Payer*: who will pay that interest rate.
    - *Payment Dates*: the dates in which the payer pays out the interest rate over the principal amount.
- For **fixed** rate **only**:
    - *Interest rate*: the agreed fixed interest rate.
- For **floating** rate **only**:
    - *Initial rate for calculation period:* the interest rate used during the initial calculations period.
    - *Interest rate option*: the derivative that tracks the exchange rate over time (e.g. `usd-libor-bba`).

#### Examples

Fixed to Floating Example

```
The parties:

	Big USA Bank, undermentioned as BUSAB; and
	Big German Bank, undermentioned as BGB.

Hereby enter in an Interest Rate Swap Transaction Agreement defined as follows:

	The parties agree on an interest rate swap transaction over the notational
	principal of USD 10,000,000.00, with an effective date as of the 1st of
	October 2021 and termination on the 1st of October 2025.

	The BUSAB will pay a fixed rate interest of 3.5% over the notational amount
	on the following dates: 15th of October 2021, 15th of October 2022, 15th of
	October 2023, 15th of October 2024 and 15th of October 2025.

	The BGB will pay a floating rate interest, initially defined as 2.4%, over
	the notational amount on the following dates: 15th of October 2021, 15th
	of October 2022, 15th of October 2023, 15th of October 2024 and 15th of
	October 2025. The floating rate option is USD-LIBOR.

Signed by BUSAB and BGB on the 15th of September 2021.
```

### Currency Swap Transaction Agreement

#### Information

- [https://www.investopedia.com/terms/c/currencyswap.asp](https://www.investopedia.com/terms/c/currencyswap.asp)
- [https://www.investopedia.com/ask/answers/051215/what-difference-between-currency-and-interest-rate-swap.asp](https://www.investopedia.com/ask/answers/051215/what-difference-between-currency-and-interest-rate-swap.asp) (Currency Swap vs. Interest Rate Swap)

#### Properties

- *Principal A:* principal provided by Payer A.
    - *Payer:* who provides the principal.
- *Principal B*: principal provided by Payer B.
    - *Payer*: who provides the principal.
- *Implied Exchange Rate*: implied exchange rate from *Principal A* and *Principal B*.
- *Maturity Date*: the date at which both parties have to swap again at either the original implied exchange rate or another pre-agreed rate.
- *(Optional) End Exchange Rate*: the pre-agreed rate to swap at maturity date. If non-specified, it is the implied exchange rate.
- These transactions can also **involve** interest payouts during the agreement time, just like the Interest Rate Swap Transaction Agreement.

#### Examples

An example where the implied exchange rate USD/EUR is 1.25.

```
The parties:

	Big USA Bank, undermentioned as BUSAB; and
	Big German Bank, undermentioned as BGB.

Hereby enter in a Currency Swap Transaction Agreement defined as follows:

	The parties agree on a currency swap transaction effective as of the
	1st of October 2021 and termination on the 1st of October 2025.

	The BUSAB will pay a principal amount of USD 10,000,000.00, and the
	BGB will pay a principal amount of EUR 12,500,000.00.

Signed by BGB and BUSAB on the 15th of September 2021.
```

With interest over the loans **and** an end exchange rate of USD/EUR 1.30:

```
The parties:

	Big USA Bank, undermentioned as BUSAB; and
	Big German Bank, undermentioned as BGB.

Hereby enter in a Currency Swap Transaction Agreement defined as follows:

	The parties agree on a currency swap transaction effective as of the
	1st of October 2021 and termination on the 1st of October 2025.

	The BUSAB will pay a principal amount of USD 10,000,000.00, and the
	BGB will pay a principal amount of EUR 12,500,000.00.
	
	The BUSAB will pay a fixed rate interest of 3.5% over the notational amount
	on the following dates: 15th of October 2021, 15th of October 2022, 15th of
	October 2023, 15th of October 2024 and 15th of October 2025.

	The BGB will pay a floating rate interest, initially defined as 2.4%, over
	the notational amount on the following dates: 15th of October 2021, 15th
	of October 2022, 15th of October 2023, 15th of October 2024 and 15th of
	October 2025. The floating rate option is USD-LIBOR.

	At maturity, the principal amounts shall be exchanged back with an interest
	rate of USD/EUR 1.30.

Signed by BGB and BUSAB on the 15th of September 2021.
```

### Other Agreements to Consider

- Spot Foreign Exchange Transaction Agreement
- Foreign Currency Exchange Agreement
- Equity (Stocks) Purchase Agreement
- Certificate of Deposit Purchase Agreement

## BNF Definition

The formal definition of the language. This will help with the implementation.

```
<Digit>  ::= 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
<Digits> ::= <Digit> | <Digits><Digit>
<Number> ::= <Digits> . <Digits>
<Letter> ::= A | B | C | D | E | F | G | H | I | J | K | L | M |
	N | O | P | Q | R | S | T | U | V | W | X | Y | Z

# Assets
<Currency> ::= <Letter><Letter><Letter>
<Money>    ::= <Currency> <Number>

<Stocks>   ::=

<CertificateOfDeposit> ::=

<Bond> ::=

<Asset> ::= <Money> | <Stocks> | <CertificateOfDeposit> | <Bond>
```
