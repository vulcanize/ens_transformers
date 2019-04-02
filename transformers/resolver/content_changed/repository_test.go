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

package content_changed_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"

	"github.com/vulcanize/ens_transformers/test_config"
	"github.com/vulcanize/ens_transformers/transformers/resolver/content_changed"
	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
	"github.com/vulcanize/ens_transformers/transformers/test_data"
	"github.com/vulcanize/ens_transformers/transformers/test_data/shared_behaviors"
)

var _ = Describe("ContentChanged repository", func() {
	var (
		hashInvalidatedRepository content_changed.ContentChangedRepository
		db                        *postgres.DB
	)

	BeforeEach(func() {
		db = test_config.NewTestDB(test_config.NewTestNode())
		test_config.CleanTestDB(db)
		hashInvalidatedRepository = content_changed.ContentChangedRepository{}
		hashInvalidatedRepository.SetDB(db)
	})

	Describe("Create", func() {
		modelWithDifferentLogIdx := test_data.ContentChangedModel
		modelWithDifferentLogIdx.LogIndex++
		inputs := shared_behaviors.CreateBehaviorInputs{
			CheckedHeaderColumnName:  constants.ContentChangedChecked,
			LogEventTableName:        "ens.content_changed",
			TestModel:                test_data.ContentChangedModel,
			ModelWithDifferentLogIdx: modelWithDifferentLogIdx,
			Repository:               &hashInvalidatedRepository,
		}

		shared_behaviors.SharedRepositoryCreateBehaviors(&inputs)

		It("persists a content_changed record", func() {
			headerRepository := repositories.NewHeaderRepository(db)
			headerID, err := headerRepository.CreateOrUpdateHeader(fakes.FakeHeader)
			Expect(err).NotTo(HaveOccurred())

			err = hashInvalidatedRepository.Create(headerID, []interface{}{test_data.ContentChangedModel})

			Expect(err).NotTo(HaveOccurred())
			var dbContentChanged content_changed.ContentChangedModel
			err = db.Get(&dbContentChanged, `SELECT resolver, node, hash, log_idx, tx_idx, raw_log FROM ens.content_changed WHERE header_id = $1`, headerID)
			Expect(err).NotTo(HaveOccurred())
			Expect(dbContentChanged.Resolver).To(Equal(test_data.ContentChangedModel.Resolver))
			Expect(dbContentChanged.Node).To(Equal(test_data.ContentChangedModel.Node))
			Expect(dbContentChanged.Hash).To(Equal(test_data.ContentChangedModel.Hash))
			Expect(dbContentChanged.LogIndex).To(Equal(test_data.ContentChangedModel.LogIndex))
			Expect(dbContentChanged.TransactionIndex).To(Equal(test_data.ContentChangedModel.TransactionIndex))
			Expect(dbContentChanged.Raw).To(MatchJSON(test_data.ContentChangedModel.Raw))
		})
	})

	Describe("MarkHeaderChecked", func() {
		inputs := shared_behaviors.MarkedHeaderCheckedBehaviorInputs{
			CheckedHeaderColumnName: constants.ContentChangedChecked,
			Repository:              &hashInvalidatedRepository,
		}

		shared_behaviors.SharedRepositoryMarkHeaderCheckedBehaviors(&inputs)
	})
})
