package mongo

import (
	"fmt"
	"testing"

	"time"

	"gopkg.in/mgo.v2/bson"
)

func TestAddUserInfo(t *testing.T) {
	OpenMongo("10.0.90.152:21117")
	err := AddUserInfo(UserInfo{Openid: "openId", SessionKey: "session"})
	fmt.Println("error : ", err)
}

func TestAddUserInfo2(t *testing.T) {
	OpenMongo("10.0.90.152:21117")
	err := mongoDB.DB(MongoDBName).C("userInfo").Update(bson.M{"openid": "openId"}, bson.M{"$set": bson.M{"unionid": "unionid"}})
	fmt.Println("error : ", err)
}

func TestFindUserInfoByOpenId(t *testing.T) {
	//OpenMongo("10.0.90.152:21117")
	//mongoDB.DB("test").C("testTime").Insert(bson.M{
	//	"myDate": time.Now(),
	//})
	//
	//type te struct {
	//	MyDate time.Time `bson:"myDate"`
	//}
	//var ss []te
	//mongoDB.DB("test").C("testTime").Find(nil).Sort("-myDate").All(&ss)
	//for index, _ := range ss {
	//	fmt.Println(ss[index].MyDate.Format("2006-01-02 15:04:05"))
	//}
	//fmt.Println("ss is : ", ss)

	//time, err := time.Parse("2016-01-02T15:04:05", "2018-07-27T20:35:24")
	//time, err := time.Parse("2016/01/02 15:04:05", "2018/7/25 12:13:19")
	fmt.Println(time.Unix(1534326467, 0).Format("2006-01-02T15:04:05"))

}
