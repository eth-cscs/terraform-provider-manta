package manta

func (w *Wrapper) AddBss(params BssParams) error {
	return openchamiRequestPost(
		"https://foobar.openchami.cluster:8443/boot/v1/bootparameters",
		w.GetAccessToken(),
		params,
	)

}

func (w *Wrapper) UpdateBss(params BssParams) error {
	return openchamiRequestPatch(
		"https://foobar.openchami.cluster:8443/boot/v1/bootparameters",
		w.GetAccessToken(),
		params,
	)
}

func (w *Wrapper) GetBssParams() ([]BssParams, error) {
	return openchamiRequestGet[[]BssParams](
		"https://foobar.openchami.cluster:8443/boot/v1/bootparameters",
		w.GetAccessToken(),
	)
}

func (w *Wrapper) DeleteBss(params BssParams) error {
	return openchamiRequestDeleteBody(
		"https://foobar.openchami.cluster:8443/boot/v1/bootparameters",
		w.GetAccessToken(),
		params,
	)
}
