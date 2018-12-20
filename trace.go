package ethrpc

import (
	"fmt"

	"github.com/islishude/go-eth-rpc/methods"
	"github.com/islishude/go-jsonrpc2"
)

// ParityTrxTrace is for parity `trace_transaction`
//   if type == suicide then `result` is nil
//   if error != nil then `result` is nil
type ParityTrxTrace struct {
	Error        *string           `json:"error"`
	Type         string            `json:"type"`
	SubTrace     int               `json:"subtrace"`
	TraceAddress []int             `json:"traceAddress"`
	TrxPosition  int               `json:"transactionPosition"`
	Result       map[string]string `json:"result"`
	Action       map[string]string `json:"action"`
}

// ParityTrxTraceResp is result of parity trace_transaction
type ParityTrxTraceResp struct {
	jr2.RespInlineData
	Result []*ParityTrxTrace `json:"result"`
}

// TraceTrxByParity call parity trace_transaction
func (c *Ethereum) TraceTrxByParity(txid string) ([]*ParityTrxTrace, error) {
	resChan, errChan := make(chan []*ParityTrxTrace, 1), make(chan error, 1)

	go func() {
		defer close(resChan)
		defer close(errChan)

		var resp ParityTrxTraceResp
		if err := c.CallOne(jr2.NewReqData(txid, methods.ParityTraceTrx, txid), &resp); err != nil {
			errChan <- err
			return
		}

		if err := resp.Error; err != nil {
			errChan <- fmt.Errorf("[EthRPC::TraceTrxByParity] Msg %v Code %v", err["message"], err["code"])
			return
		}

		resChan <- resp.Result
	}()

	return <-resChan, <-errChan
}
