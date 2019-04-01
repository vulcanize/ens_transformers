// VulcanizeDB
// Copyright © 2018 Vulcanize

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package test_data

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/vulcanize/vulcanizedb/pkg/fakes"

	"github.com/vulcanize/ens_transformers/transformers/resolver/multihash_changed"
)

const (
	TemporaryMultihashChangedBlockNumber = int64(26)
	TemporaryMultihashChangedData        = "0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000026e40101701b206c41b0e4e24593f5155277e2a1ec45545db24d424974eb71173421523db0e2a60000000000000000000000000000000000000000000000000000"
	TemporaryMultihashChangedTransaction = "0x5c698f13940a2153440c6d19660878bc90219d9298fdcf37365aa8d88d40fc42"
)

var (
	multihashChangedRawJson, _ = json.Marshal(EthMultihashChangedLog)
)

var EthMultihashChangedLog = types.Log{
	Address: common.HexToAddress(ResolverAddress),
	Topics: []common.Hash{
		common.HexToHash("0x99b5620489b6ef926d4518936cfec15d305452712b88bd59da2d9c10fb0953e8"),
		common.HexToHash("0x4554480000000000000000000000000000000000000000000000000000000000"),
	},
	Data:        hexutil.MustDecode(TemporaryMultihashChangedData),
	BlockNumber: uint64(TemporaryMultihashChangedBlockNumber),
	TxHash:      common.HexToHash(TemporaryMultihashChangedTransaction),
	TxIndex:     111,
	BlockHash:   fakes.FakeHash,
	Index:       7,
	Removed:     false,
}

var MultihashChangedEntity = multihash_changed.MultihashChangedEntity{
	Resolver:         common.HexToAddress(ResolverAddress),
	Node:             node,
	Hash:             targetHash,
	LogIndex:         EthMultihashChangedLog.Index,
	TransactionIndex: EthMultihashChangedLog.TxIndex,
	Raw:              EthMultihashChangedLog,
}

var MultihashChangedModel = multihash_changed.MultihashChangedModel{
	Resolver:         ResolverAddress,
	Node:             node.Hex(),
	Hash:             targetHash,
	LogIndex:         EthMultihashChangedLog.Index,
	TransactionIndex: EthMultihashChangedLog.TxIndex,
	Raw:              multihashChangedRawJson,
}
