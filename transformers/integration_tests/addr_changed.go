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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/ens_transformers/test_config"
	"github.com/vulcanize/ens_transformers/transformers/resolver/addr_changed"
	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
	"github.com/vulcanize/ens_transformers/transformers/test_data"
	c2 "github.com/vulcanize/vulcanizedb/libraries/shared/constants"
	"github.com/vulcanize/vulcanizedb/libraries/shared/factories/event"
	fetch "github.com/vulcanize/vulcanizedb/libraries/shared/fetcher"
	"github.com/vulcanize/vulcanizedb/libraries/shared/transformer"
)

var testAddrChangedConfig = transformer.EventTransformerConfig{
	TransformerName:     constants.AddrChangedLabel,
	ContractAddresses:   []string{test_data.ResolverAddress},
	ContractAbi:         test_data.CompleteResolverAbi,
	Topic:               test_data.AddrChangedSignature,
	StartingBlockNumber: 0,
	EndingBlockNumber:   -1,
}

var _ = Describe("AddrChanged Transformer", func() {
	It("fetches and transforms a AddrChanged event from mainnet chain", func() {
		blockNumber := int64(7347406)
		config := testAddrChangedConfig
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
			Converter:  &addr_changed.AddrChangedConverter{},
			Repository: &addr_changed.AddrChangedRepository{},
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

		var dbResult []addr_changed.AddrChangedModel
		err = db.Select(&dbResult, `SELECT resolver, node, address FROM ens.addr_changed`)
		Expect(err).NotTo(HaveOccurred())
		Expect(len(dbResult)).To(Equal(1))
		res := dbResult[0]
		Expect(res.Node).To(Equal("0x5959d9e59dbcefc34e5c15b3e39af59c29b4aaaea04b56b3ac7dd3f6b0f95a4a"))
		Expect(res.Address).To(Equal("0x91de6EF9B3998609aE0B44F4c1ea51A8E3a1c208"))
		Expect(res.Resolver).To(Equal("0x1da022710dF5002339274AaDEe8D58218e9D6AB5"))
	})

	It("unpacks an event log", func() {
		converter := addr_changed.AddrChangedConverter{}
		var eventLog = test_data.EthAddrChangedLog
		entities, err := converter.ToEntities(test_data.CompleteResolverAbi, []types.Log{eventLog})
		Expect(err).NotTo(HaveOccurred())
		Expect(len(entities)).To(Equal(1))
		entity, ok := entities[0].(addr_changed.AddrChangedEntity)
		Expect(ok).To(Equal(true))
		expectedEntity := test_data.AddrChangedEntity
		Expect(entity.Node).To(Equal(expectedEntity.Node))
		Expect(entity.Resolver).To(Equal(expectedEntity.Resolver))
		Expect(entity.A).To(Equal(expectedEntity.A))
	})

	It("rechecks header for addr_changed event", func() {
		blockNumber := int64(7347406)
		config := testAddrChangedConfig
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
			Converter:  &addr_changed.AddrChangedConverter{},
			Repository: &addr_changed.AddrChangedRepository{},
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

		var addr_changedChecked []int
		err = db.Select(&addr_changedChecked, `SELECT addr_changed_checked FROM public.checked_headers WHERE header_id = $1`, headerID)
		Expect(err).NotTo(HaveOccurred())

		Expect(addr_changedChecked[0]).To(Equal(2))
	})
})
