package ethrpc

import (
	"fmt"

	"github.com/islishude/go-eth-rpc/methods"
	"github.com/islishude/go-jsonrpc2"
)

/*
TrxPool is result of `txpool_content`

	export interface IEthTxPoolContent {
		pending: {
			[address: string]: {
				[nonce: string]: IEthTx[];
			};
		};
		queued: {
			[address: string]: {
				[nonce: string]: IEthTx[];
			};
		};
	}

Please note, there may be multiple transactions associated with the same account
and nonce. This can happen if the user broadcast multiple ones with varying gas
allowances (or even completely different transactions).
*/
type TrxPool struct {
	Pending map[string]map[string][]*TrxResult `json:"pending"`
	Queued  map[string]map[string][]*TrxResult `json:"queued"`
}

// TrxPoolResp is response of txpool_content
type TrxPoolResp struct {
	jr2.RespInlineData
	Result *TrxPool `json:"result"`
}

// GetTrxPoolConten is for `txpool_content`
func (c *Ethereum) GetTrxPoolConten() (*TrxPool, error) {
	resChan, errChan := make(chan *TrxPool, 1), make(chan error, 1)
	go func() {
		defer close(resChan)
		defer close(errChan)

		var resp TrxPoolResp
		if err := c.CallOne(jr2.NewReqData(nil, methods.TrxPoolContent), &resp); err != nil {
			errChan <- fmt.Errorf("[EthRPC::GetTrxPool::Resp] request error: %v", err)
			return
		}
		if err := resp.Error; err != nil {
			errChan <- fmt.Errorf("[EthRPC::GetTrxPool::Resp] Msg %v Code %v", err["message"], err["code"])
			return
		}
		resChan <- resp.Result
	}()
	return <-resChan, <-errChan
}
