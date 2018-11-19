package base

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"brilliance/client_e2e_test/blockchain/backend/api/asset"
	"brilliance/client_e2e_test/blockchain/backend/service"
	"brilliance/client_e2e_test/blockchain/backend/service/fabric"
	"brilliance/client_e2e_test/blockchain/chaincode/model/fish"
)

type BaseGetDby struct {
}

func (t *BaseGetDby) GetInfo(resource string) (data interface{}, productName, productId, router string, err error) {
	productName = "半边天多宝鱼"
	router = "/trace/trace"

	ss := strings.Split(resource, "?") //去除问号

	urlItems := strings.Split(ss[0], "/")

	if urlItems == nil || len(urlItems) <= 1 {
		return nil, "", "", "", errors.New("not found")
	}

	productId = urlItems[len(urlItems)-1]
	data, err = getTraceInfo(productId)
	return
}

func getTraceInfo(id string) (interface{}, error) {

	fabric, err := fabric.NewFabricClient()
	if err != nil {
		return nil, errors.New("not yet")
	}
	result, err := fabric.Query(service.InvokeConfig{
		ChannelId: "yhchannel",
		CcName:    "921108",       //sku
		CcFcn:     "getTraceInfo", //chainCodeMethod
		CcArgs:    [][]byte{[]byte(id)},
	})

	if err != nil {
		fmt.Println("server error : ", err.Error())
		return nil, err
	}

	if len(result.Payload) == 0 {
		return nil, errors.New("not yet")
	}

	var ret asset.Ret
	err = json.Unmarshal(result.Payload, &ret)
	if err != nil {
		fmt.Println("server error : ", err.Error())
		return nil, err
	}

	if ret.Code != http.StatusOK || len(ret.Data) == 0 {
		return nil, errors.New("not yet")
	}

	var traceAndQuality fish.TraceAndQuality
	err = json.Unmarshal(ret.Data, &traceAndQuality)
	if err != nil {

		return nil, err
	}
	return traceAndQuality.TraceInfo, nil
}

type Ret struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []byte `json:"data"`
}
