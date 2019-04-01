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

package hash_registered_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"

	"github.com/vulcanize/ens_transformers/test_config"
	"github.com/vulcanize/ens_transformers/transformers/registar/hash_registered"
	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
	"github.com/vulcanize/ens_transformers/transformers/test_data"
	"github.com/vulcanize/ens_transformers/transformers/test_data/shared_behaviors"
)

var _ = Describe("HashRegistered repository", func() {
	var (
		hashInvalidatedRepository hash_registered.HashRegisteredRepository
		db                        *postgres.DB
	)

	BeforeEach(func() {
		db = test_config.NewTestDB(test_config.NewTestNode())
		test_config.CleanTestDB(db)
		hashInvalidatedRepository = hash_registered.HashRegisteredRepository{}
		hashInvalidatedRepository.SetDB(db)
	})

	Describe("Create", func() {
		modelWithDifferentLogIdx := test_data.HashRegisteredModel
		modelWithDifferentLogIdx.LogIndex++
		inputs := shared_behaviors.CreateBehaviorInputs{
			CheckedHeaderColumnName:  constants.HashRegisteredChecked,
			LogEventTableName:        "ens.hash_registered",
			TestModel:                test_data.HashRegisteredModel,
			ModelWithDifferentLogIdx: modelWithDifferentLogIdx,
			Repository:               &hashInvalidatedRepository,
		}

		shared_behaviors.SharedRepositoryCreateBehaviors(&inputs)

		It("persists a hash_registered record", func() {
			headerRepository := repositories.NewHeaderRepository(db)
			headerID, err := headerRepository.CreateOrUpdateHeader(fakes.FakeHeader)
			Expect(err).NotTo(HaveOccurred())

			err = hashInvalidatedRepository.Create(headerID, []interface{}{test_data.HashRegisteredModel})

			Expect(err).NotTo(HaveOccurred())
			var dbHashRegistered hash_registered.HashRegisteredModel
			err = db.Get(&dbHashRegistered, `SELECT hash, owner, value, registration_date, log_idx, tx_idx, raw_log FROM ens.hash_registered WHERE header_id = $1`, headerID)
			Expect(err).NotTo(HaveOccurred())
			Expect(dbHashRegistered.Hash).To(Equal(test_data.HashRegisteredModel.Hash))
			Expect(dbHashRegistered.Owner).To(Equal(test_data.HashRegisteredModel.Owner))
			Expect(dbHashRegistered.Value).To(Equal(test_data.HashRegisteredModel.Value))
			Expect(dbHashRegistered.RegistrationDate).To(Equal(test_data.HashRegisteredModel.RegistrationDate))
			Expect(dbHashRegistered.LogIndex).To(Equal(test_data.HashRegisteredModel.LogIndex))
			Expect(dbHashRegistered.TransactionIndex).To(Equal(test_data.HashRegisteredModel.TransactionIndex))
			Expect(dbHashRegistered.Raw).To(MatchJSON(test_data.HashRegisteredModel.Raw))
		})
	})

	Describe("MarkHeaderChecked", func() {
		inputs := shared_behaviors.MarkedHeaderCheckedBehaviorInputs{
			CheckedHeaderColumnName: constants.HashRegisteredChecked,
			Repository:              &hashInvalidatedRepository,
		}

		shared_behaviors.SharedRepositoryMarkHeaderCheckedBehaviors(&inputs)
	})
})
