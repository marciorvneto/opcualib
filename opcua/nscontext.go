package opcua

import (
	"errors"
	"fmt"
)

type NamespaceContext struct {
	namespaces map[int]string
}

func NewNamespaceContext() NamespaceContext {
	return NamespaceContext{
		namespaces: make(map[int]string),
	}
}

func (c *NamespaceContext) AddNamespace(uri string) int {
	index := len(c.namespaces)
	c.namespaces[index] = uri
	return index
}

func (c *NamespaceContext) GetNamespaceURI(namespaceIndex int) (string, error) {
	uri, ok := c.namespaces[namespaceIndex]
	if !ok {
		return "", errors.New(fmt.Sprintf("Couldn't find a namespece with index %d", namespaceIndex))
	}
	return uri, nil
}
