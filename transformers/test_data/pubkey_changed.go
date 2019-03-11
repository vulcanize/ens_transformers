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

	"github.com/vulcanize/ens_transformers/transformers/resolver/pubkey_changed"
)

const (
	TemporaryPubkeyChangedBlockNumber = int64(26)
	TemporaryPubkeyChangedData        = "0x952e25626ff3bb77e2c896c259362efcb621b96aed268faacfbd1e6d7a539f9b3102b2896ee73d9b52ff39c72f007b46b38d8e9c142dcaf0648d9b46fb81cd77"
	TemporaryPubkeyChangedTransaction = "0x5c698f13940a2153440c6d19660878bc90219d9298fdcf37365aa8d88d40fc42"
)

var (
	pubkeyChangedRawJson, _ = json.Marshal(EthPubkeyChangedLog)
	x                       = common.BytesToHash([]byte{149, 46, 37, 98, 111, 243, 187, 119, 226, 200, 150, 194, 89, 54, 46, 252, 182, 33, 185, 106, 237, 38, 143, 170, 207, 189, 30, 109, 122, 83, 159, 155})
	y                       = common.BytesToHash([]byte{49, 2, 178, 137, 110, 231, 61, 155, 82, 255, 57, 199, 47, 0, 123, 70, 179, 141, 142, 156, 20, 45, 202, 240, 100, 141, 155, 70, 251, 129, 205, 119})
)

var EthPubkeyChangedLog = types.Log{
	Address: common.HexToAddress(ResolverAddress),
	Topics: []common.Hash{
		common.HexToHash("0x99b5620489b6ef926d4518936cfec15d305452712b88bd59da2d9c10fb0953e8"),
		common.HexToHash("0x4554480000000000000000000000000000000000000000000000000000000000"),
	},
	Data:        hexutil.MustDecode(TemporaryPubkeyChangedData),
	BlockNumber: uint64(TemporaryPubkeyChangedBlockNumber),
	TxHash:      common.HexToHash(TemporaryPubkeyChangedTransaction),
	TxIndex:     111,
	BlockHash:   fakes.FakeHash,
	Index:       7,
	Removed:     false,
}

var PubkeyChangedEntity = pubkey_changed.PubkeyChangedEntity{
	Resolver:         common.HexToAddress(ResolverAddress),
	Node:             node,
	X:                x,
	Y:                y,
	LogIndex:         EthPubkeyChangedLog.Index,
	TransactionIndex: EthPubkeyChangedLog.TxIndex,
	Raw:              EthPubkeyChangedLog,
}

var PubkeyChangedModel = pubkey_changed.PubkeyChangedModel{
	Resolver:         ResolverAddress,
	Node:             node.Hex(),
	X:                x.Hex(),
	Y:                y.Hex(),
	LogIndex:         EthPubkeyChangedLog.Index,
	TransactionIndex: EthPubkeyChangedLog.TxIndex,
	Raw:              pubkeyChangedRawJson,
}
