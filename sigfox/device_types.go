package sigfox

import (
	"context"
	"fmt"
	"net/http"
)

type DeviceTypeService service

type DeviceType struct {
	ID                 string       `json:"id,omitempty"`
	AutomaticRenewal   bool         `json:"automaticRenewal,omitempty"`
	Name               string       `json:"name,omitempty"`
	Description        string       `json:"description,omitempty"`
	KeepAlive          int64        `json:"keepAlive,omitempty"`
	PayloadType        int32        `json:"payloadType,omitempty"`
	AlertEmail         string       `json:"alertEmail,omitempty"`
	DownlinkMode       int32        `json:"downlinkMode,omitempty"`
	DownlinkDataString string       `json:"downlinkDataString,omitemptys"`
	Group              Group        `json:"group,omitempty"`
	Contract           ContractInfo `json:"contract,omitempty"`
	CreationTime       int64        `json:"creationTime,omitempty"`
	CreatedBy          string       `json:"createdBy,omitempty"`
	LastEditionTime    int64        `json:"lastEditionTime,omitempty"`
	LastEditedBy       string       `json:"lastEditedBy,omitempty"`
}

type ContractInfo struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ListDeviceTypesOptions struct {
	Name string `url:"name,omitempty"`
}

type ListDeviceTypesOutput struct {
	Data   []DeviceType `json:"data"`
	Paging Pagination   `json:"paging"`
}

func (s *DeviceTypeService) List(ctx context.Context, opt *ListDeviceTypesOptions) (*ListDeviceTypesOutput, *http.Response, error) {
	spath := fmt.Sprintf("/device-types")
	spath, err := addOptions(spath, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, nil, err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return nil, res, err
	}

	var out ListDeviceTypesOutput
	if err := decodeBody(res, &out); err != nil {
		return nil, res, err
	}
	return &out, res, nil
}

type CreateDeviceTypeInput struct {
	Name               string `json:"name,omitempty"`
	KeepAlive          int64  `json:"keepAlive,omitempty"`
	AlertEmail         string `json:"alertEmail,omitempty"`
	PayloadType        int32  `json:"payloadType,omitempty"`
	PayloadConfig      string `json:"payloadConfig,omitempty"`
	DownlinkMode       int32  `json:"downlinkMode,omitempty"`
	DownlinkDataString string `json:"downlinkDataString,omitempty"`
	Description        string `json:"description,omitempty"`
	GroupID            string `json:"groupId,omitempty"`
	ContractID         string `json:"contractId,omitempty"`
}

type CreateDeviceTypeOutput struct {
	ID string `json:"id,omitempty"`
}

func (s *DeviceTypeService) Create(ctx context.Context, input *CreateDeviceTypeInput) (*CreateDeviceTypeOutput, *http.Response, error) {
	req, err := s.client.newRequest(ctx, "POST", "/device-types", input)
	if err != nil {
		return nil, nil, err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return nil, res, err
	}

	var out CreateDeviceTypeOutput
	if err = decodeBody(res, &out); err != nil {
		return nil, res, err
	}

	return &out, res, nil
}

func (s *DeviceTypeService) Delete(ctx context.Context, deviceTypeID string) (*http.Response, error) {
	spath := fmt.Sprintf("/device-types/%s", deviceTypeID)
	req, err := s.client.newRequest(ctx, "DELETE", spath, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return res, err
	}

	return res, nil
}

type ListCallbackErrorsOptions struct {
	Since  int64 `url:"since,omitempty"`
	Before int64 `url:"before,omitempty"`
	Limit  int32 `url:"limit,omitempty"`
	Offset int32 `url:"offset,omitempty"`
}

type ListCallbackErrorsOutput struct {
	Data   []Hosts    `json:"data"`
	Paging Pagination `json:"paging"`
}

func (s *DeviceTypeService) ListCallbackErrors(ctx context.Context, deviceTypeID string, opt *ListCallbackErrorsOptions) (*ListCallbackErrorsOutput, *http.Response, error) {
	spath := fmt.Sprintf("/device-types/%s/callbacks-not-delivered", deviceTypeID)
	spath, err := addOptions(spath, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, nil, err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return nil, res, err
	}

	var out ListCallbackErrorsOutput
	if err := decodeBody(res, &out); err != nil {
		return nil, res, err
	}

	return &out, res, nil
}
