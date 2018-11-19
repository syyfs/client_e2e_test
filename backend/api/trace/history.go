package trace

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"brilliance/client_e2e_test/blockchain/database/mongo"
)

//根据用户的openid关联查询历史
func GetSearchList(c *gin.Context) {
	openId, _ := c.GetQuery("openid")

	if len(openId) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": http.StatusOK,
			"message": "openid is null",
		})
		return
	}

	//查询数据记录
	info, err := mongo.FindUserInfoByOpenId(openId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": http.StatusInternalServerError,
			"data":    "",
			"message": "not found user by openid = " + openId,
		})
		return
	}

	var his RetHistory
	for index, _ := range info.SearchList {
		record, err := mongo.FindTraceRecord(info.SearchList[index].ProductId, info.SearchList[index].TraceId)
		if err == nil {
			searchTime := record.SearchTime.Format("2006-01-02 15:04:05")

			his.SearchList = append(his.SearchList, SearchList{
				ProductName: record.ProductName,
				Count:       record.Count,
				SearchTime:  searchTime,
			})
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"success": http.StatusOK,
		"message": "successful",
		"data":    his,
	})

}

type RetHistory struct {
	SearchList []SearchList `json:"search_list"` //查询记录

}

type SearchList struct {
	ProductName string `json:"product_name"` //商品名称
	Count       int    `json:"count"`        //查询次数
	SearchTime  string `json:"search_time"`  //查询时间
}
