package trace

import (
	"net/http"

	"fmt"

	"net/url"

	"strings"

	"github.com/gin-gonic/gin"
	"brilliance/client_e2e_test/blockchain/database/mongo"
)

type SearchRet struct {
}

const BASEURL = "https://sy.yonghui.cn/trace/"

func Search(c *gin.Context) {

	/*
		多宝鱼二维码内容：http://yu.jywykjgs.com/bbt/hotel/product/syInfoTiaos/5326077199
		规范二维码内容： https://sy.yonghui.cn/trace/qr?productId=921108&traceId=123

		扫码获取二维码内容，
		1、获取二维码内容（链接）
		2、解析链接，拿到productId&traceId
		3、构建新的请求链接 https://sy.yonghui.cn/trace/html/:productId/traceId
		4、返回给前端
	*/

	qrcode := c.Query("qrcode")
	u, err := url.Parse(qrcode)
	if err != nil {
		fmt.Println("url.Parse error : ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": http.StatusInternalServerError,
			"message": "failed",
			"router":  BASEURL + "404",
		})
		return

	}

	if u.Host != "sy.yonghui.cn" {
		//不是正规二维码，可能是供应商提供的二维码
		fmt.Println("not usuals qrcode")
		router := ScanUnusualQRCode(qrcode)
		c.JSON(http.StatusOK, gin.H{
			"success": http.StatusOK,
			"message": "success",
			"router":  router,
		})

		return
	}

	mapData, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		fmt.Println("url.ParseQuery error : ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": http.StatusInternalServerError,
			"message": "failed",
			"router":  BASEURL + "404",
		})
		return
	}

	productId := mapData["productId"]
	traceId := mapData["traceId"]
	if len(productId) > 0 && len(productId[0]) > 0 && len(traceId) > 0 && len(traceId[0]) > 0 {
		router := GetRedirect(productId[0], traceId[0])
		c.JSON(http.StatusOK, gin.H{
			"success": http.StatusOK,
			"message": "success",
			"router":  router,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": http.StatusOK,
		"message": "success",
		"router":  BASEURL + "404",
	})

}

func GetRedirect(productId, traceId string) (route string) {
	route = BASEURL + "html" + "/showTraceInfo" + "/" + productId + "/" + traceId //二维码扫描暂时只有溯源查询
	return
}

func SearchByQRCode(c *gin.Context) {
	productId := c.Query("productId")
	traceId := c.Query("traceId")

	if len(productId) == 0 || len(traceId) == 0 {
		c.HTML(http.StatusOK, BASEURL, gin.H{})
		return
	}

	c.Redirect(http.StatusMovedPermanently, GetRedirect(productId, traceId))
}

func ShowTraceInfoHtml(feature, openId, traceId string, productId int64) (route string, data interface{}, err error) {
	productHtml, err := mongo.FindProductHtml(feature, productId)
	if err != nil {
		fmt.Println("productPage error : ", err.Error())
		route = "404.html"
		return
	}

	if len(productHtml.Router) == 0 {
		route = "404.html"

		return
	}

	route = productHtml.Router
	data, err = GetTraceInfo(productId, traceId)
	if err != nil {
		fmt.Println("err is : ", err.Error())
		route = "404.html"
		return
	}

	//统计记录，记录到商品名下
	err = mongo.AddTraceRecord(productId, traceId)
	if err != nil {
		fmt.Println("add traceRecord error : ", err.Error())
	}

	if len(openId) > 0 {
		err = mongo.AddSearch(openId, mongo.Search{
			ProductId: productId,
			TraceId:   traceId,
		})

		if err != nil {
			fmt.Println("add search history err is : ", err.Error())
		}
	}

	return
}

func ScanUnusualQRCode(url string) (route string) {
	if strings.Contains(url, "jywykjgs.com") {
		//兴城多宝鱼 http://yu.jywykjgs.com/bbt/hotel/product/syInfoTiaos/5326077199
		//TODO 多宝鱼的ID与商品ID的对应关系
		return GetRedirect("1", "123")
	} else {
		return GetRedirect("1", "123")
	}
}
