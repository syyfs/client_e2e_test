package cryptogen

import (
	gmx509 "brilliance/gm/x509"
	"crypto"
	"crypto/rand"
	//"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/bccsp/signer"
	"github.com/hyperledger/fabric/bccsp/utils"

	"io/ioutil"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func InitCANodeSpec(commonName, Country, Province, Locality string) NodeSpec {
	canodeSpec := &NodeSpec{}
	canodeSpec.CommonName = "ca." + commonName
	canodeSpec.Country = Country
	canodeSpec.Province = Province
	canodeSpec.Locality = Locality

	return *canodeSpec
}

func cleanup(dir string) {
	os.RemoveAll(dir)
}
/**
生成组织ca私钥和公钥 (GM)
 */
func GeneratePrivAndPub( orgDomain string, commonName string, country, province, locality, orgUnit, streetAddress, postalCode string) (privStr, pubStr string, err error) {

	priv, privsigner, err := GeneratePrivKey() // bccsp.keystore bccsp.key
	fmt.Printf("=== GeneratePrivKey err:[%#v]\n ====",err)
	if err != nil {
		return "", "", err
	}

	// get public signing certificate
	//ecPubKey, err := GetPublicKey(priv)
	ecPubKey := privsigner.Public()
	//ecPubKey, err := GetECPublicKey(priv)
	fmt.Printf("=== GetGMPublicKey err:[%#v]\n ====",err)
	fmt.Printf("=== GetGMPublicKey ecPubKey:[%#v]\n ====",ecPubKey)
	if err != nil {
		return "", "", err
	}
	//template := x509Template()
	template := x509Template()
	//this is a CA
	template.IsCA = true
	template.KeyUsage |= gmx509.KeyUsageDigitalSignature |
		gmx509.KeyUsageKeyEncipherment | gmx509.KeyUsageCertSign |
		gmx509.KeyUsageCRLSign
	template.ExtKeyUsage = []gmx509.ExtKeyUsage{gmx509.ExtKeyUsageAny}

	//set the organization for the subject
	subject := subjectTemplateAdditional(country, province, locality, orgUnit, streetAddress, postalCode)
	subject.Organization = []string{orgDomain}
	subject.CommonName = commonName

	template.Subject = subject
	template.SubjectKeyId = priv.SKI()
	certBytes, err := gmx509.CreateCertificate(rand.Reader, &template, &template, ecPubKey, privsigner)
	fmt.Printf("=== gmx509.CreateCertificate err:[%s]\n ====",err)
	if err != nil {
		return "", "", err
	}

	block := &pem.Block{Type: "CERTIFICATE", Bytes: certBytes}
	pemByte := pem.EncodeToMemory(block)
	pubStr = string(pemByte)
	privStr, err = GetPrivateKey(priv)
	fmt.Printf("GetGmPrivateKey err:[%s] \n", err)
	if err != nil {
		return "", "", err
	}
	fmt.Printf("===========================\n")
	fmt.Printf("privStr := [%s] \n", privStr)
	fmt.Printf("===========================\n")
	return privStr, pubStr, err
}


/**
生成组织ca私钥和公钥 (GM)
 */
func GeneratePrivAndPubWithPath(keystoryPath string, orgDomain string, commonName string, country, province, locality, orgUnit, streetAddress, postalCode string) (privStr, pubStr string, err error) {

	priv, privsigner, err := GeneratePrivKeyWithPath(keystoryPath) // bccsp.keystore bccsp.key
	if err != nil {
		return "", "", err
	}

	// get public signing certificate
	ecPubKey := privsigner.Public()
	if err != nil {
		return "", "", err
	}
	//template := x509Template()
	template := x509Template()
	//this is a CA
	template.IsCA = true
	template.KeyUsage |= gmx509.KeyUsageDigitalSignature |
		gmx509.KeyUsageKeyEncipherment | gmx509.KeyUsageCertSign |
		gmx509.KeyUsageCRLSign
	template.ExtKeyUsage = []gmx509.ExtKeyUsage{gmx509.ExtKeyUsageAny}

	//set the organization for the subject
	subject := subjectTemplateAdditional(country, province, locality, orgUnit, streetAddress, postalCode)
	subject.Organization = []string{orgDomain}
	subject.CommonName = commonName

	template.Subject = subject
	template.SubjectKeyId = priv.SKI()
	certBytes, err := gmx509.CreateCertificate(rand.Reader, &template, &template, ecPubKey, privsigner)
	if err != nil {
		return "", "", err
	}

	block := &pem.Block{Type: "CERTIFICATE", Bytes: certBytes}
	pemByte := pem.EncodeToMemory(block)
	pubStr = string(pemByte)
	privStr, err = GetPrivateKey(priv)
	if err != nil {
		return "", "", err
	}

	return privStr, pubStr, err
}


func GenerateTLSPrivAndPub( orgDomain string, commonName string, country, province, locality, orgUnit, streetAddress, postalCode string) (privStr, pubStr string, err error) {

	priv, privsigner, err := GenerateTLSPrivKey() // bccsp.keystore bccsp.key
	if err != nil {
		return "", "", err
	}

	// get public signing certificate
	//ecPubKey, err := GetECPublicKey(priv)
	ecPubKey, err := GetPublicKey(priv)
	if err != nil {
		return "", "", err
	}
	template := x509Template()
	//this is a CA
	template.IsCA = true
	template.KeyUsage |= gmx509.KeyUsageDigitalSignature |
		gmx509.KeyUsageKeyEncipherment | gmx509.KeyUsageCertSign |
		gmx509.KeyUsageCRLSign
	template.ExtKeyUsage = []gmx509.ExtKeyUsage{gmx509.ExtKeyUsageAny}

	//set the organization for the subject
	subject := subjectTemplateAdditional(country, province, locality, orgUnit, streetAddress, postalCode)
	subject.Organization = []string{orgDomain}
	subject.CommonName = commonName

	template.Subject = subject
	template.SubjectKeyId = priv.SKI()

	certBytes, err := gmx509.CreateCertificate(rand.Reader, &template, &template, ecPubKey, privsigner)
	if err != nil {
		return "", "", err
	}

	block := &pem.Block{Type: "CERTIFICATE", Bytes: certBytes}
	pemByte := pem.EncodeToMemory(block)
	pubStr = string(pemByte)
	privStr, err = GetPrivateKey(priv)
	//privStr, err = LoadPrivateKey(keystorePath)
	if err != nil {
		return "", "", err
	}
	fmt.Printf("===========================\n")
	fmt.Printf("org tls privStr := [%s] \n", privStr)
	fmt.Printf("===========================\n")
	return privStr, pubStr, err
}


func GenerateTLSPrivAndPubWithPath( keystorepath string , orgDomain string, commonName string, country, province, locality, orgUnit, streetAddress, postalCode string) (privStr, pubStr string, err error) {

	priv, privsigner, err := GenerateTLSPrivKeyWithPath(keystorepath) // bccsp.keystore bccsp.key
	if err != nil {
		return "", "", err
	}

	// get public signing certificate
	ecPubKey, err := GetPublicKey(priv)
	if err != nil {
		return "", "", err
	}
	template := x509Template()
	//this is a CA
	template.IsCA = true
	template.KeyUsage |= gmx509.KeyUsageDigitalSignature |
		gmx509.KeyUsageKeyEncipherment | gmx509.KeyUsageCertSign |
		gmx509.KeyUsageCRLSign
	template.ExtKeyUsage = []gmx509.ExtKeyUsage{gmx509.ExtKeyUsageAny}

	//set the organization for the subject
	subject := subjectTemplateAdditional(country, province, locality, orgUnit, streetAddress, postalCode)
	subject.Organization = []string{orgDomain}
	subject.CommonName = commonName

	template.Subject = subject
	template.SubjectKeyId = priv.SKI()

	certBytes, err := gmx509.CreateCertificate(rand.Reader, &template, &template, ecPubKey, privsigner)
	if err != nil {
		return "", "", err
	}

	block := &pem.Block{Type: "CERTIFICATE", Bytes: certBytes}
	pemByte := pem.EncodeToMemory(block)
	pubStr = string(pemByte)
	privStr, err = GetPrivateKey(priv)
	if err != nil {
		return "", "", err
	}

	return privStr, pubStr, err
}

func GetPrivateKey(privKey bccsp.Key) (privStr string , err error){
	priv , err:= privKey.Bytes()
	return string(priv), err
}

func GetSignPriveKey(privKey string) (sign crypto.Signer, err error) {

	opts := &factory.FactoryOpts{
		ProviderName: "GM",
		SwOpts: &factory.SwOpts{
			HashFamily: "SHA2",
			SecLevel:   256,
		},
	}

	csp, err := factory.GetBCCSPFromOpts(opts)
	if err != nil {
		return nil, err
	}

	p, _ := pem.Decode([]byte(privKey))

	privateKey, err := csp.KeyImport(p.Bytes, &bccsp.GMSM2PrivateKeyImportOpts{Temporary: true})
	if err != nil {
		return
	}

	sign, err = signer.New(csp, privateKey)

	return sign, err
}


func GetTLSSignPriveKey(privKey string) (sign crypto.Signer, err error) {

	opts := &factory.FactoryOpts{
		ProviderName: "SW",
		SwOpts: &factory.SwOpts{
			HashFamily: "SHA2",
			SecLevel:   256,
		},
	}

	csp, err := factory.GetBCCSPFromOpts(opts)
	if err != nil {
		return nil, err
	}

	p, _ := pem.Decode([]byte(privKey))

	privateKey, err := csp.KeyImport(p.Bytes, &bccsp.ECDSAPrivateKeyImportOpts{Temporary: true})
	if err != nil {
		return
	}

	sign, err = signer.New(csp, privateKey)

	return sign, err
}


// GeneratePrivateKey creates a private key and stores it in keystorePath
func GeneratePrivKey() (bccsp.Key, crypto.Signer, error) {

	//cleanup(keystorePath)
	var err error
	var priv bccsp.Key
	var s crypto.Signer
	opts := &factory.FactoryOpts{
		ProviderName: "GM",
		SwOpts: &factory.SwOpts{
			HashFamily: "SHA2",
			SecLevel:   256,

			//FileKeystore: &factory.FileKeystoreOpts{
			//	KeyStorePath: keystorePath,
			//},
		},
	}

	csp, err := factory.GetBCCSPFromOpts(opts)
	if err == nil {
		// generate a key  opts KeyGenOpts
		priv, err = csp.KeyGen(&bccsp.GMSM2KeyGenOpts{Temporary: true})
		if err == nil {
			// create a crypto.Signer
			s, err = signer.New(csp, priv)
		}
	}
	return priv, s, err
}

// GeneratePrivateKey creates a private key and stores it in keystorePath
func GeneratePrivKeyWithPath(keystoryPath string) (bccsp.Key, crypto.Signer, error) {

	cleanup(keystoryPath)
	var err error
	var priv bccsp.Key
	var s crypto.Signer
	opts := &factory.FactoryOpts{
		ProviderName: "GM",
		SwOpts: &factory.SwOpts{
			HashFamily: "SHA2",
			SecLevel:   256,

			FileKeystore: &factory.FileKeystoreOpts{
				KeyStorePath: keystoryPath,
			},
		},
	}

	csp, err := factory.GetBCCSPFromOpts(opts)
	if err == nil {
		// generate a key  opts KeyGenOpts
		priv, err = csp.KeyGen(&bccsp.GMSM2KeyGenOpts{Temporary: false})
		if err == nil {
			// create a crypto.Signer
			s, err = signer.New(csp, priv)
		}
	}
	return priv, s, err
}

// GeneratePrivateKey creates a private key and stores it in keystorePath
func GenerateTLSPrivKey() (bccsp.Key, crypto.Signer, error) {

	//cleanup(keystorePath)
	var err error
	var priv bccsp.Key
	var s crypto.Signer
	opts := &factory.FactoryOpts{
		ProviderName: "SW",
		SwOpts: &factory.SwOpts{
			HashFamily: "SHA2",
			SecLevel:   256,

			//FileKeystore: &factory.FileKeystoreOpts{
			//	KeyStorePath: keystorePath,
			//},
		},
	}

	csp, err := factory.GetBCCSPFromOpts(opts)
	if err == nil {
		// generate a key  opts KeyGenOpts
		priv, err = csp.KeyGen(&bccsp.ECDSAP256KeyGenOpts{Temporary: true})
		if err == nil {
			// create a crypto.Signer
			s, err = signer.New(csp, priv)
		}
	}
	return priv, s, err
}

// GeneratePrivateKey creates a private key and stores it in keystorePath
func GenerateTLSPrivKeyWithPath(keystorepath string) (bccsp.Key, crypto.Signer, error) {

	cleanup(keystorepath)
	var err error
	var priv bccsp.Key
	var s crypto.Signer
	opts := &factory.FactoryOpts{
		ProviderName: "SW",
		SwOpts: &factory.SwOpts{
			HashFamily: "SHA2",
			SecLevel:   256,

			FileKeystore: &factory.FileKeystoreOpts{
				KeyStorePath: keystorepath,
			},
		},
	}

	csp, err := factory.GetBCCSPFromOpts(opts)
	if err == nil {
		// generate a key  opts KeyGenOpts
		priv, err = csp.KeyGen(&bccsp.ECDSAP256KeyGenOpts{Temporary: false})
		if err == nil {
			// create a crypto.Signer
			s, err = signer.New(csp, priv)
		}
	}
	return priv, s, err
}

// LoadPrivateKey loads a private key from file in keystorePath
func LoadPrivateKey(keystorePath string) (string, error) {
	var err error

	var rawKey string

	walkFunc := func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, "_sk") {
			rawKeyByte, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			rawKey = string(rawKeyByte)
		}
		return nil
	}

	err = filepath.Walk(keystorePath, walkFunc)
	if err != nil {
		return "", err
	}

	return rawKey, err
}

