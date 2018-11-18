// Copyright © 2017-2018 The IPFN Developers. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cells

import (
	cid "gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"

	multihash "github.com/multiformats/go-multihash"
)

// TODO - Returns empty cell.
func TODO() *BinaryCell {
	return new(BinaryCell)
}

// Ops - Returns ops.
func Ops(ops ...Cell) []Cell {
	return ops
}

// Op - Creates new binary cell.
func Op(op ID, ops ...Cell) *BinaryCell {
	return &BinaryCell{opCode: op, children: ops}
}

// Root - Creates new binary cell.
func Root(ops []Cell) *BinaryCell {
	return &BinaryCell{children: ops}
}

// New - Creates new binary cell.
func New(op ID, memory []byte, children ...Cell) *BinaryCell {
	return &BinaryCell{
		opCode:   op,
		memory:   memory,
		children: children,
	}
}

// UnmarshalBinary - Unmarshals new binary cell.
func UnmarshalBinary(body []byte) (c *BinaryCell, err error) {
	c = new(BinaryCell)
	err = Unmarshal(c, body)
	return
}

// CellPrefix - Binary cell CID prefix.
var CellPrefix = cid.Prefix{
	Version:  1,
	Codec:    0x70bc,
	MhType:   multihash.SHA2_256,
	MhLength: 32,
}
