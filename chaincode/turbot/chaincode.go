package turbot

import (
	"fmt"

	"encoding/json"

	"net/http"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"brilliance/client_e2e_test/blockchain/chaincode/model"
	"brilliance/client_e2e_test/blockchain/chaincode/model/fish"
	"brilliance/client_e2e_test/blockchain/chaincode/util"
)

type TracingChaincode struct {
}

// Init ...
func (t *TracingChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke ...
func (t *TracingChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Printf("Invoke func: %s args: %v\n", function, args)

	if function == "putTraceInfo" {
		return t.putTraceInfo(stub, args)
	} else if function == "getTraceInfo" {
		return t.getTraceInfo(stub, args)
	} else if function == "putValue" {
		if len(args) < 2 {
			return shim.Error("The length of args less than 2.")
		}

		stub.PutState(args[0], []byte(args[1]))
		return shim.Success([]byte(args[0]))
	} else if function == "getValue" {
		if len(args) < 1 {
			return shim.Error("The length of args less than 1.")
		}

		data, err := stub.GetState(args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(data)
	}

	return shim.Error(fmt.Sprintf("Unsupported function: %s", function))
}

func (t *TracingChaincode) putTraceInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var body = args[0]
	var traceInfoUpload fish.TraceInfoUpload

	err := json.Unmarshal([]byte(body), &traceInfoUpload) //用json.unmarshal校验数据是否是json格式
	if err != nil {
		msg := "json Unmarshal error, please check the args..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)
	}

	flag, responseByte := validateData(&traceInfoUpload)
	if !flag {
		return shim.Success(responseByte)
	}

	err = stub.PutState(stub.GetTxID(), []byte(body))
	if err != nil {
		msg := "put traceinfo by txid error..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)
	}

	responseByte = util.GetResponse(http.StatusOK, util.SUCCESS, []byte{})
	return shim.Success(responseByte)
}

func (t *TracingChaincode) getTraceInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fishId := args[0]

	queryString := fmt.Sprintf("{\"selector\": {\"fish_ids\": {\"$in\": [\"%s\"]}}, \"limit\": 1}", fishId)
	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		msg := "getTraceInfo query error, please check the sql..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)
	}
	traceInfoBytes, err := util.GetResult(resultsIterator)
	if err != nil {
		msg := "getTraceInfo get traceinfo error..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)
	}

	if len(traceInfoBytes) == 0 {
		responseByte := util.GetResponse(http.StatusInternalServerError, "no traceInfo found", []byte{})
		return shim.Success(responseByte)
	}

	var traceInfo fish.TraceInfo
	err = json.Unmarshal(traceInfoBytes, &traceInfo)
	if err != nil {
		msg := "json Unmarshal by traceInfoBytes error..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)
	}

	traceInfo.Uuid = fishId

	CheckTraceInfoParam(&traceInfo)

	var bog model.BOGInfo
	bog.Uuid = traceInfo.Uuid
	bog.BatchId = traceInfo.BatchId
	bog.OrderId = traceInfo.OrderId
	bog.GoodsId = traceInfo.GoodsId
	bogByte, _ := json.Marshal(bog)

	var traceAndQuality fish.TraceAndQuality
	traceAndQuality.Uuid = traceInfo.Uuid
	traceAndQuality.TraceInfo = traceInfo

	qualityResp := stub.InvokeChaincode(util.FreshQualityChaincodeName, [][]byte{[]byte(util.FreshQualityFuncGetQualityByBOG), bogByte}, util.ChannelName)

	var response util.Response
	if len(qualityResp.Payload) > 0 {
		err = json.Unmarshal(qualityResp.GetPayload(), &response)
		if err == nil {
			var freshQuality []model.FreshQuality
			json.Unmarshal(response.Data, &freshQuality)

			traceAndQuality.Quality = freshQuality
		}
	}

	traceAndQualityBytes, err := json.Marshal(traceAndQuality)
	if err != nil {
		msg := "json Marshal traceAndQuality error..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)
	}

	responseByte := util.GetResponse(http.StatusOK, util.SUCCESS, traceAndQualityBytes)
	return shim.Success(responseByte)
}

func CheckTraceInfoParam(traceInfo *fish.TraceInfo) {
	for _, data := range traceInfo.Certs {
		CheckImageInfo(&data)
	}

	for _, data := range traceInfo.Fries {
		CheckImageInfo(&data)
	}
	for _, data := range traceInfo.Catches {
		CheckImageInfo(&data)
	}
}

func CheckImageInfo(info *model.ImageInfoUpload) {
	if len(info.Name) == 0 {
		info.Name = "图片"
	}
}

func validateData(traceInfoUpload *fish.TraceInfoUpload) (bool, []byte) {
	if len(traceInfoUpload.FishIds) == 0 {
		msg := "field fish_ids not be null, please check the data..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return false, responseByte
	}
	if len(traceInfoUpload.GoodsId) == 0 {
		msg := "field goods_id not be null, please check the data..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return false, responseByte
	}
	if len(traceInfoUpload.GoodsName) == 0 {
		msg := "field goods_name not be null, please check the data..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return false, responseByte
	}
	return true, nil
}
