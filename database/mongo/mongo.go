package mongo

import (
	"gopkg.in/mgo.v2"
)

var mongoDB *mgo.Session

const MongoDBName = "ytrace"

func OpenMongo(url string) {
	mongoDB, _ = mgo.Dial(url)
	mongoDB.SetMode(mgo.Monotonic, true)

	Init()
}

func CloseMongo() {
	mongoDB.Clone()
}

var baseProductC *mgo.Collection
var baseProductTypeC *mgo.Collection
var baseProductHtmlC *mgo.Collection
var baseUserC *mgo.Collection
var qualityC *mgo.Collection
var traceRecordC *mgo.Collection

func Init() {
	baseProductC = mongoDB.DB(MongoDBName).C(BASE_PRODUCT)
	baseProductTypeC = mongoDB.DB(MongoDBName).C(BASE_PRODUCT_TYPE)
	baseProductHtmlC = mongoDB.DB(MongoDBName).C(BASE_PRODUCT_PAGE)
	baseUserC = mongoDB.DB(MongoDBName).C(BASE_USER)
	qualityC = mongoDB.DB(MongoYunWangDbName).C(MongoCollectionFreshQualityName)
	traceRecordC = mongoDB.DB(MongoDBName).C("traceRecord")
}
