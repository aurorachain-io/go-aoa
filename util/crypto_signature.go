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
	"crypto/ecdsa"
	"github.com/Aurorachain-io/go-aoa/crypto"
	"github.com/Aurorachain-io/go-aoa/crypto/secp256k1"
	"github.com/Aurorachain-io/go-aoa/crypto/sha3"
	"strings"
)

// sign data and return sign bytes
func Sign(data []byte, privateKey *ecdsa.PrivateKey) (sign []byte, err error) {
	dataBytes := sha3.Sum256(data)
	return crypto.Sign(dataBytes[:], privateKey)
}

// check by address
func VerifySignature(address string, data, sign []byte) (bool, error) {
	dataBytes := sha3.Sum256(data)

	pubkey, err := secp256k1.RecoverPubkey(dataBytes[:], sign)

	if err != nil {
		return false, err
	}

	return strings.EqualFold(crypto.PubkeyToAddress(*crypto.ToECDSAPub(pubkey)).Hex(), address), nil

}
