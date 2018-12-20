package ethrpc

import (
	"fmt"

	"github.com/islishude/go-eth-rpc/methods"

	"github.com/islishude/go-jsonrpc2"
)

// TrxResult is rpc.call(eth_getTransactionByHash) returns
type TrxResult struct {
	Hash      string `json:"hash"`        // DATA, 32 Bytes - hash of the transaction.
	Nonce     string `json:"nonce"`       // QUANTITY - the number of transactions made by the sender prior to this one.
	Height    string `json:"blockNumber"` // QUANTITY - block number where this transaction was in. empty when its pending.
	BlockHash string `json:"blockHash"`   // DATA, 32 Bytes - hash of the block where this transaction was in. empty when its pending.
	From      string `json:"from"`        // DATA, 20 Bytes - address of the sender.
	To        string `json:"to"`          // DATA, 20 Bytes - address of the receiver. empty when its a contract creation transaction.
	Value     string `json:"value"`       // QUANTITY - value transferred in Wei.
	GasLimit  string `json:"gas"`         // QUANTITY - gas provided by the sender.
	GasPrice  string `json:"gasPrice"`    // QUANTITY - gas price provided by the sender in Wei.
	Input     string `json:"input"`       // DATA - the data send along with the transaction.
	V         string `json:"v"`           // QUANTITY - ECDSA recovery id
	R         string `json:"r"`           // DATA, 32 Bytes - ECDSA signature r
	S         string `json:"s"`           // DATA, 32 Bytes - ECDSA signature s
}

// TrxResp is rpc.call(eth_getTransactionByHash) returns
type TrxResp struct {
	jr2.RespInlineData
	Result *TrxResult `json:"result"` // A trx object, or nil when no trx was found
}

// GetTrxByHash is for `eth_getTransactionByHash`
func (c *Ethereum) GetTrxByHash(hash string) (*TrxResult, error) {
	resChan, errChan := make(chan *TrxResult, 1), make(chan error, 1)
	go func() {
		defer close(resChan)
		defer close(errChan)

		var resp TrxResp
		if err := c.CallOne(jr2.NewReqData(nil, methods.GetTrxByHash, hash), &resp); err != nil {
			errChan <- fmt.Errorf("[EthRPC::GetTrxByHash::Resp] request error: %v", err)
			return
		}

		if err := resp.Error; err != nil {
			errChan <- fmt.Errorf("[EthRPC::GetTrxByHash::Resp] Msg %v Code %v", err["message"], err["code"])
			return
		}

		resChan <- resp.Result
	}()
	return <-resChan, <-errChan
}
