package cryptogen

type Generator interface {
	CreateOrganization(commonName string) (Organization, error)
	CreateOrganizationWithPath(commonName string) (Organization, error)
	CreateNode(commonName string, org Organization) (Node, error)
	CreateUser(userName string, org Organization) (User, error)
}
