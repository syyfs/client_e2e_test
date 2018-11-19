package service

type InvokeConfig struct {
	ChannelId      string
	CcName         string
	CcFcn          string
	CcArgs         [][]byte
	CcTransientMap map[string][]byte
}

type InvokeResult struct {
	TxId    string
	Payload []byte
}

type FabricClient interface {
	Close() error
	Query(cfg InvokeConfig) (InvokeResult, error)
	Execute(cfg InvokeConfig) (InvokeResult, error)
}

var FabClient FabricClient
