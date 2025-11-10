package manta

type ClusterDefaults struct {
	CloudProvider    string   `json:"cloud_provider,omitempty"`
	Region           string   `json:"region,omitempty"`
	AvailabilityZone string   `json:"availability-zone,omitempty"`
	ClusterName      string   `json:"cluster-name,omitempty"`
	PublicKeys       []string `json:"public-keys,omitempty"`
	BaseUrl          string   `json:"base-url,omitempty"`
	BootSubnet       string   `json:"boot-subnet,omitempty"`
	WGSubnet         string   `json:"wg-subnet,omitempty"`
	ShortName        string   `json:"short-name,omitempty"`
	NidLength        int      `json:"nid-length,omitempty"`
}

func (self *ClusterDefaults) Cmp(other *ClusterDefaults) bool {
	if cmpArray(self.PublicKeys, other.PublicKeys) {
		return true
	}
	if self.BaseUrl != other.BaseUrl {
		return true
	}

	return false
}

func (w *Wrapper) CreateClusterDefault(newClusterDefaults ClusterDefaults) error {
	return openchamiRequestPost(
		"https://foobar.openchami.cluster:8443/cloud-init/admin/cluster-defaults",
		w.GetAccessToken(),
		newClusterDefaults,
	)
}

func (w *Wrapper) GetClusterDefault() (ClusterDefaults, error) {
	return openchamiRequestGet[ClusterDefaults](
		"https://foobar.openchami.cluster:8443/cloud-init/admin/cluster-defaults",
		w.GetAccessToken(),
	)
}
