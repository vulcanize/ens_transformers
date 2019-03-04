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
	TemporaryPubkeyChangedData        = "0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000005"
	TemporaryPubkeyChangedTransaction = "0x5c698f13940a2153440c6d19660878bc90219d9298fdcf37365aa8d88d40fc42"
)

var (
	pubkeyChangedRawJson, _ = json.Marshal(EthPubkeyChangedLog)
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
	LogIndex:         EthPubkeyChangedLog.Index,
	TransactionIndex: EthPubkeyChangedLog.TxIndex,
	Raw:              EthPubkeyChangedLog,
}

var PubkeyChangedModel = pubkey_changed.PubkeyChangedModel{
	LogIndex:         EthPubkeyChangedLog.Index,
	TransactionIndex: EthPubkeyChangedLog.TxIndex,
	Raw:              pubkeyChangedRawJson,
}
