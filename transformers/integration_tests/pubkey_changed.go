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
	"github.com/vulcanize/ens_transformers/transformers/resolver/pubkey_changed"
	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
	"github.com/vulcanize/ens_transformers/transformers/test_data"
	c2 "github.com/vulcanize/vulcanizedb/libraries/shared/constants"
	"github.com/vulcanize/vulcanizedb/libraries/shared/factories/event"
	fetch "github.com/vulcanize/vulcanizedb/libraries/shared/fetcher"
	"github.com/vulcanize/vulcanizedb/libraries/shared/transformer"
)

var testPubkeyChangedConfig = transformer.EventTransformerConfig{
	TransformerName:     constants.PubkeyChangedLabel,
	ContractAddresses:   []string{test_data.ResolverAddress},
	ContractAbi:         test_data.CompleteResolverAbi,
	Topic:               test_data.PubkeyChangedSignature,
	StartingBlockNumber: 0,
	EndingBlockNumber:   -1,
}

var _ = Describe("PubkeyChanged Transformer", func() {
	XIt("fetches and transforms a PubkeyChanged event from mainnet chain", func() {
		blockNumber := int64(7341716)
		config := testPubkeyChangedConfig
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
			Converter:  &pubkey_changed.PubkeyChangedConverter{},
			Repository: &pubkey_changed.PubkeyChangedRepository{},
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

		var dbResult []pubkey_changed.PubkeyChangedModel
		err = db.Select(&dbResult, `SELECT resolver, node, x, y FROM ens.pubkey_changed`)
		Expect(err).NotTo(HaveOccurred())
		Expect(len(dbResult)).To(Equal(1))
		res := dbResult[0]
		Expect(res.X).To(Equal("0x952e25626ff3bb77e2c896c259362efcb621b96aed268faacfbd1e6d7a539f9b"))
		Expect(res.Y).To(Equal("0x3102b2896ee73d9b52ff39c72f007b46b38d8e9c142dcaf0648d9b46fb81cd77"))
		Expect(res.Resolver).To(Equal("0x5FfC014343cd971B7eb70732021E26C35B744cc4"))
		Expect(res.Node).To(Equal("0x9db381b3596a1d14e5afa113274537bd5043c8a6a84d3df7cc791a00e408231d"))
	})

	It("unpacks an event log", func() {
		converter := pubkey_changed.PubkeyChangedConverter{}
		var eventLog = test_data.EthPubkeyChangedLog
		entities, err := converter.ToEntities(test_data.CompleteResolverAbi, []types.Log{eventLog})
		Expect(err).NotTo(HaveOccurred())
		Expect(len(entities)).To(Equal(1))
		entity, ok := entities[0].(pubkey_changed.PubkeyChangedEntity)
		Expect(ok).To(Equal(true))
		expectedEntity := test_data.PubkeyChangedEntity
		Expect(entity.X).To(Equal(expectedEntity.X))
		Expect(entity.Y).To(Equal(expectedEntity.Y))
		Expect(entity.Node).To(Equal(expectedEntity.Node))
		Expect(entity.Resolver).To(Equal(expectedEntity.Resolver))
	})

	XIt("rechecks header for pubkey_changed event", func() {
		blockNumber := int64(7341716)
		config := testPubkeyChangedConfig
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
			Converter:  &pubkey_changed.PubkeyChangedConverter{},
			Repository: &pubkey_changed.PubkeyChangedRepository{},
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

		var pubkey_changedChecked []int
		err = db.Select(&pubkey_changedChecked, `SELECT pubkey_changed_checked FROM public.checked_headers WHERE header_id = $1`, headerID)
		Expect(err).NotTo(HaveOccurred())

		Expect(pubkey_changedChecked[0]).To(Equal(2))
	})
})
