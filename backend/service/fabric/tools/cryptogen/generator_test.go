package cryptogen

import (
	"brilliance/fast-deploy/common/config"
	"fmt"
	"os"
	"testing"
)

var configPath = "../../../../../config/config.yaml"

func TestGenerateOrgCert(t *testing.T) {
	err := os.Setenv("FABRIC_CFG_PATH", "../../../../../config/")
	err = config.InitConfig(configPath)
	if err != nil {
		fmt.Printf("****** Set env FABRIC_CFG_PATH faild ! err is [%s]**** \n", err)
	}
	generator, _ := NewGeneratorGmCerts()
	ordererorg, err := generator.CreateOrganization("orderer.citic.com")
	if err != nil {
		fmt.Printf("****** CreateOrganization is faild ! err is [%s]**** \n", err)
	}
	fmt.Printf("***** orderer.citic.com is name = [%s] ; ******\n", ordererorg.Name)
	fmt.Printf("***** orderer.citic.com MSPCACert = [%s] ; ******\n", ordererorg.MSPCACert)
	fmt.Printf("***** orderer.citic.com MSPCAKey = [%s] ; ******\n", ordererorg.MSPCAKey)
	fmt.Printf("***** orderer.citic.com TLSCACert = [%s] ; ******\n", ordererorg.TLSCACert)
	fmt.Printf("***** orderer.citic.com TLSCAKey = [%s] ; ******\n", ordererorg.TLSCAKey)




	peerorg, err := generator.CreateOrganization("peer.citic.com")
	if err != nil {
		fmt.Printf("****** CreateOrganization is faild ! err is [%s]**** \n", err)
	}
	fmt.Printf("***** *****************************\n")
	fmt.Printf("***** peer.citic.com is name = [%s] ; ******\n", peerorg.Name)
	fmt.Printf("***** peer.citic.com MSPCACert = [%s] ; ******\n", peerorg.MSPCACert)
	fmt.Printf("***** peer.citic.com MSPCAKey = [%s] ; ******\n", peerorg.MSPCAKey)
	fmt.Printf("***** peer.citic.com TLSCACert = [%s] ; ******\n", peerorg.TLSCACert)
	fmt.Printf("***** peer.citic.com TLSCAKey = [%s] ; ******\n", peerorg.TLSCAKey)

}



func TestCreateOrganizationWithPath(t *testing.T){
	generator, _ := NewGeneratorGmCerts()
	ordererorg, err := generator.CreateOrganizationWithPath("orderer.citic.com")
	if err != nil {
		fmt.Printf("****** CreateOrganization is faild ! err is [%s]**** \n", err)
	}
	fmt.Printf("***** orderer.citic.com is name = [%s] ; ******\n", ordererorg.Name)
	fmt.Printf("***** orderer.citic.com MSPCACert = [%s] ; ******\n", ordererorg.MSPCACert)
	fmt.Printf("***** orderer.citic.com MSPCAKey = [%s] ; ******\n", ordererorg.MSPCAKey)
	fmt.Printf("***** orderer.citic.com TLSCACert = [%s] ; ******\n", ordererorg.TLSCACert)
	fmt.Printf("***** orderer.citic.com TLSCAKey = [%s] ; ******\n", ordererorg.TLSCAKey)
}

