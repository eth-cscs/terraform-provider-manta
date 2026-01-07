package manta

type CloudConfigFile struct {
	Content  []byte `json:"content"`
	Name     string `json:"filename"`
	Encoding string `json:"encoding,omitempty"`
}

type GroupData struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Data        map[string]interface{} `json:"meta-data,omitempty"`
	File        CloudConfigFile        `json:"file,omitempty"`
	Versions    map[string]string      `json:"versions,omitempty"`
}

func (self *GroupData) Cmp(other *GroupData) bool {
	if cmpArray(self.File.Content, other.File.Content) {
		return true
	}

	return false
}

func (w *Wrapper) CreateGroupData(newGroupData GroupData) error {
	return requestPost(
		"https://foobar.openchami.cluster:8443/cloud-init/admin/groups",
		w.GetAccessToken(),
		newGroupData,
	)
}

func (w *Wrapper) DeleteGroupData(groupDataName string) error {
	_, err := requestDelete[any](
		"https://foobar.openchami.cluster:8443/cloud-init/admin/groups/"+groupDataName,
		w.GetAccessToken(),
	)

	if err.Error() == "unexpected end of JSON input" {
		return nil
	}

	return err
}

func (w *Wrapper) GetGroupData(group_name string) (GroupData, error) {
	return requestGet[GroupData](
		"https://foobar.openchami.cluster:8443/cloud-init/admin/groups/"+group_name,
		w.GetAccessToken(),
	)
}
