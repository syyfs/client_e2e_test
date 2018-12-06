package fabric

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"
	fabcfg "github.com/hyperledger/fabric-sdk-go/pkg/core/config"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"
	"brilliance/client_e2e_test/blockchain/backend/service"
	"brilliance/client_e2e_test/blockchain/common/config"
	"strconv"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/core/common/ccprovider"
	"brilliance/client_e2e_test/blockchain/backend/service/fabric/chaincode"
)

var chaincodeName = "mycc1"

type BlockData struct {
	TxId   string `json:"txId"`
	TxType int32  `json:"txType"`
	TxTime string `json:"txTime"`
}

var configPath = "../../../config/config.yaml"


func TestClientE2E(t *testing.T) {
	os.Setenv("FABRIC_ARTIFACTS", "../../../")
	err := config.InitConfig(configPath)
	if err != nil {
		t.Error(err)
	}

	configProvider := fabcfg.FromFile(config.GetConfigFile())
	sdk, err := fabsdk.New(configProvider)

	if err != nil {
		t.Error(err)
	}

	client := &Client{fabSdk: sdk, user: "Admin", mspId: "Org1"}

	key_A := "a"

	key_B := "b"


	result, err := client.Execute(service.InvokeConfig{
		ChannelId: "mychannel",
		CcName:    "mycc", //mycc
		CcFcn:     "invoke",
		CcArgs:    [][]byte{
						[]byte(key_A),[]byte(key_B), []byte(strconv.Itoa(10))},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(result)

}
func TestClientQueryE2E(t *testing.T){
	os.Setenv("FABRIC_ARTIFACTS", "../../../")
	err := config.InitConfig(configPath)
	if err != nil {
		t.Error(err)
	}
	client, err := NewFabricClient()
	if err != nil {
		t.Error(err)
	}

	key_A := "a"


	result, err := client.Query(service.InvokeConfig{
		ChannelId: "mychannel",
		CcName:    "mycc", //mycc
		CcFcn:     "query",
		CcArgs:    [][]byte{
			[]byte(key_A)},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
func TestClient_DiscoveryService(t *testing.T) {
	os.Setenv("FABRIC_ARTIFACTS", "../../../")
	err := config.InitConfig(configPath)
	if err != nil {
		t.Error(err)
	}
	client, err := NewFabricClient()
	if err != nil {
		t.Error(err)
	}


	t.Log(client)

}
func TestClient_Query_GetBlockByNumber(t *testing.T) {
	os.Setenv("FABRIC_ARTIFACTS", "../../../")
	err := config.InitConfig(configPath)
	if err != nil {
		t.Error(err)
	}
	client, err := NewFabricClient()
	if err != nil {
		t.Fatal(err)
	}
	result, err := client.Query(service.InvokeConfig{
		ChannelId: "mychannel",
		CcName:    "qscc", //mycc
		CcFcn:     "GetBlockByNumber",
		CcArgs:    [][]byte{[]byte("yhchannel"), []byte("100")},
	})
	if err != nil {
		t.Fatal(err)
	}
	//解析请求结果
	dataMap := make(map[string]interface{})
	parseBlock(result.Payload, dataMap)
	fmt.Println(dataMap)
	//t.Log(result)
}
func TestClient_Query_GetChainInfo(t *testing.T) {
	os.Setenv("FABRIC_ARTIFACTS", "../../../")
	err := config.InitConfig(configPath)
	if err != nil {
		t.Error(err)
	}
	client, err := NewFabricClient()
	if err != nil {
		t.Error(err)
	}
	result, err := client.Query(service.InvokeConfig{
		ChannelId: "yhchannel",
		CcName:    "qscc", //mycc
		CcFcn:     "GetChainInfo",
		CcArgs:    [][]byte{[]byte("yhchannel")},
	})
	if err != nil {
		t.Error(err)
	}
	//解析请求结果
	dataMap := make(map[string]interface{})
	parseBlockchainInfo(result.Payload, dataMap)
	fmt.Println(dataMap)
	//t.Log(result)
}
func TestClient_Query_GetTransactionByID(t *testing.T) {
	os.Setenv("FABRIC_ARTIFACTS", "../../../")
	err := config.InitConfig(configPath)
	if err != nil {
		t.Error(err)
	}
	client, err := NewFabricClient()
	if err != nil {
		t.Error(err)
	}
	result, err := client.Query(service.InvokeConfig{
		ChannelId: "mychannel",
		CcName:    "qscc", //mycc
		CcFcn:     "GetTransactionByID",
		CcArgs:    [][]byte{[]byte("mychannel"), []byte("587a8c3d1b44236c65cdcb91b38d6e90c7111197a286829adb5610f36a9cf31a")},
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("result:[%#v]\n",result)
	//fmt.Println(result)
	//解析请求结果
	transactionEnvelope := parseTrans_1(result.Payload)

	payload := &common.Payload{}
	proto.Unmarshal(transactionEnvelope.GetPayload(), payload)
	//
	signatureHeader := &common.SignatureHeader{}
	proto.Unmarshal(payload.Header.SignatureHeader, signatureHeader)

	signatureHeader.String()
	signatureHeader.GetCreator()

	t.Logf("transactionEnvelope = [%s]\n", transactionEnvelope.String())
	t.Logf("signatureHeader.String() = [%s]\n", signatureHeader.String())


}

func TestClient_Execute(t *testing.T) {
	os.Setenv("FABRIC_ARTIFACTS", "../../../")
	err := config.InitConfig(configPath)
	if err != nil {
		t.Error(err)
	}
	client, err := NewFabricClient()
	if err != nil {
		t.Error(err)
	}

	key := "example_key"
	value := make([]byte, 10)
	rand.Read(value)

	result, err := client.Execute(service.InvokeConfig{
		ChannelId: "yhchannel",
		CcName:    "921108", //mycc
		CcFcn:     "putTraceInfo",
		CcArgs:    [][]byte{[]byte(key), value},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(result)

	result, err = client.Query(service.InvokeConfig{
		ChannelId: "yhchannel",
		CcName:    "921108", //mycc
		CcFcn:     "putTraceInfo",
		CcArgs:    [][]byte{[]byte(key)},
	})
	if err != nil {
		t.Error(err)
	}

	if bytes.Compare(result.Payload, value) != 0 {
		t.Error("The result is not correct!!!")
	}
	t.Log("Result:", result.Payload)
	client.Close()
}
func TestClient_GetTraceInfo(t *testing.T) {
	os.Setenv("FABRIC_ARTIFACTS", "../../../")
	err := config.InitConfig(configPath)
	if err != nil {
		t.Fatalf("init config error...")
	}

	client, err := NewFabricClient()
	if err != nil {
		t.Fatalf("NewFabClient error...")
	}

	result, err := client.Query(service.InvokeConfig{
		ChannelId: "yhchannel",
		CcName:    "turbot0803", //mycc
		CcFcn:     "getTraceInfo",
		CcArgs:    [][]byte{[]byte(`123`)},
	})
	t.Log(string(result.Payload))
	if err != nil {
		t.Fatal(err.Error())
	}
	client.Close()
}

/**
解析指定区块信息
*/
func parseBlock(buf []byte, result map[string]interface{}) {
	//解析请求结果
	block := &common.Block{}
	proto.Unmarshal(buf, block)
	dataHash := hex.EncodeToString(block.GetHeader().GetDataHash())
	previousHash := hex.EncodeToString(block.GetHeader().GetPreviousHash())
	var blockTime string
	var blockData []BlockData
	for _, b := range block.GetData().Data {
		envelope := &common.Envelope{}
		payload := &common.Payload{}
		channelHeader := &common.ChannelHeader{}
		proto.Unmarshal(b, envelope)
		proto.Unmarshal(envelope.GetPayload(), payload)
		proto.Unmarshal(payload.Header.ChannelHeader, channelHeader)
		blockTime = time.Unix(channelHeader.Timestamp.Seconds, 0).Format("2006-01-02 15:04:05")
		blockData = append(blockData, BlockData{channelHeader.TxId, channelHeader.Type, blockTime})
	}
	trans, _ := json.Marshal(blockData)
	result["hash"] = dataHash
	result["previousHash"] = previousHash
	result["txNum"] = len(blockData)
	result["blockTime"] = blockTime
	result["trans"] = string(trans)

}

/**
解析指定通道的区块链信息
*/
func parseBlockchainInfo(buf []byte, result map[string]interface{}) {
	//解析请求结果
	blockChainInfo := &common.BlockchainInfo{}
	proto.Unmarshal(buf, blockChainInfo)

	result["height"] = blockChainInfo.GetHeight()
	result["currentBlockHash"] = hex.EncodeToString(blockChainInfo.GetCurrentBlockHash())
	result["previousBlockHash"] = hex.EncodeToString(blockChainInfo.GetPreviousBlockHash())
}

/**
解析指定交易信息
*/
func parseTrans(buf []byte, result map[string]interface{}) {
	//解析请求结果
	processedTx := &peer.ProcessedTransaction{}
	fmt.Println(processedTx)
	proto.Unmarshal(buf, processedTx)

	envelope := processedTx.GetTransactionEnvelope()
	payload := &common.Payload{}
	proto.Unmarshal(envelope.GetPayload(), payload)

	channelHeader := &common.ChannelHeader{}
	proto.Unmarshal(payload.Header.ChannelHeader, channelHeader)
	blockTime := time.Unix(channelHeader.Timestamp.Seconds, 0).Format("2006-01-02 15:04:05")



	result["txid"] = channelHeader.GetTxId()
	result["blockTime"] = blockTime
	result["ValidationCode"] = processedTx.GetValidationCode()
	result["ValidationCodename"] = peer.TxValidationCode_name[processedTx.GetValidationCode()]

	signature := processedTx.GetTransactionEnvelope().GetSignature()

	result["signature"] = signature

}

func parseTrans_1(buf []byte) (*common.Envelope ){
	//解析请求结果
	processedTx := &peer.ProcessedTransaction{}
	fmt.Println(processedTx)
	proto.Unmarshal(buf, processedTx)

	envelope := processedTx.GetTransactionEnvelope()
	payload := &common.Payload{}
	proto.Unmarshal(envelope.GetPayload(), payload)
	//
	signatureHeader := &common.SignatureHeader{}
	proto.Unmarshal(payload.Header.SignatureHeader, signatureHeader)
	//blockTime := time.Unix(channelHeader.Timestamp.Seconds, 0).Format("2006-01-02 15:04:05")
	signatureHeader.String()
	signatureHeader.GetCreator()


	return envelope

}



func TestClient_GetCCData(t *testing.T) {
	os.Setenv("FABRIC_ARTIFACTS", "../../../")
	err := config.InitConfig(configPath)
	if err != nil {
		t.Fatalf("init config error...")
	}

	configProvider := fabcfg.FromFile(config.GetConfigFile())
	sdk, err := fabsdk.New(configProvider)

	if err != nil {
		t.Error(err)
	}

	client := &Client{fabSdk: sdk, user: "Admin", mspId: "Org2"}

	result, err := client.Query(service.InvokeConfig{
		ChannelId: "mychannel",
		CcName:    "lscc", //mycc
		CcFcn:     "getccdata",
		CcArgs:    [][]byte{[]byte("mychannel"),[]byte("mycc2")},
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	ccdata := &ccprovider.ChaincodeData{}
	proto.Unmarshal(result.Payload, ccdata)

	spe , err := chaincode.ParseSignPolicyEnvelop(ccdata.Policy)
	if err != nil {
		t.Fatal(err.Error())

	}
	noutof := spe.Rule.GetNOutOf()

	t.Logf("ccdata.noutof = [%s]\n", noutof.String())

	mspiden := spe.Identities
	for i , msp := range mspiden{
		t.Logf("The [%d] Identities [%#v] \n",i, string(msp.Principal))
	}
	t.Logf("result.TxId :[%s] \n" ,result.TxId)
	client.Close()
}


