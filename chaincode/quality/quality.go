package quality

import (
	"fmt"

	"encoding/json"

	"net/http"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"brilliance/client_e2e_test/blockchain/chaincode/model"
	"brilliance/client_e2e_test/blockchain/chaincode/util"
)

type QualityChainCode struct {
}

// Init ...
func (t *QualityChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke ...
func (t *QualityChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Printf("Invoke func: %s args: %v\n", function, args)

	if function == "putQuality" { // put fresh type quality message
		return t.putQuality(stub, args)
	} else if function == "getQualityByUUID" {
		return t.getQualityByUUID(stub, args)
	} else if function == "getQualityByBatchId" {
		return t.getQualityByBatchId(stub, args)
	} else if function == "getQualityByOrderId" {
		return t.getQualityByOrderId(stub, args)
	} else if function == "getQualityByGoodsId" {
		return t.getQualityByGoodsId(stub, args)
	} else if function == "getQualityByBOG" {
		return t.getQualityByBOG(stub, args)
	}

	return shim.Error(fmt.Sprintf("Unsupported function: %s", function))
}

func (t *QualityChainCode) putQuality(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var body = args[0]
	var quality model.FreshQuality

	err := json.Unmarshal([]byte(body), &quality)
	if err != nil {
		msg := "json.Unmarshal err, please check the args..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)

	}

	flag, responseBody := validateData(&quality)
	if !flag {
		return shim.Success(responseBody)
	}

	err = stub.PutState(stub.GetTxID(), []byte(body))
	if err != nil {
		msg := "put quality by key txid error..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)
	}

	latestKey := util.FormatStateKey("latest", quality.GoodsId)
	err = stub.PutState(latestKey, []byte(stub.GetTxID()))
	if err != nil {
		msg := "put quality by key latest error..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)
	}

	responseByte := util.GetResponse(http.StatusOK, util.SUCCESS, []byte{})
	return shim.Success(responseByte)
}

func (t *QualityChainCode) getQualityByUUID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	sampleId := args[0]
	queryString := fmt.Sprintf("{\"selector\": {\"uuid\": \"%s\"}}", sampleId)
	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		msg := "getQualityByUUID query error, please check the sql..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)
	}
	qualities, err := util.GetListResult(resultsIterator)
	if err != nil {
		msg := "getQualityByUUID get quality error..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)
	}

	qualities = addQueryType(qualities, util.QueryTypeUUID)
	return shim.Success(qualities)
}

func (t *QualityChainCode) getQualityByBatchId(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	batchId := args[0]
	queryString := fmt.Sprintf("{\"selector\": {\"batch_id\": \"%s\"}}", batchId)
	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		msg := "getQualityByBatchId query error, please check the sql..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)
	}
	qualities, err := util.GetListResult(resultsIterator)
	if err != nil {
		msg := "getQualityByBatchId get quality error..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)
	}

	qualities = addQueryType(qualities, util.QueryTypeBatchId)
	return shim.Success(qualities)
}

func (t *QualityChainCode) getQualityByOrderId(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	orderId := args[0]
	queryString := fmt.Sprintf("{\"selector\": {\"order_id\": \"%s\"}}", orderId)
	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		msg := "getQualityByOrderId query error, please check the sql..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)
	}
	qualities, err := util.GetListResult(resultsIterator)
	if err != nil {
		msg := "getQualityByOrderId get quality error..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)
	}

	qualities = addQueryType(qualities, util.QueryTypeOrderId)
	return shim.Success(qualities)
}

func (t *QualityChainCode) getQualityByGoodsId(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	goodsId := args[0]
	queryString := fmt.Sprintf("{\"selector\": {\"goods_id\": \"%s\"}}", goodsId)
	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		msg := "getQualityByGoodsId query error, please check the sql..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)
	}
	qualities, err := util.GetListResult(resultsIterator)
	if err != nil {
		msg := "getQualityByGoodsId get quality error..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		shim.Success(responseByte)
	}

	qualities = addQueryType(qualities, util.QueryTypeGoodsId)
	fmt.Println("qualities is = ", string(qualities))
	return shim.Success(qualities)
}

