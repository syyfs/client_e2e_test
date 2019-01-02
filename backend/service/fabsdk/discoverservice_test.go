package fabsdk

import (
	"brilliance/client_e2e_test/blockchain/common/config"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/common/discovery/dynamicdiscovery"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/context"
	fabcfg "github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/comm"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/discovery"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk/factory/defsvc"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk/provider/chpvdr"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	discclient "github.com/hyperledger/fabric-sdk-go/pkg/client/common/discovery"
	"testing"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"time"
)

var configPath = "../../../config/config.yaml"

const (
	peer0Org1 = "peer0.org1.example.com"
)
func TestClientDiscoveryE2E(t *testing.T) {

	os.Setenv("FABRIC_ARTIFACTS", "../../../")
	err := config.InitConfig(configPath)
	if err != nil {
		t.Error(err)
	}
	configProvider := fabcfg.FromFile(config.GetConfigFile())
	sdk , err := fabsdk.New(configProvider,fabsdk.WithServicePkg(&DynamicDiscoveryProviderFactory{}))
	require.NoError(t, err, "Failed to create new SDK")
	defer sdk.Close()

	chProvider := sdk.ChannelContext("mychannel", fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))
	//chCtx, err := chProvider()
	//require.NoError(t, err, "Error creating channel context")

	chCtx, err := chProvider()
	require.NoError(t, err, "Error creating channel context")

	discoveryService, err := chCtx.ChannelService().Discovery()
	require.NoError(t, err, "Error creating discovery service")

	peers, err := discoveryService.GetPeers()
	for _, p := range peers {
		fmt.Printf("- [%s] - MSP [%s] \n ", p.URL(), p.MSPID())
	}

	require.NoError(t, err, "Error  discovery Invoke")
}


func TestClient_DiscoveryService(t *testing.T) {
	os.Setenv("FABRIC_ARTIFACTS", "../../../")
	err := config.InitConfig(configPath)
	if err != nil {
		t.Error(err)
	}
	configProvider := fabcfg.FromFile(config.GetConfigFile())
	sdk , err := fabsdk.New(configProvider,fabsdk.WithServicePkg(&DynamicDiscoveryProviderFactory{}))
	require.NoError(t, err, "Failed to create new SDK")
	defer sdk.Close()
	chProvider := sdk.Context(fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))
	chCtx, err := chProvider()
	require.NoError(t, err, "Error creating channel context")

	var client *discovery.Client
	client, err = discovery.New(chCtx)
	require.NoError(t, err, "error creating discovery client")

	reqCtx, cancel := context.NewRequest(chCtx, context.WithTimeout(10*time.Second))
	defer cancel()


	peerCfg1, err := comm.NetworkPeerConfig(chCtx.EndpointConfig(), "peer0.org1.example.com")
	require.NoErrorf(t, err, "error getting peer config for [%s]", "peer0.org1.example.com")

	req := discclient.NewRequest().OfChannel(orgChannelID).AddPeersQuery()
	responses, err := client.Send(reqCtx, req,peerCfg1.PeerConfig)

	for _ , res := range responses {
		fmt.Printf("res.Target():[%s] ; res.ForChannel:[%#v ]\n ",res.Target(),res.ForChannel(orgChannelID))
	}

	resp := responses[0]
	chanResp := resp.ForChannel(orgChannelID)
	peers, err := chanResp.Peers()

	//chanResp.Endorsers()
	require.NoError(t, err, "error getting peers")
	require.NotEmpty(t, peers, "expecting at least one peer but got none")

	for _, peer := range peers {
		aliveMsg := peer.AliveMessage.GetAliveMsg()
		if !assert.NotNil(t, aliveMsg, "got nil AliveMessage") {
			continue
		}
		sID := &msp.SerializedIdentity{}
		err := json.Unmarshal(peer.Identity, sID.SerializedIdentity)
		if err != nil {
			errors.Wrap(err, "failed unmarshaling peer's identity")
		}
		fmt.Printf("--- Endpoint: %s \n", aliveMsg.Membership.Endpoint)
		//fmt.Printf("----- peer.Identity:[%s]\n",string(peer.Identity))
		if !assert.NotNil(t, aliveMsg.Membership, "got nil Membership") {
			continue
		}
		if !assert.NotNil(t, peer.StateInfoMessage, "got nil StateInfoMessage") {
			continue
		}
		stateInfo := peer.StateInfoMessage.GetStateInfo()
		if !assert.NotNil(t, stateInfo, "got nil stateInfo") {
			continue
		}
		if !assert.NotNil(t, stateInfo.Properties, "got nil stateInfo.Properties") {
			continue
		}
		//fmt.Printf("--- Ledger Height: %d \n", stateInfo.Properties.LedgerHeight)
		for _, cc := range stateInfo.Properties.Chaincodes {
			fmt.Printf("------Chaincodes: %s:%s \n", cc.Name, cc.Version)
		}
	}

}

