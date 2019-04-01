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

package constants

import (
	"fmt"

	"github.com/spf13/viper"
)

var initialized = false

func initConfig() {
	if initialized {
		return
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Using config file: %s\n\n", viper.ConfigFileUsed())
	} else {
		panic(fmt.Sprintf("Could not find environment file: %v", err))
	}
	initialized = true
}

func getEnvironmentString(key string) string {
	initConfig()
	value := viper.GetString(key)
	if value == "" {
		panic(fmt.Sprintf("No environment configuration variable set for key: \"%v\"", key))
	}
	return value
}

func getEnvironmentInt64(key string) int64 {
	initConfig()
	value := viper.GetInt64(key)
	if value == -1 {
		panic(fmt.Sprintf("No environment configuration variable set for key: \"%v\"", key))
	}
	return value
}

// Getters for contract addresses from environment files
func RegistryContractAddress() string { return getEnvironmentString("contract.address.registry") }
func RegistarContractAddress() string { return getEnvironmentString("contract.address.registar") }
func ResolverContractAddress() string { return getEnvironmentString("contract.address.resolver") }

func RegistryABI() string { return getEnvironmentString("contract.abi.registry") }
func RegistarABI() string { return getEnvironmentString("contract.abi.registar") }
func ResolverABI() string { return getEnvironmentString("contract.abi.resolver") }

func RegistryDeploymentBlock() int64 {
	return getEnvironmentInt64("contract.deployment-block.registry")
}
func RegistarDeploymentBlock() int64 {
	return getEnvironmentInt64("contract.deployment-block.registar")
}
func ResolverDeploymentBlock() int64 {
	return getEnvironmentInt64("contract.deployment-block.resolver")
}
