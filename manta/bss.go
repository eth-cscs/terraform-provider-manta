package manta

func (w *Wrapper) AddBss(params BssParams) error {
	return openchamiRequestPost(
		w.Base_url+"/bss/boot-parameters",
		w.GetAccessToken(),
		params,
	)

}

func (w *Wrapper) UpdateBss(params BssParams) error {
	return openchamiRequestPatch(
		w.Base_url+"/bss/boot-parameters",
		w.GetAccessToken(),
		params,
	)
}

func (w *Wrapper) GetBssParams() ([]BssParams, error) {
	return openchamiRequestGet[[]BssParams](
		w.Base_url+"/bss/boot-parameters",
		w.GetAccessToken(),
	)
}

func (w *Wrapper) DeleteBss(params BssParams) error {
	return openchamiRequestDeleteBody(
		w.Base_url+"/bss/boot-parameters",
		w.GetAccessToken(),
		params,
	)
}
