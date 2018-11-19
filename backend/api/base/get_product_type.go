package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"brilliance/client_e2e_test/blockchain/database/mongo"

)

type AddiProductType struct {
	mongo.ProductType
	Choosen int `json:"choosen"` //是否被选中
}

func GetTypes(c *gin.Context) {
	types, err := mongo.FindAllProductTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": http.StatusInternalServerError,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	if len(types) > 0 {
		infos := trans2Addi(types)

		c.JSON(http.StatusOK, gin.H{
			"success": http.StatusOK,
			"message": "successful",
			"data":    infos,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": http.StatusOK,
			"message": "not found",
			"data":    "",
		})
	}

}

func trans2Addi(types []mongo.ProductType) (infos []AddiProductType) {
	if len(types) == 0 {
		return
	}
	for index, _ := range types {
		infos = append(infos, AddiProductType{ProductType: mongo.ProductType{Id: types[index].Id, Name: types[index].Name, CreateTime: types[index].CreateTime, UpdateTime: types[index].UpdateTime}})
	}
	infos[0].Choosen = 1
	return
}
