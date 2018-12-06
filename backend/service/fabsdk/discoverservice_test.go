package fabsdk

import (
	"brilliance/client_e2e_test/blockchain/common/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/common/discovery/dynamicdiscovery"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	fabcfg "github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk/factory/defsvc"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk/provider/chpvdr"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var configPath = "../../../config/config.yaml"


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
		t.Logf("- [%s] - MSP [%s]", p.URL(), p.MSPID())

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

	chProvider := sdk.ChannelContext("mychannel", fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))
	chCtx, err := chProvider()
	require.NoError(t, err, "Error creating channel context")

	discoveryService, err := chCtx.ChannelService().Discovery()
	require.NoError(t, err, "Error creating discovery service")

	peers, err := discoveryService.GetPeers()

	for _, p := range peers {
		t.Logf("- [%s] - MSP [%s]", p.URL(), p.MSPID())
	}


	//t.Log(client)

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