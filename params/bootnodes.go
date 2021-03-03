// Copyright 2021 The go-aoa Authors
// This file is part of the go-aoa library.
//
// The the go-aoa library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The the go-aoa library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-aoa library. If not, see <http://www.gnu.org/licenses/>.

package params

import "math/big"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main eminer-pro network.
var MainnetBootnodes = []string{
	"enode://31e515e0875ba9898629a00fade15617e9df9de58d3d5238cb440eabc8f16b773a84541bc8a93b2fd828a47054478d3c68d1a76988b83fb32ba748917a574ef8@40.83.72.70:30303",
}

// TestnetBootnodes
var TestnetBootnodes = []string{
	// "enode://7baae2fac6c271737672ad6f15200b60a5b971cd802f85854999536c47bfa644e04eb9dcc8a57333dbd755d77f4797a4dadc0e8c2d0da4f38dd9f422ee593f7f@172.16.20.76:30303",
}

// block reward
var (
	AnnulProfit = 1.10
	AnnulBlockAmount = big.NewInt(3153600)
)
