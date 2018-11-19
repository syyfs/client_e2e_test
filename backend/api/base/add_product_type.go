package base

import (
	"net/http"

	"encoding/json"

	"time"

	"fmt"

	"github.com/gin-gonic/gin"
	"brilliance/client_e2e_test/blockchain/common/util"
	"brilliance/client_e2e_test/blockchain/database/mongo"
)

type ReqProductType struct {
	Infos []mongo.ProductType `json:"infos"`
}

func AddProductType(c *gin.Context) {
	requestBody, err := util.ProcessBody(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": http.StatusInternalServerError,
			"message": "server error1 ",
		})
		return
	}

	var reqType ReqProductType

	err = json.Unmarshal(requestBody, &reqType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": http.StatusInternalServerError,
			"message": "server error2 ",
		})
		return
	}

	now := time.Now()
	for index, _ := range reqType.Infos {
		//TODO 通过sequence 获取主键
		reqType.Infos[index].CreateTime = now
		reqType.Infos[index].UpdateTime = now

		err = mongo.AddProductType(reqType.Infos[index])
		fmt.Println("error is : ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": http.StatusInternalServerError,
			"message": "server error 3",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": http.StatusOK,
		"message": "successful",
	})

}
