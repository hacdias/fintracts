# Verification Financial Smart Contracts

- [Problem Statement](#problem-statement)
- [Motivation](#motivation)
- [State of The Art](#state-of-the-art)
  - [Most Common Vulnerabilities](#most-common-vulnerabilities)
  - [Existing Tools](#existing-tools)
  - [State of F*](#state-of-f)
- [Proposed Solution](#proposed-solution)
  - [Initial Solution](#initial-solution)
  - [Current Solution](#current-solution)

## Problem Statement

The goal of this project is to investigate how we can build verified financial risk-aware smart contracts that can be deployed to the blockchain. This verification should lead to less error-prone contracts, whose financial properties are verified and ensured before being added to the blockchain.

## Motivation

Advances in economy and society have led to more complex transactions and contracts. Contracts which are often not easily understood, or that leave room for different interpretations. It is important to find a secure and verifiable way of designing such contracts. To do so, we can use distributed ledger technologies.

Distributed ledger technologies have become more popular over the last years, such as the blockchain technology, which allows for a distributed computing architecture. In this model, there are many nodes participating on the network and each transaction is announced to all nodes. Then, the nodes agree on a single order for the transactions. These transactions are grouped into blocks, which, in turn, are given a timestamp and then published. Each block is linked to the previous block in the chain, making it virtually impossible to fake a block[^10.1109/MITP.2018.021921652]. To this property we call *immutability*.

One of the most known blockchains is Ethereum, which is a public blockchain. By being public, it can be inspected and used by anyone. At its core, there are smart contracts. A smart contract is a program, usually written using a high-level language such as Solidity[^solidity], and then compiled into bytecode to be executed by the blockchain. Due to its immutable nature, a smart contract should work as expected. Any bugs or errors will be permanent once a smart contract is deployed.

Over the last years, there have been many attacks to specific, well-known, smart contracts. For instance, in June 2016, the DAO smart contract suffered an attack that led to $60 million of ether to be stolen[^dao-attack]; in August 2021, the Poly network lost $600 million[^poly-attack]; among others. The impact of bugs like this in big economies can lead to massive losses and market disruptions. Then, it is important that we can deploy contracts where we have a certain degree of certainty that they are **(1)** bug free; and **(2)** do not pose a large financial risk to network as a whole.

## State of The Art

There are numerous surveys analyzing the existing tools to formally verify and validate smart contracts[^10.1145/3464421][^10.1016/j.pmcj.2020.101227][^10.1016/j.patter.2020.100179][^10.1145/3437378.3437879][^10.1007/978-3-030-78621-2_14] and to evaluate them, whether by testing against real life smart contracts[^10.1145/3377811.3380364], or by injecting bugs on contracts' code[^10.1145/3395363.3397385].

### Most Common Vulnerabilities

According to current research[^10.1007/978-3-030-78621-2_14][^dasp], the most common smart contract vulnerabilities are: integer underflow and overflow, reentrancy and access control. There are others, such as denial of services, bad randomness, time manipulation, among others, but I will not be going into detail into those.

**Integer Underflow and Overflow**, also known as Arithmetic Bug, happen when the underflow or overflow of an integer variable is not taken into account and leads to unwanted behavior. When writing a contract, one must be sure to ensure that integers stay within the limits provided by the language.

**Reentrancy**, also known as recursive call vulnerability, occurs when external calls to a contract are allowed to make new calls to the calling contract even though the initial execution did not finish yet. This means that the contract will be in the middle of execution while executing other calls, leading to an unexpected state. The already mentioned DAO smart contract is the most well-known case of this attack[^dao-attack].

**Access Control** *TODO*
### Existing Tools

*TODO*

### State of F*

The initial assignment proposed the use of the language F* to tackle this [problem](#problem-statement). F*[^fstar] is a language that combines general-purpose programming with a proof assistant, being based on dependent types. F* can be compiled to OCaml, F#, or even to C. There is some research done in F* aimed at verification of smart contracts in the Ethereum network.

**Formal Verification of Smart Contracts: Short Paper**[^10.1145/2993600.2993611]: in 2016, Bhargavan et al. introduce a framework to analyze and verify runtime safety and functional correctness of Ethereum smart contracts. To do this, they built two tools, Solidity* and EVM*, which convert Solidity and EVM bytecode to F*, respectively. The idea is to then verify the resulting F* code. However, it is noted[^10.1145/3437378.3437879] that this tools do not proide fully automatic verification since the users are required to manually define the effects in F*.

**A Semantic Framework for the Security Analysis of Ethereum Smart Contracts**[^10.1007/978-3-319-89722-6_10]: in 2018, Grishchenko et al. introduce a formalization of the EVM bytecode in F*. Among the verified properties, we can find call integrity, atomicity and independence from miner controller parameters. This tool is mostly compared to Oyente, which is neither complete nor sound. The code is open source[^ethsemantics].

**Celestial: A Smart Contracts Verification Framework**[^celestial]: in 2020, Dharanikota et al. introduce Celestial, a tool to convert contracts with functional annotated-specifications written in Solidity to F*, and verify those. After verification, the tool removes the specification to generate the Solidity code and add it to the blockchain. Unfortunately, the tool is no longer maintained, it only supported a subset of Solidity and Solidity has also had major releases since this tool was built.

We can see that most of the existing tools in F* are either outdated, or not complete. Besides, the lack of documentation[^fstar-docs] led us to remove the usage of F* as a requirement.

## Proposed Solution

### Initial Solution

Our initial proposed solution was to design a domain specific language (DSL) targeted to the financial domain, mainly to exchange any assets by other assets. This language would be human readable, while keeping a strict grammar, in order to be easily processed by a parser. In addition, each primitive of the contract would encode, in some way, an intrinsic risk that, in later stages, would be used to calculate how much risk the contract imposes to the financial system.

A simple use case for this DSL could be simply described as: **(1)** a lawyer writes a contract between Firm A and Firm B. Then, **(2)**, the contract is parsed and validated by the software and compiled into the blockchain's bytecode. Afterwards, **(3)** it would be submitted to the blockchain and there would be a validation that checks whether or not adding this contract increases the systemic risk by a certain threshold. In case it passes the checks, **(4a)** the contract is added to the blockchain and can be executed. Otherwise, **(4b)** the contract is rejected.

### Current Solution

The current solution has slightly deviated from the initial proposal, while still keeping its roots. Instead of designing a single DSL language, our solution is to create an ecosystem of tools that surround a single, common, format that represents a financial contract composed by one or more agreements.

![Ecosystem Overview](ecosystem.svg)

This tools convert from and to the common format. For example, there can be a tool to convert a human readable text to the common format, and then from the common format to Solidity, to be executed in the Ethereum blockchain. There can be a tool to convert from the common specification back to a human readable language, such as English or French.

Our contribution focuses on writing the [common specification](SPECIFICATION.md), as well as building three initial tools that can be part of this ecosystem. You can read more about on the [readme](README.md) of this project.

[^solidity]: https://soliditylang.org/

[^dao-attack]: https://www.gemini.com/cryptopedia/the-dao-hack-makerdao

[^poly-attack]: https://www.coindesk.com/markets/2021/08/10/cross-chain-defi-site-poly-network-hacked-hundreds-of-millions-potentially-lost/

[^dasp]: https://dasp.co/

[^10.1109/MITP.2018.021921652]: https://doi.org/10.1109/MITP.2018.021921652

[^10.1016/j.pmcj.2020.101227]: https://doi.org/10.1016/j.pmcj.2020.101227

[^10.1145/3437378.3437879]: https://doi.org/10.1145/3437378.3437879

[^10.1016/j.patter.2020.100179]: https://doi.org/10.1016/j.patter.2020.100179

[^10.1145/3464421]: https://doi.org/10.1145/3464421

[^10.1007/978-3-030-78621-2_14]: https://doi.org/10.1007/978-3-030-78621-2_14

[^10.1145/3377811.3380364]: https://doi.org/10.1145/3377811.3380364

[^10.1145/3395363.3397385]: https://doi.org/10.1145/3395363.3397385

[^fstar]: https://www.fstar-lang.org/

[^fstar-docs]: https://github.com/FStarLang/FStar/issues/1566

[^celestial]: https://www.microsoft.com/en-us/research/publication/celestial-a-smart-contracts-verification-framework/

[^10.1007/978-3-319-89722-6_10]: https://doi.org/10.1007/978-3-319-89722-6_10

[^ethsemantics]: https://secpriv.wien/ethsemantics/

[^10.1145/2993600.2993611]: https://doi.org/10.1145/2993600.2993611