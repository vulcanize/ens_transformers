# ENS Registar Transformer

These transformers track these events at the ENS hash registar contract:

```
event AuctionStarted(bytes32 indexed hash, uint registrationDate);
event NewBid(bytes32 indexed hash, address indexed bidder, uint deposit);
event BidRevealed(bytes32 indexed hash, address indexed owner, uint value, uint8 status);
event HashRegistered(bytes32 indexed hash, address indexed owner, uint value, uint registrationDate);
event HashReleased(bytes32 indexed hash, uint value);
event HashInvalidated(bytes32 indexed hash, string indexed name, uint value, uint registrationDate);
```