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
	"github.com/vulcanize/ens_transformers/transformers/registar/hash_registered"
	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
	"github.com/vulcanize/ens_transformers/transformers/test_data"
	c2 "github.com/vulcanize/vulcanizedb/libraries/shared/constants"
	"github.com/vulcanize/vulcanizedb/libraries/shared/factories/event"
	fetch "github.com/vulcanize/vulcanizedb/libraries/shared/fetcher"
	"github.com/vulcanize/vulcanizedb/libraries/shared/transformer"
)

var testHashRegisteredConfig = transformer.EventTransformerConfig{
	TransformerName:     constants.HashRegisteredLabel,
	ContractAddresses:   []string{test_data.RegistarAddress},
	ContractAbi:         test_data.RegistarAbi,
	Topic:               test_data.HashRegisteredSignature,
	StartingBlockNumber: 0,
	EndingBlockNumber:   -1,
}

var _ = Describe("HashRegistered Transformer", func() {
	XIt("fetches and transforms a HashRegistered event from mainnet chain", func() {
		blockNumber := int64(7381651)
		config := testHashRegisteredConfig
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
			Converter:  &hash_registered.HashRegisteredConverter{},
			Repository: &hash_registered.HashRegisteredRepository{},
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

		var dbResult []hash_registered.HashRegisteredModel
		err = db.Select(&dbResult, `SELECT hash, owner, value, registration_date FROM ens.hash_registered`)
		Expect(err).NotTo(HaveOccurred())
		Expect(len(dbResult)).To(Equal(1))
		res := dbResult[0]
		Expect(res.Hash).To(Equal("0x39a3c43e563f9353204499f192e02970512be1c963e9116967ddf8f75dc6e44c"))
		Expect(res.Owner).To(Equal("0x000000000000000000000000efb0328d3c0a8930f0a9859ff3f64186105b99b0"))
		Expect(res.Value).To(Equal("10000000000000000"))
		Expect(res.RegistrationDate).To(Equal("1552760424"))
	})

	It("unpacks an event log", func() {
		converter := hash_registered.HashRegisteredConverter{}
		var eventLog = test_data.EthHashRegisteredLog
		entities, err := converter.ToEntities(test_data.RegistarAbi, []types.Log{eventLog})
		Expect(err).NotTo(HaveOccurred())
		Expect(len(entities)).To(Equal(1))
		entity, ok := entities[0].(hash_registered.HashRegisteredEntity)
		Expect(ok).To(Equal(true))
		expectedEntity := test_data.HashRegisteredEntity
		Expect(entity.Hash).To(Equal(expectedEntity.Hash))
		Expect(entity.Owner).To(Equal(expectedEntity.Owner))
		Expect(entity.Value).To(Equal(expectedEntity.Value))
		Expect(entity.RegistrationDate).To(Equal(expectedEntity.RegistrationDate))
	})

	XIt("rechecks header for hash_registered event", func() {
		blockNumber := int64(7381651)
		config := testHashRegisteredConfig
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
			Converter:  &hash_registered.HashRegisteredConverter{},
			Repository: &hash_registered.HashRegisteredRepository{},
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

		var hash_registeredChecked []int
		err = db.Select(&hash_registeredChecked, `SELECT hash_registered_checked FROM public.checked_headers WHERE header_id = $1`, headerID)
		Expect(err).NotTo(HaveOccurred())

		Expect(hash_registeredChecked[0]).To(Equal(2))
	})
})
