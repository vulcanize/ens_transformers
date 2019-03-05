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
	"github.com/vulcanize/ens_transformers/transformers/registar/new_bid"
	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
	"github.com/vulcanize/ens_transformers/transformers/test_data"
)

var testNewBidConfig = transformer.TransformerConfig{
	TransformerName:     constants.NewBidLabel,
	ContractAddresses:   []string{test_data.RegistarAddress},
	ContractAbi:         test_data.RegistarAbi,
	Topic:               test_data.NewBidSignature,
	StartingBlockNumber: 0,
	EndingBlockNumber:   -1,
}

var _ = Describe("NewBid Transformer", func() {
	It("fetches and transforms a NewBid event from mainnet chain", func() {
		blockNumber := int64(8956422)
		config := testNewBidConfig
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
			Converter:  &new_bid.NewBidConverter{},
			Repository: &new_bid.NewBidRepository{},
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

		var dbResult []new_bid.NewBidModel
		err = db.Select(&dbResult, `SELECT hash, bidder, deposit FROM ens.new_bid`)
		Expect(err).NotTo(HaveOccurred())
		Expect(len(dbResult)).To(Equal(1))
		res := dbResult[0]
		Expect(res.Hash).To(Equal("0000"))
		Expect(res.Bidder).To(Equal(""))
		Expect(res.Deposit).To(Equal(""))
	})

	It("unpacks an event log", func() {
		address := common.HexToAddress(test_data.RegistarAddress)
		abi, err := geth.ParseAbi(test_data.RegistarAbi)
		Expect(err).NotTo(HaveOccurred())

		contract := bind.NewBoundContract(address, abi, nil, nil, nil)
		entity := &new_bid.NewBidEntity{}

		var eventLog = test_data.EthNewBidLog

		err = contract.UnpackLog(entity, "NewBid", eventLog)
		Expect(err).NotTo(HaveOccurred())

		expectedEntity := test_data.NewBidEntity
		Expect(entity.Hash).To(Equal(expectedEntity.Hash))
		Expect(entity.Bidder).To(Equal(expectedEntity.Bidder))
		Expect(entity.Deposit).To(Equal(expectedEntity.Deposit))
	})

	It("rechecks header for new_bid event", func() {
		blockNumber := int64(8956422)
		config := testNewBidConfig
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
			Converter:  &new_bid.NewBidConverter{},
			Repository: &new_bid.NewBidRepository{},
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

		var new_bidChecked []int
		err = db.Select(&new_bidChecked, `SELECT new_bid_checked FROM public.checked_headers WHERE header_id = $1`, headerID)
		Expect(err).NotTo(HaveOccurred())

		Expect(new_bidChecked[0]).To(Equal(2))
	})

	It("unpacks an event log", func() {
		address := common.HexToAddress(test_data.RegistarAddress)
		abi, err := geth.ParseAbi(test_data.RegistarAbi)
		Expect(err).NotTo(HaveOccurred())

		contract := bind.NewBoundContract(address, abi, nil, nil, nil)
		entity := &new_bid.NewBidEntity{}

		var eventLog = test_data.EthNewBidLog

		err = contract.UnpackLog(entity, "NewBid", eventLog)
		Expect(err).NotTo(HaveOccurred())

		expectedEntity := test_data.NewBidEntity
		Expect(entity.Hash).To(Equal(expectedEntity.Hash))
		Expect(entity.Bidder).To(Equal(expectedEntity.Bidder))
		Expect(entity.Deposit).To(Equal(expectedEntity.Deposit))
	})
})
