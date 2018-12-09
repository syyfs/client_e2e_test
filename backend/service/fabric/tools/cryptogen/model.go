package cryptogen

const (
	Country  = "CN"
	Province = "BeiJing"
	Locality = "BeiJing"
)

const (
	CAType    = "ca"
	TlsCAType = "tlsca"
)

type Node struct {
	Name    string
	TLSCert string
	TLSKey  string
	MSPCert string
	MSPKey  string
}

type User struct {
	Name    string
	TLSCert string
	TLSKey  string
	MSPCert string
	MSPKey  string
}

type Organization struct {
	Name      string // req -> domain orderer.citic.com ; peer.citit.com
	TLSCACert string
	TLSCAKey  string
	MSPCACert string
	MSPCAKey  string
}

type OrgSpec struct {
	//Name          string       `json:"Name"`
	CommonName string // 组织域名
	//EnableNodeOUs bool         `json:"EnableNodeOUs"` // 是否生成 节点下的msp的config.yaml配置文件
	CA NodeSpec `json:"CA"`
	//Template      NodeTemplate `json:"Template"`
	//Specs         []NodeSpec   `json:"Specs"`
	//Users UsersSpec
}

func NewOrgSpec() *OrgSpec {
	orgspec := &OrgSpec{}
	return orgspec
}

type NodeSpec struct {
	Hostname           string
	CommonName         string
	Country            string
	Province           string
	Locality           string
	OrganizationalUnit string
	StreetAddress      string
	PostalCode         string
	SANS               []string
}

type UsersSpec struct {
	Count int `json:"Count"`
}