func TestClient_DiscoveryService_GetCATlsCert(t *testing.T) {
	os.Setenv("FABRIC_ARTIFACTS", "../../../")
	err := config.InitConfig(configPath)
	if err != nil {
		t.Error(err)
	}
	configProvider := fabcfg.FromFile(config.GetConfigFile())
	sdk , err := fabsdk.New(configProvider,fabsdk.WithServicePkg(&DynamicDiscoveryProviderFactory{}))
	require.NoError(t, err, "Failed to create new SDK")
	defer sdk.Close()
	chProvider := sdk.Context(fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))
	chCtx, err := chProvider()
	require.NoError(t, err, "Error creating channel context")

	var client *discovery.Client
	client, err = discovery.New(chCtx)
	require.NoError(t, err, "error creating discovery client")

	reqCtx, cancel := context.NewRequest(chCtx, context.WithTimeout(10*time.Second))
	defer cancel()


	peerCfg1, err := comm.NetworkPeerConfig(chCtx.EndpointConfig(), "peer0.org1.example.com")
	require.NoErrorf(t, err, "error getting peer config for [%s]", "peer0.org1.example.com")

	req := discclient.NewRequest().OfChannel(orgChannelID).AddConfigQuery()
	responses, err := client.Send(reqCtx, req,peerCfg1.PeerConfig)

	for _ , res := range responses {
		fmt.Printf("res.Target():[%s] ; res.ForChannel:[%#v ]\n ",res.Target(),res.ForChannel(orgChannelID))
	}

	resp := responses[0]
	chanResp := resp.ForChannel(orgChannelID)
	//peers, err := chanResp.Peers()
	configResult , err := chanResp.Config()
	//chanResp.Endorsers()
	require.NoError(t, err, "error getting config")
	require.NotEmpty(t, configResult, "expecting at least one peer but got none")

	for _, tlscert := range configResult.Msps["Org1MSP"].TlsRootCerts {

		fmt.Printf("==== cert :[%s] \n", string(tlscert))
	}

}

type DynamicDiscoveryProviderFactory struct {
	// 外部引用
	defsvc.ProviderFactory
}

type channelProvider struct {
	fab.ChannelProvider
	services map[string]*dynamicdiscovery.ChannelService
}

type channelService struct {
	fab.ChannelService
	discovery fab.DiscoveryService
}

// 对ProviderFactory中的CreateChannelProvider()函数进行了重写
// CreateChannelProvider returns a new default implementation of channel provider
func (f *DynamicDiscoveryProviderFactory) CreateChannelProvider(config fab.EndpointConfig) (fab.ChannelProvider, error) {
	chProvider, err := chpvdr.New(config)
	if err != nil {
		return nil, err
	}
	return &channelProvider{
		ChannelProvider: chProvider,
		services:        make(map[string]*dynamicdiscovery.ChannelService),
	}, nil
}

// Close frees resources and caches. Close释放资源和缓存
func (cp *channelProvider) Close() {
	if c, ok := cp.ChannelProvider.(closable); ok {
		c.Close()
	}
	for _, discovery := range cp.services {
		discovery.Close()
	}
}

// ChannelService creates a ChannelService for an identity
func (cp *channelProvider) ChannelService(ctx fab.ClientContext, channelID string) (fab.ChannelService, error) {
	chService, err := cp.ChannelProvider.ChannelService(ctx, channelID)
	if err != nil {
		return nil, err
	}

	membership, err := chService.Membership()
	if err != nil {
		return nil, err
	}

	discovery, ok := cp.services[channelID]
	if !ok {
		discovery, err = dynamicdiscovery.NewChannelService(ctx, membership, channelID)
		if err != nil {
			return nil, err
		}
		cp.services[channelID] = discovery
	}

	return &channelService{
		ChannelService: chService,
		discovery:      discovery,
	}, nil
}

func (cs *channelService) Discovery() (fab.DiscoveryService, error) {
	return cs.discovery, nil
}




type closable interface {
	Close()
}