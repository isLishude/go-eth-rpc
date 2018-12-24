package ethrpc

import (
	"encoding/hex"
	"strconv"

	"github.com/islishude/go-eth-rpc/methods"
	"github.com/islishude/go-jsonrpc2"
)

// ERC20Sigs
const (
	ERC20BalanceOfSig    = "0x70a08231"
	ERC20DecimalsSig     = "0x313ce567"
	ERC20DECIMALSSig     = "0x2e0f2625"
	ERC20NameSig         = "0x06fdde03"
	ERC20NAMESig         = "0xa3f4df7e"
	ERC20SymbolSig       = "0x95d89b41"
	ERC20SYMBOLSig       = "0xf76f8d78"
	ERC20TransferSig     = "0xa9059cbb"
	ERC20TransferFromSig = "0x23b872dd"
	ERC20TotalSupply     = "0x18160ddd"
)

// Util
const (
	zeroPadding = "000000000000000000000000"
)

func rmHexPrefix(input string) string {
	if input[:2] == "0x" || input[:2] == "0X" {
		input = input[2:]
	}
	return input
}

func paddingAddress(address string) string {
	return zeroPadding + rmHexPrefix(address)
}

// GetERC20Balance get ERC20 Token balance
func (c *Ethereum) GetERC20Balance(contract, address string) (string, error) {
	resChan, errChan := make(chan string, 1), make(chan error, 1)
	go func() {
		defer close(resChan)
		defer close(errChan)

		reqData := CallFuncParam{
			Data: ERC20BalanceOfSig + paddingAddress(address),
			To:   contract,
		}

		var resp StrResp
		c.CallOne(jr2.NewReqData(nil, methods.Call, &reqData), &resp)
		// TODO(islishude): check error and parse result to *big.Int
		resChan <- string(resp.Result)
	}()

	return <-resChan, <-errChan
}

// GetERC20Name is
func (c *Ethereum) GetERC20Name(contract string) (string, error) {
	resChan, errChan := make(chan string, 1), make(chan error, 1)
	go func() {
		defer close(resChan)
		defer close(errChan)

		reqData := CallFuncParam{
			Data: ERC20NameSig,
			To:   contract,
		}

		var resp StrResp
		err := c.CallOne(jr2.NewReqData(nil, methods.Call, &reqData), &resp)
		result := rmHexPrefix(resp.Result)
		// offset(64) + length(64) + data(>=64) >= 192
		if len(result) < 192 {
			return
		}
		if err != nil {
			errChan <- err
			return
		}

		nameLen, err := strconv.ParseInt(result[64:128], 16, 16)
		if err != nil {
			errChan <- err
			return
		}

		data, err := hex.DecodeString(result[128 : nameLen*2])
		if err != nil {
			errChan <- err
			return
		}
		resChan <- string(data)
	}()
	return <-resChan, <-errChan
}

// GetERC20Symbol is
func (c *Ethereum) GetERC20Symbol(contract string) (string, error) {
	resChan, errChan := make(chan string, 1), make(chan error, 1)
	go func() {
		defer close(resChan)
		defer close(errChan)

		reqData := CallFuncParam{
			Data: ERC20SymbolSig,
			To:   contract,
		}

		var resp StrResp
		err := c.CallOne(jr2.NewReqData(nil, methods.Call, &reqData), &resp)

		result := rmHexPrefix(resp.Result)
		// offset(64) + length(64) + data(>=64) >= 192
		if len(result) < 192 {
			return
		}
		if err != nil {
			errChan <- err
			return
		}

		nameLen, err := strconv.ParseInt(result[64:128], 16, 16)
		if err != nil {
			errChan <- err
			return
		}

		data, err := hex.DecodeString(result[128 : nameLen*2])
		if err != nil {
			errChan <- err
			return
		}
		resChan <- string(data)
	}()
	return <-resChan, <-errChan
}

// GetERC20Decimals is
func (c *Ethereum) GetERC20Decimals(contract string) (uint8, error) {
	resChan, errChan := make(chan uint8, 1), make(chan error, 1)
	go func() {
		defer close(resChan)
		defer close(errChan)

		reqData := CallFuncParam{
			Data: ERC20DecimalsSig,
			To:   contract,
		}

		var resp StrResp
		err := c.CallOne(jr2.NewReqData(nil, methods.Call, &reqData), &resp)
		result := rmHexPrefix(resp.Result)
		if result == "" {
			return
		}
		if err != nil {
			errChan <- err
			return
		}

		decimals, err := strconv.ParseInt(result, 16, 16)
		if err != nil {
			errChan <- err
			return
		}
		resChan <- uint8(decimals)
	}()
	return <-resChan, <-errChan
}

// GetERC20TotalSupply is
func (c *Ethereum) GetERC20TotalSupply(contract string) (string, error) {
	resChan, errChan := make(chan string, 1), make(chan error, 1)
	go func() {
		defer close(resChan)
		defer close(errChan)

		reqData := CallFuncParam{
			Data: ERC20TotalSupply,
			To:   contract,
		}

		var resp StrResp
		c.CallOne(jr2.NewReqData(nil, methods.Call, &reqData), &resp)
		// TODO(islishude): check error and parse result to *big.Int
		resChan <- string(resp.Result)
	}()
	return <-resChan, <-errChan
}
