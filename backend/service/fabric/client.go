package fabric

import (
	"github.com/cloudflare/cfssl/log"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	fabcfg "github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"brilliance/client_e2e_test/blockchain/backend/service"
	"brilliance/client_e2e_test/blockchain/common/config"
)

type Client struct {
	fabSdk *fabsdk.FabricSDK
	user   string
	mspId  string
}

func NewFabricClient() (service.FabricClient, error) {
	configProvider := fabcfg.FromFile(config.GetConfigFile())
	sdk, err := fabsdk.New(configProvider)

	if err != nil {
		return nil, err
	}

	return &Client{fabSdk: sdk, user: "Admin", mspId: "Org1"}, nil
}

func (c *Client) Close() error {
	c.fabSdk.Close()
	return nil
}

func (c *Client) Query(cfg service.InvokeConfig) (result service.InvokeResult, err error) {
	clientChannelContext := c.fabSdk.ChannelContext(cfg.ChannelId, fabsdk.WithUser(c.user), fabsdk.WithOrg(c.mspId))
	client, err := channel.New(clientChannelContext)
	if err != nil {
		return
	}

	resp, err := client.Query(channel.Request{
		ChaincodeID:  cfg.CcName,
		Fcn:          cfg.CcFcn,
		Args:         cfg.CcArgs,
		TransientMap: cfg.CcTransientMap,
	}, channel.WithTargetEndpoints("peer0.org1.example.com"))

	log.Infof("Query chaincode '%s', TxId=%s", cfg.CcName, resp.TransactionID)
	result.TxId = string(resp.TransactionID)
	result.Payload = resp.Payload
	return
}

func (c *Client) Execute(cfg service.InvokeConfig) (result service.InvokeResult, err error) {
	clientChannelContext := c.fabSdk.ChannelContext(cfg.ChannelId, fabsdk.WithUser(c.user), fabsdk.WithOrg(c.mspId))
	client, err := channel.New(clientChannelContext)
	if err != nil {
		return
	}

	resp, err := client.Execute(channel.Request{
		ChaincodeID:  cfg.CcName,
		Fcn:          cfg.CcFcn,
		Args:         cfg.CcArgs,
		TransientMap: cfg.CcTransientMap,
		// channel.WithTargetEndpoints("peer0.org1.example.com"))
	}, channel.WithTargetEndpoints("peer0.org1.example.com"))

	log.Infof("Query chaincode '%s', TxId=%s", cfg.CcName, resp.TransactionID)
	result.TxId = string(resp.TransactionID)
	result.Payload = resp.Payload
	return
}
