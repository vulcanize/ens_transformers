// VulcanizeDB
// Copyright © 2019 Vulcanize

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
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/vulcanize/vulcanizedb/pkg/fakes"

	"github.com/vulcanize/ens_transformers/transformers/resolver/abi_changed"
)

const (
	TemporaryAbiChangedBlockNumber = int64(26)
	TemporaryAbiChangedTransaction = "0x5c698f13940a2153440c6d19660878bc90219d9298fdcf37365aa8d88d40fc42"
)

var (
	abiChangedRawJson, _  = json.Marshal(EthAbiChangedLog)
	node                  = common.HexToHash("0x4554480000000000000000000000000000000000000000000000000000000000")
	abiChangedContentType = big.NewInt(1)
)

var EthAbiChangedLog = types.Log{
	Address: common.HexToAddress(ResolverAddress),
	Topics: []common.Hash{
		common.HexToHash("0x99b5620489b6ef926d4518936cfec15d305452712b88bd59da2d9c10fb0953e8"),
		common.HexToHash("0x4554480000000000000000000000000000000000000000000000000000000000"),
		common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
	},
	Data:        nil,
	BlockNumber: uint64(TemporaryAbiChangedBlockNumber),
	TxHash:      common.HexToHash(TemporaryAbiChangedTransaction),
	TxIndex:     111,
	BlockHash:   fakes.FakeHash,
	Index:       7,
	Removed:     false,
}

var AbiChangedEntity = abi_changed.AbiChangedEntity{
	Resolver:         common.HexToAddress(ResolverAddress),
	Node:             node,
	ContentType:      abiChangedContentType,
	LogIndex:         EthAbiChangedLog.Index,
	TransactionIndex: EthAbiChangedLog.TxIndex,
	Raw:              EthAbiChangedLog,
}

var AbiChangedModel = abi_changed.AbiChangedModel{
	Resolver:         ResolverAddress,
	Node:             node.Hex(),
	ContentType:      abiChangedContentType.String(),
	LogIndex:         EthAbiChangedLog.Index,
	TransactionIndex: EthAbiChangedLog.TxIndex,
	Raw:              abiChangedRawJson,
}
