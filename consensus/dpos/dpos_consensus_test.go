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

package dpos

import (
	"math"
	"math/big"
	"testing"
)

func TestBlockReward(t *testing.T) {
	// begin with 100 reward
	var (
		// begin with 100 reward
		basicReward float64 = 100
		annulProfit         = 1.15
		// annulBlockAmount = big.NewInt(10)
		annulBlockAmount   = big.NewInt(3153600)
		blockReward        = big.NewInt(1e+18)
		currentBlockHeight = big.NewInt(0)
	)

	for i := 0; i < 100_000_000; i++ {
		currentBlockHeight.Add(currentBlockHeight, big.NewInt(1))
		yearNumber := currentBlockHeight.Int64() / annulBlockAmount.Int64()
		currentReward := (int64)(basicReward * math.Pow(annulProfit, float64(yearNumber)))
		if currentBlockHeight.Int64()%annulBlockAmount.Int64() == 0 {
			t.Logf("block number=%d currentReward=%d", currentBlockHeight, currentReward)
			precisionReward := new(big.Int).Mul(big.NewInt(currentReward), blockReward)
			t.Log(precisionReward)
		}
		//t.Logf("block number=%d currentReward=%d",currentBlockHeight,currentReward)
		//precisionReward := new(big.Int).Mul(big.NewInt(currentReward), blockReward)
		//t.Log(precisionReward)
	}

}
