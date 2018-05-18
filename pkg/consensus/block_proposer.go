package consensus

import (
	"context"
	"log"

	"github.com/dfinity/go-dfinity-crypto/bls"
)

// Validator validates the system transaction.
type Validator interface {
	Validate(SysTxn) bool
}

// BlockProposer produces one block proposal if it is in the block
// proposal committee in the current round.
type BlockProposer struct {
	sk     bls.SecretKey
	leader leader
}

// NewBlockProposer creates a new block proposer.
func NewBlockProposer(sk bls.SecretKey, leader leader) *BlockProposer {
	return &BlockProposer{sk: sk, leader: leader}
}

// CollectTxn collects transactions and returns a block proposal when
// the context is done.
func (b *BlockProposer) CollectTxn(ctx context.Context, txCh chan []byte, sysTxCh chan SysTxn, pendingTx chan []byte) *BlockProposal {
	var bp BlockProposal
	bp.PrevBlock = hash(b.leader.Block.Encode(true))
	bp.Round = b.leader.Block.Round + 1
	bp.Owner = hash(b.sk.GetPublicKey().Serialize()).Addr()
	transition := b.leader.State.Transition()
	sysTransition := b.leader.SysState.Transition()
	for {
		select {
		case <-ctx.Done():
			bp.SysTxns = sysTransition.Txns()
			bp.Data = transition.Encode()
			close(pendingTx)
			return &bp
		case tx := <-txCh:
			valid, future := transition.Record(tx)
			if !valid {
				log.Printf("received invalid txn, len: %d\n", len(tx))
				continue
			}

			if future {
				pendingTx <- tx
			}
		case sysTx := <-sysTxCh:
			if !sysTransition.Record(sysTx) {
				log.Println("received invalid sys txn")
				continue
			}
		}
	}
}
