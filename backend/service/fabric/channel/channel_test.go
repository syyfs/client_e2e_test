package channel

import (
	"testing"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"os"
	"github.com/stretchr/testify/require"
	"brilliance/client_e2e_test/blockchain/common/config"
	fabcfg "github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
)

// 创建channel，加入channel
var configPath = "../../../../config/config.yaml"

var ChannelConfigPath = "../../../../config/channel-artifacts/channel.tx"
var ChannelName = "mychannel"

func TestCreateChannel(t *testing.T){
	os.Setenv("FABRIC_ARTIFACTS", "../../../../")
	err := config.InitConfig(configPath)
	if err != nil {
		t.Error(err)
	}
	configProvider := fabcfg.FromFile(config.GetConfigFile())
	sdk , err := fabsdk.New(configProvider)
	require.NoError(t, err, "Failed to create new SDK")
	defer sdk.Close()
	clientContext := sdk.Context(fabsdk.WithUser("Admin"), fabsdk.WithOrg("OrdererOrg"))

	resMgmtClient, err := resmgmt.New(clientContext)


	mspClient, err := mspclient.New(sdk.Context(), mspclient.WithOrg("Org1"))
	if err != nil {
		t.Fatal(err)
	}
	adminIdentity, err := mspClient.GetSigningIdentity("Admin")
	if err != nil {
		t.Fatal(err)
	}
	req := resmgmt.SaveChannelRequest{ChannelID: ChannelName,
		ChannelConfigPath: ChannelConfigPath,
		SigningIdentities: []msp.SigningIdentity{adminIdentity}}
	txID, err := resMgmtClient.SaveChannel(req, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint("orderer.example.com"))
	require.Nil(t, err, "error should be nil")
	require.NotEmpty(t, txID, "transaction ID should be populated")
	t.Logf("txID ===> [%#v]", txID)
}

func TestJoinChannel(t *testing.T){
	os.Setenv("FABRIC_ARTIFACTS", "../../../../")
	err := config.InitConfig(configPath)
	if err != nil {
		t.Error(err)
	}
	configProvider := fabcfg.FromFile(config.GetConfigFile())
	sdk , err := fabsdk.New(configProvider)
	require.NoError(t, err, "Failed to create new SDK")
	defer sdk.Close()
	//prepare context
	adminContext := sdk.Context(fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))

	// Org resource management client
	orgResMgmt, err := resmgmt.New(adminContext)
	if err != nil {
		t.Fatalf("Failed to create new resource management client: %s", err)
	}

	// Org peers join channel
	if err = orgResMgmt.JoinChannel(ChannelName, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint("orderer.example.com")); err != nil {
		t.Fatalf("Org peers failed to JoinChannel: %s", err)
	}

}