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

// Registar
func auctionStartedMethod() string { return GetSolidityMethodSignature(RegistarABI(), "AuctionStarted") }
func bidRevealedMethod() string    { return GetSolidityMethodSignature(RegistarABI(), "BidRevealed") }
func hashInvalidatedMethod() string {
	return GetSolidityMethodSignature(RegistarABI(), "HashInvalidated")
}
func hashRegisteredMethod() string { return GetSolidityMethodSignature(RegistarABI(), "HashRegistered") }
func hashReleasedMethod() string   { return GetSolidityMethodSignature(RegistarABI(), "HashReleased") }
func newBidMethod() string         { return GetSolidityMethodSignature(RegistarABI(), "NewBid") }

// Registry
func newOwnerMethod() string    { return GetSolidityMethodSignature(RegistryABI(), "NewOwner") }
func newResolverMethod() string { return GetSolidityMethodSignature(RegistryABI(), "NewResolver") }
func newTtlMethod() string      { return GetSolidityMethodSignature(RegistryABI(), "NewTTL") }
func transferMethod() string    { return GetSolidityMethodSignature(RegistryABI(), "Transfer") }

// Resolver
func abiChangedMethod() string     { return GetSolidityMethodSignature(ResolverABI(), "AbiChanged") }
func addrChangedMethod() string    { return GetSolidityMethodSignature(ResolverABI(), "AddrChanged") }
func contentChangedMethod() string { return GetSolidityMethodSignature(ResolverABI(), "ContentChanged") }
func contenthashChangedMethod() string {
	return GetSolidityMethodSignature(ResolverABI(), "ContenthashChanged")
}
func multihashChangedMethod() string {
	return GetSolidityMethodSignature(ResolverABI(), "MultihashChanged")
}
func nameChangedMethod() string   { return GetSolidityMethodSignature(ResolverABI(), "NameChanged") }
func pubkeyChangedMethod() string { return GetSolidityMethodSignature(ResolverABI(), "PubkeyChanged") }
