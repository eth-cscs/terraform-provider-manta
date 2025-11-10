package manta

func (w *Wrapper) GetNodeId(id string) (NodeItem, error) {
	node, err := openchamiRequestGet[NodeItem](
		"https://foobar.openchami.cluster:8443/hsm/v2/State/Components/"+id,
		w.GetAccessToken(),
	)

	if err != nil {
		return NodeItem{}, err
	}

	node.State, err = w.GetPowerStatusNodeId(id)

	if err != nil {
		return NodeItem{}, err
	}

	return node, err
}
