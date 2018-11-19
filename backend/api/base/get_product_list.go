package base

import (
	"net/http"

	"strings"

	"strconv"

	"github.com/gin-gonic/gin"
	"brilliance/client_e2e_test/blockchain/database/mongo"
)

func GetList(c *gin.Context) {
	typeIdStr, _ := c.GetQuery("typeId")
	if len(typeIdStr) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": http.StatusInternalServerError,
			"message": "typeId is null",
		})
		return
	}

	typeId, _ := strconv.ParseInt(typeIdStr, 10, 64)

	products, err := mongo.FindProductListsByTypeId(typeId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	if len(products) > 0 {
		for i1, _ := range products {
			for i2, _ := range products[i1].PicUrl {
				products[i1].PicUrl[i2] = ChangeIp(products[i1].PicUrl[i2])
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"success": http.StatusOK,
			"message": "successful",
			"data":    products,
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": http.StatusOK,
			"message": "not found",
			"data":    "",
		})
	}

}

func ChangeIp(oldUrl string) (url string) {
	if len(oldUrl) == 0 {
		return oldUrl
	}
	if strings.HasPrefix(oldUrl, "http://10.0.90.53/") {
		url = strings.Replace(oldUrl, "http://10.0.90.53/", "https://sy.yonghui.cn/statics/", -1)
	}
	return
}
