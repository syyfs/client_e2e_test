package fabsdk

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/common/selection/fabricselection"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk/factory/defsvc"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"brilliance/client_e2e_test/blockchain/common/config"
	fabcfg "github.com/hyperledger/fabric-sdk-go/pkg/core/config"
)


var orgChannelID = "mychannel"

type fabricSelectionProviderFactory struct {
	defsvc.ProviderFactory
}

func TestFabricSelection(t *testing.T){
	os.Setenv("FABRIC_ARTIFACTS", "../../../")
	err := config.InitConfig(configPath)
	if err != nil {
		t.Error(err)
	}
	configProvider := fabcfg.FromFile(config.GetConfigFile())
	sdk , err := fabsdk.New(configProvider,fabsdk.WithServicePkg(&fabricSelectionProviderFactory{}))
	require.NoError(t, err, "Failed to create new SDK")
	defer sdk.Close()
	ctxProvider := sdk.ChannelContext(orgChannelID, fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))
	ctx, err := ctxProvider()
	require.NoError(t, err, "error getting channel context")

	selectionService, err := ctx.ChannelService().Selection()
	require.NoError(t, err)

	_, ok := selectionService.(*fabricselection.Service)
	if !ok {

	}
	//endorsers, err := selectionService.GetEndorsersForChaincode(chaincodes, opts...)
}
