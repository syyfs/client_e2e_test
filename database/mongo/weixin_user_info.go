package mongo

import (
	"fmt"

	"time"

	"gopkg.in/mgo.v2/bson"
)

type UserInfo struct {
	Id         int64     `bson:"_id" json:"id"`                  //主键
	Openid     string    `bson:"openid" json:"openid"`           //用户唯一标识
	SessionKey string    `bson:"session_key" json:"session_key"` //会话密钥
	Unionid    string    `bson:"unionid" json:"unionid"`         //用户在开放平台的唯一标识符
	CreateTime time.Time `bson:"create_time" json:"create_time"` //创建时间
	SearchList []Search  `bson:"search_list" json:"search_list"` //查询记录
}

type Search struct {
	ProductId int64  `bson:"product_id" json:"product_id"` //商品ID
	TraceId   string `bson:"trace_id" json:"trace_id"`     //溯源码
}

const BASE_USER = "base_user"

func AddUserInfo(info UserInfo) error {
	count, err := baseUserC.Find(bson.M{"openid": info.Openid}).Count() //检查是否已经注册
	if err != nil {
		return err
	}

	if count != 0 {
		fmt.Println("already signed")
		return nil
	}

	err = baseUserC.Insert(info)
	if err != nil {
		return err
	}

	return nil
}

func AddSearch(openId string, search Search) error {
	info, err := FindUserInfoByOpenId(openId)
	if err != nil {
		return err
	}

	foundIt := false //确定重复添加
	if len(info.SearchList) != 0 {
		for index, _ := range info.SearchList {
			if search.ProductId == info.SearchList[index].ProductId && search.TraceId == info.SearchList[index].TraceId && !foundIt {
				foundIt = true
				continue
			}
		}
	}

	if !foundIt {
		info.SearchList = append(info.SearchList, search)
		err = baseUserC.Update(bson.M{"openid": openId}, info)
		if err != nil {
			return err
		}
	}

	return nil
}

func FindUserInfoByOpenId(openId string) (info UserInfo, err error) {
	err = baseUserC.Find(bson.M{"openid": openId}).One(&info)
	if err != nil {
		return UserInfo{}, err
	}
	return
}
