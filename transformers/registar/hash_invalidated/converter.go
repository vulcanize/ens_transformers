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

package hash_invalidated

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/vulcanize/vulcanizedb/pkg/geth"
)

type HashInvalidatedConverter struct{}

func (HashInvalidatedConverter) ToEntities(contractAbi string, ethLogs []types.Log) ([]interface{}, error) {
	var entities []interface{}
	for _, ethLog := range ethLogs {
		entity := &HashInvalidatedEntity{}
		intermediateMap := map[string]interface{}{}
		address := ethLog.Address
		abi, err := geth.ParseAbi(contractAbi)
		if err != nil {
			return nil, err
		}

		contract := bind.NewBoundContract(address, abi, nil, nil, nil)

		err = contract.UnpackLogIntoMap(intermediateMap, "HashInvalidated", ethLog)
		if err != nil {
			return nil, err
		}

		entity.Hash = common.BytesToHash(intermediateMap["hash"].([]uint8))
		entity.Value = intermediateMap["value"].(*big.Int)
		entity.Name = intermediateMap["name"].(string)
		entity.RegistrationDate = intermediateMap["registrationDate"].(*big.Int)
		entity.Raw = ethLog
		entity.LogIndex = ethLog.Index
		entity.TransactionIndex = ethLog.TxIndex

		entities = append(entities, *entity)
	}

	return entities, nil
}

func (converter HashInvalidatedConverter) ToModels(entities []interface{}) ([]interface{}, error) {
	var models []interface{}
	for _, entity := range entities {
		hashEntity, ok := entity.(HashInvalidatedEntity)
		if !ok {
			return nil, fmt.Errorf("entity of type %T, not %T", entity, HashInvalidatedEntity{})
		}

		logIdx := hashEntity.LogIndex
		txIdx := hashEntity.TransactionIndex
		rawLog, err := json.Marshal(hashEntity.Raw)
		if err != nil {
			return nil, err
		}

		model := HashInvalidatedModel{
			Hash:             hashEntity.Hash.Hex(),
			Name:             hashEntity.Name,
			Value:            hashEntity.Value.String(),
			RegistrationDate: hashEntity.RegistrationDate.String(),
			LogIndex:         logIdx,
			TransactionIndex: txIdx,
			Raw:              rawLog,
		}
		models = append(models, model)
	}
	return models, nil
}
