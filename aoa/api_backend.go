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

package aoa

import (
	"context"
	"math/big"

	"github.com/Aurorachain-io/go-aoa/accounts"
	aa "github.com/Aurorachain-io/go-aoa/accounts/walletType"
	"github.com/Aurorachain-io/go-aoa/common"
	"github.com/Aurorachain-io/go-aoa/common/math"
	"github.com/Aurorachain-io/go-aoa/core"
	"github.com/Aurorachain-io/go-aoa/core/bloombits"
	"github.com/Aurorachain-io/go-aoa/core/state"
	"github.com/Aurorachain-io/go-aoa/core/types"
	"github.com/Aurorachain-io/go-aoa/core/vm"
	"github.com/Aurorachain-io/go-aoa/core/watch"
	"github.com/Aurorachain-io/go-aoa/aoa/downloader"
	"github.com/Aurorachain-io/go-aoa/aoa/gasprice"
	"github.com/Aurorachain-io/go-aoa/aoadb"
	"github.com/Aurorachain-io/go-aoa/event"
	"github.com/Aurorachain-io/go-aoa/params"
	"github.com/Aurorachain-io/go-aoa/rpc"
)

//implements emapi.Backend for full nodes
type DacApiBackend struct {
	dac *Dacchain
	gpo *gasprice.Oracle
}

func (b *DacApiBackend) ChainConfig() *params.ChainConfig {
	return b.dac.chainConfig
}

func (b *DacApiBackend) GetDelegateWalletInfoCallback() func(data *aa.DelegateWalletInfo) {
	return b.dac.protocolManager.GetAddDelegateWalletCallback()
}

func (b *DacApiBackend) CurrentBlock() *types.Block {
	return b.dac.blockchain.CurrentBlock()
}

func (b *DacApiBackend) SetHead(number uint64) {
	b.dac.protocolManager.downloader.Cancel()
	b.dac.blockchain.SetHead(number)
}

func (b *DacApiBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	// Pending block is only known by the delegate
	if blockNr == rpc.PendingBlockNumber {
		block := b.dac.dposMiner.PendingBlock()
		return block.Header(), nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.dac.blockchain.CurrentBlock().Header(), nil
	}
	return b.dac.blockchain.GetHeaderByNumber(uint64(blockNr)), nil
}

func (b *DacApiBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	// Pending block is only known by the delegate
	if blockNr == rpc.PendingBlockNumber {
		block := b.dac.dposMiner.PendingBlock()
		return block, nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.dac.blockchain.CurrentBlock(), nil
	}
	return b.dac.blockchain.GetBlockByNumber(uint64(blockNr)), nil
}

func (b *DacApiBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	// Pending state is only known by the delegate
	if blockNr == rpc.PendingBlockNumber {
		currentBlock := b.dac.blockchain.CurrentBlock()
		statedb, _ := b.dac.blockchain.StateAt(currentBlock.Root())
		return statedb, currentBlock.Header(), nil
	}
	// Otherwise resolve the block number and return its state
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, nil, err
	}
	stateDb, err := b.dac.BlockChain().StateAt(header.Root)
	return stateDb, header, err
}

func (b *DacApiBackend) GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error) {
	return b.dac.blockchain.GetBlockByHash(blockHash), nil
}

func (b *DacApiBackend) GetDelegatePoll(block *types.Block) (*map[common.Address]types.Candidate, error) {
	delegateDB, err := b.dac.blockchain.DelegateStateAt(block.DelegateRoot())
	if err != nil {
		return nil, err
	}
	delegates := delegateDB.GetDelegates()
	res := make(map[common.Address]types.Candidate)
	for _, delegate := range delegates {
		address := common.HexToAddress(delegate.Address)
		res[address] = delegate
	}
	return &res, nil
}

func (b *DacApiBackend) GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error) {
	return core.GetBlockReceipts(b.dac.chainDb, blockHash, core.GetBlockNumber(b.dac.chainDb, blockHash)), nil
}

func (b *DacApiBackend) GetTd(blockHash common.Hash) *big.Int {
	return b.dac.blockchain.GetTdByHash(blockHash)
}

func (b *DacApiBackend) GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error) {
	state.SetBalance(msg.From(), math.MaxBig256)
	vmError := func() error { return nil }

	ctext := core.NewEVMContext(msg, header, b.dac.BlockChain(), nil)
	return vm.NewEVM(ctext, state, b.dac.chainConfig, vmCfg), vmError, nil
}

func (b *DacApiBackend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	return b.dac.BlockChain().SubscribeRemovedLogsEvent(ch)
}

func (b *DacApiBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	return b.dac.BlockChain().SubscribeChainEvent(ch)
}

func (b *DacApiBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return b.dac.BlockChain().SubscribeChainHeadEvent(ch)
}

func (b *DacApiBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	return b.dac.BlockChain().SubscribeChainSideEvent(ch)
}

func (b *DacApiBackend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return b.dac.BlockChain().SubscribeLogsEvent(ch)
}

func (b *DacApiBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.dac.txPool.AddLocal(signedTx)
}

func (b *DacApiBackend) GetPoolTransactions() (types.Transactions, error) {
	pending, err := b.dac.txPool.Pending()
	if err != nil {
		return nil, err
	}
	var txs types.Transactions
	for _, batch := range pending {
		txs = append(txs, batch...)
	}
	return txs, nil
}

func (b *DacApiBackend) GetPoolTransaction(hash common.Hash) *types.Transaction {
	return b.dac.txPool.Get(hash)
}

func (b *DacApiBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return b.dac.txPool.State().GetNonce(addr), nil
}

func (b *DacApiBackend) Stats() (pending int, queued int) {
	return b.dac.txPool.Stats()
}

func (b *DacApiBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return b.dac.TxPool().Content()
}

func (b *DacApiBackend) SubscribeTxPreEvent(ch chan<- core.TxPreEvent) event.Subscription {
	return b.dac.TxPool().SubscribeTxPreEvent(ch)
}

func (b *DacApiBackend) Downloader() *downloader.Downloader {
	return b.dac.Downloader()
}

func (b *DacApiBackend) ProtocolVersion() int {
	return b.dac.EthVersion()
}

func (b *DacApiBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	return b.gpo.SuggestPrice(ctx)
}

func (b *DacApiBackend) ChainDb() aoadb.Database {
	return b.dac.ChainDb()
}

func (b *DacApiBackend) AccountManager() *accounts.Manager {
	return b.dac.AccountManager()
}

func (b *DacApiBackend) BloomStatus() (uint64, uint64) {
	sections, _, _ := b.dac.bloomIndexer.Sections()
	return params.BloomBitsBlocks, sections
}

func (b *DacApiBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	for i := 0; i < bloomFilterThreads; i++ {
		go session.Multiplex(bloomRetrievalBatch, bloomRetrievalWait, b.dac.bloomRequests)
	}
}

func (b *DacApiBackend) IsWatchInnerTxEnable() bool {
	return b.dac.config.EnableInterTxWatching
}

func (b *DacApiBackend) GetInnerTxDb() watch.InnerTxDb {
	return b.dac.BlockChain().GetInnerTxDb()
}