func GetPublicKey(priv bccsp.Key)(pub interface{}, err error){
	// get the public key
	pubKey, err := priv.PublicKey()
	if err != nil {
		return nil, err
	}
	// marshal to bytes
	pubKeyBytes, err := pubKey.Bytes()
	if err != nil {
		return nil, err
	}

	public , err := utils.DERToPublicKey(pubKeyBytes)
	if err !=  nil {
		return nil , err
	}
	return public, nil
}

// default template for X509 certificates
func x509Template() gmx509.Certificate {

	// generate a serial number
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, serialNumberLimit)

	// set expiry to around 10 years
	expiry := 3650 * 24 * time.Hour
	// backdate 5 min
	notBefore := time.Now().Add(-5 * time.Minute).UTC()

	//basic template to use
	x509 := gmx509.Certificate{
		SerialNumber:          serialNumber,
		NotBefore:             notBefore,
		NotAfter:              notBefore.Add(expiry).UTC(),
		BasicConstraintsValid: true,
	}
	return x509

}

// Additional for X509 subject
func subjectTemplateAdditional(country, province, locality, orgUnit, streetAddress, postalCode string) pkix.Name {
	name := subjectTemplate()
	if len(country) >= 1 {
		name.Country = []string{country}
	}
	if len(province) >= 1 {
		name.Province = []string{province}
	}

	if len(locality) >= 1 {
		name.Locality = []string{locality}
	}
	if len(orgUnit) >= 1 {
		name.OrganizationalUnit = []string{orgUnit}
	}
	if len(streetAddress) >= 1 {
		name.StreetAddress = []string{streetAddress}
	}
	if len(postalCode) >= 1 {
		name.PostalCode = []string{postalCode}
	}
	return name
}

