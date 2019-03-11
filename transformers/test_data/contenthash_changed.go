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

	"github.com/vulcanize/ens_transformers/transformers/resolver/contenthash_changed"
)

const (
	TemporaryContenthashChangedBlockNumber = int64(26)
	TemporaryContenthashChangedData = "0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000026e40101701b206c41b0e4e24593f5155277e2a1ec45545db24d424974eb71173421523db0e2a60000000000000000000000000000000000000000000000000000"
	TemporaryContenthashChangedTransaction = "0x5c698f13940a2153440c6d19660878bc90219d9298fdcf37365aa8d88d40fc42"
)

var (
	contenthashChangedRawJson, _ = json.Marshal(EthContenthashChangedLog)
	targetHash = []byte{228, 1, 1, 112, 27, 32, 108, 65, 176, 228, 226, 69, 147, 245, 21, 82, 119, 226, 161, 236, 69, 84, 93, 178, 77, 66, 73, 116, 235, 113, 23, 52, 33, 82, 61, 176, 226, 166}
)

var EthContenthashChangedLog = types.Log{
	Address: common.HexToAddress(ResolverAddress),
	Topics: []common.Hash{
		common.HexToHash("0x99b5620489b6ef926d4518936cfec15d305452712b88bd59da2d9c10fb0953e8"),
		common.HexToHash("0x4554480000000000000000000000000000000000000000000000000000000000"),
	},
	Data:        hexutil.MustDecode(TemporaryContenthashChangedData),
	BlockNumber: uint64(TemporaryContenthashChangedBlockNumber),
	TxHash:      common.HexToHash(TemporaryContenthashChangedTransaction),
	TxIndex:     111,
	BlockHash:   fakes.FakeHash,
	Index:       7,
	Removed:     false,
}

var ContenthashChangedEntity = contenthash_changed.ContenthashChangedEntity{
	Resolver:         common.HexToAddress(ResolverAddress),
	Node:             node,
	Hash:             targetHash,
	LogIndex:         EthContenthashChangedLog.Index,
	TransactionIndex: EthContenthashChangedLog.TxIndex,
	Raw:              EthContenthashChangedLog,
}

var ContenthashChangedModel = contenthash_changed.ContenthashChangedModel{
	Resolver:         ResolverAddress,
	Node:             node.Hex(),
	Hash:             targetHash,
	LogIndex:         EthContenthashChangedLog.Index,
	TransactionIndex: EthContenthashChangedLog.TxIndex,
	Raw:              contenthashChangedRawJson,
}
