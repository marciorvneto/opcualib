package opcua

import (
	"errors"
	"fmt"
)

type AddressSpace struct {
	namespaces NamespaceContext
	nodes      map[string]Node
}

func NewAddressSpace() AddressSpace {
	return AddressSpace{
		namespaces: NewNamespaceContext(),
		nodes:      make(map[string]Node),
	}
}

func (as *AddressSpace) AddNode(node Node) error {
	idString := node.GetId().String(as)
	_, hasValue := as.nodes[idString]
	if hasValue {
		return errors.New(fmt.Sprintf("There already exists a node with id=%s", idString))
	}
	as.nodes[idString] = node
	return nil
}

func (as *AddressSpace) GetNode(node Node) (Node, error) {
	idString := node.GetId().String(as)
	node, ok := as.nodes[idString]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Can't find node with id=%s\n", idString))
	}
	return node, nil
}
func (as *AddressSpace) AddNamespace(uri string) int {
	return as.namespaces.AddNamespace(uri)
}
func (as *AddressSpace) GetNamespaceURI(namespaceIndex int) string {
	uri, err := as.namespaces.GetNamespaceURI(namespaceIndex)
	if err != nil {
		panic(err)
	}
	return uri
}
