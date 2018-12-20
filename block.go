package ethrpc

import (
	"fmt"
	"math/big"

	"github.com/islishude/go-eth-rpc/methods"
	"github.com/islishude/go-jsonrpc2"
)

// ParseBig256 parses hex string to a 256 bits integer.
func (c *Ethereum) ParseBig256(hex string) (bigint *big.Int, ok bool) {
	if hex == "" {
		return new(big.Int), true
	}
	if len(hex) >= 2 && (hex[:2] == "0x" || hex[:2] == "0X") {
		bigint, ok = new(big.Int).SetString(hex[2:], 16)
	} else {
		bigint, ok = new(big.Int).SetString(hex, 10)
	}
	if ok && bigint.BitLen() > 256 {
		bigint, ok = nil, false
	}
	return bigint, ok
}

// GetHeight call `eth_getBlock`
func (c *Ethereum) GetHeight() (uint64, error) {
	resChan, errChan := make(chan uint64, 1), make(chan error, 1)
	go func() {
		defer close(resChan)
		defer close(errChan)

		var res StrResp
		if err := c.CallOne(jr2.NewReqData(nil, methods.GetHeight), &res); err != nil {
			errChan <- err
			return
		}

		if err := res.Error; err != nil {
			errChan <- fmt.Errorf("[EthRPC::GetHeight::Resp] Msg %v Code %v", err["message"], err["code"])
			return
		}

		if height, ok := c.ParseBig256(res.Result); ok {
			resChan <- height.Uint64()
			return
		}
		errChan <- ErrHexParse
	}()
	return <-resChan, <-errChan
}

// BlockResult is result of rpc.call(eth_getBlockByNumber)
type BlockResult struct {
	Height          string   `json:"number"`
	Hash            string   `json:"string"`
	ParentHash      string   `json:"parentHash"`
	Nonce           string   `json:"nonce"`
	SHA3Uncles      string   `json:"sha3Uncles"`
	LogsBloom       string   `json:"logsBloom"`
	TrxRoot         string   `json:"transactionRoot"`
	StateRoot       string   `json:"stateRoot"`
	Miner           string   `json:"miner"`
	Difficulty      string   `json:"difficulty"`
	TotalDifficulty string   `json:"totalDifficulty"`
	ExtraData       string   `json:"extraData"`
	Size            string   `json:"size"`
	GasLimit        string   `json:"gasLimit"`
	GasUsed         string   `json:"gasUsed"`
	Timestamp       string   `json:"timestamp"`
	Uncles          []string `json:"uncles"`
	TrxList         []string `json:"transactions"`
}

// BlockResp is rpc.call(eth_getBlockByNumber) returns
type BlockResp struct {
	jr2.RespInlineData
	Result *BlockResult `json:"result"`
}

// GetBlock is for `eth_getBlockByNumber`
func (c *Ethereum) GetBlock(height uint32) (*BlockResult, error) {
	resChan, errChan := make(chan *BlockResult, 1), make(chan error, 1)
	go func() {
		defer close(resChan)
		defer close(errChan)

		var res BlockResp
		if err := c.CallOne(jr2.NewReqData(nil, methods.GetBlockByHeight, fmt.Sprintf("%#x", height), false), &res); err != nil {
			errChan <- fmt.Errorf("[EthRPC::GetBlock::Resp] request error: %v", err)
			return
		}

		if err := res.Error; err != nil {
			errChan <- fmt.Errorf("[EthRPC::GetBlock::Resp] Msg %v Code %v", err["message"], err["code"])
			return
		}

		resChan <- res.Result
	}()
	return <-resChan, <-errChan
}
