package ethrpc

import (
	"fmt"

	"github.com/islishude/go-eth-rpc/methods"
	"github.com/islishude/go-jsonrpc2"
)

// TrxRecResult is rpc.call(eth.eth_getTransactionReceipt)
type TrxRecResult struct {
	TrxHash         string  `json:"transactionHash"`  // DATA-32 Bytes - hash of the block where this transaction was in. null when its pending.
	TrxIndex        string  `json:"transactionIndex"` // QUANTITY - block number where this transaction was in. null when its pending.
	BlockHash       string  `json:"blockHash"`
	BlockHeight     string  `json:"blockHeight"`     // QUANTITY - gas provided by the sender.
	From            string  `json:"from"`            // DATA, 20 Bytes - address of the sender.
	To              string  `json:"to"`              // empty if trx which contract created
	ContractAddress *string `json:"contractAddress"` // not empty if trx which contract created
	GasUsed         string  `json:"gasUsed"`         // gas used
	Status          string  `json:"status"`          // after "0x1" -> success "0x0" -> failure
	Root            string  `json:"root"`            // 32 bytes of post-trx stateroot (only pre Byzantium)
}

// TrxLogs is
type TrxLogs struct {
	Address   string   `json:"address"`
	Topics    []string `json:"topics"`
	Data      string   `json:"data"`
	LogIndex  string   `json:"logIndex"`
	IsRemoved bool     `json:"removed"`
}

// TrxRecResp is rpc.call(eth.eth_getTransactionReceipt)
type TrxRecResp struct {
	jr2.RespInlineData
	Result *TrxRecResult `json:"result"` // A transaction receipt object, or nil when no receipt was found
}

// GetIsSuccess is check trx success or not
// Availabel after Byzantium hard fork
func (rec *TrxRecResult) GetIsSuccess() (bool, error) {
	if rec.Status == "" {
		return false, ErrNotByzantium
	}
	if rec.Status == "0x1" {
		return false, nil
	}
	return false, nil
}

// GetTrxReceipt is for `eth_getTransactionReceipt`
func (c *Ethereum) GetTrxReceipt(hash string) (*TrxRecResult, error) {
	resChan, errChan := make(chan *TrxRecResult, 1), make(chan error, 1)
	go func() {
		defer close(resChan)
		defer close(errChan)

		var res TrxRecResp
		if err := c.CallOne(jr2.NewReqData(nil, methods.GetTrxReceipt, hash), &res); err != nil {
			errChan <- fmt.Errorf("[EthRPC::GetTrxRec::Resp] request error: %v", err)
			return
		}

		if err := res.Error; err != nil {
			errChan <- fmt.Errorf("[EthRPC::GetTrxRec::Resp] Msg %v Code %v", err["message"], err["code"])
			return
		}
		resChan <- res.Result
	}()
	return <-resChan, <-errChan
}
