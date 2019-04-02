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

package pubkey_changed_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"

	"github.com/vulcanize/ens_transformers/test_config"
	"github.com/vulcanize/ens_transformers/transformers/resolver/pubkey_changed"
	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
	"github.com/vulcanize/ens_transformers/transformers/test_data"
	"github.com/vulcanize/ens_transformers/transformers/test_data/shared_behaviors"
)

var _ = Describe("PubkeyChanged repository", func() {
	var (
		hashInvalidatedRepository pubkey_changed.PubkeyChangedRepository
		db                        *postgres.DB
	)

	BeforeEach(func() {
		db = test_config.NewTestDB(test_config.NewTestNode())
		test_config.CleanTestDB(db)
		hashInvalidatedRepository = pubkey_changed.PubkeyChangedRepository{}
		hashInvalidatedRepository.SetDB(db)
	})

	Describe("Create", func() {
		modelWithDifferentLogIdx := test_data.PubkeyChangedModel
		modelWithDifferentLogIdx.LogIndex++
		inputs := shared_behaviors.CreateBehaviorInputs{
			CheckedHeaderColumnName:  constants.PubkeyChangedChecked,
			LogEventTableName:        "ens.pubkey_changed",
			TestModel:                test_data.PubkeyChangedModel,
			ModelWithDifferentLogIdx: modelWithDifferentLogIdx,
			Repository:               &hashInvalidatedRepository,
		}

		shared_behaviors.SharedRepositoryCreateBehaviors(&inputs)

		It("persists a pubkey_changed record", func() {
			headerRepository := repositories.NewHeaderRepository(db)
			headerID, err := headerRepository.CreateOrUpdateHeader(fakes.FakeHeader)
			Expect(err).NotTo(HaveOccurred())

			err = hashInvalidatedRepository.Create(headerID, []interface{}{test_data.PubkeyChangedModel})

			Expect(err).NotTo(HaveOccurred())
			var dbPubkeyChanged pubkey_changed.PubkeyChangedModel
			err = db.Get(&dbPubkeyChanged, `SELECT resolver, node, x, y, log_idx, tx_idx, raw_log FROM ens.pubkey_changed WHERE header_id = $1`, headerID)
			Expect(err).NotTo(HaveOccurred())
			Expect(dbPubkeyChanged.Resolver).To(Equal(test_data.PubkeyChangedModel.Resolver))
			Expect(dbPubkeyChanged.Node).To(Equal(test_data.PubkeyChangedModel.Node))
			Expect(dbPubkeyChanged.X).To(Equal(test_data.PubkeyChangedModel.X))
			Expect(dbPubkeyChanged.Y).To(Equal(test_data.PubkeyChangedModel.Y))
			Expect(dbPubkeyChanged.LogIndex).To(Equal(test_data.PubkeyChangedModel.LogIndex))
			Expect(dbPubkeyChanged.TransactionIndex).To(Equal(test_data.PubkeyChangedModel.TransactionIndex))
			Expect(dbPubkeyChanged.Raw).To(MatchJSON(test_data.PubkeyChangedModel.Raw))
		})
	})

	Describe("MarkHeaderChecked", func() {
		inputs := shared_behaviors.MarkedHeaderCheckedBehaviorInputs{
			CheckedHeaderColumnName: constants.PubkeyChangedChecked,
			Repository:              &hashInvalidatedRepository,
		}

		shared_behaviors.SharedRepositoryMarkHeaderCheckedBehaviors(&inputs)
	})
})
