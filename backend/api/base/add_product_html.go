package base

import (
	"encoding/json"
	"net/http"
	"time"

	"fmt"

	"github.com/gin-gonic/gin"
	"brilliance/client_e2e_test/blockchain/common/util"
	"brilliance/client_e2e_test/blockchain/database/mongo"
)

type ReqProductHtml struct {
	Infos []mongo.ProductHtml
}

func AddProductHtml(c *gin.Context) {
	requestBody, err := util.ProcessBody(c)
	if err != nil {
		fmt.Println("error 1 is : ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": http.StatusInternalServerError,
			"message": "server error 1",
		})
		return
	}

	var reqProductHtml ReqProductHtml

	err = json.Unmarshal(requestBody, &reqProductHtml)
	if err != nil {
		fmt.Println("error 2 is : ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": http.StatusInternalServerError,
			"message": "server error 2",
		})
		return
	}

	now := time.Now()
	for index, _ := range reqProductHtml.Infos {
		reqProductHtml.Infos[index].CreateTime = now
		reqProductHtml.Infos[index].UpdateTime = now

		err = mongo.AddProductHtml(reqProductHtml.Infos[index])
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
