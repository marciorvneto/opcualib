package opcua

type Node interface {
	GetId() NodeId
	GetClass() NodeClass
}
