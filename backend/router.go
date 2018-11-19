package main

import (
	"time"

	"html/template"

	"errors"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"brilliance/client_e2e_test/blockchain/backend/api/asset"
	"brilliance/client_e2e_test/blockchain/backend/api/base"
	"brilliance/client_e2e_test/blockchain/backend/api/chain"
	"brilliance/client_e2e_test/blockchain/backend/api/quality"
	"brilliance/client_e2e_test/blockchain/backend/api/templates"
	"brilliance/client_e2e_test/blockchain/backend/api/trace"
)

func convertToTime(dateTime interface{}) (time.Time, error) {
	switch v := dateTime.(type) {
	case string:
		return time.Parse("2006-01-02 15:04:05", v)
	case time.Time:
		return v, nil
	}
	return time.Now(), errors.New("unsupported date time format")
}

func formatAsDate(dateTime interface{}) string {
	t, err := convertToTime(dateTime)
	if err != nil {
		return err.Error()
	}

	return t.Format("2006-01-02")
}

func formatAsTime(dateTime string) string {
	t, err := convertToTime(dateTime)
	if err != nil {
		return err.Error()
	}

	return t.Format("15:04:05")
}

func initFuncMap(router *gin.Engine) {
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
		"formatAsTime": formatAsTime,
	})
}

func initRouter() *gin.Engine {
	router := gin.Default()
	initFuncMap(router)

	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true

	router.Use(cors.New(cfg))

	router.POST("/asset/:product/:method", asset.PostData)
	router.GET("/asset/:product/:id", asset.GetData)
	router.GET("/chain/block/", chain.GetChainInfo)                   //GetChainInfo
	router.GET("/chain/blockByNum/:number", chain.GetBlockByNumber)   //GetBlockByNumber
	router.GET("/chain/blockByHash/:blockHash", chain.GetBlockByHash) //GetBlockByHash
	router.GET("/chain/blockByTxHash/:txHash", chain.GetBlockByTxID)  //GetBlockByHash
	router.GET("/chain/tx/:txHash", chain.GetTransactionByID)         //GetTransactionByID

	router.GET("/base/login", base.Login) //微信登录

	router.GET("/base/getProductTypes", base.GetTypes) //查询商品列别列表
	router.GET("/base/getProductLists", base.GetList)  //查询商品某个类别的列表

	router.POST("/base/addProductType", base.AddProductType) //添加商品类目
	router.POST("/base/addProduct", base.AddProduct)         //添加商品
	router.POST("/base/addProductHtml", base.AddProductHtml) //添加商品页面

	router.GET("/quality/:id", quality.GetQualityByGoodsId)       //根据商品 id 查询历史质检信息
	router.GET("/quality/:id/:page", quality.GetQualityByGoodsId) //根据商品 id 查询历史质检信息
	router.GET("/qualities", quality.GetQualities)                //查询最新的商品质检信息
	router.GET("/qualities/:page", quality.GetQualities)          //查询最新的商品质检信息

	router.GET("/trace/history", trace.GetSearchList) //查询溯源记录
	router.GET("/trace/scan", trace.Search)           //查询溯源
	router.GET("/qr", trace.SearchByQRCode)           //通过二维码内容直接打开溯源页面

	router.GET("/html/:feature/:productId/:traceId", templates.LoadHtml) //加载页面
	router.GET("/404", templates.Show404)

	router.Static("/global", "templates/global")
	router.LoadHTMLGlob("templates/html/*/*")

	return router
}
