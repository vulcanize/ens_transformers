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

package constants

// Registar
func GetAuctionStartedSignature() string  { return GetEventSignature(auctionStartedMethod()) }
func GetBidRevealedSignature() string     { return GetEventSignature(bidRevealedMethod()) }
func GetHashInvalidatedSignature() string { return GetEventSignature(hashInvalidatedMethod()) }
func GetHashRegisteredSignature() string  { return GetEventSignature(hashRegisteredMethod()) }
func GetHashReleasedSignature() string    { return GetEventSignature(hashReleasedMethod()) }
func GetNewBidSignature() string          { return GetEventSignature(newBidMethod()) }

// Registry
func GetNewOwnerSignature() string    { return GetEventSignature(newOwnerMethod()) }
func GetNewResolverSignature() string { return GetEventSignature(newResolverMethod()) }
func GetNewTtlSignature() string      { return GetEventSignature(newTtlMethod()) }
func GetTransferSignature() string    { return GetEventSignature(transferMethod()) }

// Resolver
func GetAbiChangedSignature() string         { return GetEventSignature(abiChangedMethod()) }
func GetAddrChangedSignature() string        { return GetEventSignature(addrChangedMethod()) }
func GetContentChangedSignature() string     { return GetEventSignature(contentChangedMethod()) }
func GetContenthashChangedSignature() string { return GetEventSignature(contenthashChangedMethod()) }
func GetMultihashChangedSignature() string   { return GetEventSignature(multihashChangedMethod()) }
func GetNameChangedSignature() string        { return GetEventSignature(nameChangedMethod()) }
func GetPubkeyChangedSignature() string      { return GetEventSignature(pubkeyChangedMethod()) }
func GetTextChangedSignature() string        { return GetEventSignature(textChangedMethod()) }
