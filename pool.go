package calc

import (
	"math/big"
	"sync"
)

var pool = sync.Pool{
	New: func() interface{} {
		return new(big.Float).SetPrec(big.MaxPrec)
	},
}

func reuse() *big.Float {
	return pool.Get().(*big.Float)
}

func release(xs ...*big.Float) {
	for i := range xs {
		xs[i].SetFloat64(0)
		pool.Put(xs[i])
	}
}
