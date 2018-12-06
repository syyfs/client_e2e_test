package chaincode

import (
	"brilliance/client_e2e_test/blockchain/backend/service/model"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	contextImpl "github.com/hyperledger/fabric-sdk-go/pkg/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/resource"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"
	"github.com/protobuf/proto"
	"io/ioutil"
	// /data/gopath/src/github.com/hyperledger/fabric-sdk-go/pkg/fabsdk/fabsdk.go:8
)

func GetCcPolicySignedByAnyMember(ids []string) *common.SignaturePolicyEnvelope{
	ccPolicy := cauthdsl.SignedByAnyMember(ids)
	return ccPolicy
}

func GetCcPolicySignedByAnyClient(ids []string) *common.SignaturePolicyEnvelope{
	ccPolicy := cauthdsl.SignedByAnyClient(ids)
	return ccPolicy
}

func GetCcPolicySignedByAnyPeer(ids []string) *common.SignaturePolicyEnvelope{
	ccPolicy := cauthdsl.SignedByAnyPeer(ids)
	return ccPolicy
}

func GetCcPolicySignedByAnyAdmin(ids []string) *common.SignaturePolicyEnvelope{
	ccPolicy := cauthdsl.SignedByAnyAdmin(ids)
	return ccPolicy
}

func GetCcPolicySignedByMSPAdmin(mspid string) *common.SignaturePolicyEnvelope{
	ccPolicy := cauthdsl.SignedByMspAdmin(mspid)
	return ccPolicy
}

func SignedByAssignMember(n int32, mspids []string) *common.SignaturePolicyEnvelope {
	ccPolicy := cauthdsl.SignedByAssignMember(n , mspids)
	return ccPolicy
}

func SignedByGivenRoleMP( mspids []string) *common.SignaturePolicyEnvelope {
	ccPolicy := cauthdsl.SignedByGivenRoleMP( mspids)
	return ccPolicy
}

func ParseSignPolicyEnvelop(data []byte ) (*common.SignaturePolicyEnvelope, error){
	spe := &common.SignaturePolicyEnvelope{}
	err := proto.Unmarshal(data,spe)
	if err != nil {
		return nil , err
	}

	return spe, err
}


func ReadChaincodePkg() ([]byte, error){

	code, err := ioutil.ReadFile("/data/gopath/src/brilliance/client_e2e_test/blockchain/chaincode/src.tar.gz")
	if err != nil {
		return nil , err
	}

	return code , err
}

func GetCCPkg() (*resource.CCPackage, error){
	code , err := ReadChaincodePkg()
	if err != nil{
		return nil ,err
	}

	pkg := &resource.CCPackage{
		Type: pb.ChaincodeSpec_GOLANG,
		Code: code,
	}

	return pkg ,err
}



func QueryChaincodes(sdk *fabsdk.FabricSDK)(chaincodes []model.ChaincodeInfo){
	//prepare context
	ctxProvider := sdk.Context(fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))
	ctx, err := ctxProvider()
	if err != nil {
		return
	}
	reqCtx, cancel := contextImpl.NewRequest(ctx, contextImpl.WithTimeoutType(fab.PeerResponse))
	defer cancel()
	peersCfg, _ := ctx.EndpointConfig().PeersConfig("Org1")
	if len(peersCfg) < 1 {
		err = errors.New(fmt.Sprintf("can't find peer in 'Org1'", ))
		return
	}

	chaincodemap := make(map[string]model.ChaincodeInfo)
	for _, peerCfg := range peersCfg {
		peer, e := ctx.InfraProvider().CreatePeerFromConfig(&fab.NetworkPeer{PeerConfig: peerCfg})
		if e != nil {
			fmt.Errorf("InfraProvider().CreatePeerFromConfig faild [%s]\n",err)
			continue
		}
		chaincodeQueryResponse, e := resource.QueryInstalledChaincodes(reqCtx, peer, resource.WithRetry(retry.DefaultResMgmtOpts))
		if e != nil {
			fmt.Errorf("QueryInstalledChaincodes faild [%s]\n",err)
			continue
		}
		for _, c := range chaincodeQueryResponse.Chaincodes {
			_ , ok := chaincodemap[c.Name]
			fmt.Printf("\n %c[1;40;32m ok: %#v %c[0m\n\n", 0x1B, ok,0x1B)
			if  !ok {
				chaincodemap[c.Name] =  model.ChaincodeInfo{
					Name:    c.Name,
					Version: c.Version,
					Path:    c.Path,
				}
				chaincodes = append(chaincodes, model.ChaincodeInfo{
					Name:    c.Name,
					Version: c.Version,
					Path:    c.Path,
				})
			}
		}
		break
	}
	fmt.Printf("\n %c[1;40;32m chaincodeMap: %#v %c[0m\n\n", 0x1B, chaincodemap,0x1B)
	fmt.Printf("\n %c[1;40;31m chaincodes: %#v %c[0m\n\n", 0x1B, chaincodes,0x1B)
	return
}