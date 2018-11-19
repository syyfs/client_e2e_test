package turbot

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"brilliance/client_e2e_test/blockchain/chaincode/model/fish"
)

/*func TestPutQuality (t *testing.T) {
	code := new(TracingChaincode)
	stub := shim.NewMockStub("mockStub", code)
	args := []string {`{ "area_code": "350100","area_name": "福建省/福州市","create_time": "2018-07-23 17:12:56","deliver_region": "","farmer": "001","goods_no": "","goosd_name": "本地黑李","id": 1189727,"qualified": "1","sample_no": "043","status": 1,"supplier_name": "核果","supplier_no": "","test_item_name": "有机磷、氨基甲酸酯类农药残留","test_person": "永辉超市","test_site": "永辉超市食品安全云网","test_time": "2016-08-05 15:48:06","test_unit": "%","test_value": "10.6","trace_code": "20160805350100A01043","update_time": "2016-08-05 15:56:29"}`}

	stub.MockTransactionStart("1111")
	code.putQuality(stub, args)
	stub.MockTransactionEnd("1111")

	var quality Quality
	qualityBytes, err := stub.GetState("latest")
	if err != nil {
		t.Fatalf("getState error...")
	}

	err = json.Unmarshal(qualityBytes, &quality)
	if err != nil {
		t.Fatalf("json Unmarshal error...")
	}
	fmt.Println(&quality)
}*/

