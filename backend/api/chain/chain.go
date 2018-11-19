package chain

import "github.com/op/go-logging"

var logger *logging.Logger

const (
	ChannelId = "yhchannel"
	CcName    = "qscc"
)

type BlockData struct {
	TxId   string `json:"txId"`
	TxType int32  `json:"txType"`
	TxTime string `json:"txTime"`
}

func init() {
	logger = logging.MustGetLogger("chain")
}
