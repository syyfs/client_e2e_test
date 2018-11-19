package fish

import (
	"time"

	"brilliance/client_e2e_test/blockchain/chaincode/model"
)

type TraceInfo struct {
	model.BOGInfo
	GoodsName       string                  `json:"goods_name"`       //商品名称
	RunCompany      string                  `json:"run_company"`      //经营企业
	OrderTime       time.Time               `json:"order_time"`       //出单时间
	CheckOrg        string                  `json:"check_org"`        //检测单位
	CheckTime       time.Time               `json:"check_time"`       //检测时间
	CheckResult     string                  `json:"check_result"`     //检测结果
	CheckPerson     string                  `json:"check_person"`     //检测人
	BreedCompany    string                  `json:"breed_company"`    //养殖企业
	CatchTime       time.Time               `json:"catch_time"`       //捕捞时间
	Temperature     string                  `json:"temperature"`      //水温
	Salinity        string                  `json:"salinity"`         //盐度
	Growth          string                  `json:"growth"`           //生长周期
	BreedPerson     string                  `json:"breed_person"`     //养殖操作人
	BreedContact    string                  `json:"breed_contact"`    //养殖人联系方式
	TransportType   string                  `json:"transport_type"`   //公司/个体（运输方式）
	TransportTime   time.Time               `json:"transport_time"`   //装车时间
	PlateNumber     string                  `json:"plate_number"`     //车牌号
	TransportPerson string                  `json:"transport_person"` //运输操作人
	DeliveryCompany string                  `json:"delivery_company"` //经销商
	DeliveryTime    time.Time               `json:"delivery_time"`    //到达时间
	DeliveryPerson  string                  `json:"delivery_person"`  //经销商操作人
	TradeNumber     string                  `json:"trade_number"`     //交易码
	CreateTime      time.Time               `json:"create_time"`      //创建时间
	UpdateTime      time.Time               `json:"update_time"`      //更新时间
	Status          string                  `json:"status"`           //状态 default : 1
	Certs           []model.ImageInfoUpload `json:"certs"`            //检测报告
	Fries           []model.ImageInfoUpload `json:"fries"`            //鱼苗信息
	Catches         []model.ImageInfoUpload `json:"catches"`          //捕捞信息
}

type TraceInfoUpload struct {
	TraceInfo
	FishIds []string `json:"fish_ids"`
}

type TraceAndQuality struct {
	Uuid      string               `json:"uuid"`
	TraceInfo TraceInfo            `json:"traceInfo"`
	Quality   []model.FreshQuality `json:"quality"`
}
