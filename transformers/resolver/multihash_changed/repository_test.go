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

package multihash_changed_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"

	"github.com/vulcanize/ens_transformers/test_config"
	"github.com/vulcanize/ens_transformers/transformers/resolver/multihash_changed"
	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
	"github.com/vulcanize/ens_transformers/transformers/test_data"
	"github.com/vulcanize/ens_transformers/transformers/test_data/shared_behaviors"
)

var _ = Describe("MultihashChanged repository", func() {
	var (
		hashInvalidatedRepository multihash_changed.MultihashChangedRepository
		db                        *postgres.DB
	)

	BeforeEach(func() {
		db = test_config.NewTestDB(test_config.NewTestNode())
		test_config.CleanTestDB(db)
		hashInvalidatedRepository = multihash_changed.MultihashChangedRepository{}
		hashInvalidatedRepository.SetDB(db)
	})

	Describe("Create", func() {
		modelWithDifferentLogIdx := test_data.MultihashChangedModel
		modelWithDifferentLogIdx.LogIndex++
		inputs := shared_behaviors.CreateBehaviorInputs{
			CheckedHeaderColumnName:  constants.MultihashChangedChecked,
			LogEventTableName:        "ens.multihash_changed",
			TestModel:                test_data.MultihashChangedModel,
			ModelWithDifferentLogIdx: modelWithDifferentLogIdx,
			Repository:               &hashInvalidatedRepository,
		}

		shared_behaviors.SharedRepositoryCreateBehaviors(&inputs)

		It("persists a multihash_changed record", func() {
			headerRepository := repositories.NewHeaderRepository(db)
			headerID, err := headerRepository.CreateOrUpdateHeader(fakes.FakeHeader)
			Expect(err).NotTo(HaveOccurred())

			err = hashInvalidatedRepository.Create(headerID, []interface{}{test_data.MultihashChangedModel})

			Expect(err).NotTo(HaveOccurred())
			var dbMultihashChanged multihash_changed.MultihashChangedModel
			err = db.Get(&dbMultihashChanged, `SELECT resolver, node, hash, log_idx, tx_idx, raw_log FROM ens.multihash_changed WHERE header_id = $1`, headerID)
			Expect(err).NotTo(HaveOccurred())
			Expect(dbMultihashChanged.Resolver).To(Equal(test_data.MultihashChangedModel.Resolver))
			Expect(dbMultihashChanged.Node).To(Equal(test_data.MultihashChangedModel.Node))
			Expect(dbMultihashChanged.Hash).To(Equal(test_data.MultihashChangedModel.Hash))
			Expect(dbMultihashChanged.LogIndex).To(Equal(test_data.MultihashChangedModel.LogIndex))
			Expect(dbMultihashChanged.TransactionIndex).To(Equal(test_data.MultihashChangedModel.TransactionIndex))
			Expect(dbMultihashChanged.Raw).To(MatchJSON(test_data.MultihashChangedModel.Raw))
		})
	})

	Describe("MarkHeaderChecked", func() {
		inputs := shared_behaviors.MarkedHeaderCheckedBehaviorInputs{
			CheckedHeaderColumnName: constants.MultihashChangedChecked,
			Repository:              &hashInvalidatedRepository,
		}

		shared_behaviors.SharedRepositoryMarkHeaderCheckedBehaviors(&inputs)
	})
})
