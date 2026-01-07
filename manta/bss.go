package manta

func (w *Wrapper) AddBss(params BssParams) error {
	return requestPost(
		w.Base_url+"/bss/boot-parameters",
		w.GetAccessToken(),
		params,
	)

}

func (w *Wrapper) UpdateBss(params BssParams) error {
	return requestPatch(
		w.Base_url+"/bss/boot-parameters",
		w.GetAccessToken(),
		params,
	)
}

func (w *Wrapper) GetBssParams() ([]BssParams, error) {
	return requestGet[[]BssParams](
		w.Base_url+"/bss/boot-parameters",
		w.GetAccessToken(),
	)
}

func (w *Wrapper) DeleteBss(params BssParams) error {
	return requestDeleteBody(
		w.Base_url+"/bss/boot-parameters",
		w.GetAccessToken(),
		params,
	)
}
