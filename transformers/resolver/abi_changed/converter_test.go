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

package abi_changed_test

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vulcanize/ens_transformers/transformers/resolver/abi_changed"
	"github.com/vulcanize/ens_transformers/transformers/test_data"
)

var _ = Describe("AbiChanged Converter", func() {
	var converter = abi_changed.AbiChangedConverter{}

	Describe("ToEntity", func() {
		It("converts an eth log to a AbiChanged entity", func() {
			entities, err := converter.ToEntities(test_data.ResolverAbi, []types.Log{test_data.EthAbiChangedLog})

			Expect(err).NotTo(HaveOccurred())
			Expect(len(entities)).To(Equal(1))
			entity := entities[0]
			Expect(entity).To(Equal(test_data.AbiChangedEntity))
		})

		It("returns an error if converting log to entity fails", func() {
			_, err := converter.ToEntities("error abi", []types.Log{test_data.EthAbiChangedLog})

			Expect(err).To(HaveOccurred())
		})
	})

	Describe("ToModel", func() {
		var emptyEntity = abi_changed.AbiChangedEntity{}

		It("converts an Entity to a Model", func() {
			models, err := converter.ToModels([]interface{}{test_data.AbiChangedEntity})

			Expect(err).NotTo(HaveOccurred())
			Expect(len(models)).To(Equal(1))
			model := models[0]
			Expect(model).To(Equal(test_data.AbiChangedModel))
		})

		It("returns an error if the entity type is wrong", func() {
			_, err := converter.ToModels([]interface{}{test_data.WrongEntity{}})

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("entity of type test_data.WrongEntity, not abi_changed.AbiChangedEntity"))
		})

		It("handles nil values", func() {
			emptyLog, err := json.Marshal(types.Log{})
			Expect(err).NotTo(HaveOccurred())
			var temp *big.Int
			expectedModel := abi_changed.AbiChangedModel{
				Resolver:         "0x0000000000000000000000000000000000000000",
				Node:             "0x0000000000000000000000000000000000000000000000000000000000000000",
				ContentType:      temp.String(),
				TransactionIndex: 0,
				Raw:              emptyLog,
			}
			models, err := converter.ToModels([]interface{}{emptyEntity})

			Expect(err).NotTo(HaveOccurred())
			Expect(len(models)).To(Equal(1))
			model := models[0]
			Expect(model).To(Equal(expectedModel))
		})
	})
})
