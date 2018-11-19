package quality

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"brilliance/client_e2e_test/blockchain/backend/api"
	"brilliance/client_e2e_test/blockchain/chaincode/model"
	"brilliance/client_e2e_test/blockchain/database/mongo"
)

func GetQualityByGoodsId(c *gin.Context) {
	id := c.Param("id")
	page := c.Param("page")
	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		pageNumber = 1
	}
	if pageNumber < 1 {
		pageNumber = 1
	}

	qualities, err := mongo.GetFreshQualitiesById(id, pageNumber)
	if err != nil {
		fmt.Println("server error : ", err.Error())
		c.JSON(http.StatusInternalServerError, api.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	retQualities := FormatQualities(qualities)

	c.JSON(http.StatusOK, api.Response{
		Code: http.StatusOK,
		Data: retQualities,
	})
}

func GetQualities(c *gin.Context) {
	page := c.Param("page")
	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		pageNumber = 1
	}
	if pageNumber < 1 {
		pageNumber = 1
	}

	qualities, err := mongo.GetFreshQualities(pageNumber)
	if err != nil {
		fmt.Println("server error : ", err.Error())
		c.JSON(http.StatusInternalServerError, api.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	retQualities := FormatQualities(qualities)

	c.JSON(http.StatusOK, api.Response{
		Code: http.StatusOK,
		Data: retQualities,
	})
}

func FormatQualities(freshQualities []model.FreshQuality) (qualities []RetQuality) {
	for index, _ := range freshQualities {
		var retQuality RetQuality

		retQuality.FreshQuality = freshQualities[index]
		retQuality.TestTime = freshQualities[index].TestTime.Format("2006-01-02 15:04:05")
		retQuality.CreateTime = freshQualities[index].CreateTime.Format("2006-01-02 15:04:05")
		retQuality.UpdateTime = freshQualities[index].UpdateTime.Format("2006-01-02 15:04:05")

		qualities = append(qualities, retQuality)
	}
	return
}

type RetQuality struct {
	model.FreshQuality
	TestTime   string `bson:"test_time" json:"test_time"`     //检测日期  类型为： timestamp
	CreateTime string `bson:"create_time" json:"create_time"` //创建时间
	UpdateTime string `bson:"update_time" json:"update_time"` //更新时间
}
