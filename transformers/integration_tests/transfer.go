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
	"github.com/vulcanize/ens_transformers/transformers/registry/transfer"
	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
	"github.com/vulcanize/ens_transformers/transformers/test_data"
)

var testTransferConfig = transformer.TransformerConfig{
	TransformerName:     constants.TransferLabel,
	ContractAddresses:   []string{test_data.RegistryAddress},
	ContractAbi:         test_data.RegistryAbi,
	Topic:               test_data.TransferSignature,
	StartingBlockNumber: 0,
	EndingBlockNumber:   -1,
}

var _ = Describe("Transfer Transformer", func() {
	It("fetches and transforms a Transfer event from mainnet chain", func() {
		blockNumber := int64(8956422)
		config := testTransferConfig
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
			Converter:  &transfer.TransferConverter{},
			Repository: &transfer.TransferRepository{},
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

		var dbResult []transfer.TransferModel
		err = db.Select(&dbResult, `SELECT art, iart, ilk, ink, nflip, tab, urn from maker.transfer`)
		Expect(err).NotTo(HaveOccurred())
		Expect(len(dbResult)).To(Equal(1))
		res := dbResult[0]
		Expect(res.Node).To(Equal("0000"))
		Expect(res.Owner).To(Equal(""))
		Expect(res.TransactionIndex).To(Equal(0))
		Expect(res.LogIndex).To(Equal(0))
	})

	It("unpacks an event log", func() {
		address := common.HexToAddress(test_data.RegistryAddress)
		abi, err := geth.ParseAbi(test_data.RegistryAbi)
		Expect(err).NotTo(HaveOccurred())

		contract := bind.NewBoundContract(address, abi, nil, nil, nil)
		entity := &transfer.TransferEntity{}

		var eventLog = test_data.EthTransferLog

		err = contract.UnpackLog(entity, "Transfer", eventLog)
		Expect(err).NotTo(HaveOccurred())

		expectedEntity := test_data.TransferEntity
		Expect(entity.Node).To(Equal(expectedEntity.Node))
		Expect(entity.Owner).To(Equal(expectedEntity.Owner))
	})

	It("rechecks header for transfer event", func() {
		blockNumber := int64(8956422)
		config := testTransferConfig
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
			Converter:  &transfer.TransferConverter{},
			Repository: &transfer.TransferRepository{},
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

		var transferChecked []int
		err = db.Select(&transferChecked, `SELECT ens_transfer_checked FROM public.checked_headers WHERE header_id = $1`, headerID)
		Expect(err).NotTo(HaveOccurred())

		Expect(transferChecked[0]).To(Equal(2))
	})

	It("unpacks an event log", func() {
		address := common.HexToAddress(test_data.RegistryAddress)
		abi, err := geth.ParseAbi(test_data.RegistryAbi)
		Expect(err).NotTo(HaveOccurred())

		contract := bind.NewBoundContract(address, abi, nil, nil, nil)
		entity := &transfer.TransferEntity{}

		var eventLog = test_data.EthTransferLog

		err = contract.UnpackLog(entity, "Transfer", eventLog)
		Expect(err).NotTo(HaveOccurred())

		expectedEntity := test_data.TransferEntity
		Expect(entity.Node).To(Equal(expectedEntity.Node))
		Expect(entity.Owner).To(Equal(expectedEntity.Owner))
	})
})
