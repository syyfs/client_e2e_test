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
	fmt.Printf("***** org is name = [%s] ; ******\n", ordererorg.Name)
	fmt.Printf("***** org MSPCACert = [%s] ; ******\n", ordererorg.MSPCACert)
	fmt.Printf("***** org MSPCAKey = [%s] ; ******\n", ordererorg.MSPCAKey)
	fmt.Printf("***** org TLSCACert = [%s] ; ******\n", ordererorg.TLSCACert)
	fmt.Printf("***** org TLSCAKey = [%s] ; ******\n", ordererorg.TLSCAKey)

}

func TestGenerateNodeCertByOrg(t *testing.T) {
	err := os.Setenv("FABRIC_CFG_PATH", "../../../../../config/")
	err = config.InitConfig(configPath)
	if err != nil {
		fmt.Printf("****** Set env FABRIC_CFG_PATH faild ! err is [%s]**** \n", err)
	}
	generator, _ := NewGeneratorGmCerts()
	peerorg, err := generator.CreateOrganization("peer.citic.com")
	fmt.Printf("***** org is name = [%s] ; ******\n", peerorg.Name)
	fmt.Printf("***** org MSPCACert = [%s] ; ******\n", peerorg.MSPCACert)
	fmt.Printf("***** org MSPCAKey = [%s] ; ******\n", peerorg.MSPCAKey)
	fmt.Printf("***** org TLSCACert = [%s] ; ******\n", peerorg.TLSCACert)
	fmt.Printf("***** org TLSCAKey = [%s] ; ******\n", peerorg.TLSCAKey)

	peernode, err := generator.CreateNode("node0.peer.citic.com", peerorg)
	if err != nil {
		fmt.Printf("****** CreateOrganization is faild ! err is [%s]**** \n", err)
	}

	fmt.Printf("***** org is name = [%s] ; ******\n", peernode.Name)
	fmt.Printf("***** org MSPCACert = [%s] ; ******\n", peernode.MSPCert)
	fmt.Printf("***** org MSPCAKey = [%s] ; ******\n", peernode.MSPKey)
	fmt.Printf("***** org TLSCACert = [%s] ; ******\n", peernode.TLSCert)
	fmt.Printf("***** org TLSCAKey = [%s] ; ******\n", peernode.TLSKey)

}

