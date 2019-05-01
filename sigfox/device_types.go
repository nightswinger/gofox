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

// List retrieve a list of device types according to visibility permissions and request filters.
func (s *DeviceTypeService) List(opt *ListDeviceTypesOptions) (*ListDeviceTypesOutput, *http.Response, error) {
	return s.ListContext(context.Background(), opt)
}

// ListContext retrieve a list of device types according to visibility permissions and request filters with context.
func (s *DeviceTypeService) ListContext(ctx context.Context, opt *ListDeviceTypesOptions) (*ListDeviceTypesOutput, *http.Response, error) {
	spath := fmt.Sprintf("/device-types")
	spath, err := addOptions(spath, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, nil, err
	}

	res, err := s.client.Do(ctx, req)
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

// Create a new device type.
func (s *DeviceTypeService) Create(input *CreateDeviceTypeInput) (*CreateDeviceTypeOutput, *http.Response, error) {
	return s.CreateContext(context.Background(), input)
}

// CreateContext a new device type with context.
func (s *DeviceTypeService) CreateContext(ctx context.Context, input *CreateDeviceTypeInput) (*CreateDeviceTypeOutput, *http.Response, error) {
	req, err := s.client.newRequest(ctx, "POST", "/device-types", input)
	if err != nil {
		return nil, nil, err
	}

	res, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, res, err
	}

	var out CreateDeviceTypeOutput
	if err = decodeBody(res, &out); err != nil {
		return nil, res, err
	}

	return &out, res, nil
}

// Info retrieve information about a device type.
func (s *DeviceTypeService) Info(deviceTypeID string, params ...QueryParam) (*DeviceType, *http.Response, error) {
	return s.InfoContext(context.Background(), deviceTypeID, params...)
}

// InfoContext retrieve information abount a device type with context.
func (s *DeviceTypeService) InfoContext(ctx context.Context, deviceTypeID string, params ...QueryParam) (*DeviceType, *http.Response, error) {
	opt := &QueryParams{}
	for _, param := range params {
		param(opt)
	}

	spath := fmt.Sprintf("device-types/%s", deviceTypeID)
	spath, err := addOptions(spath, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, nil, err
	}

	res, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, res, err
	}

	var out DeviceType
	if err := decodeBody(res, &out); err != nil {
		return nil, res, err
	}

	return &out, res, nil
}

// Delete a device type.
func (s *DeviceTypeService) Delete(ctx context.Context, deviceTypeID string) (*http.Response, error) {
	spath := fmt.Sprintf("/device-types/%s", deviceTypeID)
	req, err := s.client.newRequest(ctx, "DELETE", spath, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(ctx, req)
	if err != nil {
		return res, err
	}

	return res, nil
}

type ListMessagesForDeviceTypeOutput struct {
	Data   []Message  `json:"data"`
	Paging Pagination `json:"paging"`
}

// ListMessages retrieve a list of messages for a given device types with a 3-day history.
func (s *DeviceTypeService) ListMessages(ctx context.Context, deviceTypeID string, params ...QueryParam) (*ListMessagesForDeviceTypeOutput, *http.Response, error) {
	spath := fmt.Sprintf("/device-types/%s/messages", deviceTypeID)

	opt := &QueryParams{}
	for _, param := range params {
		param(opt)
	}
	spath, err := addOptions(spath, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, nil, err
	}

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, res, err
	}

	var out ListMessagesForDeviceTypeOutput
	if err := decodeBody(res, &out); err != nil {
		return nil, res, err
	}

	return &out, res, nil
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

// ListCallbackErrors retrieve a list of undelivered callback messages for a given device types.
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

	res, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, res, err
	}

	var out ListCallbackErrorsOutput
	if err := decodeBody(res, &out); err != nil {
		return nil, res, err
	}

	return &out, res, nil
}

type ListCallbacksOutput struct {
	Data []Callbacks `json:"data"`
}

type Callbacks struct {
	ID              string            `json:"id"`
	Channel         string            `json:"channel"`
	CallbackType    int32             `json:"callbackType"`
	CallbackSubtype int32             `json:"callbackSubtype"`
	PayloadConfig   string            `json:"payloadConfig,omitempty"`
	Enabled         bool              `json:"enabled"`
	SendDuplicate   bool              `json:"sendDuplicate"`
	Dead            bool              `json:"dead,omitempty"`
	URL             string            `json:"url,omitempty"`
	HTTPMethod      string            `json:"httpMethod,omitempty"`
	DownlinkHook    bool              `json:"downlinkHook,omitempty"`
	Headers         map[string]string `json:"headers,omitempty"`
	SendSni         bool              `json:"sendSni,omitempty"`
	BodyTemplate    string            `json:"bodyTemplate,omitempty"`
	LinePattern     string            `json:"linePattern,omitempty"`
	Subject         string            `json:"subject,omitempty"`
	Recipient       string            `json:"recipient,omitempty"`
	Message         string            `json:"message,omitempty"`
}

// ListCallbacks retrieve a list of callbacks for a given device type according to visibility permissions and request filters.
func (s *DeviceTypeService) ListCallbacks(ctx context.Context, deviceTypeID string) (*ListCallbacksOutput, *http.Response, error) {
	spath := fmt.Sprintf("/device-types/%s/callbacks", deviceTypeID)

	req, err := s.client.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, nil, err
	}

	res, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, res, err
	}

	var out ListCallbacksOutput
	if err := decodeBody(res, &out); err != nil {
		return nil, res, err
	}

	return &out, res, nil
}

type CreateCallbackInput struct {
	Callbacks
	ContentType string `json:"contentType"`
}

type CreateCallbackOutput struct {
	ID string `json:"id,omitempty"`
}

// CreateCallback create a new callback for a given device type.
func (s *DeviceTypeService) CreateCallback(ctx context.Context, input *CreateCallbackInput) (*CreateCallbackOutput, *http.Response, error) {
	spath := fmt.Sprintf("/device-types/%s/callbacks", input.ID)
	req, err := s.client.newRequest(ctx, "POST", spath, input)
	if err != nil {
		return nil, nil, err
	}

	res, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, res, err
	}

	var out CreateCallbackOutput
	if err = decodeBody(res, &out); err != nil {
		return nil, res, err
	}

	return &out, res, nil
}

type UpdateCallbackInput struct {
	Callbacks
	ContentType string `json:"contentType"`
}

// UpdateCallback update a callback for a given device type.
func (s *DeviceTypeService) UpdateCallback(ctx context.Context, deviceTypeID, callbackID string, input *UpdateCallbackInput) (*http.Response, error) {
	spath := fmt.Sprintf("/device-types/%s/callbacks/%s", deviceTypeID, callbackID)
	req, err := s.client.newRequest(ctx, "PUT", spath, input)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(ctx, req)
	if err != nil {
		return res, err
	}

	return res, nil
}
