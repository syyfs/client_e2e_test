package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"brilliance/client_e2e_test/blockchain/chaincode/blackPig"
)

func main() {
	err := shim.Start(new(blackPig.BpTracingChaincode))
	if err != nil {
		fmt.Printf("Error starting blackpig chaincode: %s", err)
	}
}
