package ethrpc

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"

	"github.com/islishude/go-eth-rpc/methods"
	"github.com/islishude/go-jsonrpc2"
)

// ErrorList
var (
	ErrNotERC20 = errors.New("Not standard ERC20 Token")
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
		if err := c.CallOne(jr2.NewReqData(nil, methods.Call, &reqData), &resp); err != nil {
			errChan <- err
			return
		}

		if err := resp.Error; err != nil {
			// error check for parity
			if code, ok := err["code"].(int); ok && code == -32015 {
				return
			}

			errChan <- fmt.Errorf("[GetERC20Balance response error] msg: %v code: %v", err["message"], err["code"])
			return
		}

		if resp.Result == "0x" {
			errChan <- ErrNotERC20
			return
		}

		// TODO(islishude): parse result to *big.Int
		resChan <- resp.Result
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
		if err := c.CallOne(jr2.NewReqData(nil, methods.Call, &reqData), &resp); err != nil {
			errChan <- err
			return
		}

		if err := resp.Error; err != nil {
			// error check for parity
			if code, ok := err["code"].(int); ok && code == -32015 {
				errChan <- ErrNotERC20
				return
			}

			errChan <- fmt.Errorf("[GetERC20Name response error] msg: %v code: %v", err["message"], err["code"])
			return
		}

		result := rmHexPrefix(resp.Result)
		// offset(64) + length(64) + data(>=64) >= 192
		if len(result) < 192 {
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
		if err := c.CallOne(jr2.NewReqData(nil, methods.Call, &reqData), &resp); err != nil {
			errChan <- err
			return
		}

		if err := resp.Error; err != nil {
			if code, ok := err["code"].(int); ok && code == -32015 {
				errChan <- ErrNotERC20
				return
			}

			errChan <- fmt.Errorf("[GetERC20Symbol response error] msg: %v code: %v", err["message"], err["code"])
			return
		}

		result := rmHexPrefix(resp.Result)
		// offset(64) + length(64) + data(>=64) >= 192
		if len(result) < 192 {
			errChan <- ErrNotERC20
			return
		}

		nameLen, err := strconv.ParseInt(result[64:128], 16, 16)
		if err != nil {
			errChan <- ErrNotERC20
			return
		}

		data, err := hex.DecodeString(result[128 : nameLen*2])
		if err != nil {
			errChan <- ErrNotERC20
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
		if err := c.CallOne(jr2.NewReqData(nil, methods.Call, &reqData), &resp); err != nil {
			errChan <- err
			return
		}

		if err := resp.Error; err != nil {
			// error check for parity
			if code, ok := err["code"].(int); ok && code == -32015 {
				errChan <- ErrNotERC20
				return
			}

			errChan <- fmt.Errorf("[GetERC20Decimals response error] msg: %v code: %v", err["message"], err["code"])
			return
		}

		result := rmHexPrefix(resp.Result)
		if result == "" {
			errChan <- ErrNotERC20
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
		if err := c.CallOne(jr2.NewReqData(nil, methods.Call, &reqData), &resp); err != nil {
			errChan <- err
			return
		}
		// TODO(islishude): parse result to *big.Int
		if err := resp.Error; err != nil {
			// error check for parity
			if code, ok := err["code"].(int); ok && code == -32015 {
				errChan <- ErrNotERC20
				return
			}
			errChan <- fmt.Errorf("[GetERC20TotalSupply response error] msg: %v code: %v", err["message"], err["code"])
			return
		}

		if resp.Result == "0x" {
			errChan <- ErrNotERC20
			return
		}
		resChan <- resp.Result
	}()
	return <-resChan, <-errChan
}
