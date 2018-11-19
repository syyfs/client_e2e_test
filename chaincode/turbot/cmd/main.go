package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"brilliance/client_e2e_test/blockchain/chaincode/turbot"
)

func main() {
	err := shim.Start(new(turbot.TracingChaincode))
	if err != nil {
		fmt.Printf("Error starting turbot chaincode: %s", err)
	}
}
