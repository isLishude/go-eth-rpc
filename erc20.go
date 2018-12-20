package ethrpc

import (
	"encoding/hex"
	"strings"
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
)

// Util
var (
	zeroPadding = strings.Repeat("0", 24)
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

// GetERC20Balance is 获取 ERC20 代币余额（未处理单位）
func (c *Ethereum) GetERC20Balance(contract, address string) (string, error) {
	resChan, errChan := make(chan string, 1), make(chan error, 1)
	go func() {
		defer close(resChan)
		defer close(errChan)

		reqData := CallFuncParam{
			Data: ERC20BalanceOfSig + paddingAddress(address),
			To:   contract,
		}

		result, err := c.CallFunc(&reqData)
		if err != nil {
			// TODO add canIgnoreError
		}

		resChan <- result
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
			Data: ERC20BalanceOfSig,
			To:   contract,
		}

		result, err := c.CallFunc(&reqData)
		if err != nil {
			// TODO add canIgnoreError
		}

		// TODO decode string in storage
		// https://solidity-cn.readthedocs.io/zh/develop/miscellaneous.html#storage
		// data, err :=
		hex.DecodeString(rmHexPrefix(result))

		resChan <- result
	}()

	return <-resChan, <-errChan
}
