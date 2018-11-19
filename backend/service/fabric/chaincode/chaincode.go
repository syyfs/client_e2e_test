package chaincode

import (
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/resource"
	"fmt"
	"io/ioutil"
	"os"
)

func GetCcPolicy() *common.SignaturePolicyEnvelope{
	ccPolicy := cauthdsl.SignedByAnyMember([]string{"Org1MSP"})
	return ccPolicy
}


func ReadChaincodePkg(fileName string) ([]byte, error){

	files , err := ioutil.ReadDir("../../../../chaincode/")
	for _, file := range files{

		fmt.Printf("file.Name() ===> [%s]\n", file.Name())
		fmt.Printf("file.Size() ===> [%d]\n", file.Size())
	}
	file , err := os.OpenFile(fileName,os.O_RDWR|os.O_CREATE, 0755)
	//file, err := os.Open(fileName)
	//if err != nil {
	//	fmt.Printf("open file ./e2ecc.tar err : %v\n", err)
	//	return nil , err
	//}
	fmt.Printf("file.Name() ===> [%s]\n", file.Name())

	//defer file.Close()
	//fmt.Printf("file ===> [%s]")
	//read := tar.NewReader(file)
	//hdr, err := read.Next()
	//fmt.Printf("hdr.Name ===> [%s]\n", hdr.Name)
	//fmt.Printf("hdr.Size ===> [%d]\n", hdr.Size)
	//
	//var code = make([]byte, hdr.Size)
	//
	//_, err = file.Read(code)
	//if err != nil {
	//	fmt.Printf("read err : %v\n", err)
	//	return nil , err
	//}
	return nil , err
}

func GetCCPkg(fileName string) (*resource.CCPackage, error){
	code , err := ReadChaincodePkg(fileName)
	if err != nil{
		return nil ,err
	}

	pkg := &resource.CCPackage{
		Type: pb.ChaincodeSpec_GOLANG,
		Code: code,
	}

	return pkg ,err
}