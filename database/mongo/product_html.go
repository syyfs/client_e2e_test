package mongo

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type ProductHtml struct {
	Id         int64     `bson:"_id" json:"id"`                  //	主键
	ProductId  int64     `bson:"product_id" json:"product_id"`   //商品名称
	Feature    string    `bson:"feature" json:"feature"`         //对应功能
	Router     string    `bson:"router" json:"router"`           //对应路由
	CreateTime time.Time `bson:"create_time" json:"create_time"` //创建时间
	UpdateTime time.Time `bson:"update_time" json:"update_time"` //更新时间
}

const BASE_PRODUCT_PAGE = "base_product_html"

func BatchAddproductHtml(docs ...interface{}) error {
	err := baseProductHtmlC.Insert(docs...)
	if err != nil {
		return err
	}

	return nil
}

func AddProductHtml(html ProductHtml) error {
	err := baseProductHtmlC.Insert(html)
	if err != nil {
		return err
	}

	return nil
}

func FindProductHtml(feature string, productId int64) (info ProductHtml, err error) {
	err = baseProductHtmlC.Find(bson.M{"product_id": productId, "feature": feature}).One(&info) //TODO 商品名称和chaincode名称对应关系
	if err != nil {
		return ProductHtml{}, err
	}
	return
}
