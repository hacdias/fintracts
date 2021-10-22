# Verification Financial Smart Contracts

<!-- 

TODO:
  - Refresh section numbers
  - Refresh TOC
  - Proof read
-->

- [1. Problem Statement](#1-problem-statement)
- [2. Motivation](#2-motivation)
- [3. State of The Art](#3-state-of-the-art)
  - [3.1. Common Vulnerabilities](#31-common-vulnerabilities)
  - [3.2. Existing Tools](#32-existing-tools)
  - [3.3. State of F*](#33-state-of-f)
- [4. Proposed Solution](#4-proposed-solution)
  - [4.1. Initial Solution](#41-initial-solution)
  - [4.2. Current Solution](#42-current-solution)

## 1. Problem Statement

The goal of this project is to investigate how we can build verified financial risk-aware smart contracts that can be deployed to the blockchain. This verification should lead to less error-prone contracts, whose financial properties are verified and ensured before being added to the blockchain.

## 2. Motivation

Advances in economy and society have led to more complex transactions and contracts. Contracts which are often not easily understood, or that leave room for different interpretations. It is important to find a secure and verifiable way of designing such contracts. To do so, we can use distributed ledger technologies.

Distributed ledger technologies have become more popular over the last years, such as the blockchain technology, which allows for a distributed computing architecture. In this model, there are many nodes participating on the network and each transaction is announced to all nodes. Then, the nodes agree on a single order for the transactions. These transactions are grouped into blocks, which, in turn, are given a timestamp and then published. Each block is linked to the previous block in the chain, making it virtually impossible to fake a block[^10.1109/MITP.2018.021921652]. To this property we call *immutability*.

One of the most known blockchains is Ethereum, which is a public blockchain. By being public, it can be inspected and used by anyone. At its core, there are smart contracts. A smart contract is a program, usually written using a high-level language such as Solidity[^solidity], and then compiled into bytecode to be executed by the blockchain. Due to its immutable nature, a smart contract should work as expected. Any bugs or errors will be permanent once a smart contract is deployed.

Over the last years, there have been many attacks to specific, well-known, smart contracts. For instance, in June 2016, the DAO smart contract suffered an attack that led to $60 million of ether to be stolen[^dao-attack]; in August 2021, the Poly network lost $600 million[^poly-attack]; among others. The impact of bugs like this in big economies can lead to massive losses and market disruptions. Then, it is important that we can deploy contracts where we have a certain degree of certainty that they are **(1)** bug free; and **(2)** do not pose a large financial risk to network as a whole.

## 3. State of The Art

### 3.1. Common Vulnerabilities

According to current research[^10.1007/978-3-030-78621-2_14][^dasp], the most common smart contract vulnerabilities are: integer underflow and overflow, reentrancy and access control. There are others, such as denial of services, bad randomness, time manipulation, among others, but I will not be going into detail into those.

**Integer Underflow and Overflow**, also known as Arithmetic Bug, happen when the underflow or overflow of an integer variable is not taken into account and leads to unwanted behavior. When writing a contract, one must be sure to ensure that integers stay within the limits provided by the language.

**Reentrancy**, also known as recursive call vulnerability, occurs when external calls to a contract are allowed to make new calls to the calling contract even though the initial execution did not finish yet. This means that the contract will be in the middle of execution while executing other calls, leading to an unexpected state. The already mentioned DAO smart contract is the most well-known case of this attack[^dao-attack].

**Access Control** is a type of issue that is common in almost all programs. Issues can arise when the access control is not properly handled, i.e., when users can change or access the smart contracts' functionality without having the right access. A common case here is not protect the initialization function against multiple calls, letting the contract owner be overwritten by someone else.

### 3.2. Existing Tools

There are numerous surveys analyzing the existing tools to formally verify and validate smart contracts[^10.1145/3464421][^10.1016/j.pmcj.2020.101227][^10.1016/j.patter.2020.100179][^10.1145/3437378.3437879][^10.1007/978-3-030-78621-2_14][^tools-overview] and to evaluate them, whether by testing against real life smart contracts[^10.1145/3377811.3380364], or by injecting bugs on contracts' code[^10.1145/3395363.3397385]. Some of this tools are briefly explained below.

**Oyente**[^oyente] is a symbolic execution tool that checks for various patterns, such as: transaction ordering dependency, timestamp dependency, mishandled exceptions and reentrancy. Even though it verifies against important bugs, it is incomplete[^10.1145/3437378.3437879] and throws too many false positives[^10.1145/3377811.3380364].

**Maian**[^maian], similarly to Oyente, does symbolic analysis. This tool is highly specific and validates sequence of invocations to detect fund locking and leaking, as well as to detect whether or not a contract can be killed.

Other tools that also target a specific small amount of vulnerabilities are **Mythril**[^mythril], **Slither**[^slither] and **FSolidM**[^fsolidm]. Mythril uses symoblic analysis, Slither static analysis, and FSolidM uses an automata-based approach with finite state machines.

**ReJection**[^10.1007/978-981-15-3418-8_5] is a tool that detects reentrancy vulnerabilities on a given Solidity smart contract through its abstract syntax tree (AST). This tool is implemented on top of Slither, mention above, by adding more functionality.

**Zeus**[^10.14722/ndss.2018.23082] translates Solidity to LLVM bytecode. It checks for certain types of safety vulnerabilities, mentioned in the paper. The authors indicate that this tool has zero false negatives and a low positive rate. It is not fully automatic as it requires a XACML-like file to be provided[^10.1145/3437378.3437879].

**VeriSol**[^verisol] targets a limited amount of vulnerabilities and supports limited functionality[^10.1145/3437378.3437879].

**solc-verify**[^10.1007/978-3-030-41600-3_11] verifies smart contracts against common vulnerabilities given their Solidity code and a formalization of Solidity's memory model[^10.1007/978-3-030-44914-8_9]. This could should be annotated with their specification. Otherwise, the verification is limited.

**Eth2Vec**[^10.1145/3457337.3457841] is a static analysis tool based on Machine Learning techniques that identifies vulnerabilities in smart contracts. This tool is trained on learning smart contract code via their EVM bytecode, Assembly code and AST. According to research, it has high throughput and accuracy, being also quite resistant to code rewrites and refactoring.

**Solidifier**[^10.1145/3412841.3442051] encodes Solidity using Boogie for verification. It captures Solidity's memory model and uses lazy blockchain exploration and memory-precise verification harnesses. The author's evaluation shows that Solidifier provides a better speed-precision compromise than similar tools. 

**VeriSmart**[^10.1109/SP40000.2020.00032] is a tool that ensures arithmetic safety of smart contracts written in Solidity. This algorithm can infer hidden transaction invariants and leverage them during the verification process.

There is also a small set of tools surround Coq[^coq], a proof assistant, that are worth mention. Firstly, there is ConCert[^10.1145/3372885.3373829], a smart contract certification framework, not targeted at a particular blockchain. On top of ConCert, a tool was built[^10.1145/3437992.3439934] to extract smart contracts written in Liquidity, Midland and Elm. There's also a hybrid formal verification system in Coq that uses static analysis to scan for vulnerabilities and symbolic execution for verification[^10.1109/ACCESS.2020.2969437].

On the legal and financial domain, there are also some interesting tools. **Catala**[^2106.04826] is an open-source[^catala] domain-specific language that can be used to derive algorithms for legislative texts. There is also **Findel**[^10.1016/j.jlamp.2021.100665], which is a composable domain specific language for financial derivatives that can be executed in the blockchain. However, it is limited by supporting only two parties, no loops and there is no guarantee the contracts are executed.

### 3.3. State of F*

The initial assignment proposed the use of the language F* to tackle this [problem](#problem-statement). F*[^fstar] is a language that combines general-purpose programming with a proof assistant, being based on dependent types. F* can be compiled to OCaml, F#, or even to C. There is some research done in F* aimed at verification of smart contracts in the Ethereum network.

In 2016, Bhargavan et al.[^10.1145/2993600.2993611] introduce a framework to analyze and verify runtime safety and functional correctness of Ethereum smart contracts. To do this, they built two tools, Solidity* and EVM*, which convert Solidity and EVM bytecode to F*, respectively. The idea is to then verify the resulting F* code. However, it is noted[^10.1145/3437378.3437879] that this tools do not proide fully automatic verification since the users are required to manually define the effects in F*.

In 2018, Grishchenko et al.[^10.1007/978-3-319-89722-6_10] introduce a formalization of the EVM bytecode in F*. Among the verified properties, we can find call integrity, atomicity and independence from miner controller parameters. This tool is mostly compared to Oyente, which is neither complete nor sound. The code is open source[^ethsemantics].

In 2020, Dharanikota et al.[^celestial] introduce Celestial, a tool to convert contracts with functional annotated-specifications written in Solidity to F*, and verify those. After verification, the tool removes the specification to generate the Solidity code and add it to the blockchain. Unfortunately, the tool is no longer maintained, it only supported a subset of Solidity and Solidity has also had major releases since this tool was built.

We can see that most of the existing tools in F* are either outdated, or not complete. Besides, the lack of documentation[^fstar-docs] led us to remove the usage of F* as a requirement.

## 4. Proposed Solution

### 4.1. Initial Solution

Our initial proposed solution was to design a domain specific language (DSL) targeted to the financial domain, mainly to exchange any assets by other assets. This language would be human readable, while keeping a strict grammar, in order to be easily processed by a parser. In addition, each primitive of the contract would encode, in some way, an intrinsic risk that, in later stages, would be used to calculate how much risk the contract imposes to the financial system.

A simple use case for this DSL could be simply described as: **(1)** a lawyer writes a contract between Firm A and Firm B. Then, **(2)**, the contract is parsed and validated by the software and compiled into the blockchain's bytecode. Afterwards, **(3)** it would be submitted to the blockchain and there would be a validation that checks whether or not adding this contract increases the systemic risk by a certain threshold. In case it passes the checks, **(4a)** the contract is added to the blockchain and can be executed. Otherwise, **(4b)** the contract is rejected.

### 4.2. Current Solution

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

[^10.1007/978-3-030-41600-3_11]: https://doi.org/10.1007/978-3-030-41600-3_11

[^tools-overview]: https://github.com/leonardoalt/ethereum_formal_verification_overview

[^10.1145/3457337.3457841]: https://doi.org/10.1145/3457337.3457841

[^10.1007/978-3-030-44914-8_9]: https://doi.org/10.1007/978-3-030-44914-8_9

[^mythril]: https://github.com/ConsenSys/mythril

[^slither]: https://github.com/crytic/slither

[^maian]: https://github.com/ivicanikolicsg/MAIAN

[^oyente]: https://oyente.tech/

[^10.1145/3412841.3442051]: https://doi.org/10.1145/3412841.3442051

[^10.1109/SP40000.2020.00032]: https://doi.org/10.1109/SP40000.2020.00032

[^10.14722/ndss.2018.23082]: https://dx.doi.org/10.14722/ndss.2018.23082

[^verisol]: https://github.com/microsoft/verisol

[^fsolidm]: https://github.com/anmavrid/smart-contracts

[^coq]: https://coq.inria.fr/

[^10.1145/3372885.3373829]: https://doi.org/10.1145/3372885.3373829

[^10.1109/ACCESS.2020.2969437]: https://doi.org/10.1109/ACCESS.2020.2969437

[^10.1145/3437992.3439934]: https://doi.org/10.1145/3437992.3439934

[^10.1007/978-981-15-3418-8_5]: https://doi.org/10.1007/978-981-15-3418-8_5

[^2106.04826]: https://arxiv.org/abs/2106.04826

[^catala]: https://github.com/CatalaLang/catala

[^10.1016/j.jlamp.2021.100665]: https://doi.org/10.1016/j.jlamp.2021.100665

<!-- 

## Smart Contracts, Real-Virtual World Convergence and Economic Implications

- **Date**: Aug 3, 2021
- **URL**: https://dx.doi.org/10.2139/ssrn.3898144

## Blockchain-Based Business Processes: A Solidity-to-CPN Formal Verification Approach

- **Date**: May 30, 2021
- **URL**: https://doi.org/10.1007/978-3-030-76352-7_7

Proposes a translation algorithm of Solidity contracts into colored petri nets for verification. Concluded as not good.

## Towards automated verification of smart contract fairness

- **Date**: Nov 8, 2020
- **URL**: https://doi.org/10.1145/3368089.3409740

## eThor: Practical and Provably Sound Static Analysis of Ethereum Smart Contracts

- **Date**: Oct 30, 2020
- **URL**: https://doi.org/10.1145/3372297.3417250

## Formal Verification of Ethereum Smart Contracts Using Isabelle/HOL

- **Date**: Oct 28, 2020
- **URL**: https://doi.org/10.1007/978-3-030-62077-6_7

## The Good, The Bad and The Ugly: Pitfalls and Best Practices in Automated Sound Static Analysis of Ethereum Smart Contracts

- **Date**: Oct 27, 2020
- **URL**: https://doi.org/10.1007/978-3-030-61467-6_14

## Accurate Smart Contract Verification Through Direct Modeling

- **Date**: Oct 27, 2020
- **URL**: https://doi.org/10.1007/978-3-030-61467-6_12

## ÆGIS: Shielding Vulnerable Smart Contracts Against Attacks

- **Date**: Oct 5, 2020
- **URL**: https://doi.org/10.1145/3320269.3384756

## Verified Development and Deployment of Multiple Interacting Smart Contracts with VeriSolid

- **Date**: Aug 17, 2020
- **URL**: https://doi.org/10.1109/ICBC48266.2020.9169428

## Deductive Proof of Industrial Smart Contracts Using Why3

- **Date**: Aug 13, 2020
- **URL**: https://doi.org/10.1007/978-3-030-54994-7_22

Show that Why3 is suitable for writing and verifying smart contract programs by reproducing the behavior of Solidity functions with Why3.

## Verifying Smart Contracts with Cubicle

- **Date**: Aug 13, 2020
- **URL**: https://doi.org/10.1007/978-3-030-54994-7_23

Presents Cubicle: a model checker for smart contracts. It implements the model of the smart contract itself and the blockchain transaction mechanism behind it.

## Securing smart contract with runtime validation

- **Date**: Jun 11, 2020
- **URL**: https://doi.org/10.1145/3385412.3385982

Solythesis: source to source runtime validation tool, can enforce global invariants with quantifiers, runtime validation.

The authors mention that there's new languages being designed that can eliminate certain types of errors during compile time by limiting language expressiveness.

Conclude that runtime validation is effective and efficient for smart contracts as it doesn't add a lot of overhead.

## Towards a Formally Verified EVM in Production Environment

- **Date**: Jun 10, 2020
- **URL**: https://doi.org/10.1007/978-3-030-50029-0_21

Define/formalize the EVM behavior in Why3. Later translated to OCaml, and interfaced with Rust in order to execute smart contracts on it.

## PASO: A Web-Based Parser for Solidity Language Analysis

- **Date**: Mar 30, 2020
- **URL**: https://doi.org/10.1109/IWBOSE50093.2020.9050263

Present PASO, a web-based tool able to compute smart contract metrics (payable, mappings, modifiers, addresses, events, contracts, ABI and Bytecode size). First web-based tool of the kind.

## ContractWard: Automated Vulnerability Detection Models for Ethereum Smart Contracts

- **Date**: Jan 23, 2020
- **URL**: https://doi.org/10.1109/TNSE.2020.2968505

Targets specific vulnerabilities (6). Good accuracy.

## Securify: Practical Security Analysis of Smart Contracts

- **Date**: Oct 15, 2018
- **URL**: https://doi.org/10.1145/3243734.3243780

Very well cited.

## Towards Verifying Ethereum Smart Contract Bytecode in Isabelle/HOL

- **Date**: Jan 8, 2018
- **URL**: https://doi.org/10.1145/3167084

Verification of Ethereum smart contracts at the EVM byte code level. They built a formal EVM model in Isabelle/HOL. Smart contracts need to be manually written in Isabel/HOL (I think).

## Kind

- **URL**: https://github.com/uwu-tech/kind

"Modern proof language". Aim to support writing smart contracts soon.

## Others

- [Bond issues: step-by-step guide by Practical Law Finance](https://uk.practicallaw.thomsonreuters.com/1-505-0428)
- [Marlowe: financial contracts on blockchain](https://iohk.io/en/research/library/papers/marlowefinancial-contracts-on-blockchain/)
- [The Extended UTXO Model](https://iohk.io/en/research/library/papers/the-extended-utxo-model/)
- [Rabobank executes real-time commercial paper transaction using blockchain technology](https://www.finextra.com/newsarticle/37389/rabobank-executes-real-time-commercial-paper-transaction-using-blockchain-technology)

-->
