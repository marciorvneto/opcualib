package opcua

func (as *AddressSpace) NewReference(namespaceIndex int, identifierType IdentifierType, identifier string) error {

	r := &Reference{
		NODE_CLASS_REFERENCE,
		NodeId{
			namespaceIndex: namespaceIndex,
			identifierType: identifierType,
			identifier:     identifier,
		},
	}
	err := as.AddNode(r)
	return err
}
