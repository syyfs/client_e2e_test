package mongo

import (
	"fmt"

	"time"

	"gopkg.in/mgo.v2/bson"
)

type TraceRecord struct {
	ProductId   int64     `bson:"product_id" json:"product_id"`     //商品名称
	TraceId     string    `bson:"trace_id" json:"trace_id"`         //商品ID
	ProductName string    `bson:"product_name" json:"product_name"` //商品名称
	Count       int       `bson:"count" json:"count"`               //查询次数
	SearchTime  time.Time `bson:"search_time" json:"search_time"`   //查询时间
}

func AddTraceRecord(productId int64, traceId string) error {
	exists, err := traceRecordC.Find(bson.M{"product_id": productId, "trace_id": traceId}).Count()
	if err != nil {
		fmt.Println("get record count error : ", err.Error())
		return err
	}

	if exists == 0 {
		fmt.Println("a new record to be added")
		//获取商品名称
		product, err := FindProductByProductId(productId)
		if err != nil {
			fmt.Println("get productName error : ", err.Error())
			return err
		}

		record := TraceRecord{
			ProductId:   productId,
			TraceId:     traceId,
			ProductName: product.Name + "-" + traceId,
			Count:       1,
			SearchTime:  time.Now(),
		}

		err = traceRecordC.Insert(record)

	} else {
		fmt.Println("a record got another scan")
		info, err := FindTraceRecord(productId, traceId)
		if err != nil {
			fmt.Println("find trace record error : ", err.Error())
			return err
		}
		info.Count += 1
		info.SearchTime = time.Now()

		err = traceRecordC.Update(bson.M{"product_id": productId, "trace_id": traceId}, info)
	}

	if err != nil {
		fmt.Println("update error : ", err.Error())
		return err
	}

	return nil
}

func FindTraceRecord(productId int64, traceId string) (info TraceRecord, err error) {
	err = traceRecordC.Find(bson.M{"product_id": productId, "trace_id": traceId}).One(&info)
	if err != nil {
		fmt.Println("error is : ", err.Error())
		return TraceRecord{}, err
	}
	return
}

func FindProductByProductId(productId int64) (info Product, err error) {
	err = baseProductC.Find(bson.M{"_id": productId}).One(&info)
	if err != nil {
	}
	return
}
