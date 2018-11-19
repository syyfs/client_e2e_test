package blackPig

import "brilliance/client_e2e_test/blockchain/chaincode/model"

type TraceInfo struct {
	model.Batch
	model.Transport
	QuarantineTicket1 model.ImageInfoUpload `json:"quarantine_ticket1"` //
	QuarantineTicket2 model.ImageInfoUpload `json:"quarantine_ticket2"`
	ExaminingReport   model.ImageInfoUpload `json:"examining_report"`
	OrderTime         string                `json:"order_time"`
	Total             string                `json:"total"`
	Weight            string                `json:"weight"`
	InsertTime        string                `json:"insert_time"`
	ModifyTime        string                `json:"modify_time"`
	Mark              string                `json:"mark"`
}
