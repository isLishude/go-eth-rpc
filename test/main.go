package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

func rmHexPrefix(input string) string {
	if strings.HasPrefix(input, "0x") {
		input = input[2:]
	}
	return input
}

func main() {
	raw := "0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000003424e420000000000000000000000000000000000000000000000000000000000"
	data, err := hex.DecodeString(rmHexPrefix(raw))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(data))
	fmt.Printf("%q\n", string(data))
}
