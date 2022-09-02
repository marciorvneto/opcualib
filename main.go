package main

import (
	"fmt"
	"marciorvneto/opclib/opcua"

	"github.com/google/uuid"
)

func main() {

	addrspc := opcua.NewAddressSpace()
	defaultIndex := addrspc.AddNamespace("default")

	uuid, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}

	addrspc.NewReference(defaultIndex, opcua.ID_STRING, "Reference")
	addrspc.NewReference(defaultIndex, opcua.ID_STRING, "Reference2")
	addrspc.NewReference(defaultIndex, opcua.ID_GUID, uuid.String())

	fmt.Println(defaultIndex)
	fmt.Println(addrspc)

}
