package mongo

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const BASE_PRODUCT = "base_product"

type Product struct {
	Id            int64     `bson:"_id" json:"id"`                          //主键 TODO 临时自己生成,time.Now().Unix()
	Name          string    `bson:"name" json:"name"`                       //商品名称
	Url           string    `bson:"url" json:"url"`                         //商品链接
	Money         string    `bson:"money" json:"money"`                     //商品价格
	ProductTypeId int64     `bson:"product_type_id" json:"product_type_id"` //商品类目ID
	Tag           string    `bson:"tag" json:"tag"`                         //商品标签
	PicUrl        []string  `bson:"pic_url" json:"pic_url"`                 //商品图片
	Desc          string    `bson:"desc" json:"desc"`                       //商品描述
	Other         string    `bson:"other" json:"other"`                     //扩展信息
	CreateTime    time.Time `bson:"create_time" json:"create_time"`         //创建时间
	UpdateTime    time.Time `bson:"update_time" json:"update_time"`         //更新时间
	SupplierId    string    `bson:"supplier_id" json:"supplier_id"`         //供应商ID
	SupplierName  string    `bson:"supplier_name" json:"supplier_name"`     //供应商
	CcName        string    `bson:"cc_name" json:"cc_name"`                 //chainCode name

	Yunwang YunWang `bson:"yunwang" json:"yunwang"` //云网商品信息
}

type YunWang struct {
	ProductList []YunwangProduct `bson:"product_list" json:"product_list"`
}

type YunwangProduct struct {
	YWBuyUrl    string `bson:"yw_buy_url" json:"yw_buy_url"`
	YWName      string `bson:"yw_name" json:"yw_name"`
	YWProductId string `bson:"yw_product_id" json:"yw_product_id"`
}

func BatchAddproducts(docs ...interface{}) error {
	err := baseProductC.Insert(docs...)
	if err != nil {
		return err
	}

	return nil
}

func AddProduct(product Product) error {
	err := baseProductC.Insert(product)
	if err != nil {
		return err
	}

	return nil
}

func FindProductListsByTypeId(typeId int64) (infos []Product, err error) {
	err = baseProductC.Find(bson.M{"product_type_id": typeId}).All(&infos)
	if err != nil {
	}
	return
}

func FindProductById(id int64) (info Product, err error) {
	err = baseProductC.Find(bson.M{"_id": id}).One(&info)
	if err != nil {
	}
	return
}
