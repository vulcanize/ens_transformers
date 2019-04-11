# [VulcanizeDB](https://github.com/vulcanize/vulcanizedb) Transformers for [ENS](https://ens.domains)

[![Build Status](https://travis-ci.com/vulcanize/ens_transformers.svg?branch=master)](https://travis-ci.com/vulcanize/ens_transformers)
[![Go Report Card](https://goreportcard.com/badge/github.com/vulcanize/ens_transformers)](https://goreportcard.com/report/github.com/vulcanize/ens_transformers)

## Description

This repository contains transformers for the [Ethereum Name Service](https://ens.domains) contracts whose TransformerInitializers
can be composed and executed over using vulcanizeDB's [composeAndExecute](https://github.com/vulcanize/vulcanizedb/blob/master/documentation/composeAndExecute.md) command.

Event transformers for the individual [Registry](https://github.com/vulcanize/ens_transformers/tree/master/transformers/registry),
[Resolver](https://github.com/vulcanize/ens_transformers/tree/master/transformers/resolver),
and [Registar](https://github.com/vulcanize/ens_transformers/tree/master/transformers/registar) contract events are available.


Additionally, there is an [ENS domain record transformer](https://github.com/vulcanize/ens_transformers/blob/working/transformers/domain_records/DOCUMENTATION.md)
which processes both Registry and Resolver events into domain records. It uses `NewResolver(bytes32 indexed node, address resolver);` events emitted from the Registry
contract to configure and track new Resolver addresses as they arise.

## Setup

These transformers are run as plugins to the [core VulcanizeDB software](https://github.com/vulcanize/vulcanizedb),
as described in the main [README](https://github.com/vulcanize/vulcanizedb/blob/master/README.md#usage).

## Configuration

The event transformers require additional configuration variables to set their starting block, contract address, and abi. An example
of such a config file is provided [here](https://github.com/vulcanize/ens_transformers/blob/master/environments/composeAndExecuteEventTransformers.toml).

The domain record transformer is not configured with any additional config variables outside [what is needed to load it as a plugin](https://github.com/vulcanize/vulcanizedb/blob/master/documentation/composeAndExecute.md#configuration),
as seen in this [config for the domain transformer on mainnet](https://github.com/vulcanize/ens_transformers/blob/master/environments/composeAndExecuteDomainRecordsTransformer.toml).
