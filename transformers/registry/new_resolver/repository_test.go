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

package new_resolver_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"

	"github.com/vulcanize/ens_transformers/test_config"
	"github.com/vulcanize/ens_transformers/transformers/registry/new_resolver"
	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
	"github.com/vulcanize/ens_transformers/transformers/test_data"
	"github.com/vulcanize/ens_transformers/transformers/test_data/shared_behaviors"
)

var _ = Describe("NewResolver repository", func() {
	var (
		hashInvalidatedRepository new_resolver.NewResolverRepository
		db                        *postgres.DB
	)

	BeforeEach(func() {
		db = test_config.NewTestDB(test_config.NewTestNode())
		test_config.CleanTestDB(db)
		hashInvalidatedRepository = new_resolver.NewResolverRepository{}
		hashInvalidatedRepository.SetDB(db)
	})

	Describe("Create", func() {
		modelWithDifferentLogIdx := test_data.NewResolverModel
		modelWithDifferentLogIdx.LogIndex++
		inputs := shared_behaviors.CreateBehaviorInputs{
			CheckedHeaderColumnName:  constants.NewResolverChecked,
			LogEventTableName:        "ens.new_resolver",
			TestModel:                test_data.NewResolverModel,
			ModelWithDifferentLogIdx: modelWithDifferentLogIdx,
			Repository:               &hashInvalidatedRepository,
		}

		shared_behaviors.SharedRepositoryCreateBehaviors(&inputs)

		It("persists a new_resolver record", func() {
			headerRepository := repositories.NewHeaderRepository(db)
			headerID, err := headerRepository.CreateOrUpdateHeader(fakes.FakeHeader)
			Expect(err).NotTo(HaveOccurred())

			err = hashInvalidatedRepository.Create(headerID, []interface{}{test_data.NewResolverModel})

			Expect(err).NotTo(HaveOccurred())
			var dbNewResolver new_resolver.NewResolverModel
			err = db.Get(&dbNewResolver, `SELECT node, resolver, log_idx, tx_idx, raw_log FROM ens.new_resolver WHERE header_id = $1`, headerID)
			Expect(err).NotTo(HaveOccurred())
			Expect(dbNewResolver.Node).To(Equal(test_data.NewResolverModel.Node))
			Expect(dbNewResolver.Resolver).To(Equal(test_data.NewResolverModel.Resolver))
			Expect(dbNewResolver.LogIndex).To(Equal(test_data.NewResolverModel.LogIndex))
			Expect(dbNewResolver.TransactionIndex).To(Equal(test_data.NewResolverModel.TransactionIndex))
			Expect(dbNewResolver.Raw).To(MatchJSON(test_data.NewResolverModel.Raw))
		})
	})

	Describe("MarkHeaderChecked", func() {
		inputs := shared_behaviors.MarkedHeaderCheckedBehaviorInputs{
			CheckedHeaderColumnName: constants.NewResolverChecked,
			Repository:              &hashInvalidatedRepository,
		}

		shared_behaviors.SharedRepositoryMarkHeaderCheckedBehaviors(&inputs)
	})
})