// default template for X509 subject
func subjectTemplate() pkix.Name {
	return pkix.Name{
		Country:  []string{"US"},
		Locality: []string{"San Francisco"},
		Province: []string{"California"},
	}
}

type CA struct {
	Name               string
	Country            string
	Province           string
	Locality           string
	OrganizationalUnit string
	StreetAddress      string
	PostalCode         string
	Signer             crypto.Signer
	SignCert           *gmx509.Certificate
}

func (ca *CA) SignCertificate(name string, ous, sans []string, pub interface{},
	ku gmx509.KeyUsage, eku []gmx509.ExtKeyUsage) ([]byte, error) {

	template := x509Template()
	template.KeyUsage = ku
	template.ExtKeyUsage = eku

	//set the organization for the subject
	subject := subjectTemplateAdditional(ca.Country, ca.Province, ca.Locality, ca.OrganizationalUnit, ca.StreetAddress, ca.PostalCode)
	subject.CommonName = name

	subject.OrganizationalUnit = append(subject.OrganizationalUnit, ous...)

	template.Subject = subject
	for _, san := range sans {
		// try to parse as an IP address first
		ip := net.ParseIP(san)
		if ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, san)
		}
	}

	//create the x509 public cert
	certBytes, err := gmx509.CreateCertificate(rand.Reader, &template, ca.SignCert, pub, ca.Signer)
	if err != nil {
		return nil, err
	}

	return certBytes, nil
}

