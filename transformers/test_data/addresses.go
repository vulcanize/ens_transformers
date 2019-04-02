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

const (
	RegistryAddress        = "0x314159265dD8dbb310642f98f50C066173C1259b"
	RopstenRegistryAddress = "0x112234455C3a32FD11230C42E7Bccd4A84e02010" //starts at block 25409
	ResolverAddress        = "0x1da022710dF5002339274AaDEe8D58218e9D6AB5"
	ResolverAddress2       = "0xD3ddcCDD3b25A8a7423B5bEe360a42146eb4Baf3" //starts at 6705947
	ResolverAddress3       = "0x5FfC014343cd971B7eb70732021E26C35B744cc4" //starts at 3733668
	RopstenResolverAddress = "0xcAcbE14d88380F8eb37ec0d7788ec226EE7b3434" //starts at block 4115489
	RegistarAddress        = "0x6090A6e47849629b7245Dfa1Ca21D94cd15878Ef"
	RopstenRegistarAddress = "0xC68De5B43C3d980B0C110A77a5F78d3c4c4d63B4" // starts at block 25461
)

/*

0x1da022710dF5002339274AaDEe8D58218e9D6AB5:
	event AddrChanged(bytes32 indexed node, address a);
    event ContentChanged(bytes32 indexed node, bytes32 hash);
    event NameChanged(bytes32 indexed node, string name);
    event ABIChanged(bytes32 indexed node, uint256 indexed contentType);
    event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y);

0xD3ddcCDD3b25A8a7423B5bEe360a42146eb4Baf3:
	event AddrChanged(bytes32 indexed node, address a);
    event NameChanged(bytes32 indexed node, string name);
    event ABIChanged(bytes32 indexed node, uint256 indexed contentType);
    event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y);
    event TextChanged(bytes32 indexed node, string indexedKey, string key);
    event ContenthashChanged(bytes32 indexed node, bytes hash);

0x5FfC014343cd971B7eb70732021E26C35B744cc4:
	event AddrChanged(bytes32 indexed node, address a);
    event ContentChanged(bytes32 indexed node, bytes32 hash);
    event NameChanged(bytes32 indexed node, string name);
    event ABIChanged(bytes32 indexed node, uint256 indexed contentType);
    event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y);
    event TextChanged(bytes32 indexed node, string indexed indexedKey, string key);


0xcAcbE14d88380F8eb37ec0d7788ec226EE7b3434:
	event AddrChanged(bytes32 indexed node, address a);
    event ContentChanged(bytes32 indexed node, bytes32 hash);
    event NameChanged(bytes32 indexed node, string name);
    event ABIChanged(bytes32 indexed node, uint256 indexed contentType);
    event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y);
    event TextChanged(bytes32 indexed node, string indexedKey, string key);
    event MultihashChanged(bytes32 indexed node, bytes hash);
*/
