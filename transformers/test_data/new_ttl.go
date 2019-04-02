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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/vulcanize/vulcanizedb/pkg/fakes"

	"github.com/vulcanize/ens_transformers/transformers/registry/new_ttl"
)

const (
	TemporaryNewTtlBlockNumber = int64(26)
	TemporaryNewTtlData        = "0x0000000000000000000000000000000000000000000000000000000000000004"
	TemporaryNewTtlTransaction = "0x5c698f13940a2153440c6d19660878bc90219d9298fdcf37365aa8d88d40fc42"
)

var (
	newTtlRawJson, _ = json.Marshal(EthNewTtlLog)
)

var EthNewTtlLog = types.Log{
	Address: common.HexToAddress(RegistryAddress),
	Topics: []common.Hash{
		common.HexToHash("0x99b5620489b6ef926d4518936cfec15d305452712b88bd59da2d9c10fb0953e8"),
		common.HexToHash("0x4554480000000000000000000000000000000000000000000000000000000000"),
	},
	Data:        hexutil.MustDecode(TemporaryNewTtlData),
	BlockNumber: uint64(TemporaryNewTtlBlockNumber),
	TxHash:      common.HexToHash(TemporaryNewTtlTransaction),
	TxIndex:     111,
	BlockHash:   fakes.FakeHash,
	Index:       7,
	Removed:     false,
}

var NewTtlEntity = new_ttl.NewTTLEntity{
	Node:             node,
	Ttl:              4,
	LogIndex:         EthNewTtlLog.Index,
	TransactionIndex: EthNewTtlLog.TxIndex,
	Raw:              EthNewTtlLog,
}

var NewTtlModel = new_ttl.NewTTLModel{
	Node:             node.Hex(),
	Ttl:              4,
	LogIndex:         EthNewTtlLog.Index,
	TransactionIndex: EthNewTtlLog.TxIndex,
	Raw:              newTtlRawJson,
}