func TestGenerateNodeCertByOrg(t *testing.T) {
	err := os.Setenv("FABRIC_CFG_PATH", "../../../../../config/")
	err = config.InitConfig(configPath)
	if err != nil {
		fmt.Printf("****** Set env FABRIC_CFG_PATH faild ! err is [%s]**** \n", err)
	}
	generator, _ := NewGeneratorGmCerts()
	ordererOrg, err := generator.CreateOrganization("orderer.citic.com")
	fmt.Printf("*******************************************\n")
	fmt.Printf("***** ordererOrg is name = [%s] ; ******\n", ordererOrg.Name)
	fmt.Printf("***** ordererOrg MSPCACert = [%s] ; ******\n", ordererOrg.MSPCACert)
	fmt.Printf("***** ordererOrg MSPCAKey = [%s] ; ******\n", ordererOrg.MSPCAKey)
	fmt.Printf("***** ordererOrg TLSCACert = [%s] ; ******\n", ordererOrg.TLSCACert)
	fmt.Printf("***** ordererOrg TLSCAKey = [%s] ; ******\n", ordererOrg.TLSCAKey)

	orderernode, err := generator.CreateNode("node0.orderer.citic.com", ordererOrg)
	if err != nil {
		fmt.Printf("****** CreateOrganization is faild ! err is [%s]**** \n", err)
	}

	fmt.Printf("*****************************************************\n")
	fmt.Printf("***** orderernode  is name = [%s] ; ******\n", orderernode.Name)
	fmt.Printf("***** orderernode MSPCACert = [%s] ; ******\n", orderernode.MSPCert)
	fmt.Printf("***** orderernode MSPCAKey = [%s] ; ******\n", orderernode.MSPKey)
	fmt.Printf("***** orderernode TLSCACert = [%s] ; ******\n", orderernode.TLSCert)
	fmt.Printf("***** orderernode TLSCAKey = [%s] ; ******\n", orderernode.TLSKey)

	ordereradmin, err := generator.CreateUser("Admin@orderer.citic.com", ordererOrg)
	if err != nil {
		fmt.Printf("****** CreateOrganization is faild ! err is [%s]**** \n", err)
	}
	fmt.Printf("*****************************************************\n")
	fmt.Printf("***** ordereradmin  is name = [%s] ; ******\n", ordereradmin.Name)
	fmt.Printf("***** ordereradmin MSPCACert = [%s] ; ******\n", ordereradmin.MSPCert)
	fmt.Printf("***** ordereradmin MSPCAKey = [%s] ; ******\n", ordereradmin.MSPKey)
	fmt.Printf("***** ordereradmin TLSCACert = [%s] ; ******\n", ordereradmin.TLSCert)
	fmt.Printf("***** ordereradmin TLSCAKey = [%s] ; ******\n", ordereradmin.TLSKey)



	peerorg, err := generator.CreateOrganization("peer.citic.com")
	fmt.Printf("******************peer*************************\n")
	fmt.Printf("***** peerorg is name = [%s] ; ******\n", peerorg.Name)
	fmt.Printf("***** peerorg MSPCACert = [%s] ; ******\n", peerorg.MSPCACert)
	fmt.Printf("***** peerorg MSPCAKey = [%s] ; ******\n", peerorg.MSPCAKey)
	fmt.Printf("***** peerorg TLSCACert = [%s] ; ******\n", peerorg.TLSCACert)
	fmt.Printf("***** peerorg TLSCAKey = [%s] ; ******\n", peerorg.TLSCAKey)

	peernode, err := generator.CreateNode("node0.peer.citic.com", peerorg)
	if err != nil {
		fmt.Printf("****** CreateOrganization is faild ! err is [%s]**** \n", err)
	}

	fmt.Printf("*****************************************************\n")
	fmt.Printf("***** node0  is name = [%s] ; ******\n", peernode.Name)
	fmt.Printf("***** node0 MSPCACert = [%s] ; ******\n", peernode.MSPCert)
	fmt.Printf("***** node0 MSPCAKey = [%s] ; ******\n", peernode.MSPKey)
	fmt.Printf("***** node0 TLSCACert = [%s] ; ******\n", peernode.TLSCert)
	fmt.Printf("***** node0 TLSCAKey = [%s] ; ******\n", peernode.TLSKey)

	peeradmin, err := generator.CreateUser("Admin@peer.citic.com", peerorg)
	if err != nil {
		fmt.Printf("****** CreateOrganization is faild ! err is [%s]**** \n", err)
	}

	fmt.Printf("*****************************************************\n")
	fmt.Printf("***** peeradmin  is name = [%s] ; ******\n", peeradmin.Name)
	fmt.Printf("***** peeradmin MSPCACert = [%s] ; ******\n", peeradmin.MSPCert)
	fmt.Printf("***** peeradmin MSPCAKey = [%s] ; ******\n", peeradmin.MSPKey)
	fmt.Printf("***** peeradmin TLSCACert = [%s] ; ******\n", peeradmin.TLSCert)
	fmt.Printf("***** peeradmin TLSCAKey = [%s] ; ******\n", peeradmin.TLSKey)

}

