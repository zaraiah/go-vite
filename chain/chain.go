package chain

import (
	"github.com/vitelabs/go-vite/chain_db"
	"github.com/vitelabs/go-vite/compress"
	"github.com/vitelabs/go-vite/config"
	"github.com/vitelabs/go-vite/log15"
	"github.com/vitelabs/go-vite/trie"
	"path/filepath"
	"sync"
)

type Chain struct {
	log        log15.Logger
	chainDb    *chain_db.ChainDb
	compressor *compress.Compressor

	trieNodePool  *trie.TrieNodePool
	stateTriePool *StateTriePool

	createAccountLock sync.Mutex

	needSnapshotCache *NeedSnapshotCache
}

func NewChain(cfg *config.Config) *Chain {
	chain := &Chain{
		log: log15.New("module", "chain"),
	}

	chain.stateTriePool = NewStateTriePool(chain)

	chainDb := chain_db.NewChainDb(filepath.Join(cfg.DataDir, "chain"))
	if chainDb == nil {
		chain.log.Error("NewChain failed")
		return nil
	}
	chain.chainDb = chainDb

	compressor := compress.NewCompressor()
	chain.compressor = compressor

	chain.trieNodePool = trie.NewTrieNodePool()

	chain.needSnapshotCache = NewNeedSnapshotContent(chain)

	return chain
}

func (c *Chain) Start() {
	// Start compress in the background
	c.compressor.Start()
}

func (c *Chain) Stop() {
	// Stop compress
	c.compressor.Stop()
}
