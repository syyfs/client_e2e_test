package blackPig

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"net/http"
	"brilliance/client_e2e_test/blockchain/chaincode/model/blackPig"
	"brilliance/client_e2e_test/blockchain/chaincode/util"
)

type BpTracingChaincode struct {
}

// Init ...
func (t *BpTracingChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke ...
func (t *BpTracingChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Printf("Invoke func: %s args: %v\n", function, args)

	if function == "putTraceInfo" {
		return t.putTraceInfo(stub, args)
	} else if function == "getTraceInfo" {
		return t.getTraceInfo(stub, args)
	}
	return shim.Error(fmt.Sprintf("Unsupported function: %s", function))
}

//put blackPig traceInfo
func (t *BpTracingChaincode) putTraceInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var body = args[0]
	var traceInfo blackPig.TraceInfo
	err := json.Unmarshal([]byte(body), &traceInfo)
	if err != nil {
		msg := "json Unmarshal error, please check the args..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)
	}

	flag, responseByte := validateData(&traceInfo)
	if !flag {
		return shim.Success(responseByte)
	}

	err = stub.PutState(traceInfo.BatchCode, []byte(body))
	if err != nil {
		msg := "put traceInfo by BatchCode error..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)
	}

	responseByte = util.GetResponse(http.StatusOK, util.SUCCESS, []byte{})
	return shim.Success(responseByte)
}

//get blackPig traceInfo
func (t *BpTracingChaincode) getTraceInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	batchCode := args[0]
	queryString := fmt.Sprintf("{\"selector\": {\"batch_code\": {\"$eq\": \"%s\"}}}", batchCode)
	resultIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		msg := "getTraceInfo query error, please check the sql..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)
	}
	traceInfoByte, err := util.GetResult(resultIterator)
	if err != nil {
		msg := "getTraceInfo get traceInfo error..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)
	}

	responseByte := util.GetResponse(http.StatusOK, util.SUCCESS, traceInfoByte)
	return shim.Success(responseByte)
}

func validateData(traceInfo *blackPig.TraceInfo) (bool, []byte) {
	if len(traceInfo.BatchCode) == 0 {
		msg := "field batch_code not be null, please check the data..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return false, responseByte
	}
	return true, nil
}
