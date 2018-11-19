package chaincode

import (
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/resource"
	"io/ioutil"
)

func GetCcPolicy() *common.SignaturePolicyEnvelope{
	ccPolicy := cauthdsl.SignedByAnyMember([]string{"Org1MSP"})
	return ccPolicy
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