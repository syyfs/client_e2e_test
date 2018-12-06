package chaincode

import (
	"testing"
	"os"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/stretchr/testify/require"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"brilliance/client_e2e_test/blockchain/common/config"
	fabcfg "github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/test/integration"
)

// 安装chaincode； 实例化chaincode
var configPath = "../../../../config/config.yaml"
var chaincodeName = "mycc1"
var version = "1.1"
var chaincodePath = "example02/cmd"
var channelID = "mychannel"

var filename = "/data/gopath/src/brilliance/client_e2e_test/blockchain/chaincode/src.tar.gz"


func TestInstallChaincode(t *testing.T){
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

	// TODO 根据组织是如何获取自己组织的peer节点的
	// Org resource management client
	orgResMgmt, err := resmgmt.New(adminContext)
	if err != nil {
		t.Fatalf("Failed to create new resource management client: %s", err)
	}

	//ccPkg, err := packager.NewCCPackage("example/e2e_cc/example02/cmd", "/data/gopath/src/brilliance/client_e2e_test/blockchain/chaincode")
	ccPkg, err := GetCCPkg()
	if err != nil {
		t.Fatal(err)
	}
	// Install example cc to org peers
	installCCReq := resmgmt.InstallCCRequest{Name: chaincodeName, Path:chaincodePath , Version: version, Package: ccPkg}
	_, err = orgResMgmt.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		t.Fatal(err)
	}

}


func TestInstantiateCC(t *testing.T)  {
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

	//mspids := []string{"Org1MSP","Org2MSP","Org3MSP"}
	//clientids := []string{"Org1"}
	//Set up chaincode policy
	//ccPolicy := GetCcPolicySignedByAnyMember(mspids)
	// todo 未测试通过
	//ccPolicy := GetCcPolicySignedByAnyClient(clientids)
	//ccPolicy := GetCcPolicySignedByAnyPeer(mspids)
	//ccPolicy := GetCcPolicySignedByAnyAdmin(mspids)
	//ccPolicy := GetCcPolicySignedByMSPAdmin("Org2MSP")
	//ccPolicy := SignedByAssignMember(2,[]string{"Org1MSP","Org2MSP","Org3MSP"})
	ccPolicy := SignedByGivenRoleMP([]string{"Org1MSP","Org2MSP","Org3MSP"})
	// Org resource manager will instantiate 'example_cc' on channel
	resp, err := orgResMgmt.InstantiateCC(
		channelID,
		resmgmt.InstantiateCCRequest{Name: chaincodeName, Path: chaincodePath, Version: version, Args: integration.ExampleCCInitArgs(), Policy: ccPolicy},
		resmgmt.WithRetry(retry.DefaultResMgmtOpts),
	)
	require.Nil(t, err, "error should be nil")
	require.NotEmpty(t, resp, "transaction response should be populated")
}

func TestQueryChaincodes(t *testing.T) {
	os.Setenv("FABRIC_ARTIFACTS", "../../../../")
	err := config.InitConfig(configPath)
	if err != nil {
		t.Error(err)
	}
	configProvider := fabcfg.FromFile(config.GetConfigFile())
	sdk , err := fabsdk.New(configProvider)
	require.NoError(t, err, "Failed to create new SDK")
	defer sdk.Close()
	chaincodes := QueryChaincodes(sdk)
	t.Logf("response : %#v \n",chaincodes)
}