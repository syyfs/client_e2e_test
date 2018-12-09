package cryptogen

import (
	"brilliance/gm/x509"
	"encoding/pem"
)

type GeneratorGmCertsImpl struct {
}

func NewGeneratorGmCerts() (Generator, error) {

	return &GeneratorGmCertsImpl{}, nil
}

func (g *GeneratorGmCertsImpl) CreateOrganization(commonName string) (org Organization, err error) {

	//cakeystorepath, tlscakeystorepath, err := GenerateCaDir()

	orgSpec := NewOrgSpec()
	orgSpec.CommonName = commonName
	canodeSpec := InitCANodeSpec(commonName, Country, Province, Locality) //
	orgSpec.CA = canodeSpec

	caprivStr, capubStr, err := GeneratePrivAndPub( orgSpec.CommonName, orgSpec.CA.CommonName, orgSpec.CA.Country, orgSpec.CA.Province, orgSpec.CA.Locality, orgSpec.CA.OrganizationalUnit, orgSpec.CA.StreetAddress, orgSpec.CA.PostalCode)
	if err != nil {
		return org, err
	}
	tlsprivStr, tlspubStr, err := GeneratePrivAndPub( orgSpec.CommonName, "tls"+orgSpec.CA.CommonName, orgSpec.CA.Country, orgSpec.CA.Province, orgSpec.CA.Locality, orgSpec.CA.OrganizationalUnit, orgSpec.CA.StreetAddress, orgSpec.CA.PostalCode)
	if err != nil {
		return org, err
	}

	org.Name = orgSpec.CommonName
	org.MSPCAKey = caprivStr
	org.MSPCACert = capubStr
	org.TLSCAKey = tlsprivStr
	org.TLSCACert = tlspubStr

	return org, err
}

func (g *GeneratorGmCertsImpl) CreateNode(commonName string, org Organization) (node Node, err error) {
	// commonName = node1.peer.citic.com
	//nodekeystory, nodetlskeystory, err := GenerateNodeDir()
	pemCert, _ := pem.Decode([]byte(org.MSPCACert))
	casignerCert, err := x509.ParseCertificate(pemCert.Bytes)
	casigner, err := GetSignPriveKey(org.MSPCAKey)
	nodeMspCaKeyStr, nodeMspCaCertStr, err := GetNodeCACert(commonName, casignerCert, casigner)

	tlspemCert, _ := pem.Decode([]byte(org.TLSCACert))
	tlscasignerCert, err := x509.ParseCertificate(tlspemCert.Bytes)
	tlscasigner, err := GetSignPriveKey(org.TLSCAKey)
	// 设置节点sans
	nodeTlsCaKeyStr, nodeTlsCaCertStr, err := GetNodeTlsCACert(commonName, tlscasignerCert, tlscasigner)

	node.Name = commonName
	node.MSPCert = nodeMspCaCertStr
	node.MSPKey = nodeMspCaKeyStr
	node.TLSCert = nodeTlsCaCertStr
	node.TLSKey = nodeTlsCaKeyStr

	return node, err
}

func (g *GeneratorGmCertsImpl) CreateUser(userName string, org Organization) (user User, err error) {
	node, err := g.CreateNode(userName, org)
	if err != nil {
		return user, nil
	}
	user.Name = userName
	user.MSPKey = node.MSPKey
	user.MSPCert = node.MSPCert
	user.TLSKey = node.TLSKey
	user.TLSCert = node.TLSCert
	return user, err
}
