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

package new_owner

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/vulcanize/vulcanizedb/pkg/geth"

	"github.com/vulcanize/ens_transformers/transformers/shared"
)

type NewOwnerConverter struct{}

func (NewOwnerConverter) ToEntities(contractAbi string, ethLogs []types.Log) ([]interface{}, error) {
	var entities []interface{}
	for _, ethLog := range ethLogs {
		entity := &NewOwnerEntity{}
		address := ethLog.Address
		abi, err := geth.ParseAbi(contractAbi)
		if err != nil {
			return nil, err
		}

		contract := bind.NewBoundContract(address, abi, nil, nil, nil)

		err = contract.UnpackLog(entity, "NewOwner", ethLog)
		if err != nil {
			return nil, err
		}

		entity.Raw = ethLog
		entity.LogIndex = ethLog.Index
		entity.TransactionIndex = ethLog.TxIndex

		entities = append(entities, *entity)
	}

	return entities, nil
}

func (converter NewOwnerConverter) ToModels(entities []interface{}) ([]interface{}, error) {
	var models []interface{}
	for _, entity := range entities {
		ownerEntity, ok := entity.(NewOwnerEntity)
		if !ok {
			return nil, fmt.Errorf("entity of type %T, not %T", entity, NewOwnerEntity{})
		}

		logIdx := ownerEntity.LogIndex
		txIdx := ownerEntity.TransactionIndex
		rawLog, err := json.Marshal(ownerEntity.Raw)
		if err != nil {
			return nil, err
		}

		parentHash := ownerEntity.Node.Hex()
		labelHash := ownerEntity.Label.Hex()
		subnode := shared.CreateSubnode(parentHash, labelHash)

		model := NewOwnerModel{
			Node:             parentHash,
			Label:            labelHash,
			Owner:            ownerEntity.Owner.Hex(),
			Subnode:          subnode,
			LogIndex:         logIdx,
			TransactionIndex: txIdx,
			Raw:              rawLog,
		}
		models = append(models, model)
	}
	return models, nil
}
