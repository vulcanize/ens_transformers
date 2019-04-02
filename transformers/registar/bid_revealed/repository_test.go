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

package bid_revealed_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"

	"github.com/vulcanize/ens_transformers/test_config"
	"github.com/vulcanize/ens_transformers/transformers/registar/bid_revealed"
	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
	"github.com/vulcanize/ens_transformers/transformers/test_data"
	"github.com/vulcanize/ens_transformers/transformers/test_data/shared_behaviors"
)

var _ = Describe("BidRevealed repository", func() {
	var (
		bidRevealedRepository bid_revealed.BidRevealedRepository
		db                    *postgres.DB
	)

	BeforeEach(func() {
		db = test_config.NewTestDB(test_config.NewTestNode())
		test_config.CleanTestDB(db)
		bidRevealedRepository = bid_revealed.BidRevealedRepository{}
		bidRevealedRepository.SetDB(db)
	})

	Describe("Create", func() {
		modelWithDifferentLogIdx := test_data.BidRevealedModel
		modelWithDifferentLogIdx.LogIndex++
		inputs := shared_behaviors.CreateBehaviorInputs{
			CheckedHeaderColumnName:  constants.BidRevealedChecked,
			LogEventTableName:        "ens.bid_revealed",
			TestModel:                test_data.BidRevealedModel,
			ModelWithDifferentLogIdx: modelWithDifferentLogIdx,
			Repository:               &bidRevealedRepository,
		}

		shared_behaviors.SharedRepositoryCreateBehaviors(&inputs)

		It("persists a bid_revealed record", func() {
			headerRepository := repositories.NewHeaderRepository(db)
			headerID, err := headerRepository.CreateOrUpdateHeader(fakes.FakeHeader)
			Expect(err).NotTo(HaveOccurred())

			err = bidRevealedRepository.Create(headerID, []interface{}{test_data.BidRevealedModel})

			Expect(err).NotTo(HaveOccurred())
			var dbBidRevealed bid_revealed.BidRevealedModel
			err = db.Get(&dbBidRevealed, `SELECT hash, owner, value, status, log_idx, tx_idx, raw_log FROM ens.bid_revealed WHERE header_id = $1`, headerID)
			Expect(err).NotTo(HaveOccurred())
			Expect(dbBidRevealed.Hash).To(Equal(test_data.BidRevealedModel.Hash))
			Expect(dbBidRevealed.Owner).To(Equal(test_data.BidRevealedModel.Owner))
			Expect(dbBidRevealed.Value).To(Equal(test_data.BidRevealedModel.Value))
			Expect(dbBidRevealed.Status).To(Equal(test_data.BidRevealedModel.Status))
			Expect(dbBidRevealed.LogIndex).To(Equal(test_data.BidRevealedModel.LogIndex))
			Expect(dbBidRevealed.TransactionIndex).To(Equal(test_data.BidRevealedModel.TransactionIndex))
			Expect(dbBidRevealed.Raw).To(MatchJSON(test_data.BidRevealedModel.Raw))
		})
	})

	Describe("MarkHeaderChecked", func() {
		inputs := shared_behaviors.MarkedHeaderCheckedBehaviorInputs{
			CheckedHeaderColumnName: constants.BidRevealedChecked,
			Repository:              &bidRevealedRepository,
		}

		shared_behaviors.SharedRepositoryMarkHeaderCheckedBehaviors(&inputs)
	})
})
