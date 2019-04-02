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

package new_bid_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"

	"github.com/vulcanize/ens_transformers/test_config"
	"github.com/vulcanize/ens_transformers/transformers/registar/new_bid"
	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
	"github.com/vulcanize/ens_transformers/transformers/test_data"
	"github.com/vulcanize/ens_transformers/transformers/test_data/shared_behaviors"
)

var _ = Describe("NewBid repository", func() {
	var (
		hashInvalidatedRepository new_bid.NewBidRepository
		db                        *postgres.DB
	)

	BeforeEach(func() {
		db = test_config.NewTestDB(test_config.NewTestNode())
		test_config.CleanTestDB(db)
		hashInvalidatedRepository = new_bid.NewBidRepository{}
		hashInvalidatedRepository.SetDB(db)
	})

	Describe("Create", func() {
		modelWithDifferentLogIdx := test_data.NewBidModel
		modelWithDifferentLogIdx.LogIndex++
		inputs := shared_behaviors.CreateBehaviorInputs{
			CheckedHeaderColumnName:  constants.NewBidChecked,
			LogEventTableName:        "ens.new_bid",
			TestModel:                test_data.NewBidModel,
			ModelWithDifferentLogIdx: modelWithDifferentLogIdx,
			Repository:               &hashInvalidatedRepository,
		}

		shared_behaviors.SharedRepositoryCreateBehaviors(&inputs)

		It("persists a new_bid record", func() {
			headerRepository := repositories.NewHeaderRepository(db)
			headerID, err := headerRepository.CreateOrUpdateHeader(fakes.FakeHeader)
			Expect(err).NotTo(HaveOccurred())

			err = hashInvalidatedRepository.Create(headerID, []interface{}{test_data.NewBidModel})

			Expect(err).NotTo(HaveOccurred())
			var dbNewBid new_bid.NewBidModel
			err = db.Get(&dbNewBid, `SELECT hash, bidder, deposit, log_idx, tx_idx, raw_log FROM ens.new_bid WHERE header_id = $1`, headerID)
			Expect(err).NotTo(HaveOccurred())
			Expect(dbNewBid.Hash).To(Equal(test_data.NewBidModel.Hash))
			Expect(dbNewBid.Bidder).To(Equal(test_data.NewBidModel.Bidder))
			Expect(dbNewBid.Deposit).To(Equal(test_data.NewBidModel.Deposit))
			Expect(dbNewBid.LogIndex).To(Equal(test_data.NewBidModel.LogIndex))
			Expect(dbNewBid.TransactionIndex).To(Equal(test_data.NewBidModel.TransactionIndex))
			Expect(dbNewBid.Raw).To(MatchJSON(test_data.NewBidModel.Raw))
		})
	})

	Describe("MarkHeaderChecked", func() {
		inputs := shared_behaviors.MarkedHeaderCheckedBehaviorInputs{
			CheckedHeaderColumnName: constants.NewBidChecked,
			Repository:              &hashInvalidatedRepository,
		}

		shared_behaviors.SharedRepositoryMarkHeaderCheckedBehaviors(&inputs)
	})
})
