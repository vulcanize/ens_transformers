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

package hash_invalidated_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"

	"github.com/vulcanize/ens_transformers/test_config"
	"github.com/vulcanize/ens_transformers/transformers/registar/hash_invalidated"
	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
	"github.com/vulcanize/ens_transformers/transformers/test_data"
	"github.com/vulcanize/ens_transformers/transformers/test_data/shared_behaviors"
)

var _ = Describe("HashInvalidated repository", func() {
	var (
		hashInvalidatedRepository hash_invalidated.HashInvalidatedRepository
		db                        *postgres.DB
	)

	BeforeEach(func() {
		db = test_config.NewTestDB(test_config.NewTestNode())
		test_config.CleanTestDB(db)
		hashInvalidatedRepository = hash_invalidated.HashInvalidatedRepository{}
		hashInvalidatedRepository.SetDB(db)
	})

	Describe("Create", func() {
		modelWithDifferentLogIdx := test_data.HashInvalidatedModel
		modelWithDifferentLogIdx.LogIndex++
		inputs := shared_behaviors.CreateBehaviorInputs{
			CheckedHeaderColumnName:  constants.HashInvalidatedChecked,
			LogEventTableName:        "ens.hash_invalidated",
			TestModel:                test_data.HashInvalidatedModel,
			ModelWithDifferentLogIdx: modelWithDifferentLogIdx,
			Repository:               &hashInvalidatedRepository,
		}

		shared_behaviors.SharedRepositoryCreateBehaviors(&inputs)

		It("persists a hash_invalidated record", func() {
			headerRepository := repositories.NewHeaderRepository(db)
			headerID, err := headerRepository.CreateOrUpdateHeader(fakes.FakeHeader)
			Expect(err).NotTo(HaveOccurred())

			err = hashInvalidatedRepository.Create(headerID, []interface{}{test_data.HashInvalidatedModel})

			Expect(err).NotTo(HaveOccurred())
			var dbHashInvalidated hash_invalidated.HashInvalidatedModel
			err = db.Get(&dbHashInvalidated, `SELECT hash, name, value, registration_date, log_idx, tx_idx, raw_log FROM ens.hash_invalidated WHERE header_id = $1`, headerID)
			Expect(err).NotTo(HaveOccurred())
			Expect(dbHashInvalidated.Hash).To(Equal(test_data.HashInvalidatedModel.Hash))
			Expect(dbHashInvalidated.Name).To(Equal(test_data.HashInvalidatedModel.Name))
			Expect(dbHashInvalidated.Value).To(Equal(test_data.HashInvalidatedModel.Value))
			Expect(dbHashInvalidated.RegistrationDate).To(Equal(test_data.HashInvalidatedModel.RegistrationDate))
			Expect(dbHashInvalidated.LogIndex).To(Equal(test_data.HashInvalidatedModel.LogIndex))
			Expect(dbHashInvalidated.TransactionIndex).To(Equal(test_data.HashInvalidatedModel.TransactionIndex))
			Expect(dbHashInvalidated.Raw).To(MatchJSON(test_data.HashInvalidatedModel.Raw))
		})
	})

	Describe("MarkHeaderChecked", func() {
		inputs := shared_behaviors.MarkedHeaderCheckedBehaviorInputs{
			CheckedHeaderColumnName: constants.HashInvalidatedChecked,
			Repository:              &hashInvalidatedRepository,
		}

		shared_behaviors.SharedRepositoryMarkHeaderCheckedBehaviors(&inputs)
	})
})
