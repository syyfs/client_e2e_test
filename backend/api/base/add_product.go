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

type ReqProduct struct {
	Infos []mongo.Product
}

//根据用户的openid关联查询历史
func AddProduct(c *gin.Context) {
	requestBody, err := util.ProcessBody(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": http.StatusInternalServerError,
			"message": "server error 1",
		})
		return
	}

	var reqProduct ReqProduct

	err = json.Unmarshal(requestBody, &reqProduct)
	if err != nil {
		fmt.Println("error is : ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": http.StatusInternalServerError,
			"message": "server error 2",
		})
		return
	}

	now := time.Now()
	for index, _ := range reqProduct.Infos {
		reqProduct.Infos[index].CreateTime = now
		reqProduct.Infos[index].UpdateTime = now

		err = mongo.AddProduct(reqProduct.Infos[index])
		if err != nil {
			fmt.Println("error is : ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": http.StatusInternalServerError,
				"message": "server error 3",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": http.StatusOK,
		"message": "successful",
	})

}
