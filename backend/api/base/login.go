package base

import (
	"net/http"

	"encoding/json"
	"io/ioutil"

	"time"

	"fmt"

	"github.com/gin-gonic/gin"
	"brilliance/client_e2e_test/blockchain/database/mongo"
)

type WXRet struct {
	Openid     string `json:"openid"`      //用户唯一标识
	SessionKey string `json:"session_key"` //会话密钥
	Unionid    string `json:"unionid"`     //用户在开放平台的唯一标识符
}

func Login(c *gin.Context) {
	code, _ := c.GetQuery("code")

	if len(code) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": http.StatusInternalServerError,
			"message": "login failed,no code",
		})
		return
	}

	resp, err := http.Post("https://api.weixin.qq.com/sns/jscode2session?appid=wxf2b6d3ecff3e8ec4&secret=f3b6768718757f4200e058b89a8b91ed&js_code="+code+"&grant_type=authorization_code", "", nil)

	if err != nil {
		fmt.Println("post to weixin's login service got error, and the error is : ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": http.StatusInternalServerError,
			"message": "login failed 1",
		})
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("weixin return is : ", string(body))

	if len(body) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": http.StatusInternalServerError,
			"message": "login failed 2",
		})
		return
	}

	var ret WXRet
	err = json.Unmarshal(body, &ret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": http.StatusInternalServerError,
			"message": "login failed 3 ",
		})
		return
	}

	if len(ret.Openid) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": http.StatusInternalServerError,
			"message": "login failed 3.5 ",
		})
		return
	}

	//添加用户信息
	now := time.Now()
	err = mongo.AddUserInfo(mongo.UserInfo{
		Id:         now.Unix(),
		Openid:     ret.Openid,
		SessionKey: ret.SessionKey,
		Unionid:    ret.Unionid,
		CreateTime: now,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": http.StatusInternalServerError,
			"message": "login failed 4",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":       http.StatusOK,
		"message":       "success",
		"openid":        ret.Openid,
		"unionid":       ret.Unionid,
		"local_session": "", //业务判断是否是登录状态,
		"sessoin_key":   ret.SessionKey,
	})

}
