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

// +build none

package main

import (
	"bytes"
	"flag"
	"github.com/Aurorachain-io/go-aoa/accounts/keystore"
	"go/build"
	"go/format"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	pkgInfo  *build.Package
	scriptN  = keystore.StandardScryptN
	scriptP  = keystore.StandardScryptP
	nodeName = "em-node"
)

var (
	agentsNumber = flag.Int("number", 101, "number")
	password     = flag.String("password", "yujian", "password")
	keyStoreDir  = flag.String("keystore-dir", ".", "keystore-dir")
)

func main() {
	flag.Parse()

	keystore := keystore.NewKeyStore(*keyStoreDir, scriptN, scriptP)
	addressList := make([]string, 0)
	for i := 0; i < *agentsNumber; i++ {
		acc, err := keystore.NewAccount(*password)
		if err != nil {
			log.Fatal(err)
		}
		addressList = append(addressList, strings.ToLower(acc.Address.Hex()))
	}
	src := genString(addressList)
	outputName := "mkAgents.go"
	err := ioutil.WriteFile(outputName, src, 0644)
	if err != nil {
		log.Fatal(err)
	}
	genGenesisJson(addressList)
}

func genGenesisJson(addressLists []string) {
	balance := "0x422ca8b0a00a425000000"
	f, err := os.Create("./genesis.json")
	if err != nil {
		log.Fatal(err)
	}
	jsonString := `{"config": 
		{"chainId": 1,
		"homesteadBlock": 0,
		"eip155Block": 0,
		"eip158Block": 0
		},  
	"coinbase": "0x0000000000000000000000000000000000000000",
	"difficulty": "0x02000000",
	"extraData": "",
	"gasLimit": "0x2fefd8",
	"nonce": "0x0000000000000042",
	"mixhash": "0x0000000000000000000000000000000000000000000000000000000000000000",
	"parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
	"timestamp": "0x00",
	"alloc" : {`
	for _, address := range addressLists {
		addressBalance := `"` + address + `":{"balance": "` + balance + `"},` + "\n"
		jsonString = jsonString + addressBalance
	}
	jsonString = jsonString + `}, "agents" : [`
	for index, address := range addressLists {
		add := `{"address": "` + address + `",`
		vote := `"vote":2000000,`
		nickname := `"nickname" : "node` + strconv.Itoa(index) + `",`
		registerTime := `"registerTime" : 1492009146 },` + "\n"
		jsonString = jsonString + add + vote + nickname + registerTime
	}
	jsonString = jsonString + `]}`
	_, err = io.WriteString(f, jsonString)
	if err != nil {
		log.Fatal(err)
	}
}

func genString(addressLists []string) []byte {
	const strTmp = `
// Copyright 2018 The go-dacchain Authors
// This file is part of the go-dacchain library.
//
// The go-dacchain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-dacchain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-dacchain library. If not, see <http://www.gnu.org/licenses/>.

	// +build none

/*

   The mkalloc tool creates the genesis allocation constants in genesis_alloc.go
   It outputs a const declaration that contains an RLP-encoded list of (address, balance) tuples.

       go run mkalloc.go genesis.json

*/

	package {{.pkg}}
	
	import (
		"fmt"
		"github.com/Aurorachain-io/go-aoa/core/types"
 		"strconv"
		"github.com/Aurorachain-io/go-aoa/rlp"
	)

	type genesisAgents []types.Candidate
	
	func main() {
		var list genesisAgents
		candidateList := mainTestNetAgents()
		list = append(list, candidateList...)

		data, err := rlp.EncodeToBytes(list)
		if err != nil {
			panic(err)
		}
		result := strconv.QuoteToASCII(string(data))
		fmt.Println("const agentData =", result)
	}

	func mainNetAgents() []types.Candidate {
		return []types.Candidate{ {{range $index, $address := .addressLists}}
			{"{{$address}}",uint64(2000000), "em-node{{$index}}", 1492009146, }, {{end}}
		}
	}

	func mainTestNetAgents() []types.Candidate {
		candidateList := []types.Candidate{
			{"0x34f6feaa439ea2e92438365933067acaff5e3b7c", uint64(1), "node1-1", 1492009146}, // 172.16.134.100
			{"0xb34822fea9f8aaae7c7f64a097f64e5dffb6f344", uint64(1), "node1-2", 1492009146}, // 172.16.134.100
			{"0x678fe1cef127de901cc40a4fd3b608eb2b2a8b24", uint64(1), "node1-3", 1492009146}, // 172.16.134.100
			{"0x7d8d03a2b6674b8ab89b69e118508e7a7e75c2dc", uint64(1), "node2-1", 1492009146}, // 172.16.134.101
			{"0x671b399681d3bc8e27f874a52983ca559dde35ba", uint64(1), "node2-2", 1492009146}, // 172.16.134.101
		}
		return candidateList
	}
	`

	pkgName := "main"
	data := map[string]interface{}{
		"pkg":          pkgName,
		"addressLists": addressLists,
	}

	t, err := template.New("").Parse(strTmp)
	if err != nil {
		log.Fatal(err)
	}
	buff := bytes.NewBufferString("")
	err = t.Execute(buff, data)
	if err != nil {
		log.Fatal(err)
	}
	src, err := format.Source(buff.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	return src
}
