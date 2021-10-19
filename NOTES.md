# Fintracts: Research Notes

- [Problem Statement](#problem-statement)
- [Motivation](#motivation)
- [State of the Art](#state-of-the-art)
- [How the Project Started](#how-the-project-started)
- [What Is the Current Goal](#what-is-the-current-goal)
- [State of the Art](#state-of-the-art-1)
  - [Existent Tools](#existent-tools)
- [Proposed Solution](#proposed-solution)
- [Raw Research Notes](#raw-research-notes)
  - [Papers](#papers)
  - [Tools](#tools)
  - [Others](#others)

## Problem Statement

## Motivation

## State of the Art

---

## How the Project Started

Technological advances in economy and society have led to more complex transactions and contract systems. Contracts which are often complex and not easily understood, or that leave room for different interpretations. It is important to find a sustainable, yet secure, trusted and verifiable way of designing contracts. This is where blockchain technology and smart contracts come into play.

Distributed ledger technologies have become increasingly more popular over the years. One of the most promising is blockchain technology, which allow for a distributed computing architecture. In this model, there are many nodes participating on the network and each transaction is announced to all nodes. Then, the nodes have to agree on a single order for the transactions. These transactions are grouped into blocks, which, in turn, are given a timestamp and them published. Each block is linked to the previous block in the chain, making it very difficult to change or fake a block, providing immutability.

One of the most known blockchains is Ethereum. Ethereum is a public blockchain maintained by a network of nodes. By being public, it can be inspected and verified by anyone. One of the main features of Ethereum is the presence of smart contracts. A smart contract is program, usually written using high-level languages such as Solidity, and then compiled into byte code so it can be executed in the blockchain. Due to the immutability of the blockchain, a smart contract should work as expected. Any bugs or errors will be permanent once the smart contract is published.

Over the years, there have been many attacks on the Ethereum network that exploited smart contract bugs, leading to massive losses. In June 2016, the DAO smart contract suffered an attack due to a reentrancy bug. The attacker was able to continuously withdraw funds without actually completing the withdrawal function of the contract, leading to a hard fork - practical roll back - of the Ethereum network. In August 2021, the Poly network was attacked and $600m were stolen. The hacker did it "for fun" to show the vulnerability of the network, later returning the funds.

Many other attacks have happened. The impact of bugs in big economies, where smart contracts can be used within the financial field, can lead to massive losses. Thus, it is important to make sure the contracts are free of issues. One of the ways of verifying the correctness of a contract is through formal verification. However, languages used for formal verification are practically hard for non-technical people, which, in many cases, are the ones interested in writing the contracts. Think about lawyers, bankers, who may want to design a contract for a certain financial application.

The goal is to design a domain specific language (DSL), targeted into the financial domain, which can be used to write a generic contract to exchange an asset by any other asset. This language must be risk-aware, in the sense that by automatically analyzing a contract and the current blockchain state, we will be able to determine the risk of adding the new contract to the blockchain.

In later stages, the DSL can be compiled to any other language and verified. From there, a compiler to the Ethereum Virtual Machine (EVM) byte code can be built in order to be able to execute the actual contract.

## What Is the Current Goal

The current goal has shifted from the initial goal, but the motivation remains. The current goal is to build an ecosystem of tools surrounding a common JSON contract format. This tools can be used to parse contracts to and from English, or from English to a functional smart contract in Solidity, or, perhaps, in EVM. Read more about it on the [README](./README.md).

Nevertheless, the research done on this file regards the current state of the art of formal verification of smart contracts and the different techniques to achieve such goal.

## State of the Art

This section will be written in bullet points for briefness:

- Most contracts are written in [specific languages](https://ethereum.org/en/developers/docs/smart-contracts/languages/), being the most popular Solidity, which compiles down to EVM byte code.
- Regarding F*:
  - F* is a prominent language to verify Solidity smart contracts and there's some research done. However, the tools from the research are mostly closed source and not maintained; and F* is not yet mature enough to be considered as a real solution. It is currently lacking documentation and a manual.
  - Most of the tools presented in the papers are outdated because they were strictly built for the research of those papers and never updated afterwards.
- Most common smart contract vulnerabilities:
  - Integer Underflow and Overflow
  - Reentrancy Attacks
  - Greedy Contracts
  - More in https://dasp.co/
  - More in [The Vulnerabilities in Smart Contracts: A Survey](#the-vulnerabilities-in-smart-contracts-a-survey)
- Existing tools:
  - See table in ["A comprehensive survey on smart contract construction and execution: paradigms, tools, and systems"](#a-comprehensive-survey-on-smart-contract-construction-and-execution-paradigms-tools-and-systems)
  - Section 6 of ["solc-verify: A Modular Verifier for Solidity Smart Contracts"](#solc-verify-a-modular-verifier-for-solidity-smart-contracts)
  - See https://github.com/leonardoalt/ethereum_formal_verification_overview

### Existent Tools

#### Vulnerability Pattern-Based Approaches

- **Oyente**: symbolic execution tool, checks for various patterns (incl. transaction ordering dependency, timestamp dependency, mishandled exceptions and reentrancy). Verifies important bugs, incomplete, may contain false-positives.
- **Maian**: "symbolic analysis with concrete validation over a sequence of invocations to detect fund locking, fund leaking and contracts that can be killed.".
- **Mythril**: "symbolic analysis to detect a variety of security vulnerabilities".
- **Slyther**: "static analysis framework with dedicated vulnerability checkers".
- **Solidity, EVM**: type and effect system to check for vulnerable patterns and gas boundedness.

#### Theorem Prover-Based Approaches

- **Kevm**: "is an executable formal semantics of EVM based on the K framework including a deductive program verifier to check contracts against given specifications".
- **Hirai**: formalization of the EVM in Lem.
- **Scilla**: intermediate language between smart contracts and bytecode which uses Coq.

#### Automata-Based Approaches

- **FSolidM**: specific targeted vulnerabilities.

#### SMT-Based Approaches (Satisfiability Modulo Theories)

- **Zeus**: translates Solidity to LLVM bytecode. Requires user specified policy in XACML-like file.
- **VeriSol**: targets a limited amount of vulnerabilities and supports limited functionality.
- **solc-verify**: verified smart contracts given written in Solidity annotated with their specifications. Does not need specifications, but then verification is limited. Checks common vulnerabilities.

#### Other Tools

- SmartBugs: aggregator of many other Solidity verification tools through Docker images: [https://github.com/smartbugs/smartbugs](https://github.com/smartbugs/smartbugs)

---

More in [Raw Research Notes](#raw-research-notes).

## Proposed Solution

The initial proposed solution was to design a DSL targeted to the financial domain, mainly to exchange of assets by any other assets. This language would be as human readable as possible, while keeping a strict and rigid grammar, allowing it to be processed by a computer easily. A possible use case with this DSL would be:

1. Lawyer writes contract in DSL.
2. Contract is parsed and validated.
3. A implicit risk is calculated from the contract.
4. While adding the contract to the block on the blockchain, verify if adding the new contract does not increase the systematic risk by a pre-defined threshold.
5. Contract in blockchain!

The proposed solution has evolved into what this repository currently contains: an ecosystem of tools surrounding a common format. With the right tools built, the same goals can be achieved. For the initial steps of this projects, we will build:

1. A tool to parse from an English contract to the common JSON format.
2. A web GUI generate a contract in the JSON format.

## Raw Research Notes

### Papers

#### Smart Contracts, Real-Virtual World Convergence and Economic Implications

- **Date**: Aug 3, 2021
- **URL**: https://dx.doi.org/10.2139/ssrn.3898144

#### A Survey of Smart Contract Formal Specification and Verification

- **Date**: Jul 18, 2021
- **URL**: https://doi.org/10.1145/3464421

Review of State of the Art in July 2021.

Section 4.4 is program verification (F* for example). Issues with program verification: need to include precisely abstract components of the smart contract execution and the memory model; no gas mechanism consideration (there are other tools for this).

Usually, what's done is Solidity (or another smart contract lang) to verification language (F*) and not the opposite.

Conclusion: combo of contract-level models and specifications with model checking are the most common strategies; usually program-level representations only analyse security properties through symbolic execution, theorem proving,

#### The Vulnerabilities in Smart Contracts: A Survey

- **Date**: Jun 29, 2021
- **URL**: https://doi.org/10.1007/978-3-030-78621-2_14

Explains most common smart contract vulnerabilities and how the referenced tools act in such environments. However, they don't seem to actually experiment, but only to see what other papers say.

#### Catala: A Programming Language for the Law

- **Date**: Jun 9, 2021
- **URL**: https://arxiv.org/abs/2106.04826

Catala is a DSL and a compiler for programming law specification.

#### Certifying Findel derivatives for blockchain

- **Date**: Jun 1, 2021
- **URL**: https://doi.org/10.1016/j.jlamp.2021.100665

Findel is a composable DSL for dinancial derivatives that can be executed in the blockchain. Limitations: no more than two parties, no loops, no default refund mechanism, contract execution not guaranteed.

#### Blockchain-Based Business Processes: A Solidity-to-CPN Formal Verification Approach

- **Date**: May 30, 2021
- **URL**: https://doi.org/10.1007/978-3-030-76352-7_7

Proposes a translation algorithm of Solidity contracts into colored petri nets for verification. Concluded as not good.

#### Eth2Vec: Learning Contract-Wide Code Representations for Vulnerability Detection on Ethereum Smart Contracts

- **Date**: May 24, 2021
- **URL**: https://doi.org/10.1145/3457337.3457841

Eth2Vec is a static analysis tool based on ML that identifies vulnerabilities in smart contracts by learning smart contract code via their EVM bytecode, assembly code and AST. Seems to have high throughput and accuracy, resistant to code rewrites.

Quite interesting, but not formal verification.

#### Solidifier: bounded model checking solidity using lazy contract deployment and precise memory modeling

- **Date**: Mar 22, 2021
- **URL**: https://doi.org/10.1145/3412841.3442051

No F*. Encodes Solidity using Boogie for verification. Captures Solidity's memory model, lazy blockchain exploration and memory-precise verification harnesses. The author's evaluation shows that Solidifier provides a better speed-precision compromise than similar tools.

#### A comprehensive survey on smart contract construction and execution: paradigms, tools, and systems

- **Date**: Feb 12, 2021
- **URL**: https://doi.org/10.1016/j.patter.2020.100179

Analyse the state of art in February 2021. Table 5 compares existing tools to verify smart contracts with different methods.

#### A Survey on Formal Verification for Solidity Smart Contracts

- **Date**: Feb 1, 2021
- **URL**: https://doi.org/10.1145/3437378.3437879

Analyses state of the art in February 2021 for verifying Solidity smart contracts through some selected formal approaches. They note that Solidity* and EVM* do not provide automatic verification since it requires the user to define the effects in F*.

- FSolidM and VeriSolid: specific targeted vulnerabilities. VeriSolid extends FSolidM but not same vulnerabilities. Formal verification done by NuXmv symbolic checker. Does not take into account variables.
- Zeus: requires user specified policy in XACML-like file.
- Oyente: verifies important bugs, incomplete, may contain false-positives.
- Osiris: targets integer vulnerabilities.

#### Extracting smart contracts tested and verified in Coq

- **Date**: Jan 17, 2021
- **URL**: https://doi.org/10.1145/3437992.3439934

They use ConCert to extract smart contracts written in Liquidity, Midlang and Elm and verify them in Coq.

#### Celestial: A Smart Contracts Verification Framework

- **Date**: Dec 1, 2020
- **URL**: https://www.microsoft.com/en-us/research/publication/celestial-a-smart-contracts-verification-framework/

Celestial converts contracts with functional annotated-specifications written in Solidity to F* and verify the code. After verification, the tool performs erasure of the specification to generate Solidity code ready to execute on the blockchain. Does not take into account gas.

It would be great if I had access to the code.

#### Towards automated verification of smart contract fairness

- **Date**: Nov 8, 2020
- **URL**: https://doi.org/10.1145/3368089.3409740

#### eThor: Practical and Provably Sound Static Analysis of Ethereum Smart Contracts

- **Date**: Oct 30, 2020
- **URL**: https://doi.org/10.1145/3372297.3417250

#### Formal Verification of Ethereum Smart Contracts Using Isabelle/HOL

- **Date**: Oct 28, 2020
- **URL**: https://doi.org/10.1007/978-3-030-62077-6_7

#### The Good, The Bad and The Ugly: Pitfalls and Best Practices in Automated Sound Static Analysis of Ethereum Smart Contracts

- **Date**: Oct 27, 2020
- **URL**: https://doi.org/10.1007/978-3-030-61467-6_14

#### Accurate Smart Contract Verification Through Direct Modeling

- **Date**: Oct 27, 2020
- **URL**: https://doi.org/10.1007/978-3-030-61467-6_12

#### ÆGIS: Shielding Vulnerable Smart Contracts Against Attacks

- **Date**: Oct 5, 2020
- **URL**: https://doi.org/10.1145/3320269.3384756

#### Verification of smart contracts: A survey

- **Date**: Sep 1, 2020
- **URL**: https://doi.org/10.1016/j.pmcj.2020.101227

They compare and review different papers/tools for verification of smart contracts, whether by formal verification or runtime verification. There is also a good list of common vulnerabilities on section 4.

Similar to (by same authors): On the Verification of Smart Contracts: A Systematic Review, https://doi.org/10.1007/978-3-030-59638-5_7

#### Verified Development and Deployment of Multiple Interacting Smart Contracts with VeriSolid

- **Date**: Aug 17, 2020
- **URL**: https://doi.org/10.1109/ICBC48266.2020.9169428

#### Deductive Proof of Industrial Smart Contracts Using Why3

- **Date**: Aug 13, 2020
- **URL**: https://doi.org/10.1007/978-3-030-54994-7_22

Show that Why3 is suitable for writing and verifying smart contract programs by reproducing the behavior of Solidity functions with Why3.

#### Verifying Smart Contracts with Cubicle

- **Date**: Aug 13, 2020
- **URL**: https://doi.org/10.1007/978-3-030-54994-7_23

Presents Cubicle: a model checker for smart contracts. It implements the model of the smart contract itself and the blockchain transaction mechanism behind it.

#### VERISMART: A Highly Precise Safety Verifier for Ethereum Smart Contracts

- **Date**: Jul 30, 2020
- **URL**: https://doi.org/10.1109/SP40000.2020.00032

Propose a new algorithm and tool (VeriSmart) to ensure arithmetic safety of smart contracts written in Solidity. This algorithm can infer hidden transaction invariants and leverage them during the verification process.

#### How effective are smart contract analysis tools? evaluating smart contract static analysis tools using bug injection

- **Date**: Jul 18, 2020
- **URL**: https://doi.org/10.1145/3395363.3397385

Introduces a tool to evaluate static smart contract analysis tools by injecting bugs on the contract's AST.

#### End-to-End Formal Verification of Ethereum 2.0 Deposit Smart Contract

- **Date**: Jul 14, 2020
- **URL**: https://doi.org/10.1007/978-3-030-53288-8_8

Not 100% related to the topic. Interesting. They show how they did a E2E formal verification of the new deposit contract. They used the K framework and its verification infrastructure. Verifying this single contract took 7+2+5 person-weeks excluding discussions, reporting bugs and follow ups.

#### Empirical review of automated analysis tools on 47,587 Ethereum smart contracts

- **Date**: Jun 27, 2020
- **URL**: https://doi.org/10.1145/3377811.3380364

Evaluated many contracts using common tools. Not all vulnerabilities were found and some tools threw too many false positives (Oyente). Bad Randomness and Short Addresses vulnerabilities are not found. See DASP10.

#### Securing smart contract with runtime validation

- **Date**: Jun 11, 2020
- **URL**: https://doi.org/10.1145/3385412.3385982

Solythesis: source to source runtime validation tool, can enforce global invariants with quantifiers, runtime validation.

The authors mention that there's new languages being designed that can eliminate certain types of errors during compile time by limiting language expressiveness.

Conclude that runtime validation is effective and efficient for smart contracts as it doesn't add a lot of overhead.

#### Towards a Formally Verified EVM in Production Environment

- **Date**: Jun 10, 2020
- **URL**: https://doi.org/10.1007/978-3-030-50029-0_21

Define/formalize the EVM behavior in Why3. Later translated to OCaml, and interfaced with Rust in order to execute smart contracts on it.

#### SMT-Friendly Formalization of the Solidity Memory Model

- **Date**: Apr 18, 2020
- **URL**: https://doi.org/10.1007/978-3-030-44914-8_9

Presents the formalization of Solidity's memory model used in the tool solc-verify

#### PASO: A Web-Based Parser for Solidity Language Analysis

- **Date**: Mar 30, 2020
- **URL**: https://doi.org/10.1109/IWBOSE50093.2020.9050263

Present PASO, a web-based tool able to compute smart contract metrics (payable, mappings, modifiers, addresses, events, contracts, ABI and Bytecode size). First web-based tool of the kind.

#### solc-verify: A Modular Verifier for Solidity Smart Contracts

- **Date**: Mar 14, 2020
- **URL**: https://doi.org/10.1007/978-3-030-41600-3_11

solc-verify verifies smart contracts given written in Solidity annotated with  their specification. Uses Boogie as intermediate language. Seems like a nice tool to verify the contract against some common vulnerabilities.

Very good Related Work section (6).

#### ReJection: A AST-Based Reentrancy Vulnerability Detection Method

- **Date**: Feb 20, 2020
- **URL**: https://doi.org/10.1007/978-981-15-3418-8_5

Propose ReJection: tool that inputs Solidity and detects reentrancy vulnerabilities on contract's AST. Implemented on Slither, an open-source vulnerability detection tool.

#### A Hybrid Formal Verification System in Coq for Ensuring the Reliability and Security of Ethereum-Based Service Smart Contracts

- **Date**: Jan 27, 2020
- **URL**: https://doi.org/10.1109/ACCESS.2020.2969437

Use a hybrid system for verification: static analysis to scan for vulnerabilities, symbolic execution for verification and debugging mechanisms. Table 5 contains good properties of existent tools. Does not support all Solidity features.

#### ContractWard: Automated Vulnerability Detection Models for Ethereum Smart Contracts

- **Date**: Jan 23, 2020
- **URL**: https://doi.org/10.1109/TNSE.2020.2968505

Targets specific vulnerabilities (6). Good accuracy.

#### ConCert: a smart contract certification framework in Coq

- **Date**: Jan 20, 2020
- **URL**: https://doi.org/10.1145/3372885.3373829

They introduce ConCert, which is a smart contract verification framework in Coq. Seems to be a general tool, not targeted to any specific smart contract language.

#### Securify: Practical Security Analysis of Smart Contracts

- **Date**: Oct 15, 2018
- **URL**: https://doi.org/10.1145/3243734.3243780

Very well cited.

#### A Semantic Framework for the Security Analysis of Ethereum Smart Contracts

- **Date**: Apr 14, 2018
- **URL**: https://doi.org/10.1007/978-3-319-89722-6_10

Formalizes EVM byte code in F*. Properties verified of smart contracts: call integrity, atomicity, independence from miner controlled params.

Compared to Oyente, whose verification conditions: not complete nor sound.
Solidity code nor formalized, only EVM.

Code: https://secpriv.wien/ethsemantics/

#### Towards Verifying Ethereum Smart Contract Bytecode in Isabelle/HOL

- **Date**: Jan 8, 2018
- **URL**: https://doi.org/10.1145/3167084

Verification of Ethereum smart contracts at the EVM byte code level. They built a formal EVM model in Isabelle/HOL. Smart contracts need to be manually written in Isabel/HOL (I think).

#### Formal Verification of Smart Contracts: Short Paper

- **Date**: Oct 24, 2016
- **URL**: https://doi.org/10.1145/2993600.2993611

Implemented Solidity* (Solidity to F*) and EVM* (EVM byte code to F*) in OCaml. Not fully automatic, requires user to define effects in F* (perhaps not a problem within a closed domain as ours as they can be pre-defined?). Probably tools not up to date as Solidity has changed a lot since 2016.

### Tools

#### Kind

- **URL**: https://github.com/uwu-tech/kind

"Modern proof language". Aim to support writing smart contracts soon.

### Others

- [Bond issues: step-by-step guide by Practical Law Finance](https://uk.practicallaw.thomsonreuters.com/1-505-0428)
- [Marlowe: financial contracts on blockchain](https://iohk.io/en/research/library/papers/marlowefinancial-contracts-on-blockchain/)
- [The Extended UTXO Model](https://iohk.io/en/research/library/papers/the-extended-utxo-model/)
- [Rabobank executes real-time commercial paper transaction using blockchain technology](https://www.finextra.com/newsarticle/37389/rabobank-executes-real-time-commercial-paper-transaction-using-blockchain-technology)
- [Program verification with F*](https://danel.ahman.ee/teaching/eutypes2018/index.html)
- [Verified programming in F*](https://fstar-lang.org/tutorial/tutorial.html)
