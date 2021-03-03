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

/*

   The mkalloc tool creates the genesis allocation constants in genesis_alloc.go
   It outputs a const declaration that contains an RLP-encoded list of (address, balance) tuples.

       go run mkalloc.go genesis.json

*/
package main

import (
	"fmt"
	"github.com/Aurorachain-io/go-aoa/core/types"
	"github.com/Aurorachain-io/go-aoa/rlp"
	"strconv"
)

type genesisAgents []types.Candidate

const (
	// aoa 10 billion
	firstDelegateVoteNumber  = uint64(1000_000_000)
	secondDelegateVoteNumber = uint64(500_000_000)
	thirdDelegateVoteNumber  = uint64(250_000_000)
	fourthDelegateVoteNumber = uint64(100_000_000)
	fifthDelegateVoteNumber  = uint64(50_000_000)
	sixthDelegateVoteNumber  = uint64(30_000_000)
)

func main() {
	var list genesisAgents
	candidateList := mainNetAgents()
	list = append(list, candidateList...)

	data, err := rlp.EncodeToBytes(list)
	if err != nil {
		panic(err)
	}
	result := strconv.QuoteToASCII(string(data))
	fmt.Println("const agentData =", result)
	// fmt.Println(len(list))
}

func mainNetAgents() []types.Candidate {
	return []types.Candidate{
		{"AOAd9f5038ca3908212d5a13c3b48a4df7c1dfd5a54", firstDelegateVoteNumber, "Aurora", 1614652166},
		{"AOA05e5097e0df07aaed1d3ac0ee1a6d71891a52773", firstDelegateVoteNumber, "Gaea", 1614652166},
		{"AOAe0d82b073e01a110cd9b0bfb514543fc3ee9beb0", firstDelegateVoteNumber, "Uranus", 1614652166},
		{"AOA3e1ab1de317215d1091844f147bf3aceefd20545", firstDelegateVoteNumber, "Titans", 1614652166},
		{"AOA1fe71b9116d2d67cf6c51387349fc83454ef91b7", firstDelegateVoteNumber, "Cronus", 1614652166},
		{"AOA6858033e149852f84353a69a1427f428713536fd", firstDelegateVoteNumber, "Rhea", 1614652166},
		{"AOA99948e917dba65b74b1b0c538fcb1e151bb6595d", firstDelegateVoteNumber, "Prometheus", 1614652166},
		{"AOA302bfa9cdb0a2121834ff7fc7f0c0551298f228c", firstDelegateVoteNumber, "Epimetheus", 1614652166},
		{"AOA00161f81fc254c87b70d40c6952f8e48cfe5a0c7", firstDelegateVoteNumber, "Apollo", 1614652166},
		{"AOA90ec53b83366b55bb3d7a806c4d5f0af9265270f", firstDelegateVoteNumber, "Zeus", 1614652166},

		{"AOA4b52db84e63227aac18fed08f4dcc4e3e1bbb013", secondDelegateVoteNumber, "Olympus", 1614652166},
		{"AOA3009e7c59d7d58530c800f5073af7bf73b4439bb", secondDelegateVoteNumber, "Venus", 1614652166},
		{"AOA3664c0963b944e158b0c3a05633e88f67f34851a", secondDelegateVoteNumber, "Mars", 1614652166},
		{"AOA105f5cfe913bf7bfca07f400a7aa53bb6b518317", secondDelegateVoteNumber, "Diana", 1614652166},
		{"AOAae97155a2db2691eb8ba5ae36ac085caec5709ac", secondDelegateVoteNumber, "Minerva", 1614652166},
		{"AOAafcde920ea8de6886ffa7e3726334cce7e31d464", secondDelegateVoteNumber, "Ceres", 1614652166},
		{"AOA831f61e8d060ca8632c1feb685f2d37a4dca0c30", secondDelegateVoteNumber, "Vulcanus", 1614652166},
		{"AOA8e42705e50c5b088bcd7a314d981535ef7b72574", secondDelegateVoteNumber, "Juno", 1614652166},
		{"AOA4c41c8ee18f10a7ebfdd1e5697296aec1cf2607b", secondDelegateVoteNumber, "Mercury", 1614652166},
		{"AOA444baa079e08f6c8911fb31453e8813187b5ce14", secondDelegateVoteNumber, "Vesta", 1614652166},
		{"AOAc78400f3035f5ad4fcb7514da10ea5b4e0380e1f", secondDelegateVoteNumber, "Neptunus", 1614652166},
		{"AOA2a146ea409304bed02f86ed77d25626382f33de4", secondDelegateVoteNumber, "Jupiter", 1614652166},
		{"AOA88ddc2781c44ddfea1119f57689af1f85138544c", secondDelegateVoteNumber, "Aether", 1614652166},
		{"AOAbef02d1bfead42e5187be4c1ff0566619fe7e386", secondDelegateVoteNumber, "Necessitas", 1614652166},
		{"AOA14c3174620df0b3cdf8c785c30b3b9e2f28d93d4", secondDelegateVoteNumber, "Erebus", 1614652166},
		{"AOA079ab109bdb6b94cbeb7ab15430e97d31b42ec0f", secondDelegateVoteNumber, "Terra", 1614652166},
		{"AOA6463fed2e07d1aeacdbeb14fe41e5809b382620e", secondDelegateVoteNumber, "HAOAera", 1614652166},
		{"AOA1049ce632bd01b1d3e9848882e61f99e7e7a61eb", secondDelegateVoteNumber, "Chaos", 1614652166},
		{"AOAa0a912eab7cf58ee63b1987cc3d64cbe18518cba", secondDelegateVoteNumber, "Chronus", 1614652166},
		{"AOA3cdbb42231d3019d5a0b331b532e4aed835897c5", secondDelegateVoteNumber, "Nesoi", 1614652166},

		{"AOAbb393b9dd72b852b4ec679345d4aac766b573f35", thirdDelegateVoteNumber, "Nox", 1614652166},
		{"AOA23b583a72512591a664267208521a6425958ad64", thirdDelegateVoteNumber, "Caelus", 1614652166},
		{"AOAb111e5dea09e53badbf083b3c66852bbf82fdb40", thirdDelegateVoteNumber, "Ourea", 1614652166},
		{"AOAebd2ed244c9806384bfa0e8fa7a0df04872aec82", thirdDelegateVoteNumber, "Phanes", 1614652166},
		{"AOAdeb997d23fc4c5db5f3972f4700a1303d41f20f1", thirdDelegateVoteNumber, "Pontus", 1614652166},
		{"AOA7d483a1f7a29f3f862a57107d1c5281a57ece205", thirdDelegateVoteNumber, "Tartarus", 1614652166},
		{"AOAf0c6dcd03c4dbee89f79589a7a6412d767ea0d2c", thirdDelegateVoteNumber, "Thalassa", 1614652166},
		{"AOA38fa44f22f4760c2e6e86ed6ca7c411c0a8d0ad5", thirdDelegateVoteNumber, "Coma", 1614652166},
		{"AOA93c2648a3f18c236b3a4ebd8c162967dfc25e860", thirdDelegateVoteNumber, "Corona Australis", 1614652166},
		{"AOA7aaf1fd1b4e4067c93fc0bc81748f5b642992a9c", thirdDelegateVoteNumber, "Corona Borealis", 1614652166},
		{"AOAf75e915bc188ae6fdb119602c546258b6ace7556", thirdDelegateVoteNumber, "Corvus", 1614652166},
		{"AOAd0caebfccd4df3c2540c3943f206ca1c5c253288", thirdDelegateVoteNumber, "Crater", 1614652166},
		{"AOAbab085b6a1deccd427c0415aff5f4c9d293ced56", thirdDelegateVoteNumber, "Crux", 1614652166},
		{"AOA80f2d52e1d7d4ce2ba7809e66e4bb6d413e12dd0", thirdDelegateVoteNumber, "Cygnus", 1614652166},
		{"AOA21c8d6bbc242ccdc3ae612868aada10662a9e4b6", thirdDelegateVoteNumber, "Delphinus", 1614652166},
		{"AOAad52bc2fb027171f5007f084af10ccc670d7db9d", thirdDelegateVoteNumber, "Dorado", 1614652166},
		{"AOA9fd52e067d56efa39f0f9ed5a2178d92b0a5124e", thirdDelegateVoteNumber, "Draco", 1614652166},
		{"AOAe49dccd60d1a0c77ce39bbe51f39774347e65f11", thirdDelegateVoteNumber, "Equuleus", 1614652166},
		{"AOA40a5e0ba7b1d0600fe1feeec47cedabf7374bd96", thirdDelegateVoteNumber, "Eridanus", 1614652166},
		{"AOA81c355c59d54f57c8987bff21f626d71a9b583b1", thirdDelegateVoteNumber, "Fornax", 1614652166},
		{"AOAd2b0988b7a334c3bf467b1db299b47638888bd69", thirdDelegateVoteNumber, "GAOAini", 1614652166},
		{"AOAf31c4c9e8803fdadd57afd40708a86e6d6e1660e", thirdDelegateVoteNumber, "Grus", 1614652166},
		{"AOA4d242152adb20671204632433cda86dd0f9c2849", thirdDelegateVoteNumber, "Hercules", 1614652166},
		{"AOA46f0bd68175fbee6b7788697ee759bf4116e48fb", thirdDelegateVoteNumber, "Horologium", 1614652166},
		{"AOAf5844de8a3ce5ed6744622379704dc907dcb50c0", thirdDelegateVoteNumber, "Hydra", 1614652166},
		{"AOAd7b7209351e7d9d17165475433e3cffcdd993072", thirdDelegateVoteNumber, "Hydrus", 1614652166},
		{"AOAc739e6ed7c7a6ef7234272c718ccca4bc16017e2", thirdDelegateVoteNumber, "Indus", 1614652166},
		{"AOAc40ca043b98803b0555bc985a3ac91f334a76575", thirdDelegateVoteNumber, "Lacerta", 1614652166},
		{"AOA58ba640ac5b45cee8bd49007eb7f65a9a45ddf4a", thirdDelegateVoteNumber, "Leo", 1614652166},
		{"AOA0482f37d11230a40d0c9b140a96fcf9f3d80cc1e", thirdDelegateVoteNumber, "Leo Minor", 1614652166},
		{"AOA291dd0a855f3fb9a01bf4135bd41ace98f6b34ae", thirdDelegateVoteNumber, "Lepus", 1614652166},
		{"AOAb4389ffcfdf333890b32cc136ed90c350c4a9779", thirdDelegateVoteNumber, "Libra", 1614652166},
		{"AOAb12f7f710c4a7e675a6f18cee2927b1e6801adf2", thirdDelegateVoteNumber, "Lupus", 1614652166},
		{"AOA6429514b4359a5b7fd489c4fdeb8b2bc61558312", thirdDelegateVoteNumber, "Lynx", 1614652166},
		{"AOA8a4bf668d67c13492051e98c560c8ea38194dfa3", thirdDelegateVoteNumber, "Lyra", 1614652166},

		{"AOA0b8558ae589ab1f64a6bc27743c8902514a7a49b", fourthDelegateVoteNumber, "Mensa", 1614652166},
		{"AOA527814176ebce1b92febf717ea5ad496859a52d7", fourthDelegateVoteNumber, "Microscopium", 1614652166},
		{"AOA38cbd2d929ee95fb47061b79e0aad9944e43cccb", fourthDelegateVoteNumber, "Monoceros", 1614652166},
		{"AOA9b1760fca73fb216064176a6d722376f57036e1f", fourthDelegateVoteNumber, "Musca", 1614652166},
		{"AOA439e4b69f54f54ca5d9f54e4f2b17b639acddc78", fourthDelegateVoteNumber, "Norma", 1614652166},
		{"AOA42dbfb861b76027cb03cd05dfea9c3506b8d5f37", fourthDelegateVoteNumber, "Octans", 1614652166},
		{"AOAf929998007d2aa396c29b4e989fab2ef5a323c56", fourthDelegateVoteNumber, "Ophiuchus", 1614652166},
		{"AOA20f5bb478b50c100ca4cad11e32d5a5925432fd2", fourthDelegateVoteNumber, "Orion", 1614652166},
		{"AOAbb901bbfcb9913f3514984b8e5377c9448462b78", fourthDelegateVoteNumber, "Pavo", 1614652166},
		{"AOAb8577688343f7cc0536aca5b77f3bfdb8ca286fd", fourthDelegateVoteNumber, "Pegasus", 1614652166},
		{"AOAd1ad836ea79fc5540d07f6647778f6e22ab2a344", fourthDelegateVoteNumber, "Perseus", 1614652166},
		{"AOA9e2d01365b24b0ed4d19648d824fdafc6fa1e63c", fourthDelegateVoteNumber, "Phoenix", 1614652166},
		{"AOA2a25afe9baa947ccc02ec21a5fadb8f3ac0e950b", fourthDelegateVoteNumber, "Pictor", 1614652166},
		{"AOA4c078fbfcaf306f59a37ab4600b00a82f3be1440", fourthDelegateVoteNumber, "Pisces", 1614652166},
		{"AOA073cc35f9401a7cd5c3fd73f582ed9e1d308a040", fourthDelegateVoteNumber, "Piscis Austrinus", 1614652166},
		{"AOA12c1494d898ea6313928c23f39cb5f10f89a1209", fourthDelegateVoteNumber, "Puppis", 1614652166},
		{"AOAafad8dc99c29ab46487796d5bc8612b6a3f11cc8", fourthDelegateVoteNumber, "Pyxis", 1614652166},
		{"AOA2147e6a51641ce9bde559a8970f7e7c5c511a944", fourthDelegateVoteNumber, "Reticulum", 1614652166},
		{"AOA485c0c67cf8dd3b856399d45c7a329638613fd85", fourthDelegateVoteNumber, "Sagitta", 1614652166},
		{"AOAb2854242138638b9a18c5d97f875f5d3c8008423", fourthDelegateVoteNumber, "Sagittarius", 1614652166},
		{"AOA0b4a309638e8c0fdddd5e9f211acc7b339d90be6", fourthDelegateVoteNumber, "Scorpius", 1614652166},
		{"AOAa9dada13682d7b3a1b9e47e2c31a53d83c7919df", fourthDelegateVoteNumber, "Sculptor", 1614652166},
		{"AOA3282f7b09b63e04e0233c001a72ab868cf277d74", fourthDelegateVoteNumber, "Scutum", 1614652166},
		{"AOAef5fb0059d2cfa18b91feb6c925fc33473dcc526", fourthDelegateVoteNumber, "Serpens", 1614652166},
		{"AOA3f7810e652af6b6e69859488335ab9715790bd6f", fourthDelegateVoteNumber, "Sextans", 1614652166},
		{"AOA627b6c290ceb0194301bc4daf0cc860fe18de6c3", fourthDelegateVoteNumber, "Taurus", 1614652166},
		{"AOA08fab6429391e0fa5085a43104869f740776c973", fourthDelegateVoteNumber, "Telescopium", 1614652166},
		{"AOA920fa233dd4c7c153f49df2c1682cb59d9bcdc54", fourthDelegateVoteNumber, "Triangulum", 1614652166},
		{"AOA83a4f104427b3c761119aaf7c4997da4764ccded", fourthDelegateVoteNumber, "Triangulum Australe", 1614652166},
		{"AOAe6df17eb49c225aee429537c6a6c03d672d6084b", fourthDelegateVoteNumber, "Tucana", 1614652166},

		{"AOAaeb4a828f2889f274ce82e3651336f493b2f7afc", fifthDelegateVoteNumber, "Ursa Major", 1614652166},
		{"AOA123db04e875763f06503c6b88117917f4ef1f642", fifthDelegateVoteNumber, "Ursa Minor", 1614652166},
		{"AOAe4f4e1dea4127e3b3f466b21f77775458da98b21", fifthDelegateVoteNumber, "Vela", 1614652166},
		{"AOAcd0465c3917366acb0da0256f98b1c56e7c4e8ac", fifthDelegateVoteNumber, "Virgo", 1614652166},
		{"AOAf502bc76e5aee3f3f5fa21dce14f07414e66667c", fifthDelegateVoteNumber, "Volans", 1614652166},
		{"AOA99f436f8544b191fbf2b096164eb76ebadbf6555", fifthDelegateVoteNumber, "Vulpecula", 1614652166},
	}
}

