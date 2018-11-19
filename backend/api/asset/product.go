package asset

import (
	"net/http"

	"fmt"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"brilliance/client_e2e_test/blockchain/backend/service"
	"brilliance/client_e2e_test/blockchain/common/util"
)

/**
/asset/:product/:method
*/

type Ret struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []byte `json:"data"`
}

func PostData(c *gin.Context) {
	product := c.Param("product")
	method := c.Param("method")
	if product == "" || method == "" {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code":    http.StatusMethodNotAllowed,
			"message": "request param is wrong",
		})

		return
	}

	requestBody, err := util.ProcessBody(c)
	if err != nil {
		fmt.Println("server error : ", err.Error())
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code":    http.StatusMethodNotAllowed,
			"message": "request param is wrong",
		})

		return
	}

	logger.Debugf("exec %s chainCode's method : %s \n the data is : %s", product, method, string(requestBody))

	if len(requestBody) == 0 {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code":    http.StatusMethodNotAllowed,
			"message": "request param is wrong",
		})

		return
	}

	result, err := service.FabClient.Execute(service.InvokeConfig{
		ChannelId: "yhchannel",
		CcName:    product, //sku，合约的名字，对应接口中的product参数
		CcFcn:     method,  //chainCodeMethod,对应接口中method参数 ，执行合约只需要一个参数，具体的业务逻辑判断均由合约执行
		CcArgs:    [][]byte{[]byte(requestBody)},
	})

	if err != nil {
		fmt.Println("server error : ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})

		return
	}

	fmt.Println("result is : ", string(result.Payload))
	if len(result.Payload) > 0 {
		var ret Ret
		err = json.Unmarshal(result.Payload, &ret)
		if err != nil {
			fmt.Println("server error: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "server error",
			})
			return
		}

		fmt.Println("result data is : ", string(ret.Data))

		if ret.Code == http.StatusOK && len(ret.Data) > 0 {
			c.JSON(http.StatusOK, Resp{
				Code:    http.StatusOK,
				Message: "succeed",
				Data:    ret.Data,
			})
		} else {
			c.JSON(ret.Code, gin.H{
				"code":    ret.Code,
				"message": ret.Message,
			})
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "unknown error",
		})
	}
	return
}

func GetData(c *gin.Context) {
	fmt.Println("request server is : ", c.ClientIP())
	product := c.Param("product")
	id := c.Param("id")
	logger.Debugf("query %s, id: %s", product, id)

	result, err := service.FabClient.Query(service.InvokeConfig{
		ChannelId: "yhchannel",
		CcName:    product,        //sku
		CcFcn:     "getTraceInfo", //chainCodeMethod
		CcArgs:    [][]byte{[]byte(id)},
	})

	if err != nil {
		fmt.Println("server error : ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	if len(result.Payload) > 0 {
		var ret Ret
		err = json.Unmarshal(result.Payload, &ret)
		if err != nil {
			fmt.Println("server error : ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "server error",
			})
			return
		}

		if ret.Code == http.StatusOK && len(ret.Data) > 0 {
			c.JSON(http.StatusOK, Resp{
				Code:    http.StatusOK,
				Message: "succeed",
				Data:    ret.Data,
			})
		} else {
			c.JSON(ret.Code, gin.H{
				"code":    ret.Code,
				"message": ret.Message,
			})
		}

		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "unknown error",
		})

		return
	}

}

type Resp struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}
