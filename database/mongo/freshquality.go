package mongo

import (
	"gopkg.in/mgo.v2/bson"
	"brilliance/client_e2e_test/blockchain/chaincode/model"
	"brilliance/client_e2e_test/blockchain/chaincode/util"
)

const (
	MongoYunWangDbName              = "ytrace"
	MongoCollectionFreshQualityName = "yw_jian_ce_data"
)

func GetFreshQualities(page int) (qualities []model.FreshQuality, err error) {
	err = qualityC.Find(nil).Limit(util.PageSize).Skip((page - 1) * util.PageSize).Sort("-createtime").All(&qualities)
	return
}

func GetFreshQualitiesById(goodsId string, page int) (qualities []model.FreshQuality, err error) {
	err = qualityC.Find(bson.M{"boginfo.goodsid": goodsId}).Limit(util.PageSize).Skip((page - 1) * util.PageSize).Sort("-createtime").All(&qualities)
	return
}