func GetNodeCACert(commonName string,
	casignerCert *gmx509.Certificate, casigner crypto.Signer) (nodeMspCaKeyStr string, nodeMspCaCertStr string, err error) {

	signCA := CA{
		Country:  Country,
		Province: Province,
		Locality: Locality,
		Signer:   casigner,
		SignCert: casignerCert,
	}
	nodepriv, signer, err := GeneratePrivKey()
	if err != nil {
		return "", "", err
	}

	//nodeecPubKey, err := GetGMPublicKey(nodepriv)
	//nodeecPubKey, err := GetECPublicKey(nodepriv)
	nodeecPubKey := signer.Public()
	if err != nil {
		return "", "", err
	}

	nodeMspCaKeyStr, err = GetPrivateKey(nodepriv)
	//nodeMspCaKeyStr, err = LoadPrivateKey(keystoryPath)
	if err != nil {
		return "", "", err
	}

	nodeMspCaCert, err := signCA.SignCertificate(commonName, nil, nil, nodeecPubKey, gmx509.KeyUsageDigitalSignature, []gmx509.ExtKeyUsage{})
	if err != nil {
		return "", "", err
	}

	block := &pem.Block{Type: "CERTIFICATE", Bytes: nodeMspCaCert}
	pemByte := pem.EncodeToMemory(block)
	nodeMspCaCertStr = string(pemByte)

	return nodeMspCaKeyStr, nodeMspCaCertStr, err
}

