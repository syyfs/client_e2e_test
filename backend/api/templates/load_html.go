package templates

import (
	"net/http"

	"fmt"

	"strconv"

	"github.com/gin-gonic/gin"
	"brilliance/client_e2e_test/blockchain/backend/api/trace"
	"brilliance/client_e2e_test/blockchain/database/mongo"
)

var Default_route = "404.html"

func LoadHtml(c *gin.Context) {
	feature := c.Param("feature")
	productIdStr := c.Param("productId")
	traceId := c.Param("traceId")
	openId := c.Query("openid")

	productId, _ := strconv.ParseInt(productIdStr, 10, 64)

	var data interface{}

	if feature == "showTraceInfo" {
		route, data, err := trace.ShowTraceInfoHtml(feature, openId, traceId, productId)
		if err != nil {
			fmt.Println("error is : ", err.Error())
			c.HTML(http.StatusOK, Default_route, gin.H{
				"data": data,
			})
			return
		}
		if data == nil {
			fmt.Println("data is nil")
			c.HTML(http.StatusOK, Default_route, gin.H{
				"data": data,
			})
			return
		}

		c.HTML(http.StatusOK, route, gin.H{
			"data": data,
		})
	} else if feature == "showGoodsInfo" {
		route := ShowGoodsInfo(feature, productId)

		c.HTML(http.StatusOK, route, gin.H{
			"data": data,
		})
	} else {
		c.HTML(http.StatusOK, Default_route, gin.H{})
	}

}

func Show404(c *gin.Context) {
	c.HTML(http.StatusOK, Default_route, gin.H{})
}

func ShowGoodsInfo(feature string, productId int64) (route string) {
	productPage, err := mongo.FindProductHtml(feature, productId)
	if err != nil {
		fmt.Println("productPage error : ", err.Error())
		route = Default_route
		return
	}

	route = productPage.Router
	return
}
