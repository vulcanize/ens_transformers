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
package config

import (
	"github.com/vulcanize/vulcanizedb/pkg/config"
)

var MainnetENSConfig = config.ContractConfig{
	Name:    "ENS-mainnet",
	Network: "",
	Addresses: map[string]bool{
		"0x314159265dD8dbb310642f98f50C066173C1259b": true,
	},
	Abis: map[string]string{
		"0x314159265dD8dbb310642f98f50C066173C1259b": `[{"constant":true,"inputs":[{"name":"node","type":"bytes32"}],"name":"resolver","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"node","type":"bytes32"}],"name":"owner","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"node","type":"bytes32"},{"name":"label","type":"bytes32"},{"name":"owner","type":"address"}],"name":"setSubnodeOwner","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"node","type":"bytes32"},{"name":"ttl","type":"uint64"}],"name":"setTTL","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"node","type":"bytes32"}],"name":"ttl","outputs":[{"name":"","type":"uint64"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"node","type":"bytes32"},{"name":"resolver","type":"address"}],"name":"setResolver","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"node","type":"bytes32"},{"name":"owner","type":"address"}],"name":"setOwner","outputs":[],"payable":false,"type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"node","type":"bytes32"},{"indexed":false,"name":"owner","type":"address"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"node","type":"bytes32"},{"indexed":true,"name":"label","type":"bytes32"},{"indexed":false,"name":"owner","type":"address"}],"name":"NewOwner","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"node","type":"bytes32"},{"indexed":false,"name":"resolver","type":"address"}],"name":"NewResolver","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"node","type":"bytes32"},{"indexed":false,"name":"ttl","type":"uint64"}],"name":"NewTTL","type":"event"}]`,
	},
	Events: map[string][]string{
		"0x314159265dD8dbb310642f98f50C066173C1259b": []string{},
	},
	EventArgs: map[string][]string{
		"0x314159265dD8dbb310642f98f50C066173C1259b": []string{},
	},
	Methods: map[string][]string{
		"0x314159265dD8dbb310642f98f50C066173C1259b": []string{},
	},
	MethodArgs: map[string][]string{
		"0x314159265dD8dbb310642f98f50C066173C1259b": []string{},
	},
	StartingBlocks: map[string]int64{
		"0x314159265dD8dbb310642f98f50C066173C1259b": 3327417,
	},
	Piping: map[string]bool{
		"0x314159265dD8dbb310642f98f50C066173C1259b": false,
	},
}

var RopstenENSConfig = config.ContractConfig{
	Name:    "ENS-ropsten",
	Network: "ropsten",
	Addresses: map[string]bool{
		"0x112234455C3a32FD11230C42E7Bccd4A84e02010": true,
	},
	Abis: map[string]string{
		"0x112234455C3a32FD11230C42E7Bccd4A84e02010": `[{"constant":true,"inputs":[{"name":"node","type":"bytes32"}],"name":"resolver","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"node","type":"bytes32"}],"name":"owner","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"node","type":"bytes32"},{"name":"label","type":"bytes32"},{"name":"owner","type":"address"}],"name":"setSubnodeOwner","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"node","type":"bytes32"},{"name":"ttl","type":"uint64"}],"name":"setTTL","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"node","type":"bytes32"}],"name":"ttl","outputs":[{"name":"","type":"uint64"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"node","type":"bytes32"},{"name":"resolver","type":"address"}],"name":"setResolver","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"node","type":"bytes32"},{"name":"owner","type":"address"}],"name":"setOwner","outputs":[],"payable":false,"type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"node","type":"bytes32"},{"indexed":false,"name":"owner","type":"address"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"node","type":"bytes32"},{"indexed":true,"name":"label","type":"bytes32"},{"indexed":false,"name":"owner","type":"address"}],"name":"NewOwner","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"node","type":"bytes32"},{"indexed":false,"name":"resolver","type":"address"}],"name":"NewResolver","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"node","type":"bytes32"},{"indexed":false,"name":"ttl","type":"uint64"}],"name":"NewTTL","type":"event"}]`,
	},
	Events: map[string][]string{
		"0x112234455C3a32FD11230C42E7Bccd4A84e02010": []string{},
	},
	EventArgs: map[string][]string{
		"0x112234455C3a32FD11230C42E7Bccd4A84e02010": []string{},
	},
	Methods: map[string][]string{
		"0x112234455C3a32FD11230C42E7Bccd4A84e02010": []string{},
	},
	MethodArgs: map[string][]string{
		"0x112234455C3a32FD11230C42E7Bccd4A84e02010": []string{},
	},
	StartingBlocks: map[string]int64{
		"0x112234455C3a32FD11230C42E7Bccd4A84e02010": 25409,
	},
	Piping: map[string]bool{
		"0x112234455C3a32FD11230C42E7Bccd4A84e02010": false,
	},
}
