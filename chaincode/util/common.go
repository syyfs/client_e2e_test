package util

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"bytes"
)

func GetResult(result shim.StateQueryIteratorInterface) ([]byte, error) {
	defer result.Close()
	for result.HasNext() {
		queryResponse, err := result.Next()
		if err != nil {
			return nil, err
		}
		return queryResponse.Value, nil
	}
	return nil, nil
}

func GetListResult(resultsIterator shim.StateQueryIteratorInterface) ([]byte, error) {
	defer resultsIterator.Close()
	var buffer bytes.Buffer

	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	return buffer.Bytes(), nil
}