func GetNodeTlsCACert(commonName string,
	casignerCert *gmx509.Certificate, casigner crypto.Signer) (nodeMspCaKeyStr string, nodeMspCaCertStr string, err error) {

	TlsCA := CA{
		Country:  Country,
		Province: Province,
		Locality: Locality,
		Signer:   casigner,
		SignCert: casignerCert,
	}
	tlspriv, tlssigner, err := GenerateTLSPrivKey()
	if err != nil {
		return "", "", err
	}
	tlsecPubKey := tlssigner.Public()
	//tlsecPubKey, err := GetECPublicKey(tlspriv)
	if err != nil {
		return "", "", err
	}

	nodeMspCaKeyStr, err = GetPrivateKey(tlspriv)
	//nodeMspCaKeyStr, err = LoadPrivateKey(keystoryPath)
	if err != nil {
		return "", "", err
	}
	sans := []string{commonName}
	nodeMspCaCert, err := TlsCA.SignCertificate(commonName, nil, sans, tlsecPubKey,
		gmx509.KeyUsageDigitalSignature|gmx509.KeyUsageKeyEncipherment, []gmx509.ExtKeyUsage{gmx509.ExtKeyUsageServerAuth, gmx509.ExtKeyUsageClientAuth})
	if err != nil {
		return "", "", err
	}

	block := &pem.Block{Type: "CERTIFICATE", Bytes: nodeMspCaCert}
	pemByte := pem.EncodeToMemory(block)
	nodeMspCaCertStr = string(pemByte)

	return nodeMspCaKeyStr, nodeMspCaCertStr, err
}

