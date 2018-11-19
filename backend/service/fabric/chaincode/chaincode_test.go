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
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"fmt"
)

// 安装chaincode； 实例化chaincode
var configPath = "../../../../config/config.yaml"
var chaincodeName = "mycc"
//var chaincodePath = "/data/gopath/src/yonghui.cn/blockchain/ytrace/chaincode/"
//var chaincodePath = "yonghui.cn/blockchain/ytrace/chaincode/src/example/e2e_cc/example02/cmd"
var fileName = "../../../../chaincode/e2ecc.tar"
// yonghui.cn/blockchain/ytrace/chaincode/src/example/e2e_cc/example02/cmd
var channelID = "mychannel"

var ccPath ="example/e2e_cc/example02/cmd"
//var ccPath1 ="example/e2e_cc/example02"

func TestReadChaincodePkg(t *testing.T){

	pkg, err := ReadChaincodePkg(fileName)
	if err != nil{
		t.Errorf("ReadChaincodePkg Faile! err ===> [%s]\n", err)
	}

	t.Logf("pkg ===> [%v]\n", pkg)
}

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

	// Org resource management client
	orgResMgmt, err := resmgmt.New(adminContext)
	if err != nil {
		t.Fatalf("Failed to create new resource management client: %s", err)
	}

	ccPkg, err := packager.NewCCPackage("example/e2e_cc/example02/cmd", "/data/gopath/src/brilliance/client_e2e_test/blockchain/chaincode")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("ccPkg ===> [%#v]\n", ccPkg)
	// Install example cc to org peers
	installCCReq := resmgmt.InstallCCRequest{Name: chaincodeName, Path: "example/e2e_cc/example02/cmd", Version: "1.0", Package: ccPkg}
	_, err = orgResMgmt.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		t.Fatal(err)
	}
	// Set up chaincode policy
	//ccPolicy := GetCcPolicy()
	//// Org resource manager will instantiate 'example_cc' on channel
	//resp, err := orgResMgmt.InstantiateCC(
	//	channelID,
	//	resmgmt.InstantiateCCRequest{Name: "mycc", Path: "github.com/example_cc", Version: "1.0", Args: integration.ExampleCCInitArgs(), Policy: ccPolicy},
	//	resmgmt.WithRetry(retry.DefaultResMgmtOpts),
	//)
	//require.Nil(t, err, "error should be nil")
	//require.NotEmpty(t, resp, "transaction response should be populated")


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


	//Set up chaincode policy
	ccPolicy := GetCcPolicy()
	// Org resource manager will instantiate 'example_cc' on channel
	resp, err := orgResMgmt.InstantiateCC(
		channelID,
		resmgmt.InstantiateCCRequest{Name: chaincodeName, Path: "example/e2e_cc/example02/cmd", Version: "1.0", Args: integration.ExampleCCInitArgs(), Policy: ccPolicy},
		resmgmt.WithRetry(retry.DefaultResMgmtOpts),
	)
	require.Nil(t, err, "error should be nil")
	require.NotEmpty(t, resp, "transaction response should be populated")
}