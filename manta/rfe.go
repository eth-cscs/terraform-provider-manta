package manta

import (
	"errors"
	"time"
)

func (w *Wrapper) AddRfe(rfeItem RfeItem) (RfeItem, error) {
	var rfes RedfishEndpointArray
	rfes.RedfishEndpoints = append(rfes.RedfishEndpoints, rfeItem)

	ochamiError := openchamiRequestPostOchami(
		w.Base_url+"/redfish",
		w.GetAccessToken(),
		rfes,
	)

	if ochamiError.Status != 0 {
		return RfeItem{}, errors.New("RFE not added")
	}

	if ochamiError.Status == 400 {
		return RfeItem{}, errors.New(ochamiError.Detail)
	}

	rfeReturn, _ := w.GetRfeId(rfeItem.ID)

	if rfeReturn.Enabled {
		for rfeReturn.DiscoveryInfo.LastStatus != "DiscoverOK" {
			if rfeReturn.DiscoveryInfo.LastStatus == "ChildVerificationFailed" {
				break
			}
			if rfeReturn.DiscoveryInfo.LastStatus == "HTTPsGetFailed" {
				break
			}
			time.Sleep(1 * time.Second)
			rfeReturn, _ = w.GetRfeId(rfeReturn.ID)
		}
	}

	return rfeReturn, nil
}

func (w *Wrapper) DeleteRfe(rfeID string) error {
	_, err := openchamiRequestDelete[any](
		w.Base_url+"/redfish/"+rfeID,
		w.GetAccessToken(),
	)
	return err
}

func (w *Wrapper) GetRfe() ([]RfeItem, error) {
	getrfe, err := openchamiRequestDelete[RedfishEndpointArray](
		w.Base_url+"/redfish",
		w.GetAccessToken(),
	)

	if err != nil {
		return make([]RfeItem, 0), err

	}

	return getrfe.RedfishEndpoints, nil
}

func (w *Wrapper) GetRfeId(rfeID string) (RfeItem, error) {
	rfe, err := openchamiRequestGet[RfeItem](
		w.Base_url+"/redfish/"+rfeID,
		w.GetAccessToken(),
	)

	if err != nil {
		return RfeItem{}, nil

	}

	return rfe, nil
}