func TestPutTraceInfo(t *testing.T) {
	code := new(TracingChaincode)
	stub := shim.NewMockStub("mockStub", code)
	args := []string{`{"batch_id": "6455340500", "breed_company": "刘宁养殖场","breed_contact": "13358803987","breed_person": "刘宁","catch_time": "2018/6/1 19:14:16","catches": [{"create_time": "2018-07-23 16:47:19","id": 6404,"name": "","status": "1","tracing_id": 6404,"update_time": "2018-07-04 13:14:34","url": "http://yu.jywykjgs.com/bbt/upload/20180111140651805.jpg"}],"certs": [{"create_time": "2018-07-23 16:47:19","id": 12807,"name": "鱼苗信息","status": "1","tracing_id": 6404,"update_time": "2018-07-04 13:14:34","url": "http://yu.jywykjgs.com/bbt/waterimg.jsp?/bbt/upload/20180602001317849.jpg"},{"create_time": "2018-07-23 16:47:19","id": 12808,"name": "水质监测","status": "1","tracing_id": 6404,"update_time": "2018-07-04 13:14:34","url": "http://yu.jywykjgs.com/bbt/waterimg.jsp?/bbt/upload/20180602001314641.jpg"}],"check_org": "兴城市佳盈伟业商贸有限公司质检部","check_person": "杨娇","check_result": "质检合格","check_time": "2018/5/28 11:52:20","create_time": "2018-07-23 16:47:19","delivery_company": "永辉超市股份有限公司","delivery_person": "郑辉","delivery_time": "2018-06-02T06:16:11","fries": [{"create_time": "2018-07-23 16:47:19","id": 12807,"name": "小鱼苗","status": "1","tracing_id": 6404,"update_time": "2018-07-04 13:14:34","url": "http://bbt.jywykjgs.com/duobaoyu/UpFiles/6b7978bb-29bf-465d-9c39-7be421fac782%E5%B0%8F%E9%B1%BC%E8%8B%97.jpg"},{"create_time": "2018-07-23 16:47:19","id": 12808,"name": "小鱼苗","status": "1","tracing_id": 6404,"update_time": "2018-07-04 13:14:34","url": "http://bbt.jywykjgs.com/duobaoyu/UpFiles/a59b213c-d5a9-4ec9-93c5-007a8a066053%E5%B0%8F%E9%B1%BC%E8%8B%971.jpg"}],"goods_name": "半边天多宝鱼","growth": "0.5-0.7斤1年，0.8-1斤14月，1-1.5斤18月，3斤以上3年","id": 6404,"order_time": "2018/5/28 8:11:41","plate_number": "辽PD1315","run_company": "福建永辉现代农业发展有限公司","salinity": "24","sample_ids": ["111"],"status": "1","temperature": "14","trade_number": "","transport_person": "叶彪","transport_time": "2018-06-01T20:33:40","transport_type": "个体","update_time": "2018-07-04 13:14:34"}`}

	//args1 := []string {`{"batch_id": "6455340500", "breed_company": "刘宁养殖场","breed_contact": "13358803987","breed_person": "刘宁","catch_time": "2018/6/1 19:14:16","catches": [{"create_time": "2018-07-23 16:47:19","id": 6404,"name": "","status": "1","tracing_id": 6404,"update_time": "2018-07-04 13:14:34","url": "http://yu.jywykjgs.com/bbt/upload/20180111140651805.jpg"}],"certs": [{"create_time": "2018-07-23 16:47:19","id": 12807,"name": "鱼苗信息","status": "1","tracing_id": 6404,"update_time": "2018-07-04 13:14:34","url": "http://yu.jywykjgs.com/bbt/waterimg.jsp?/bbt/upload/20180602001317849.jpg"},{"create_time": "2018-07-23 16:47:19","id": 12808,"name": "水质监测","status": "1","tracing_id": 6404,"update_time": "2018-07-04 13:14:34","url": "http://yu.jywykjgs.com/bbt/waterimg.jsp?/bbt/upload/20180602001314641.jpg"}],"check_org": "兴城市佳盈伟业商贸有限公司质检部","check_person": "杨娇","check_result": "质检合格","check_time": "2018/5/28 11:52:20","create_time": "2018-07-23 16:47:19","delivery_company": "永辉超市股份有限公司","delivery_person": "郑辉","delivery_time": "2018-06-02T06:16:11","fries": [{"create_time": "2018-07-23 16:47:19","id": 12807,"name": "小鱼苗","status": "1","tracing_id": 6404,"update_time": "2018-07-04 13:14:34","url": "http://bbt.jywykjgs.com/duobaoyu/UpFiles/6b7978bb-29bf-465d-9c39-7be421fac782%E5%B0%8F%E9%B1%BC%E8%8B%97.jpg"},{"create_time": "2018-07-23 16:47:19","id": 12808,"name": "小鱼苗","status": "1","tracing_id": 6404,"update_time": "2018-07-04 13:14:34","url": "http://bbt.jywykjgs.com/duobaoyu/UpFiles/a59b213c-d5a9-4ec9-93c5-007a8a066053%E5%B0%8F%E9%B1%BC%E8%8B%971.jpg"}],"goods_name": "半边天多宝鱼","growth": "0.5-0.7斤1年，0.8-1斤14月，1-1.5斤18月>，3斤以上3年","id": 6404,"order_time": "2018/5/28 8:11:41","plate_number": "辽PD1315","run_company": "福建永辉现代农业发展有限公司","salinity": "24","sample_ids": ["111"],"status": "1","temperature": "14","trade_number": "","transport_person": "叶彪","transport_time": "2018-06-01T20:33:40","transport_type": "个体","update_time": "2018-07-04 13:14:34"}`}

	//args := []string{`{ "id": 11,"goods_name": "半边天多宝鱼","run_company": "福建永辉现代农业发展有限公司","order_time": "2018/1/8 9:04:53","check_org": "兴城市佳盈伟业商贸有限公司质检部","check_time": "2018/1/8 13:15:51","check_result": "质检合格","check_person": "陆文爽","breed_company": "刘宁养殖场","catch_time": "2018/1/11 12:08:36","temperature": "","salinity": "24%","growth": "0.5-0.7斤1年，0.8-1斤14月，1-1.5斤18月，3斤以上3年","breed_person": "刘宁","breed_contact": "13358803987","transport_type": "个体","transport_time": "2018-01-11 13:00:49","plate_number": "辽PD1315","transport_person": "徐光灿","delivery_company": "北京--北京永辉超市有限公司","delivery_time": "2018-01-11 22:54:43","delivery_person": "刘彪","trade_number": "111111","create_time": "","update_time": "","status": 200,"certs": [{"name": "鱼苗信3","url": "http://yu.jywykjgs.com/bbt/fishimg.jsp?/bbt/upload/20180111140539408.jpg"},{"name": "水质检测111","url": "http://yu.jywykjgs.com/bbt/waterimg.jsp?/bbt/upload/20180111140524105.jpg"}],"fries": [{"name": "小鱼苗","url": "http://bbt.jywykjgs.com/duobaoyu/UpFiles/6b7978bb-29bf-465d-9c39-7be421fac782%E5%B0%8F%E9%B1%BC%E8%8B%97.jpg"},{"name": "小鱼苗","url": "http://bbt.jywykjgs.com/duobaoyu/UpFiles/a59b213c-d5a9-4ec9-93c5-007a8a066053%E5%B0%8F%E9%B1%BC%E8%8B%971.jpg"}],"catches": [{"name": "捕捞照片","url": "http://yu.jywykjgs.com/bbt/upload/20180111140651805.jpg"}],"sample_ids": ["11111"],"batch_id": "11111","order_id": "11111","goods_id": "11111"}`}
	stub.MockTransactionStart("2222")
	response := code.putTraceInfo(stub, args)
	//code.putTraceInfo(stub, args1)
	fmt.Println(response.GetPayload())

	var traceInfo fish.TraceInfo
	traceInfoBytes, err := stub.GetState("2222")
	if err != nil {
		t.Fatalf("getState error...")
	}

	err = json.Unmarshal(traceInfoBytes, &traceInfo)
	if err != nil {
		t.Fatalf("json Unmarshal error...")
	}
	fmt.Println(&traceInfo)
	stub.MockTransactionEnd("2222")

	/*var queryCode string
	queryCodeByte, err := stub.GetState("111")
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = json.Unmarshal(queryCodeByte, &queryCode)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(queryCode)*/
}

