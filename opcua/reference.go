package opcua

type Reference struct {
	NodeClass
	NodeId
}

func (r *Reference) GetId() NodeId {
	return r.NodeId
}

func (r *Reference) GetClass() NodeClass {
	return r.NodeClass
}
