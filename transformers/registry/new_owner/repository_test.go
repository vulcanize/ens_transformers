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

package new_owner_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"

	"github.com/vulcanize/ens_transformers/test_config"
	"github.com/vulcanize/ens_transformers/transformers/registry/new_owner"
	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
	"github.com/vulcanize/ens_transformers/transformers/test_data"
	"github.com/vulcanize/ens_transformers/transformers/test_data/shared_behaviors"
)

var _ = Describe("NewOwner repository", func() {
	var (
		newOwnerRepository new_owner.NewOwnerRepository
		db                 *postgres.DB
	)

	BeforeEach(func() {
		db = test_config.NewTestDB(test_config.NewTestNode())
		test_config.CleanTestDB(db)
		newOwnerRepository = new_owner.NewOwnerRepository{}
		newOwnerRepository.SetDB(db)
	})

	Describe("Create", func() {
		modelWithDifferentLogIdx := test_data.NewOwnerModel
		modelWithDifferentLogIdx.LogIndex++
		inputs := shared_behaviors.CreateBehaviorInputs{
			CheckedHeaderColumnName:  constants.NewOwnerChecked,
			LogEventTableName:        "ens.new_owner",
			TestModel:                test_data.NewOwnerModel,
			ModelWithDifferentLogIdx: modelWithDifferentLogIdx,
			Repository:               &newOwnerRepository,
		}

		shared_behaviors.SharedRepositoryCreateBehaviors(&inputs)

		It("persists a new_owner record", func() {
			headerRepository := repositories.NewHeaderRepository(db)
			headerID, err := headerRepository.CreateOrUpdateHeader(fakes.FakeHeader)
			Expect(err).NotTo(HaveOccurred())

			err = newOwnerRepository.Create(headerID, []interface{}{test_data.NewOwnerModel})

			Expect(err).NotTo(HaveOccurred())
			var dbNewOwner new_owner.NewOwnerModel
			err = db.Get(&dbNewOwner, `SELECT node, label, owner, log_idx, tx_idx, raw_log FROM ens.new_owner WHERE header_id = $1`, headerID)
			Expect(err).NotTo(HaveOccurred())
			Expect(dbNewOwner.Node).To(Equal(test_data.NewOwnerModel.Node))
			Expect(dbNewOwner.Label).To(Equal(test_data.NewOwnerModel.Label))
			Expect(dbNewOwner.Owner).To(Equal(test_data.NewOwnerModel.Owner))
			Expect(dbNewOwner.LogIndex).To(Equal(test_data.NewOwnerModel.LogIndex))
			Expect(dbNewOwner.TransactionIndex).To(Equal(test_data.NewOwnerModel.TransactionIndex))
			Expect(dbNewOwner.Raw).To(MatchJSON(test_data.NewOwnerModel.Raw))
		})
	})

	Describe("MarkHeaderChecked", func() {
		inputs := shared_behaviors.MarkedHeaderCheckedBehaviorInputs{
			CheckedHeaderColumnName: constants.NewOwnerChecked,
			Repository:              &newOwnerRepository,
		}

		shared_behaviors.SharedRepositoryMarkHeaderCheckedBehaviors(&inputs)
	})
})
