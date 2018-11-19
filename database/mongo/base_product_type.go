package mongo

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const BASE_PRODUCT_TYPE = "base_product_type"

type ProductType struct {
	Id         int64     `bson:"_id" json:"id"`                  //商品类目ID
	Name       string    `bson:"name" json:"name"`               //商品类目名称
	CreateTime time.Time `bson:"create_time" json:"create_time"` //创建时间
	UpdateTime time.Time `bson:"update_time" json:"update_time"` //更新时间
}

func BatchAddproductTypes(docs ...interface{}) error {
	err := baseProductTypeC.Insert(docs...)
	if err != nil {
		return err
	}

	return nil
}

func AddProductType(productType ProductType) error {
	err := baseProductTypeC.Insert(productType)
	if err != nil {
		return err
	}

	return nil
}

func FindProductTypeInfoById(typeId int64) (info ProductType, err error) {
	err = baseProductTypeC.Find(bson.M{"_id": typeId}).One(&info)
	if err != nil {
		return ProductType{}, err
	}
	return
}

func FindAllProductTypes() (infos []ProductType, err error) {
	err = baseProductTypeC.Find(bson.M{}).All(&infos)
	if err != nil {
		return nil, err
	}
	return
}

func FindFirstProductType() (info ProductType, err error) {
	err = baseProductTypeC.Find(bson.M{}).One(&info)
	if err != nil {
		return ProductType{}, err
	}
	return
}
