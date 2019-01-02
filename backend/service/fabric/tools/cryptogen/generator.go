package cryptogen

import (
	gmx509 "brilliance/gm/x509"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
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
	tlsprivStr, tlspubStr, err := GenerateTLSPrivAndPub(  orgSpec.CommonName, "tls"+orgSpec.CA.CommonName, orgSpec.CA.Country, orgSpec.CA.Province, orgSpec.CA.Locality, orgSpec.CA.OrganizationalUnit, orgSpec.CA.StreetAddress, orgSpec.CA.PostalCode)
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

func (g *GeneratorGmCertsImpl) CreateOrganizationWithPath(commonName string) (org Organization, err error) {
	//cakeystorepath, tlscakeystorepath, err := GenerateCaDir()
	dir, _ := filepath.Abs("./")
	baseDir := filepath.Join(dir, "temp")
	orgDir := filepath.Join(baseDir, "peerOrganizations", commonName)
	caDir := filepath.Join(orgDir, "ca")
	tlsCADir := filepath.Join(orgDir, "tlsca")

	//mspDir := filepath.Join(orgDir, "msp")
	//peersDir := filepath.Join(orgDir, "peers")
	//usersDir := filepath.Join(orgDir, "users")
	//adminCertsDir := filepath.Join(mspDir, "admincerts")


	orgSpec := NewOrgSpec()
	orgSpec.CommonName = commonName
	canodeSpec := InitCANodeSpec(commonName, Country, Province, Locality) //
	orgSpec.CA = canodeSpec
	caprivStr, capubStr, err := GeneratePrivAndPubWithPath( caDir, orgSpec.CommonName, orgSpec.CA.CommonName, orgSpec.CA.Country, orgSpec.CA.Province, orgSpec.CA.Locality, orgSpec.CA.OrganizationalUnit, orgSpec.CA.StreetAddress, orgSpec.CA.PostalCode)
	if err != nil {
		return org, err
	}
	tlsprivStr, tlspubStr, err := GenerateTLSPrivAndPubWithPath( tlsCADir, orgSpec.CommonName, "tls"+orgSpec.CA.CommonName, orgSpec.CA.Country, orgSpec.CA.Province, orgSpec.CA.Locality, orgSpec.CA.OrganizationalUnit, orgSpec.CA.StreetAddress, orgSpec.CA.PostalCode)
	if err != nil {
		return org, err
	}

	// 生成组织msp目录
	//err = createFolderStructure(mspDir, false)
	//if err != nil {
	//	return org, err
	//}
	//if err == nil {
	//	// the signing CA certificate goes into cacerts
	//	err = x509Export(filepath.Join(baseDir, "cacerts", x509Filename(commonName)), capubStr)
	//	if err != nil {
	//		return org, err
	//	}
	//	// the TLS CA certificate goes into tlscacerts
	//	err = x509Export(filepath.Join(baseDir, "tlscacerts", x509Filename(tlsCA.Name)), tlsCA.SignCert)
	//	if err != nil {
	//		return org, err
	//	}
	//}
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
	casignerCert, err := gmx509.ParseCertificate(pemCert.Bytes)
	casigner, err := GetSignPriveKey(org.MSPCAKey)
	nodeMspCaKeyStr, nodeMspCaCertStr, err := GetNodeCACert(commonName, casignerCert, casigner)

	tlspemCert, _ := pem.Decode([]byte(org.TLSCACert))
	tlscasignerCert, err := gmx509.ParseCertificate(tlspemCert.Bytes)
	tlscasigner, err := GetTLSSignPriveKey(org.TLSCAKey)
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

func x509Filename(name string) string {
	return name + "-cert.pem"
}
func x509Export(path string, cert *x509.Certificate) error {
	return pemExport(path, "CERTIFICATE", cert.Raw)
}

func pemExport(path, pemType string, bytes []byte) error {
	//write pem out to file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return pem.Encode(file, &pem.Block{Type: pemType, Bytes: bytes})
}
func createFolderStructure(rootDir string, local bool) error {

	var folders []string
	// create admincerts, cacerts, keystore and signcerts folders
	folders = []string{
		filepath.Join(rootDir, "admincerts"),
		filepath.Join(rootDir, "cacerts"),
		filepath.Join(rootDir, "tlscacerts"),
	}
	if local {
		folders = append(folders, filepath.Join(rootDir, "keystore"),
			filepath.Join(rootDir, "signcerts"))
	}

	for _, folder := range folders {
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
