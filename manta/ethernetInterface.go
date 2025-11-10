package manta

import (
	"errors"
	"strings"
)

func (w *Wrapper) DeleteEthernetInterface(mac string) error {
	ethInterfaceID := strings.Replace(mac, ":", "", -1)

	ochamiError, err := openchamiRequestDelete[OchamiError](
		"https://foobar.openchami.cluster:8443/hsm/v2/Inventory/EthernetInterfaces/"+ethInterfaceID,
		w.GetAccessToken(),
	)

	if ochamiError.Status == 404 {
		return errors.New(ochamiError.Detail)
	}

	return err
}

func (w *Wrapper) GetEthernetInterface(mac string) (NodeInterface, OchamiError) {
	ethInterfaceID := strings.Replace(mac, ":", "", -1)

	ni, _, oe := openchamiRequestOchami[NodeInterface](
		"https://foobar.openchami.cluster:8443/hsm/v2/Inventory/EthernetInterfaces/"+ethInterfaceID,
		w.GetAccessToken(),
		"GET",
	)

	return ni, oe
}

func (w *Wrapper) AddEthernetInterface(nodeInterface NodeInterface) error {
	_, ochamiError := w.GetEthernetInterface(nodeInterface.MAC)

	if ochamiError.Status != 404 {
		return w.PatchEthernetInterface(
			NodeInterfacePatch{
				MAC: nodeInterface.MAC,
				IPs: nodeInterface.IPs,
			},
		)
	}

	return openchamiRequestPost(
		"https://foobar.openchami.cluster:8443/hsm/v2/Inventory/EthernetInterfaces/",
		w.GetAccessToken(),
		nodeInterface,
	)
}

func (w *Wrapper) PatchEthernetInterface(nodeInterface NodeInterfacePatch) error {
	ethInterfaceID := strings.Replace(nodeInterface.MAC, ":", "", -1)

	return openchamiRequestPatch(
		"https://foobar.openchami.cluster:8443/hsm/v2/Inventory/EthernetInterfaces/"+ethInterfaceID,
		w.GetAccessToken(),
		nodeInterface,
	)
}
