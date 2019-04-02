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
	"github.com/vulcanize/ens_transformers/transformers/registry/new_owner"
	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
	"github.com/vulcanize/ens_transformers/transformers/test_data"
	c2 "github.com/vulcanize/vulcanizedb/libraries/shared/constants"
	"github.com/vulcanize/vulcanizedb/libraries/shared/factories/event"
	fetch "github.com/vulcanize/vulcanizedb/libraries/shared/fetcher"
	"github.com/vulcanize/vulcanizedb/libraries/shared/transformer"
)

var testNewOwnerConfig = transformer.EventTransformerConfig{
	TransformerName:     constants.NewOwnerLabel,
	ContractAddresses:   []string{test_data.RegistryAddress},
	ContractAbi:         test_data.RegistryAbi,
	Topic:               test_data.NewOwnerSignature,
	StartingBlockNumber: 0,
	EndingBlockNumber:   -1,
}

var _ = Describe("NewOwner Transformer", func() {
	It("fetches and transforms a NewOwner event from mainnet chain", func() {
		blockNumber := int64(7483567)
		config := testNewOwnerConfig
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
			Converter:  &new_owner.NewOwnerConverter{},
			Repository: &new_owner.NewOwnerRepository{},
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

		var dbResult []new_owner.NewOwnerModel
		err = db.Select(&dbResult, `SELECT node, label, owner, subnode FROM ens.new_owner`)
		Expect(err).NotTo(HaveOccurred())
		Expect(len(dbResult)).To(Equal(1))
		res := dbResult[0]
		Expect(res.Node).To(Equal("0x79aa1ec377dbfd1edf87d526cd5c116ac6ec4444e23da2a8b8ae0e9db9f46ec9"))
		Expect(res.Label).To(Equal("0xf06ea683845d8994338d98e7296850612317647cdb1db3d1051f24bafdefe9f7"))
		Expect(res.Owner).To(Equal("0xA964ed4077aD3ba1946D118ce90544657bB4003B"))
		Expect(res.Subnode).To(Equal("0xbb87bd9021ba9da3248899e6fdd901a68efb0e15ac691ac9ce5cc88ebcb306de"))
	})

	It("unpacks an event log", func() {
		converter := new_owner.NewOwnerConverter{}
		var eventLog = test_data.EthNewOwnerLog
		entities, err := converter.ToEntities(test_data.RegistryAbi, []types.Log{eventLog})
		Expect(err).NotTo(HaveOccurred())
		Expect(len(entities)).To(Equal(1))
		entity, ok := entities[0].(new_owner.NewOwnerEntity)
		Expect(ok).To(Equal(true))
		expectedEntity := test_data.NewOwnerEntity
		Expect(entity.Node).To(Equal(expectedEntity.Node))
		Expect(entity.Label).To(Equal(expectedEntity.Label))
		Expect(entity.Owner).To(Equal(expectedEntity.Owner))
	})

	It("rechecks header for new_owner event", func() {
		blockNumber := int64(7483567)
		config := testNewOwnerConfig
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
			Converter:  &new_owner.NewOwnerConverter{},
			Repository: &new_owner.NewOwnerRepository{},
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

		var new_ownerChecked []int
		err = db.Select(&new_ownerChecked, `SELECT new_owner_checked FROM public.checked_headers WHERE header_id = $1`, headerID)
		Expect(err).NotTo(HaveOccurred())

		Expect(new_ownerChecked[0]).To(Equal(2))
	})
})
