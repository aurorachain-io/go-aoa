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

package util

import (
	"fmt"
	"github.com/Aurorachain-io/go-aoa/core/types"
	"testing"
	"time"
)

func TestShuffle(t *testing.T) {
	var delegateNumber = 4
	for i := 1; i < 100; i++ {
		fmt.Println(Shuffle(int64(i), delegateNumber))
		if i%delegateNumber == 0 {
			fmt.Println("=======================")
		}
	}
}

func TestShuffleNewRound(t *testing.T) {
	var initDelegate = []types.Candidate{
		{Address: "EM70715a2a44255ddce2779d60ba95968b770fc751", Nickname: "node1"},
		{Address: "EMfd48a829397a16b3bc6c319a06a47cd2ce6b3f52", Nickname: "node2"},
		{Address: "EM612d018cc7db4137366a08075333a634c07e31b3", Nickname: "node3"},
		{Address: "EM612d018cc7db4137366a08075333a634c07e31b4", Nickname: "node4"},
		{Address: "EM612d018cc7db4137366a08075333a634c07e31b5", Nickname: "node5"},
		{Address: "EM612d018cc7db4137366a08075333a634c07e31b6", Nickname: "node6"},
		{Address: "EM612d018cc7db4137366a08075333a634c07e31b7", Nickname: "node7"},
		{Address: "EM612d018cc7db4137366a08075333a634c07e31b8", Nickname: "node8"},
		{Address: "EM612d018cc7db4137366a08075333a634c07e31b9", Nickname: "node9"},
		{Address: "EM612d018cc7db4137366a08075333a634c07e31b0", Nickname: "node10"},
	}

	lastBlockTime := time.Now().Unix()
	newRound := ShuffleNewRound(lastBlockTime, 10, initDelegate)
	for _, v := range newRound {
		fmt.Println(v)
	}

}

func TestCalShuffleTimeByHeaderTime(t *testing.T) {
	shuffleTime := CalShuffleTimeByHeaderTime(3030, 2040)
	fmt.Println(shuffleTime)

}
