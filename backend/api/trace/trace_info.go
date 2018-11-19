package trace

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/pkg/errors"

	"brilliance/client_e2e_test/blockchain/backend/api/asset"
	"brilliance/client_e2e_test/blockchain/backend/service"
	"brilliance/client_e2e_test/blockchain/backend/service/fabric"
	"brilliance/client_e2e_test/blockchain/database/mongo"
)

func GetTraceInfo(productId int64, traceId string) (map[string]interface{}, error) {
	//根据productId查询商品信息，获取chaincodeName
	productInfo, err := mongo.FindProductById(productId)
	if err != nil {
		fmt.Println("get product error : ", err.Error())
		return nil, err
	}
	if len(productInfo.CcName) == 0 {
		return nil, errors.New("chaincode not found")
	}

	payload, err := QueryFabric(productInfo.CcName, traceId)

	var ret asset.Ret

	err = json.Unmarshal(payload, &ret)
	if err != nil {
		fmt.Println("server error : ", err.Error())
		return nil, err
	}

	if len(ret.Data) == 0 || ret.Code != http.StatusOK {
		return nil, errors.New("no trace info found")
	}

	var data map[string]interface{}

	err = json.Unmarshal(ret.Data, &data)
	if err != nil {
		fmt.Println("json.unmarshal error : ", err.Error())
		return nil, err
	}

	return data, nil
}

func QueryFabric(ccName, traceId string) ([]byte, error) {
	client, err := fabric.NewFabricClient()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	result, err := client.Query(service.InvokeConfig{
		ChannelId: "yhchannel",
		CcName:    ccName,
		CcFcn:     "getTraceInfo",
		CcArgs:    [][]byte{[]byte(traceId)},
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return result.Payload, nil
}
