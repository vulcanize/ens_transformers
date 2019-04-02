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

package auction_started_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"

	"github.com/vulcanize/ens_transformers/test_config"
	"github.com/vulcanize/ens_transformers/transformers/registar/auction_started"
	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
	"github.com/vulcanize/ens_transformers/transformers/test_data"
	"github.com/vulcanize/ens_transformers/transformers/test_data/shared_behaviors"
)

var _ = Describe("AuctionStarted repository", func() {
	var (
		auctionStartedRepository auction_started.AuctionStartedRepository
		db                       *postgres.DB
	)

	BeforeEach(func() {
		db = test_config.NewTestDB(test_config.NewTestNode())
		test_config.CleanTestDB(db)
		auctionStartedRepository = auction_started.AuctionStartedRepository{}
		auctionStartedRepository.SetDB(db)
	})

	Describe("Create", func() {
		modelWithDifferentLogIdx := test_data.AuctionStartedModel
		modelWithDifferentLogIdx.LogIndex++
		inputs := shared_behaviors.CreateBehaviorInputs{
			CheckedHeaderColumnName:  constants.AuctionStartedChecked,
			LogEventTableName:        "ens.auction_started",
			TestModel:                test_data.AuctionStartedModel,
			ModelWithDifferentLogIdx: modelWithDifferentLogIdx,
			Repository:               &auctionStartedRepository,
		}

		shared_behaviors.SharedRepositoryCreateBehaviors(&inputs)

		It("persists a auction_started record", func() {
			headerRepository := repositories.NewHeaderRepository(db)
			headerID, err := headerRepository.CreateOrUpdateHeader(fakes.FakeHeader)
			Expect(err).NotTo(HaveOccurred())

			err = auctionStartedRepository.Create(headerID, []interface{}{test_data.AuctionStartedModel})

			Expect(err).NotTo(HaveOccurred())
			var dbAuctionStarted auction_started.AuctionStartedModel
			err = db.Get(&dbAuctionStarted, `SELECT hash, registration_date, log_idx, tx_idx, raw_log FROM ens.auction_started WHERE header_id = $1`, headerID)
			Expect(err).NotTo(HaveOccurred())
			Expect(dbAuctionStarted.Hash).To(Equal(test_data.AuctionStartedModel.Hash))
			Expect(dbAuctionStarted.RegistrationDate).To(Equal(test_data.AuctionStartedModel.RegistrationDate))
			Expect(dbAuctionStarted.LogIndex).To(Equal(test_data.AuctionStartedModel.LogIndex))
			Expect(dbAuctionStarted.TransactionIndex).To(Equal(test_data.AuctionStartedModel.TransactionIndex))
			Expect(dbAuctionStarted.Raw).To(MatchJSON(test_data.AuctionStartedModel.Raw))
		})
	})

	Describe("MarkHeaderChecked", func() {
		inputs := shared_behaviors.MarkedHeaderCheckedBehaviorInputs{
			CheckedHeaderColumnName: constants.AuctionStartedChecked,
			Repository:              &auctionStartedRepository,
		}

		shared_behaviors.SharedRepositoryMarkHeaderCheckedBehaviors(&inputs)
	})
})
