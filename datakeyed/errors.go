package datakeyed

// xlOverlay_go/datakeyed/keyQueue.go

import (
	e "errors"
)

var (
	NilCallBack = e.New("nil CallBack parameter")
	NilNodeID   = e.New("nil NodeID parameter")
	NilMemCache = e.New("nil MemCache parameter")
)
