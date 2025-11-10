package manta

import (
	"errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (w *Wrapper) GetPowerStatusNodeId(id string) (string, error) {
	pcs, err := openchamiRequestGet[PcsStatus](
		"https://foobar.openchami.cluster:8443/power-control/v1/power-status?xname="+id,
		w.GetAccessToken(),
	)

	if err != nil {
		return "", err
	}

	if pcs.Status[0].Error == `Component not found in component map.` {
		return "", errors.New(pcs.Status[0].Error)
	}

	var c = cases.Title(language.English)
	return c.String(pcs.Status[0].PowerState), err
}
