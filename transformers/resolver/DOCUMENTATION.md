# ENS Resolver Transformer

These transformers track these events at the ENS resolver contracts:

```
event AddrChanged(bytes32 indexed node, address a);
event ContentChanged(bytes32 indexed node, bytes32 hash);
event NameChanged(bytes32 indexed node, string name);
event ABIChanged(bytes32 indexed node, uint256 indexed contentType);
event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y);
event TextChanged(bytes32 indexed node, string indexedKey, string key);
event MultihashChanged(bytes32 indexed node, bytes hash);
event ContenthashChanged(bytes32 indexed node, bytes hash);
```