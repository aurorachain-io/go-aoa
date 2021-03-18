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
	"enode://9ecce36d18f04efc4a092e9c3afcd49219e3f28cf5ce226a42f3ef0556f820541af3ed51b02fa9e0fad041efa3eeb8a90a00abfcaf80ec5eb7564bd15fec643e@40.83.72.70:30303",
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