/*func TestGetTraceInfo (t *testing.T) {
	code := new(TracingChaincode)
	stub := shim.NewMockStub("mockStub", code)
	args := []string {"111"}
	args1 := []string {`{"batch_id": "6455340500", "breed_company": "刘宁养殖场","breed_contact": "13358803987","breed_person": "刘宁","catch_time": "2018/6/1 19:14:16","catches": [{"create_time": "2018-07-23 16:47:19","id": 6404,"name": "","status": "1","tracing_id": 6404,"update_time": "2018-07-04 13:14:34","url": "http://yu.jywykjgs.com/bbt/upload/20180111140651805.jpg"}],"certs": [{"create_time": "2018-07-23 16:47:19","id": 12807,"name": "鱼苗信息","status": "1","tracing_id": 6404,"update_time": "2018-07-04 13:14:34","url": "http://yu.jywykjgs.com/bbt/waterimg.jsp?/bbt/upload/20180602001317849.jpg"},{"create_time": "2018-07-23 16:47:19","id": 12808,"name": "水质监测","status": "1","tracing_id": 6404,"update_time": "2018-07-04 13:14:34","url": "http://yu.jywykjgs.com/bbt/waterimg.jsp?/bbt/upload/20180602001314641.jpg"}],"check_org": "兴城市佳盈伟业商贸有限公司质检部","check_person": "杨娇","check_result": "质检合格","check_time": "2018/5/28 11:52:20","create_time": "2018-07-23 16:47:19","delivery_company": "永辉超市股份有限公司","delivery_person": "郑辉","delivery_time": "2018-06-02T06:16:11","fries": [{"create_time": "2018-07-23 16:47:19","id": 12807,"name": "小鱼苗","status": "1","tracing_id": 6404,"update_time": "2018-07-04 13:14:34","url": "http://bbt.jywykjgs.com/duobaoyu/UpFiles/6b7978bb-29bf-465d-9c39-7be421fac782%E5%B0%8F%E9%B1%BC%E8%8B%97.jpg"},{"create_time": "2018-07-23 16:47:19","id": 12808,"name": "小鱼苗","status": "1","tracing_id": 6404,"update_time": "2018-07-04 13:14:34","url": "http://bbt.jywykjgs.com/duobaoyu/UpFiles/a59b213c-d5a9-4ec9-93c5-007a8a066053%E5%B0%8F%E9%B1%BC%E8%8B%971.jpg"}],"goods_name": "半边天多宝鱼","growth": "0.5-0.7斤1年，0.8-1斤14月，1-1.5斤18月，3斤以上3年","id": 6404,"order_time": "2018/5/28 8:11:41","plate_number": "辽PD1315","run_company": "福建永辉现代农业发展有限公司","salinity": "24","sample_ids": ["111"],"status": "1","temperature": "14","trade_number": "","transport_person": "叶彪","transport_time": "2018-06-01T20:33:40","transport_type": "个体","update_time": "2018-07-04 13:14:34"}`}
	args2 := []string {`{ "area_code": "350100","area_name": "福建省/福州市","create_time": "2018-07-23 17:12:56","deliver_region": "","farmer": "001","goods_no": "","goosd_name": "本地黑李","id": 1189727,"qualified": "1","sample_no": "043","status": 1,"supplier_name": "核果","supplier_no": "","test_item_name": "有机磷、氨基甲酸酯类农药残留","test_person": "永辉超市","test_site": "永辉超市食品安全云网","test_time": "2016-08-05 15:48:06","test_unit": "%","test_value": "10.6","trace_code": "20160805350100A01043","update_time": "2016-08-05 15:56:29"}`}

	stub.MockTransactionStart("3333")

	code.putTraceInfo(stub, args1)

	code.putQuality(stub, args2)

	traceAndQuality := code.getTraceInfo(stub, args)
	fmt.Println(traceAndQuality)
	var tAndQuality TraceAndQuality
	err := json.Unmarshal(traceAndQuality.GetPayload(), &tAndQuality)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(tAndQuality.Quality)
	fmt.Println(tAndQuality.TraceInfos)
	stub.MockTransactionEnd("3333")
}*/
