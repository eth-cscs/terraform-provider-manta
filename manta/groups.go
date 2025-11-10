package manta

type MemberAddBody struct {
	ID string `json:"id"`
}

type Members struct {
	IDs []string `json:"ids,omitempty"`
}

type Group struct {
	Label          string   `json:"label"`
	Description    string   `json:"description"`
	Tags           []string `json:"tags,omitempty"`
	ExclusiveGroup string   `json:"exclusiveGroup,omitempty"`
	Members        `json:"members,omitempty"`
}

func (w *Wrapper) AddMemberToGroup(xname, group_label string) error {
	return openchamiRequestPost(
		"https://foobar.openchami.cluster:8443/hsm/v2/groups/"+group_label+"/members",
		w.GetAccessToken(),
		MemberAddBody{ID: xname},
	)
}

func (w *Wrapper) DeleteMemberToGroup(xname, group_label string) error {
	_, err := openchamiRequestDelete[any](
		"https://foobar.openchami.cluster:8443/hsm/v2/groups/"+group_label+"/members/"+xname,
		w.GetAccessToken(),
	)
	return err
}

func (w *Wrapper) UpdateGroup(newGroup Group) error {
	_, err := openchamiRequestDelete[any](
		"https://foobar.openchami.cluster:8443/hsm/v2/groups/"+newGroup.Label,
		w.GetAccessToken(),
	)

	if err != nil {
		return err
	}

	return openchamiRequestPost(
		"https://foobar.openchami.cluster:8443/hsm/v2/groups",
		w.GetAccessToken(),
		newGroup,
	)
}

func (w *Wrapper) CreateGroup(newGroup Group) error {
	return openchamiRequestPost(
		"https://foobar.openchami.cluster:8443/hsm/v2/groups",
		w.GetAccessToken(),
		newGroup,
	)
}

func (w *Wrapper) DeleteGroup(group_label string) error {
	_, err := openchamiRequestDelete[any](
		"https://foobar.openchami.cluster:8443/hsm/v2/groups/"+group_label,
		w.GetAccessToken(),
	)
	return err
}
