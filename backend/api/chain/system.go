package chain

import (
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"
	"brilliance/client_e2e_test/blockchain/backend/service"
)

func GetBlock(c *gin.Context) {
	numberOrHash := c.Param("numberOrHash")
	if len(numberOrHash) < 1 {
		logger.Info("GetChainInfo")

	} else if len(numberOrHash) < 10 {
		logger.Info("GetBlockByNumber", numberOrHash)

	} else {
		logger.Info("GetBlockByHash", numberOrHash)

	}
}

func GetTransaction(c *gin.Context) {

}

func GetTransactionByID(c *gin.Context) {
	txHash := c.Param("txHash")
	logger.Info("GetTransactionByID")
	result, err := service.FabClient.Execute(service.InvokeConfig{
		ChannelId: ChannelId,
		CcName:    CcName,               //qscc
		CcFcn:     "GetTransactionByID", //qscc chainCodeMethod
		CcArgs:    [][]byte{[]byte(ChannelId), []byte(txHash)},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"data": err.Error(),
		})
		return
	}

	if len(result.Payload) > 0 {
		mapData := make(map[string]interface{})
		parseTransaction(result.Payload, mapData)
		tid := mapData["txid"]
		if tid == nil {
			mapData["txid"] = txHash
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": mapData,
		})

		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"data": "unknown error", //TODO 规范合约返回信息，约定所有请求的返回
		})

		return
	}
}

func GetBlockByNumber(c *gin.Context) {
	number := c.Param("number")
	logger.Info("GetBlockByNumber")
	result, err := service.FabClient.Execute(service.InvokeConfig{
		ChannelId: ChannelId,
		CcName:    CcName,             //qscc
		CcFcn:     "GetBlockByNumber", //qscc chainCodeMethod
		CcArgs:    [][]byte{[]byte(ChannelId), []byte(number)},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"data": err.Error(),
		})
		return
	}

	if len(result.Payload) > 0 {
		mapData := make(map[string]interface{})
		parseBlock(result.Payload, mapData)
		//添加区块高度
		mapData["height"] = number
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": mapData,
		})

		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"data": "unknown error", //TODO 规范合约返回信息，约定所有请求的返回
		})

		return
	}
}
func GetBlockByHash(c *gin.Context) {
	logger.Info("GetBlockByHash")
	blockHash := c.Param("blockHash")
	//传入参数是数组转码成16进制后的string 需要解码后使用
	byteHash, er := hex.DecodeString(blockHash)
	if er != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"data": er.Error(),
		})
		return
	}
	result, err := service.FabClient.Execute(service.InvokeConfig{
		ChannelId: ChannelId,
		CcName:    CcName,           //qscc
		CcFcn:     "GetBlockByHash", //qscc chainCodeMethod
		CcArgs:    [][]byte{[]byte(ChannelId), byteHash},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"data": err.Error(),
		})
		return
	}

	if len(result.Payload) > 0 {
		mapData := make(map[string]interface{})
		parseBlock(result.Payload, mapData)
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": mapData,
		})

		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"data": "unknown error", //TODO 规范合约返回信息，约定所有请求的返回
		})

		return
	}
}
func GetChainInfo(c *gin.Context) {
	logger.Info("GetChainInfo")
	result, err := service.FabClient.Execute(service.InvokeConfig{
		ChannelId: ChannelId,
		CcName:    CcName,         //qscc
		CcFcn:     "GetChainInfo", //qscc chainCodeMethod
		CcArgs:    [][]byte{[]byte(ChannelId)},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"data": err.Error(),
		})
		return
	}

	if len(result.Payload) > 0 {
		mapData := make(map[string]interface{})
		parseBlockchainInfo(result.Payload, mapData)
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": mapData,
		})

		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"data": "unknown error", //TODO 规范合约返回信息，约定所有请求的返回
		})

		return
	}
}

func GetBlockByTxID(c *gin.Context) {
	txHash := c.Param("txHash")
	logger.Info("GetBlockByTxID")
	result, err := service.FabClient.Execute(service.InvokeConfig{
		ChannelId: ChannelId,
		CcName:    CcName,           //qscc
		CcFcn:     "GetBlockByTxID", //qscc chainCodeMethod
		CcArgs:    [][]byte{[]byte(ChannelId), []byte(txHash)},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"data": err.Error(),
		})
		return
	}

	if len(result.Payload) > 0 {
		mapData := make(map[string]interface{})
		parseBlock(result.Payload, mapData)
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": mapData,
		})

		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"data": "unknown error", //TODO 规范合约返回信息，约定所有请求的返回
		})

		return
	}
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
	result["dataHash"] = dataHash //这个是区块体的hash数据  并不是区块的唯一表示hash  所以不能使用这个获取区块
	result["previousHash"] = previousHash
	result["txNum"] = len(blockData)
	result["blockTime"] = blockTime
	result["trans"] = string(trans)

}

/**
解析指定交易信息
*/
func parseTransaction(buf []byte, result map[string]interface{}) {
	//解析请求结果
	processedTx := &peer.ProcessedTransaction{}
	proto.Unmarshal(buf, processedTx)
	envelope := processedTx.GetTransactionEnvelope()
	payload := &common.Payload{}
	proto.Unmarshal(envelope.GetPayload(), payload)
	channelHeader := &common.ChannelHeader{}
	proto.Unmarshal(payload.Header.ChannelHeader, channelHeader)
	blockTime := time.Unix(channelHeader.Timestamp.Seconds, 0).Format("2006-01-02 15:04:05")

	result["txid"] = channelHeader.GetTxId()
	result["blockTime"] = blockTime
	result["validationcode"] = processedTx.GetValidationCode()
	result["validcodename"] = peer.TxValidationCode_name[processedTx.GetValidationCode()]
}
