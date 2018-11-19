package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"brilliance/client_e2e_test/blockchain/chaincode/quality"
)

func main() {
	err := shim.Start(new(quality.QualityChainCode))
	if err != nil {
		fmt.Printf("Error starting quality chaincode: %s", err)
	}
}
