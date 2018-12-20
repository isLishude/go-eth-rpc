package ethrpc

import (
	"fmt"

	"github.com/islishude/go-eth-rpc/methods"
	"github.com/islishude/go-jsonrpc2"
)

// StrResp is string result response of jsonrpc
type StrResp struct {
	jr2.RespInlineData
	Result string `json:"result"`
}

// GetBalance is call `eth_getBalance`
func (c *Ethereum) GetBalance(param ...string) (string, error) {
	resChan, errChan := make(chan string, 1), make(chan error, 1)

	go func() {
		defer close(resChan)
		defer close(errChan)

		var req *jr2.RequestData
		switch len(param) {
		case 1:
			req = jr2.NewReqData(nil, methods.GetBalance, param[0], "latest")
		case 2:
			req = jr2.NewReqData(nil, methods.GetBalance, param[0], param[1])
		default:
			errChan <- ErrParamLen
			return
		}

		var res StrResp
		if err := c.CallOne(req, &res); err != nil {
			errChan <- err
			return
		}
		if err := res.Error; err != nil {
			errChan <- fmt.Errorf("[Eth::GetBalance] response error code %v msg: %v", err["code"], err["message"])
			return
		}

		resChan <- res.Result
	}()

	return <-resChan, <-errChan
}

// GetNonce get address nonce
func (c *Ethereum) GetNonce(param ...string) (string, error) {
	resChan, errChan := make(chan string, 1), make(chan error, 1)
	go func() {
		defer close(resChan)
		defer close(errChan)

		var req *jr2.RequestData
		switch len(param) {
		case 1:
			req = jr2.NewReqData(nil, methods.GetNonce, param[0], "latest")
		case 2:
			req = jr2.NewReqData(nil, methods.GetNonce, param[0], param[1])
		default:
			errChan <- ErrParamLen
			return
		}

		var res StrResp
		if err := c.CallOne(req, &res); err != nil {
			errChan <- err
			return
		}
		if err := res.Error; err != nil {
			errChan <- fmt.Errorf("[Eth::GetNonce] response error code %v msg: %v", err["code"], err["message"])
			return
		}
		resChan <- res.Result
	}()
	return <-resChan, <-errChan
}

// CallFuncParam is param for `eth_call`
type CallFuncParam struct {
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Gas      string `json:"gas,omitempty"`
	GasPrice string `json:"gasPrice,omitempty"`
	Value    string `json:"value,omitempty"`
	Data     string `json:"data,omitempty"`
}

// CallFunc is for `eth_call`
func (c *Ethereum) CallFunc(input *CallFuncParam) (string, error) {
	resChan, errChan := make(chan string, 1), make(chan error, 1)

	go func() {
		defer close(resChan)
		defer close(errChan)

		var resp StrResp
		if err := c.CallOne(jr2.NewReqData(nil, methods.Call, input), &resp); err != nil {
			errChan <- err
			return
		}
		if err := resp.Error; err != nil {
			errChan <- fmt.Errorf("[Eth::CallFunc] response error code %v msg: %v", err["code"], err["message"])
			return
		}
		resChan <- resp.Result
	}()

	return <-resChan, <-errChan
}
