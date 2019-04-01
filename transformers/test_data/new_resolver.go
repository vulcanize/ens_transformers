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

	"github.com/vulcanize/ens_transformers/transformers/registry/new_resolver"
)

const (
	TemporaryNewResolverBlockNumber = int64(26)
	TemporaryNewResolverData        = "0x0000000000000000000000000000d8b4147eda80fec7122ae16da2479cbd7ffb"
	TemporaryNewResolverTransaction = "0x5c698f13940a2153440c6d19660878bc90219d9298fdcf37365aa8d88d40fc42"
)

var (
	newResolverRawJson, _ = json.Marshal(EthNewResolverLog)
)

var EthNewResolverLog = types.Log{
	Address: common.HexToAddress(RegistryAddress),
	Topics: []common.Hash{
		common.HexToHash("0x99b5620489b6ef926d4518936cfec15d305452712b88bd59da2d9c10fb0953e8"),
		common.HexToHash("0x4554480000000000000000000000000000000000000000000000000000000000"),
	},
	Data:        hexutil.MustDecode(TemporaryNewResolverData),
	BlockNumber: uint64(TemporaryNewResolverBlockNumber),
	TxHash:      common.HexToHash(TemporaryNewResolverTransaction),
	TxIndex:     111,
	BlockHash:   fakes.FakeHash,
	Index:       7,
	Removed:     false,
}

var NewResolverEntity = new_resolver.NewResolverEntity{
	Node:             node,
	Resolver:         owner,
	LogIndex:         EthNewResolverLog.Index,
	TransactionIndex: EthNewResolverLog.TxIndex,
	Raw:              EthNewResolverLog,
}

var NewResolverModel = new_resolver.NewResolverModel{
	Node:             node.Hex(),
	Resolver:         owner.Hex(),
	LogIndex:         EthNewResolverLog.Index,
	TransactionIndex: EthNewResolverLog.TxIndex,
	Raw:              newResolverRawJson,
}
