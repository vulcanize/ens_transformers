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

package test_data

import "github.com/vulcanize/vulcanizedb/pkg/contract_watcher/shared/helpers"

var (
	// Resolver
	AddrChangedSignature        = helpers.GenerateSignature("AddrChanged(bytes32,address)")
	ContentChangedSignature     = helpers.GenerateSignature("ContentChanged(bytes32,bytes32)")
	NameChangedSignature        = helpers.GenerateSignature("NameChanged(bytes32,string)")
	AbiChangedSignature         = helpers.GenerateSignature("ABIChanged(bytes32,uint256)")
	PubkeyChangedSignature      = helpers.GenerateSignature("PubkeyChanged(bytes32,bytes32,bytes32)")
	TextChangedSignature        = helpers.GenerateSignature("TextChanged(bytes32,string,string)")
	MultihashChangedSignature   = helpers.GenerateSignature("MultihashChanged(bytes32,bytes)")
	ContenthashChangedSignature = helpers.GenerateSignature("ContenthashChanged(bytes32,bytes)")
	// Registry
	NewOwnerSignature    = helpers.GenerateSignature("NewOwner(bytes32,bytes32,address)")
	TransferSignature    = helpers.GenerateSignature("Transfer(bytes32,address)")
	NewResolverSignature = helpers.GenerateSignature("NewResolver(bytes32,address)")
	NewTtlSignature      = helpers.GenerateSignature("NewTTL(bytes32,uint64)")
	// Registar
	AuctionStartedSignature  = helpers.GenerateSignature("AuctionStarted(bytes32,uint)")
	NewBidSignature          = helpers.GenerateSignature("NewBid(bytes32,address,uint)")
	BidRevealedSignature     = helpers.GenerateSignature("BidRevealed(bytes32,address,uint,uint8)")
	HashRegisteredSignature  = helpers.GenerateSignature("HashRegistered(bytes32,address,uint,uint)")
	HashReleasedSignature    = helpers.GenerateSignature("HashReleased(bytes32,uint)")
	HashInvalidatedSignature = helpers.GenerateSignature("HashInvalidated(bytes32,string,uint,uint)")
)
