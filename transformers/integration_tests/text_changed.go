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

package integration_tests

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	c2 "github.com/vulcanize/vulcanizedb/libraries/shared/constants"
	"github.com/vulcanize/vulcanizedb/libraries/shared/factories"
	fetch "github.com/vulcanize/vulcanizedb/libraries/shared/fetcher"
	"github.com/vulcanize/vulcanizedb/libraries/shared/transformer"
	"github.com/vulcanize/vulcanizedb/pkg/geth"

	"github.com/vulcanize/ens_transformers/test_config"
	"github.com/vulcanize/ens_transformers/transformers/resolver/text_changed"
	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
	"github.com/vulcanize/ens_transformers/transformers/test_data"
)

var testTextChangedConfig = transformer.TransformerConfig{
	TransformerName:     constants.TextChangedLabel,
	ContractAddresses:   []string{test_data.ResolverAddress},
	ContractAbi:         test_data.CompleteResolverAbi,
	Topic:               test_data.TextChangedSignature,
	StartingBlockNumber: 0,
	EndingBlockNumber:   -1,
}

var _ = Describe("TextChanged Transformer", func() {
	It("fetches and transforms a TextChanged event from mainnet chain", func() {
		blockNumber := int64(8956422)
		config := testTextChangedConfig
		config.StartingBlockNumber = blockNumber
		config.EndingBlockNumber = blockNumber

		rpcClient, ethClient, err := getClients(ipc)
		Expect(err).NotTo(HaveOccurred())
		blockChain, err := getBlockChain(rpcClient, ethClient)
		Expect(err).NotTo(HaveOccurred())

		db := test_config.NewTestDB(blockChain.Node())
		test_config.CleanTestDB(db)

		header, err := persistHeader(db, blockNumber, blockChain)
		Expect(err).NotTo(HaveOccurred())

		initializer := factories.Transformer{
			Config:     config,
			Converter:  &text_changed.TextChangedConverter{},
			Repository: &text_changed.TextChangedRepository{},
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

		var dbResult []text_changed.TextChangedModel
		err = db.Select(&dbResult, `SELECT resolver, node, indexed_key, key FROM ens.text_changed`)
		Expect(err).NotTo(HaveOccurred())
		Expect(len(dbResult)).To(Equal(1))
		res := dbResult[0]
		Expect(res.IndexedKey).To(Equal("0000"))
		Expect(res.Key).To(Equal("0000"))
		Expect(res.Resolver).To(Equal(""))
		Expect(res.Node).To(Equal(""))
		Expect(res.TransactionIndex).To(Equal(0))
		Expect(res.LogIndex).To(Equal(0))
	})

	It("unpacks an event log", func() {
		address := common.HexToAddress(test_data.ResolverAddress)
		abi, err := geth.ParseAbi(test_data.CompleteResolverAbi)
		Expect(err).NotTo(HaveOccurred())

		contract := bind.NewBoundContract(address, abi, nil, nil, nil)
		entity := &text_changed.TextChangedEntity{}

		var eventLog = test_data.EthTextChangedLog

		err = contract.UnpackLog(entity, "TextChanged", eventLog)
		Expect(err).NotTo(HaveOccurred())

		expectedEntity := test_data.TextChangedEntity
		Expect(entity.IndexedKey).To(Equal(expectedEntity.IndexedKey))
		Expect(entity.Key).To(Equal(expectedEntity.Key))
		Expect(entity.Node).To(Equal(expectedEntity.Node))
		Expect(entity.Resolver).To(Equal(expectedEntity.Resolver))
	})

	It("rechecks header for text_changed event", func() {
		blockNumber := int64(8956422)
		config := testTextChangedConfig
		config.StartingBlockNumber = blockNumber
		config.EndingBlockNumber = blockNumber

		rpcClient, ethClient, err := getClients(ipc)
		Expect(err).NotTo(HaveOccurred())
		blockChain, err := getBlockChain(rpcClient, ethClient)
		Expect(err).NotTo(HaveOccurred())

		db := test_config.NewTestDB(blockChain.Node())
		test_config.CleanTestDB(db)

		header, err := persistHeader(db, blockNumber, blockChain)
		Expect(err).NotTo(HaveOccurred())

		initializer := factories.Transformer{
			Config:     config,
			Converter:  &text_changed.TextChangedConverter{},
			Repository: &text_changed.TextChangedRepository{},
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

		var text_changedChecked []int
		err = db.Select(&text_changedChecked, `SELECT text_changed_checked FROM public.checked_headers WHERE header_id = $1`, headerID)
		Expect(err).NotTo(HaveOccurred())

		Expect(text_changedChecked[0]).To(Equal(2))
	})

	It("unpacks an event log", func() {
		address := common.HexToAddress(test_data.ResolverAddress)
		abi, err := geth.ParseAbi(test_data.CompleteResolverAbi)
		Expect(err).NotTo(HaveOccurred())

		contract := bind.NewBoundContract(address, abi, nil, nil, nil)
		entity := &text_changed.TextChangedEntity{}

		var eventLog = test_data.EthTextChangedLog

		err = contract.UnpackLog(entity, "TextChanged", eventLog)
		Expect(err).NotTo(HaveOccurred())

		expectedEntity := test_data.TextChangedEntity
		Expect(entity.IndexedKey).To(Equal(expectedEntity.IndexedKey))
		Expect(entity.Key).To(Equal(expectedEntity.Key))
		Expect(entity.Node).To(Equal(expectedEntity.Node))
		Expect(entity.Resolver).To(Equal(expectedEntity.Resolver))
	})
})
