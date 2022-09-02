package opcua

import "fmt"

type IdentifierType string

const (
	ID_NUMERIC IdentifierType = "i"
	ID_STRING                 = "s"
	ID_GUID                   = "g"
	ID_OPAQUE                 = "b"
)

type NodeId struct {
	namespaceIndex int
	identifierType IdentifierType
	identifier     string
}

func (nid NodeId) String(as *AddressSpace) string {
	namespaceURI := as.GetNamespaceURI(nid.namespaceIndex)
	return fmt.Sprintf("ns=%s;%s=%s", namespaceURI, nid.identifierType, nid.identifier)
}
