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
	"github.com/vulcanize/ens_transformers/transformers/shared"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/vulcanize/vulcanizedb/pkg/fakes"

	"github.com/vulcanize/ens_transformers/transformers/registry/new_owner"
)

const (
	TemporaryNewOwnerBlockNumber = int64(26)
	TemporaryNewOwnerData        = "0x000000000000000000000000fDb33f8AC7ce72d7D4795Dd8610E323B4C122fbB"
	TemporaryNewOwnerTransaction = "0x5c698f13940a2153440c6d19660878bc90219d9298fdcf37365aa8d88d40fc42"
)

var (
	newOwnerRawJson, _ = json.Marshal(EthNewOwnerLog)
	label              = common.HexToHash("0x0000000000000000000000000000d8b4147eda80fec7122ae16da2479cbd7ffb")
	ownerAddr          = common.HexToAddress("0x000000000000000000000000fDb33f8AC7ce72d7D4795Dd8610E323B4C122fbB")
	subnode            = shared.CreateSubnode(node.Hex(), label.Hex())
)

var EthNewOwnerLog = types.Log{
	Address: common.HexToAddress(RegistryAddress),
	Topics: []common.Hash{
		common.HexToHash("0x99b5620489b6ef926d4518936cfec15d305452712b88bd59da2d9c10fb0953e8"),
		common.HexToHash("0x4554480000000000000000000000000000000000000000000000000000000000"),
		common.HexToHash("0x0000000000000000000000000000d8b4147eda80fec7122ae16da2479cbd7ffb"),
	},
	Data:        hexutil.MustDecode(TemporaryNewOwnerData),
	BlockNumber: uint64(TemporaryNewOwnerBlockNumber),
	TxHash:      common.HexToHash(TemporaryNewOwnerTransaction),
	TxIndex:     111,
	BlockHash:   fakes.FakeHash,
	Index:       7,
	Removed:     false,
}

var NewOwnerEntity = new_owner.NewOwnerEntity{
	Node:             node,
	Label:            label,
	Owner:            ownerAddr,
	LogIndex:         EthNewOwnerLog.Index,
	TransactionIndex: EthNewOwnerLog.TxIndex,
	Raw:              EthNewOwnerLog,
}

var NewOwnerModel = new_owner.NewOwnerModel{
	Node:             node.Hex(),
	Label:            label.Hex(),
	Owner:            ownerAddr.Hex(),
	Subnode:          subnode,
	LogIndex:         EthNewOwnerLog.Index,
	TransactionIndex: EthNewOwnerLog.TxIndex,
	Raw:              newOwnerRawJson,
}
