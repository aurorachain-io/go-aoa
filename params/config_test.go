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

import (
	"math/big"
	"reflect"
	"testing"
)

func TestCheckCompatible(t *testing.T) {
	type test struct {
		stored, new *ChainConfig
		head        uint64
		wantErr     *ConfigCompatError
	}
	tests := []test{
		{stored: AllDacchainProtocolChanges, new: AllDacchainProtocolChanges, head: 0, wantErr: nil},
		{stored: AllDacchainProtocolChanges, new: AllDacchainProtocolChanges, head: 100, wantErr: nil},
		{
			stored:  &ChainConfig{},
			new:     &ChainConfig{ByzantiumBlock: big.NewInt(20)},
			head:    9,
			wantErr: nil,
		},
	}

	for _, test := range tests {
		err := test.stored.CheckCompatible(test.new, test.head)
		if !reflect.DeepEqual(err, test.wantErr) {
			t.Errorf("error mismatch:\nstored: %v\nnew: %v\nhead: %v\nerr: %v\nwant: %v", test.stored, test.new, test.head, err, test.wantErr)
		}
	}
}

func TestConfigRules(t *testing.T) {
	c := &ChainConfig{
		LondonBlock:  new(big.Int),
		ShanghaiTime: newUint64(450),
	}
	var stamp uint64
	if r := c.Rules(big.NewInt(1), true, stamp); r.IsShanghai {
		t.Errorf("expected %v to not be", stamp)
	}
	stamp = 450
	if r := c.Rules(big.NewInt(1), true, stamp); !r.IsShanghai {
		t.Errorf("expected %v to be", stamp)
	}
	stamp = math.MaxInt64
	if r := c.Rules(big.NewInt(1), true, stamp); !r.IsShanghai {
		t.Errorf("expected %v to be", stamp)
	}
}
