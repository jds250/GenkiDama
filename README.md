# GenkiDama

This repository contains the supplementary code for the paper "GENKIDAMA: Cross-Chain Rollup for Scalable Blockchain Interoperability".

## Table of Contents

- [Introduction](#introduction)
- [Requirements](#requirements)
- [Setup](#setup)
- [Usage](#usage)
- [Example](#example)

## Introduction

Our rollup contract can be used on any blockchain that supports Solidity contracts. This project demonstrates how to execute cross-chain transactions using our FullRollup contract.

## Requirements

To run this project, you will need the following dependencies:

- Go (version > 1.20)
- Geth (version 1.12)
- Evmos (latest version)
- Gnark (version 0.8.1)

## Setup

1. **Clone the repository:**

    ```sh
    git clone https://github.com/jds250/GenkiDama.git
    cd GenkiDama
    ```

2. **Install Go:**

    Ensure you have Go installed and the version is greater than 1.20. You can download and install it from [Go's official site](https://golang.org/dl/).

3. **Install Geth:**

    Install Geth (Go-Ethereum) version 1.12. Instructions can be found on the [Geth installation page](https://geth.ethereum.org/docs/install-and-build/installing-geth).

4. **Install Evmos:**

    Follow the instructions on the [Evmos GitHub repository](https://github.com/tharsis/evmos) to install the latest version.

5. **Install Gnark:**

    Install Gnark version 0.8.1 by following the instructions on the [Gnark GitHub repository](https://github.com/ConsenSys/gnark).

## Usage

### Deploy Contracts

To execute a cross-chain transaction, you first need to deploy the FullRollup contract on two target chains.

1. **Deploy FullRollup Contract:**

    Deploy the contract using your preferred Solidity development environment (e.g., Remix, Hardhat, Truffle).

2. **Generate Go-Bindings:**

    Use the `abigen` tool to generate Go bindings for the deployed contract:

    ```sh
    abigen --sol=path/to/FullRollup.sol --pkg=yourpackage --out=fullrollup.go
    ```

### Run Tests

1. **Configure Parameters:**

    Edit the parameters in the `Rollup-Offchain/example/test.go` file to match your deployment setup.

2. **Run the Example Test:**

    ```sh
    go run example/test.go
    ```
