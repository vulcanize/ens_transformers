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

package integration_tests

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/ens_transformers/test_config"
	"github.com/vulcanize/ens_transformers/transformers/registar/hash_invalidated"
	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
	"github.com/vulcanize/ens_transformers/transformers/test_data"
	c2 "github.com/vulcanize/vulcanizedb/libraries/shared/constants"
	"github.com/vulcanize/vulcanizedb/libraries/shared/factories/event"
	fetch "github.com/vulcanize/vulcanizedb/libraries/shared/fetcher"
	"github.com/vulcanize/vulcanizedb/libraries/shared/transformer"
)

var testHashInvalidatedConfig = transformer.EventTransformerConfig{
	TransformerName:     constants.HashInvalidatedLabel,
	ContractAddresses:   []string{test_data.RegistarAddress},
	ContractAbi:         test_data.RegistarAbi,
	Topic:               test_data.HashInvalidatedSignature,
	StartingBlockNumber: 0,
	EndingBlockNumber:   -1,
}

var _ = Describe("HashInvalidated Transformer", func() {
	XIt("fetches and transforms a HashInvalidated event from mainnet chain", func() {
		blockNumber := int64(8956422)
		config := testHashInvalidatedConfig
		config.StartingBlockNumber = blockNumber
		config.EndingBlockNumber = blockNumber

		rpcClient, ethClient, err := getClients(ipc)
		Expect(err).NotTo(HaveOccurred())
		blockChain, err := getBlockChain(rpcClient, ethClient)
		Expect(err).NotTo(HaveOccurred())

		db := test_config.NewTestDB(blockChain.Node())
		defer test_config.CleanTestDB(db)

		header, err := persistHeader(db, blockNumber, blockChain)
		Expect(err).NotTo(HaveOccurred())

		initializer := event.Transformer{
			Config:     config,
			Converter:  &hash_invalidated.HashInvalidatedConverter{},
			Repository: &hash_invalidated.HashInvalidatedRepository{},
		}
		transformer := initializer.NewTransformer(db)

		fetcher := fetch.NewFetcher(blockChain)
		logs, err := fetcher.FetchLogs(
			[]common.Address{common.HexToAddress(config.ContractAddresses[0])},
			[]common.Hash{common.HexToHash(config.Topic)},
			header)
		Expect(err).NotTo(HaveOccurred())

		err = transformer.Execute(logs, header, c2.HeaderMissing)
		Expect(err).NotTo(HaveOccurred())

		var dbResult []hash_invalidated.HashInvalidatedModel
		err = db.Select(&dbResult, `SELECT hash, name, value, registration_date FROM ens.hash_invalidated`)
		Expect(err).NotTo(HaveOccurred())
		Expect(len(dbResult)).To(Equal(1))
		res := dbResult[0]
		Expect(res.Hash).To(Equal("0000"))
		Expect(res.Name).To(Equal(""))
		Expect(res.Value).To(Equal(""))
		Expect(res.RegistrationDate).To(Equal(""))
	})

	It("unpacks an event log", func() {
		converter := hash_invalidated.HashInvalidatedConverter{}
		var eventLog = test_data.EthHashInvalidatedLog
		entities, err := converter.ToEntities(test_data.RegistarAbi, []types.Log{eventLog})
		Expect(err).NotTo(HaveOccurred())
		Expect(len(entities)).To(Equal(1))
		entity, ok := entities[0].(hash_invalidated.HashInvalidatedEntity)
		Expect(ok).To(Equal(true))
		expectedEntity := test_data.HashInvalidatedEntity
		Expect(entity.Hash).To(Equal(expectedEntity.Hash))
		Expect(entity.Name).To(Equal(expectedEntity.Name))
		Expect(entity.Value).To(Equal(expectedEntity.Value))
	})

	XIt("rechecks header for hash_invalidated event", func() {
		blockNumber := int64(8956422)
		config := testHashInvalidatedConfig
		config.StartingBlockNumber = blockNumber
		config.EndingBlockNumber = blockNumber

		rpcClient, ethClient, err := getClients(ipc)
		Expect(err).NotTo(HaveOccurred())
		blockChain, err := getBlockChain(rpcClient, ethClient)
		Expect(err).NotTo(HaveOccurred())

		db := test_config.NewTestDB(blockChain.Node())
		defer test_config.CleanTestDB(db)

		header, err := persistHeader(db, blockNumber, blockChain)
		Expect(err).NotTo(HaveOccurred())

		initializer := event.Transformer{
			Config:     config,
			Converter:  &hash_invalidated.HashInvalidatedConverter{},
			Repository: &hash_invalidated.HashInvalidatedRepository{},
		}
		transformer := initializer.NewTransformer(db)

		fetcher := fetch.NewFetcher(blockChain)
		logs, err := fetcher.FetchLogs(
			[]common.Address{common.HexToAddress(config.ContractAddresses[0])},
			[]common.Hash{common.HexToHash(config.Topic)},
			header)
		Expect(err).NotTo(HaveOccurred())

		err = transformer.Execute(logs, header, c2.HeaderMissing)
		Expect(err).NotTo(HaveOccurred())

		err = transformer.Execute(logs, header, c2.HeaderRecheck)
		Expect(err).NotTo(HaveOccurred())

		var headerID int64
		err = db.Get(&headerID, `SELECT id FROM public.headers WHERE block_number = $1`, blockNumber)
		Expect(err).NotTo(HaveOccurred())

		var hash_invalidatedChecked []int
		err = db.Select(&hash_invalidatedChecked, `SELECT hash_invalidated_checked FROM public.checked_headers WHERE header_id = $1`, headerID)
		Expect(err).NotTo(HaveOccurred())

		Expect(hash_invalidatedChecked[0]).To(Equal(2))
	})
})
