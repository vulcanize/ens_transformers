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

package abi_changed

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/vulcanize/vulcanizedb/pkg/geth"
)

type AbiChangedConverter struct{}

func (AbiChangedConverter) ToEntities(contractAbi string, ethLogs []types.Log) ([]interface{}, error) {
	var entities []interface{}
	for _, ethLog := range ethLogs {
		entity := &AbiChangedEntity{}
		entity.Resolver = ethLog.Address
		abi, err := geth.ParseAbi(contractAbi)
		if err != nil {
			return nil, err
		}

		contract := bind.NewBoundContract(entity.Resolver, abi, nil, nil, nil)

		err = contract.UnpackLog(entity, "AbiChanged", ethLog)
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

func (converter AbiChangedConverter) ToModels(entities []interface{}) ([]interface{}, error) {
	var models []interface{}
	for _, entity := range entities {
		abiEntity, ok := entity.(AbiChangedEntity)
		if !ok {
			return nil, fmt.Errorf("entity of type %T, not %T", entity, AbiChangedEntity{})
		}

		logIdx := abiEntity.LogIndex
		txIdx := abiEntity.TransactionIndex
		rawLog, err := json.Marshal(abiEntity.Raw)
		if err != nil {
			return nil, err
		}

		model := AbiChangedModel{
			Resolver:         abiEntity.Resolver.Hex(),
			Node:             abiEntity.Node.Hex(),
			ContentType:      abiEntity.ContentType.String(),
			LogIndex:         logIdx,
			TransactionIndex: txIdx,
			Raw:              rawLog,
		}
		models = append(models, model)
	}
	return models, nil
}
