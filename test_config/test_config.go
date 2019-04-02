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

package test_config

import (
	"fmt"
	"os"

	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/vulcanize/vulcanizedb/pkg/config"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
)

var TestConfig *viper.Viper
var DBConfig config.Database
var TestClient config.Client
var Infura *viper.Viper
var InfuraClient config.Client
var ABIFilePath string

func init() {
	setTestConfig()
	setInfuraConfig()
	setABIPath()
}

func setTestConfig() {
	TestConfig = viper.New()
	TestConfig.SetConfigName("private")
	TestConfig.AddConfigPath("$GOPATH/src/github.com/vulcanize/ens_transformers/environments/")
	err := TestConfig.ReadInConfig()
	ipc := TestConfig.GetString("client.ipcPath")
	if err != nil {
		log.Fatal(err)
	}
	hn := TestConfig.GetString("database.hostname")
	port := TestConfig.GetInt("database.port")
	name := TestConfig.GetString("database.name")
	DBConfig = config.Database{
		Hostname: hn,
		Name:     name,
		Port:     port,
	}
	TestClient = config.Client{
		IPCPath: ipc,
	}
}

func setInfuraConfig() {
	Infura = viper.New()
	Infura.SetConfigName("infura")
	Infura.AddConfigPath("$GOPATH/src/github.com/vulcanize/ens_transformers/environments/")
	err := Infura.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	ipc := Infura.GetString("client.ipcpath")

	// If we don't have an ipc path in the config file, check the env variable
	if ipc == "" {
		Infura.BindEnv("url", "INFURA_URL")
		ipc = Infura.GetString("url")
	}
	if ipc == "" {
		log.Fatal("infura.toml IPC path or $INFURA_URL env variable need to be set")
	}

	InfuraClient = config.Client{
		IPCPath: ipc,
	}
}

func setABIPath() {
	gp := os.Getenv("GOPATH")
	ABIFilePath = gp + "/src/github.com/vulcanize/vulcanizedb/pkg/geth/testing/"
}

func NewTestDB(node core.Node) *postgres.DB {
	db, err := postgres.NewDB(DBConfig, node)
	if err != nil {
		panic(fmt.Sprintf("Could not create new test db: %v", err))
	}
	return db
}

func CleanTestDB(db *postgres.DB) {
	db.MustExec("DELETE FROM blocks")
	db.MustExec("DELETE FROM headers")
	db.MustExec("DELETE FROM checked_headers")
	db.MustExec("DELETE FROM log_filters")
	db.MustExec("DELETE FROM logs")
	db.MustExec("DELETE FROM receipts")
	db.MustExec("DELETE FROM light_sync_transactions")
	db.MustExec("DELETE FROM full_sync_transactions")
	db.MustExec("DELETE FROM watched_contracts")
	db.MustExec("DELETE FROM ens.auction_started")
	db.MustExec("DELETE FROM ens.bid_revealed")
	db.MustExec("DELETE FROM ens.hash_invalidated")
	db.MustExec("DELETE FROM ens.hash_registered")
	db.MustExec("DELETE FROM ens.hash_released")
	db.MustExec("DELETE FROM ens.new_bid")
	db.MustExec("DELETE FROM ens.new_owner")
	db.MustExec("DELETE FROM ens.new_resolver")
	db.MustExec("DELETE FROM ens.new_ttl")
	db.MustExec("DELETE FROM ens.transfer")
	db.MustExec("DELETE FROM ens.abi_changed")
	db.MustExec("DELETE FROM ens.addr_changed")
	db.MustExec("DELETE FROM ens.content_changed")
	db.MustExec("DELETE FROM ens.contenthash_changed")
	db.MustExec("DELETE FROM ens.multihash_changed")
	db.MustExec("DELETE FROM ens.name_changed")
	db.MustExec("DELETE FROM ens.pubkey_changed")
	db.MustExec("DELETE FROM ens.text_changed")
}

// Returns a new test node, with the same ID
func NewTestNode() core.Node {
	return core.Node{
		GenesisBlock: "GENESIS",
		NetworkID:    1,
		ID:           "b6f90c0fdd8ec9607aed8ee45c69322e47b7063f0bfb7a29c8ecafab24d0a22d24dd2329b5ee6ed4125a03cb14e57fd584e67f9e53e6c631055cbbd82f080845",
		ClientName:   "Geth/v1.7.2-stable-1db4ecdc/darwin-amd64/go1.9",
	}
}

func NewTestBlock(blockNumber int64, repository repositories.BlockRepository) (blockId int64) {
	blockId, err := repository.CreateOrUpdateBlock(core.Block{Number: blockNumber})
	Expect(err).NotTo(HaveOccurred())

	return blockId
}
