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
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/vulcanize/vulcanizedb/pkg/fakes"

	"github.com/vulcanize/ens_transformers/transformers/registar/new_bid"
)

const (
	TemporaryNewBidBlockNumber = int64(26)
	TemporaryNewBidData        = "0x0000000000000000000000000000000000000000000000000000000000000005"
	TemporaryNewBidTransaction = "0x5c698f13940a2153440c6d19660878bc90219d9298fdcf37365aa8d88d40fc42"
)

var (
	newBidRawJson, _ = json.Marshal(EthNewBidLog)
	deposit          = big.NewInt(5)
)

var EthNewBidLog = types.Log{
	Address: common.HexToAddress(RegistarAddress),
	Topics: []common.Hash{
		common.HexToHash("0x99b5620489b6ef926d4518936cfec15d305452712b88bd59da2d9c10fb0953e8"),
		common.HexToHash("0x4554480000000000000000000000000000000000000000000000000000000000"),
		common.HexToHash("0x0000000000000000000000000000d8b4147eda80fec7122ae16da2479cbd7ffb"),
	},
	Data:        hexutil.MustDecode(TemporaryNewBidData),
	BlockNumber: uint64(TemporaryNewBidBlockNumber),
	TxHash:      common.HexToHash(TemporaryNewBidTransaction),
	TxIndex:     111,
	BlockHash:   fakes.FakeHash,
	Index:       7,
	Removed:     false,
}

var NewBidEntity = new_bid.NewBidEntity{
	Hash:             node,
	Bidder:           owner,
	Deposit:          deposit,
	LogIndex:         EthNewBidLog.Index,
	TransactionIndex: EthNewBidLog.TxIndex,
	Raw:              EthNewBidLog,
}

var NewBidModel = new_bid.NewBidModel{
	Hash:             node.Hex(),
	Bidder:           owner.Hex(),
	Deposit:          deposit.String(),
	LogIndex:         EthNewBidLog.Index,
	TransactionIndex: EthNewBidLog.TxIndex,
	Raw:              newBidRawJson,
}
