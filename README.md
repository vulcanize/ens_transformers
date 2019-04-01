# [VulcanizeDB](https://github.com/vulcanize/vulcanizedb) Transformers for [ENS](https://ens.domains)

[![Join the chat at https://gitter.im/vulcanizeio/VulcanizeDB](https://badges.gitter.im/vulcanizeio/VulcanizeDB.svg)](https://gitter.im/vulcanizeio/VulcanizeDB?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

[![Build Status](https://travis-ci.com/vulcanize/ens_transformers.svg?branch=master)](https://travis-ci.com/vulcanize/ens_transformers)

This repository contains transformers for the [Ethereum Name Service](https://ens.domains) contracts whose [TransformerInitializers](https://github.com/vulcanize/maker-vulcanizedb/blob/compose_and_execute/libraries/shared/transformer/event_transformer.go#L33)
can be composed and executed over using vulcanizeDB's [composeAndExecute](https://github.com/vulcanize/maker-vulcanizedb/blob/compose_and_execute/cmd/composeAndExecute.go) command.

Transformers for the below events are provided:
```
// Registry events
event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner);
event Transfer(bytes32 indexed node, address owner);
event NewResolver(bytes32 indexed node, address resolver);
event NewTTL(bytes32 indexed node, uint64 ttl);

// Resolver events
event AddrChanged(bytes32 indexed node, address a);
event ContentChanged(bytes32 indexed node, bytes32 hash);
event NameChanged(bytes32 indexed node, string name);
event ABIChanged(bytes32 indexed node, uint256 indexed contentType);
event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y);
event TextChanged(bytes32 indexed node, string indexedKey, string key);
event MultihashChanged(bytes32 indexed node, bytes hash);
event ContenthashChanged(bytes32 indexed node, bytes hash);

// Registar events
event AuctionStarted(bytes32 indexed hash, uint registrationDate);
event NewBid(bytes32 indexed hash, address indexed bidder, uint deposit);
event BidRevealed(bytes32 indexed hash, address indexed owner, uint value, uint8 status);
event HashRegistered(bytes32 indexed hash, address indexed owner, uint value, uint registrationDate);
event HashReleased(bytes32 indexed hash, uint value);
event HashInvalidated(bytes32 indexed hash, string indexed name, uint value, uint registrationDate);
```