func mainTestNetAgents() []types.Candidate {
	candidateList := []types.Candidate{
		{"0xae6951c73938d009835ffe56649083a24b823c42", uint64(2000000), "aoa-node1", 1544421807},
		{"0x8f0ccc912c0ec8ce7238044c6976abe146914ff5", uint64(2000000), "aoa-node2", 1544421807},
		{"0xf37be5b5eaff9cf82ae59b7bb2c106e9a5cb1523", uint64(2000000), "aoa-node3", 1544421807},
		{"0x5844650159455c08ebad1e719c0ed85113d9a91a", uint64(2000000), "aoa-node4", 1544421807},
		{"0xb78ec5b82019b6adb3026443a850a3f25d66e813", uint64(2000000), "aoa-node5", 1544421807},
		{"0xfe07c4020a86a9f902bd24c762fc02df1e78312f", uint64(2000000), "aoa-node6", 1544421807},
		{"0x8e218afb0aeb913c3b9aefa03a3194b95f0be2a8", uint64(2000000), "aoa-node7", 1544421807},
		{"0xb3ccd4bd10a19f0f0b68d86ecf8c8ba707472d47", uint64(2000000), "aoa-node8", 1544421807},
		{"0xf848390987333c7ee7fd0b63d5ef1b3fa7d1cca2", uint64(2000000), "aoa-node9", 1544421807},
		{"0x75de393d3e2b4671257e130b9f37b6e55de1d116", uint64(2000000), "aoa-node10", 1544421807},
		{"0xc3fd54f8031a068aee81be222d68cc31a7a93c5f", uint64(2000000), "aoa-node11", 1544421807},
		{"0x9c2490702a9f779ff93697c3935c1686d0ec481a", uint64(2000000), "aoa-node12", 1544421807},
		{"0x0a484d0a330f662769846873a3985192febdbbda", uint64(2000000), "aoa-node13", 1544421807},
		{"0x62ca04d697f23b86d57d9d41d904b7a5e808873d", uint64(2000000), "aoa-node14", 1544421807},
		{"0x9b5bc9b6047e090837b1bdc7a5ce9ab3305e211d", uint64(2000000), "aoa-node15", 1544421807},
		{"0x5b04e79a854435d3c4c849f1deac066404a3b266", uint64(2000000), "aoa-node16", 1544421807},
		{"0x21e8e874b78e150f5978b2ad72e3e52b120b8c79", uint64(2000000), "aoa-node17", 1544421807},
		{"0x3f1cd2621dd662ed2a49f76fd3e2c9377ff8e1f8", uint64(2000000), "aoa-node18", 1544421807},
		{"0xc71c00c652208621b60339a5c4da14edd48b297c", uint64(2000000), "aoa-node19", 1544421807},
		{"0x58cc16fe47a25898102e9aa7bacd9d743f3834a7", uint64(2000000), "aoa-node20", 1544421807},
		{"0x290183e728c68bda990939faca666f71fdb0729a", uint64(2000000), "aoa-node21", 1544421807},
		{"0x9c218da1360bc9a6953f7b1bb403908a94204375", uint64(2000000), "aoa-node22", 1544421807},
		{"0xe32c0383f7b945a1c8b85a07135f376fe0e0a6d6", uint64(2000000), "aoa-node23", 1544421807},
		{"0xb4cfd6e65d5b77f5c5f1764cbe6ffca3f734f4ea", uint64(2000000), "aoa-node24", 1544421807},
		{"0x5e2e87067d5428b7d2a281ef0e27dcd889cb87a5", uint64(2000000), "aoa-node25", 1544421807},
		{"0x4aabbbcc6f07e88396f85eb99e777bb25c540720", uint64(2000000), "aoa-node26", 1544421807},
		{"0xf99c1fd103b9beeb6bf5c4d76568812898f7a7b9", uint64(2000000), "aoa-node27", 1544421807},
		{"0xe27636f648505609ecfd622c9cdef371f396a60b", uint64(2000000), "aoa-node28", 1544421807},
		{"0xeaf37aa9392b32287d6a168ff7986e33169c39cd", uint64(2000000), "aoa-node29", 1544421807},
		{"0xe6b8c9bdb76e30eb11fc890718413f1eb32bdf8e", uint64(2000000), "aoa-node30", 1544421807},
		{"0x7a6b0b4b34e76122a9a0085a3e9e4feafec11470", uint64(2000000), "aoa-node31", 1544421807},
		{"0x689d2223743547a3e9c016cefcbac6d04de009ef", uint64(2000000), "aoa-node32", 1544421807},
		{"0x1b6b92e6891cbaed468fcb21db47f1d6a330bb2c", uint64(2000000), "aoa-node33", 1544421807},
		{"0x1f9cf1b6d66130de4cef4b4ba2c27427d1d281be", uint64(2000000), "aoa-node34", 1544421807},
		{"0x255d1f8eeb0bc29937915ea5adf6774c7df220e4", uint64(2000000), "aoa-node35", 1544421807},
		{"0xebbbac11a3861b70678a83b48a3f96403232fcb9", uint64(2000000), "aoa-node36", 1544421807},
		{"0x2a4e995b17a703e65df494d151c7b6582ef2732e", uint64(2000000), "aoa-node37", 1544421807},
		{"0xdc6a079ad488dbd2b2aa73165bc8a3af33cb0996", uint64(2000000), "aoa-node38", 1544421807},
		{"0x6a00b9059a095473ea410d9aa185025cf395e2aa", uint64(2000000), "aoa-node39", 1544421807},
		{"0x91bf699159786134843764861967de8bfafb6323", uint64(2000000), "aoa-node40", 1544421807},
		{"0x434829c35f3ed551eacc238f6d9aee757dc4d14e", uint64(2000000), "aoa-node41", 1544421807},
		{"0xb69fb7df43a9b5131c0e9ab25dd005530cacdd86", uint64(2000000), "aoa-node42", 1544421807},
		{"0xef8c732b3ba657da11b00a7ee0871cda50e433dd", uint64(2000000), "aoa-node43", 1544421807},
		{"0xdecfb4a2af72474c60705f7fae82ee4c6ed306f4", uint64(2000000), "aoa-node44", 1544421807},
		{"0x48165ab216a3436cfa798279bd13ead7c66aa508", uint64(2000000), "aoa-node45", 1544421807},
		{"0xfb9b90ef7d29ca05a9dbf51382708bbe6b1cab1d", uint64(2000000), "aoa-node46", 1544421807},
		{"0xc3025d450debcf8d88d8f2109fdc6a672c102789", uint64(2000000), "aoa-node47", 1544421807},
		{"0x056d0911ad19e1f8ceeafb097364d699c59cc64e", uint64(2000000), "aoa-node48", 1544421807},
		{"0xffd7b90cd18aee781ecb824e4bff5af57f76a734", uint64(2000000), "aoa-node49", 1544421807},
		{"0x72823a6016ea77d30238df3bbe9f22cd429d4ae8", uint64(2000000), "aoa-node50", 1544421807},
		{"0x7db8ff022d12273e95d4d9e232c4b3cdf7158dfe", uint64(2000000), "aoa-node51", 1544421807},
		{"0xf205583e5f7cac53336f2288b8db1b247978df88", uint64(2000000), "aoa-node52", 1544421807},
		{"0xca6fa5fb62c68758e9790005412c36aa3e315d2e", uint64(2000000), "aoa-node53", 1544421807},
		{"0xde5d2f5ae46d9339b796c1dad47ae25a917ab5a3", uint64(2000000), "aoa-node54", 1544421807},
		{"0xb32cfef47545fdcfe109543971cef6626528aea4", uint64(2000000), "aoa-node55", 1544421807},
		{"0x3e9b3ce3a11f00331ec39cf3feb4b6dc42ba3671", uint64(2000000), "aoa-node56", 1544421807},
		{"0xa584bbe89bd46d86925294701bcbd1b4158abf12", uint64(2000000), "aoa-node57", 1544421807},
		{"0x13042f9b0f6da4f598e56a9bb8249baacb9f24ae", uint64(2000000), "aoa-node58", 1544421807},
		{"0xccdaf20553d819faf64e4947fd4e94f65f20df5d", uint64(2000000), "aoa-node59", 1544421807},
		{"0x967062fff19489237fdce95432e0d8c0a67ebae4", uint64(2000000), "aoa-node60", 1544421807},
		{"0xc5f4ae494d7bd558b2035b63318b2a2cc47b0c2b", uint64(2000000), "aoa-node61", 1544421807},
		{"0xc1407b1974cc0ed2f6c1006c222b4b9893174323", uint64(2000000), "aoa-node62", 1544421807},
		{"0xb4008ab135ec7ced2b5b05216ef74bf1cd835170", uint64(2000000), "aoa-node63", 1544421807},
		{"0xdfd00dba4589b55dbaf812b5634b374a774b02ab", uint64(2000000), "aoa-node64", 1544421807},
		{"0x580322da3b0878a9b9bf70d746c681d407058952", uint64(2000000), "aoa-node65", 1544421807},
		{"0xb85a21534edc46fc331f7d3f51ea15c219b9ac31", uint64(2000000), "aoa-node66", 1544421807},
		{"0xf9bc1507ae2ec3b937d5f6fca42afa3f8cf57e7d", uint64(2000000), "aoa-node67", 1544421807},
		{"0xc36e16ab11b20c67bce7801247eecc196e13be76", uint64(2000000), "aoa-node68", 1544421807},
		{"0x7428eff1d1037357983af28829fb59c4fee02d2e", uint64(2000000), "aoa-node69", 1544421807},
		{"0x2ce4b50f320f81167719d3ba559960172f0f6fe4", uint64(2000000), "aoa-node70", 1544421807},
		{"0x0e67571505632bb30cab56c915593a3cd2b309f2", uint64(2000000), "aoa-node71", 1544421807},
		{"0x092fed360946dd1f11de7da0d8a3b731ebd4e5e7", uint64(2000000), "aoa-node72", 1544421807},
		{"0x69fe4725fb71142725814687aca7190c2dd138d8", uint64(2000000), "aoa-node73", 1544421807},
		{"0xe409807b63f317d6f223fcab05373413e334f7cf", uint64(2000000), "aoa-node74", 1544421807},
		{"0xa82c697e57898cc14e9d4270aec4456393b28ae0", uint64(2000000), "aoa-node75", 1544421807},
		{"0xbf93f54597b481c1b17fe06b7e2caec4e06c938c", uint64(2000000), "aoa-node76", 1544421807},
		{"0x9e824f10c206d6f6ab9e1c1afd9d00ff8db4bdc8", uint64(2000000), "aoa-node77", 1544421807},
		{"0xea278feefa5697a35e71da435521f3276086e16b", uint64(2000000), "aoa-node78", 1544421807},
		{"0xf118a7719460a99c7222554ffdd44e9463aea674", uint64(2000000), "aoa-node79", 1544421807},
		{"0x62de91d4286dfdd36cbc06c99b39005bb493c552", uint64(2000000), "aoa-node80", 1544421807},
		{"0x116d54909c42b7588b006a1bdb6bbb5186c8eaf7", uint64(2000000), "aoa-node81", 1544421807},
		{"0x3386d282ef9334aedc6c664999f5b6a7e7084a2e", uint64(2000000), "aoa-node82", 1544421807},
		{"0x5454c44dfbc5d60f2575231069b1168c99e4dabe", uint64(2000000), "aoa-node83", 1544421807},
		{"0x56232f4a299d343aba2620cbc11f7ab731169f96", uint64(2000000), "aoa-node84", 1544421807},
		{"0xbeb755aae22eace8d8dbbe92e229edf9d086caef", uint64(2000000), "aoa-node85", 1544421807},
		{"0x26c452286fd9f25251213e056ef1f8533b5468cd", uint64(2000000), "aoa-node86", 1544421807},
		{"0xa0afac3fc553f41cdf901060db66ef8493a355e2", uint64(2000000), "aoa-node87", 1544421807},
		{"0x98f78fe0cb1b3f5643d1d5bf54d9fcd2241f6a42", uint64(2000000), "aoa-node88", 1544421807},
		{"0x649d293e167658e899d8c96cf5f5ac5c5624309e", uint64(2000000), "aoa-node89", 1544421807},
		{"0xabb431653084c2b7ccac452f94a373eb976aa08f", uint64(2000000), "aoa-node90", 1544421807},
		{"0xa5f3394b2f11b2e288fa2f4e436cb30dcde521c7", uint64(2000000), "aoa-node91", 1544421807},
		{"0xd5f15ade48ecfd7425807544b82ee0eb6e272e2d", uint64(2000000), "aoa-node92", 1544421807},
		{"0xd9d05e7dae8727bb90e3a51e5231cb28aa0b4c2f", uint64(2000000), "aoa-node93", 1544421807},
		{"0xcd45eecc7c899054fb7ffda9b03d27ad339fd3c5", uint64(2000000), "aoa-node94", 1544421807},
		{"0x17cf6befd3ebc0ebc4815ad1675f0a17e3ef0e3a", uint64(2000000), "aoa-node95", 1544421807},
		{"0xb2afa200ef65564393f3446dc779ba8127a793b8", uint64(2000000), "aoa-node96", 1544421807},
		{"0x09e3b312613cd9ccbe041cc01b3d4e2ac78e6cb4", uint64(2000000), "aoa-node97", 1544421807},
		{"0xcd75375984cacf5bb9bd03e505ba2a97e4f7a37c", uint64(2000000), "aoa-node98", 1544421807},
		{"0x409ae4e6c465026bb67e665d7f49cb9bafdf3033", uint64(2000000), "aoa-node99", 1544421807},
		{"0x1858803562fc7ff10b3f224d7b6d538fa2a69523", uint64(2000000), "aoa-node100", 1544421807},
		{"0xac735a15e1ec9ff3dda829202a4359dcac4b0cfb", uint64(2000000), "aoa-node101", 1544421807},
	}
	return candidateList
}