func (t *QualityChainCode) getQualityByBOG(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//参数：model.BOGInfo
	bogByte := args[0]

	//解析出 model.BOGInfo
	var bog model.BOGInfo
	err := json.Unmarshal([]byte(bogByte), &bog)
	if err != nil {
		msg := "json Unmarshal BOGInfo by bogType error..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return shim.Success(responseByte)
	}

	//构建 map[string]qualityUpload ，以itemId 为主键
	mapQuality := make(map[string]model.FreshQuality)

	//根据 uuid 查出质检list,便利list，如果map中不存在itemId为主见的quality，则添加到map中
	if len(bog.Uuid) > 0 {
		qualitys := t.getQualityByUUID(stub, []string{bog.Uuid})
		if len(qualitys.Payload) > 0 {

			var qualities []model.FreshQuality
			err = json.Unmarshal(qualitys.Payload, &qualities)
			if err != nil {
				msg := "getQualityByUUID json Unmarshal qualities by qualities.Payload error..."
				responseBody := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
				shim.Success(responseBody)
			}
			for index := range qualities {
				_, ok := mapQuality[qualities[index].TestItemId]
				if !ok {
					mapQuality[qualities[index].TestItemId] = qualities[index]
				}
			}
		}
	}

	//根据 batchId 查出质检list,便利list，如果map中不存在itemId为主见的quality，则添加到map中
	if len(bog.BatchId) > 0 {
		qualitys := t.getQualityByBatchId(stub, []string{bog.BatchId})
		if len(qualitys.Payload) > 0 {

			var qualities []model.FreshQuality
			err = json.Unmarshal(qualitys.Payload, &qualities)
			if err != nil {
				msg := "getQualityByBatchId json Unmarshal qualities by qualities.Payload error..."
				responseBody := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
				shim.Success(responseBody)
			}
			for index := range qualities {
				_, ok := mapQuality[qualities[index].TestItemId]
				if !ok {
					mapQuality[qualities[index].TestItemId] = qualities[index]
				}
			}
		}

	}
	//根据 OrderId 查出质检list,便利list，如果map中不存在itemId为主见的quality，则添加到map中
	if len(bog.OrderId) > 0 {
		qualitys := t.getQualityByOrderId(stub, []string{bog.OrderId})
		if len(qualitys.Payload) > 0 {

			var qualities []model.FreshQuality
			err = json.Unmarshal(qualitys.Payload, &qualities)
			if err != nil {
				msg := "getQualityByOrderId json Unmarshal qualities by qualities.Payload error..."
				responseBody := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
				shim.Success(responseBody)
			}
			for index := range qualities {
				_, ok := mapQuality[qualities[index].TestItemId]
				if !ok {
					mapQuality[qualities[index].TestItemId] = qualities[index]
				}
			}
		}

	}
	//根据 goodsId 查出质检list,便利list，如果map中不存在itemId为主见的quality，则添加到map中
	if len(bog.GoodsId) > 0 {
		qualitys := t.getQualityByGoodsId(stub, []string{bog.GoodsId})
		if len(qualitys.Payload) > 0 {

			var qualities []model.FreshQuality
			err = json.Unmarshal(qualitys.Payload, &qualities)
			if err != nil {
				msg := "getQualityByGoodsId json Unmarshal qualities by qualities.Payload error..."
				responseBody := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
				shim.Success(responseBody)
			}
			for index := range qualities {
				_, ok := mapQuality[qualities[index].TestItemId]
				if !ok {
					mapQuality[qualities[index].TestItemId] = qualities[index]
				}
			}
		}

	}

	if len(mapQuality) > 0 {
		var listQuality []model.FreshQuality
		for key := range mapQuality {
			listQuality = append(listQuality, mapQuality[key])
		}

		//返回json的quality list
		listQualityByte, err := json.Marshal(listQuality)
		if err != nil {
			msg := "json Marshal listQuality error..."
			responseBody := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
			shim.Success(responseBody)
		}

		responseByte := util.GetResponse(http.StatusOK, util.SUCCESS, listQualityByte)
		return shim.Success(responseByte)

	}
	responseByte := util.GetResponse(http.StatusOK, util.SUCCESS, []byte{})
	return shim.Success(responseByte)
}

func validateData(quality *model.FreshQuality) (bool, []byte) {
	if len(quality.BOGInfo.GoodsId) == 0 {
		msg := "field goods_id not be null, please check the data..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return false, responseByte
	}
	if len(quality.GoosdName) == 0 {
		msg := "field goods_name not be null, please check the data..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return false, responseByte
	}
	if len(quality.TestItemId) == 0 {
		msg := "field test_item_id not be null, please check the data..."
		responseByte := util.GetResponse(http.StatusInternalServerError, msg, []byte{})
		return false, responseByte
	}
	return true, nil
}

func addQueryType(qualitiesByte []byte, queryType int) []byte {
	var qualities []model.FreshQuality
	json.Unmarshal(qualitiesByte, &qualities)

	for index := range qualities {
		qualities[index].QueryType = queryType
	}

	qualitiesByte, _ = json.Marshal(&qualities)
	fmt.Println("qualitiesByte is = " + string(qualitiesByte))
	return qualitiesByte
}
