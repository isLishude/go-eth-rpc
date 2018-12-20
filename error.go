package ethrpc

import "errors"

// Error list
var (
	ErrParamLen     = errors.New("Param length should be one or two")
	ErrHexParse     = errors.New("Parse hex string to bigInt error")
	ErrNotByzantium = errors.New("Use GetIsSuccess post-Byzantium fork(4,370,000 height)")
)
