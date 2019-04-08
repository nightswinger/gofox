package sigfox

import (
	"context"
	"fmt"
)

type DeviceTypeService service

type DeviceType struct {
	ID                 string `json:"id,omitempty"`
	AutomaticRenewal   bool   `json:"automaticRenewal,omitempty"`
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
	KeepAlive          int64  `json:"keepAlive,omitempty"`
	PayloadType        int32  `json:"payloadType,omitempty"`
	AlertEmail         string `json:"alertEmail,omitempty"`
	DownlinkMode       int32  `json:"downlinkMode,omitempty"`
	DownlinkDataString string `json:"downlinkDataString,omitemptys"`
	Group              struct {
		ID string `json:"id,omitempty"`
	} `json:"group,omitempty"`
	Contract struct {
		ID string `json:"id,omitempty"`
	} `json:"contract,omitempty"`
	CreationTime    int64  `json:"creationTime,omitempty"`
	CreatedBy       string `json:"createdBy,omitempty"`
	LastEditionTime int64  `json:"lastEditionTime,omitempty"`
	LastEditedBy    string `json:"lastEditedBy,omitempty"`
}

type ListDeviceTypesOptions struct {
	Name string `url:"name,omitempty"`
}

type ListDeviceTypesOutput struct {
	Data   []DeviceType `json:"data"`
	Paging Pagination   `json:"paging"`
}

func (s *DeviceTypeService) List(ctx context.Context, opt *ListDeviceTypesOptions) (out *ListDeviceTypesOutput, err error) {
	spath := fmt.Sprintf("/device-types")
	spath, err = addOptions(spath, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if err := decodeBody(res, &out); err != nil {
		return nil, err
	}
	return out, nil
}