# ENS Domain Records Transformer

This transformer dynamically tracks resolver contracts whose addresses are seen emitted from the registry's NewResolver events.

It uses data collected from these events:
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
```

To populate ENS domain records in a Postgres table of this form:
```postgresql
CREATE TABLE public.domain_records (
  id                    SERIAL PRIMARY KEY,
  block_number          BIGINT NOT NULL,
  name_hash             VARCHAR(66) NOT NULL,
  label_hash            VARCHAR(66) NOT NULL,
  parent_hash           VARCHAR(66) NOT NULL,
  owner_addr            VARCHAR(66) NOT NULL,
  resolver_addr         VARCHAR(66),
  points_to_addr        VARCHAR(66),
  resolved_name         VARCHAR(66),
  content_              VARCHAR(66),
  content_type          TEXT,
  pub_key_x             VARCHAR(66),
  pub_key_y             VARCHAR(66),
  ttl                   TEXT,
  text_key              TEXT,
  indexed_text_key      TEXT,
  multihash             TEXT,
  contenthash           TEXT,
  UNIQUE (block_number, name_hash)
);
```

Note that the database inserts an updated record only every time the record changes state (a new event occurs for that namehash)
This means the sequence of records for a given name_hash will have large block_number gaps where the state of the domain in those gaps has not changed since the previous record. 
This removes a lot of redundancy that would otherwise exist in the database, reducing the storage used and greatly reducing the number of database writes performed during sync.
But, this also affects how queries against the database must be structured to extract certain information.