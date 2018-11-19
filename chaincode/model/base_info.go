package model

type BOGInfo struct {
	Uuid    string `bson:"uuid" json:"uuid"`
	BatchId string `bson:"batch_id" json:"batch_id"` //批次编码
	OrderId string `bson:"order_id" json:"order_id"` //订单编码
	GoodsId string `bson:"goods_id" json:"goods_id"` //商品编码
}

/**
检测报告、鱼苗信息、捕捞信息
*/
type ImageInfoUpload struct {
	Name string `json:"name"` //名称
	Url  string `json:"url"`  //图片URL
}

type Transport struct {
	Id              string `json:"id"`
	TransportType   string `json:"transport_type"`
	TransportPerson string `json:"transport_person"`
	TransportTime   string `json:"transport_time"`
	PlateNumber     string `json:"plate_number"`
}

type Batch struct {
	Id             string `json:"id"`
	GoodsId        string `json:"goods_id"`
	GoodsName      string `json:"goods_name"`
	SupplierId     string `json:"supplier_id"`
	SupplierName   string `json:"supplier_name"`
	StoreId        string `json:"store_id"`
	StoreName      string `json:"store_name"`
	DeliveryTime   string `json:"delivery_time"`
	DeliveryPerson string `json:"delivery_person"`
	ReceiveTime    string `json:"receive_time"`
	ReceivePerson  string `json:"receive_person"`
	CreateTime     string `json:"create_time"`
	BatchCode      string `json:"batch_code"`
}
