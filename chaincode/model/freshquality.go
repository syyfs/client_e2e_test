package model

import "time"

type FreshQuality struct {
	BOGInfo                 //BatchId,orderId,GoodsId
	TraceCode     string    `bson:"trace_code" json:"trace_code"`         //追溯码
	AreaCode      string    `bson:"area_code" json:"area_code"`           //产地编码
	AreaName      string    `bson:"area_name" json:"area_name"`           //产地名称
	Farmer        string    `bson:"farmer" json:"farmer"`                 //农户
	SupplierNo    string    `bson:"supplier_no" json:"supplier_no"`       //供应商编码
	SupplierName  string    `bson:"supplier_name" json:"supplier_name"`   //供应商名称
	GoosdName     string    `bson:"goosd_name" json:"goosd_name"`         //商品名称  单词写错
	TestItemId    string    `bson:"test_item_id" json:"test_item_id"`     //检测项目编号
	TestItemName  string    `bson:"test_item_name" json:"test_item_name"` //检测项目名称
	TestSite      string    `bson:"test_site" json:"test_site"`           //监测站点
	TestPerson    string    `bson:"test_person" json:"test_person"`       //检测人
	TestValue     string    `bson:"test_value" json:"test_value"`         //检测值
	TestUnit      string    `bson:"test_unit" json:"test_unit"`           //检测单位
	TestTime      time.Time `bson:"test_time" json:"test_time"`           //检测日期  类型为： timestamp
	Qualified     string    `bson:"qualified" json:"qualified"`           //是否合格（0：不合格；1：合格）
	DeliverRegion string    `bson:"deliver_region" json:"deliver_region"` //配送区域
	CreateTime    time.Time `bson:"create_time" json:"create_time"`       //创建时间
	UpdateTime    time.Time `bson:"update_time" json:"update_time"`       //更新时间
	Status        int       `bson:"status" json:"status"`                 //状态
	QueryType     int       `bson:"query_type" json:"query_type"`         //通过哪个字段进行查询获取结果
}
