package sigfox

import (
	"context"
	"fmt"
)

type DevicesService service

type Device struct {
	ID                  string `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	Prototype           bool   `json:"prototype,omitempty"`
	PAC                 string `json:"pac,omitempty"`
	SequenceNumber      int32  `json:"sequenceNumber,omitempty"`
	TrashSequenceNumber int32  `json:"trashSequenceNumber,omitempty"`
	LastCom             int    `json:"lastCom,omitempty"`
	Lqi                 int32  `json:"lqi,omitempty"`
	AverageSnr          string `json:"averageSnr,omitempty"`
	AverageRssi         string `json:"averageRssi,omitempty"`
	ActivationTime      int64  `json:"activationTime,omitempty"`
	CreationTime        int64  `json:"creationTime,omitempty"`
	State               int32  `json:"state,omitempty"`
	ComState            int32  `json:"comState,omitempty"`
}

type DeviceListOptions struct {
	ID           string   `url:"id"`
	GroupIds     []string `url:"groupIds"`
	Deep         bool     `url:"deep"`
	DeviceTypeID string   `url:"deviceTypeId"`
	OperatorID   string   `url:"operatorId"`
	Sort         string   `url:"sort"`
	MinID        string   `url:"minId"`
	MaxID        string   `url:"maxId"`
	Fields       []string `url:"fields"`
	Limit        int32    `url:"limit"`
	Offset       int32    `url:"offset"`
	PageId       string   `url:"pageId"`
}

type ListDevices struct {
	Data   []Device   `json:"data"`
	Paging Pagination `json:"paging"`
}

func (s *DevicesService) List(ctx context.Context, opt *DeviceListOptions) (*ListDevices, error) {
	spath := fmt.Sprintf("/devices")
	spath, err := addOptions(spath, opt)
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

	var listDevices ListDevices
	if err := decodeBody(res, &listDevices); err != nil {
		return nil, err
	}

	return &listDevices, nil
}

type CreateDeviceBody struct {
	ID                    string  `json:"id"`
	Name                  string  `json:"name"`
	DeviceTypeID          string  `json:"deviceTypeId"`
	PAC                   string  `json:"pac"`
	Lat                   float64 `json:"lat,omitempty"`
	Lng                   float64 `json:"lng,omitempty"`
	ProductCertificateKey string  `json:"productCertificateKey,omitempty"`
	Prototype             bool    `json:"prototype,omitempty"`
	AutomaticRenewal      bool    `json:"automaticRenewal,omitempty"`
	Activable             bool    `json:"activable,omitempty"`
}

type CreateDeviceOutput struct {
	ID string `json:"id"`
}

func (s *DevicesService) Create(ctx context.Context, body *CreateDeviceBody) (*CreateDeviceOutput, error) {
	req, err := s.client.newRequest(ctx, "POST", "/devices", body)
	if err != nil {
		return nil, err
	}

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var output CreateDeviceOutput
	if err := decodeBody(res, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (s *DevicesService) GetInfo(ctx context.Context, deviceID string) (*Device, error) {
	spath := fmt.Sprintf("/devices/%s", deviceID)
	req, err := s.client.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var device Device
	if err := decodeBody(res, &device); err != nil {
		return nil, err
	}

	return &device, nil
}

type UpdateDeviceBody struct {
	Name                  string  `json:"name,omitempty"`
	Lat                   float64 `json:"lat,omitempty"`
	Lng                   float64 `json:"lng,omitempty"`
	ProductCertificateKey string  `json:"productCertificateKey,omitempty"`
	Prototype             bool    `json:"prototype,omitempty"`
	AutomaticRenewal      bool    `json:"automaticRenewal,omitempty"`
	Activable             bool    `json:"activable,omitempty"`
}

func (s *DevicesService) UpdateDevice(ctx context.Context, deviceID string, body *UpdateDeviceBody) error {
	spath := fmt.Sprintf("/devices/%s", deviceID)

	req, err := s.client.newRequest(ctx, "PUT", spath, body)
	if err != nil {
		return err
	}

	_, err = s.client.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

type UndeliveredCallbacksOptions struct {
	Since  int64 `url:"since,omitempty"`
	Before int64 `url:"before,omitempty"`
	Limit  int32 `url:"limit,omitempty"`
	Offset int32 `url:"offset,omitempty"`
}

type UndeliveredCallbacks struct {
	Data   []*Hosts   `json:"data"`
	Paging Pagination `json:"paging"`
}

type Hosts struct {
	Device     string          `json:"device"`
	DeviceURL  string          `json:"deviceUrl"`
	DeviceType string          `json:"deviceType"`
	Time       int64           `json:"time"`
	Data       string          `json:"data"`
	Snr        string          `json:"snr"`
	Status     string          `json:"status"`
	Message    string          `json:"message"`
	Callback   *CallbackMedium `json:"callbackMedium"`
	//Parameters
}

type CallbackMedium struct {
	Subject string `json:"subject"`
	Message string `json:"message"`
	URL     string `json:"url"`
	//Headers
	Body        string `json:"body"`
	ContentType string `json:"contentType"`
	Method      string `json:"method"`
	Error       string `json:"error"`
}

func (s *DevicesService) ListUndeliveredCallbacks(ctx context.Context, deviceID string, opt *UndeliveredCallbacksOptions) (*UndeliveredCallbacks, error) {
	spath := fmt.Sprintf("/devices/%s/callbacks-not-delivered", deviceID)
	spath, err := addOptions(spath, opt)
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

	var listUndelivered UndeliveredCallbacks
	if err := decodeBody(res, &listUndelivered); err != nil {
		return nil, err
	}

	return &listUndelivered, nil
}

func (s *DevicesService) DisengageSequenceNumber(ctx context.Context, deviceID string) error {
	spath := fmt.Sprintf("/devices/%s/disengage", deviceID)

	req, err := s.client.newRequest(ctx, "POST", spath, nil)
	if err != nil {
		return err
	}
	_, err = s.client.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

type DeviceMessagesOptions struct {
	Fields string `url:"fields,omitempty"`
	Since  int64  `url:"since,omitempty"`
	Before int64  `url:"before,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Offset int32  `url:"offset,omitempty"`
}

type DeviceMessages struct {
	Data   []Message  `json:"data,omitempty"`
	Paging Pagination `json:"paging,omitempty"`
}

type Message struct {
	Device       Device `json:"device,omitempty"`
	Time         int64  `json:"time,omitempty"`
	Data         string `json:"data,omitempty"`
	AckRequired  bool   `json:"ackRequired,omitempty"`
	Lqi          int32  `json:"lqi,omitempty"`
	LqiRepeaters int32  `json:"lqiRepeaters,omitempty"`
	SeqNumber    int32  `json:"seqNumber,omitempty"`
	NbFrames     int32  `json:"nbFrames,omitempty"`
	//ComputedLocation
	Rinfos []Rinfo `json:"rinfos,omitempty"`
	//DownlinkAnserStatus
}

type Rinfo struct {
	BaseStation   MinBaseStation `json:"baseStation,omitempty"`
	Rssi          string         `json:"rssi,omitempty"`
	RssiRepeaters string         `json:"rssiRepeaters"`
	Lat           string         `json:"lat,omitempty"`
	Lng           string         `json:"lng,omitempty"`
	Snr           string         `json:"snr,omitempty"`
	SnrRepeaters  string         `json:"snrRepeaters,omitempty"`
	Freq          float64        `json:"freq,omitempty"`
	FreqRepeaters float64        `json:"freqReaters"`
	Rep           int32          `json:"rep,omitempty"`
	CbStatus      []CbStatus     `json:"cbStatus,omitempty"`
}

type MinBaseStation struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CbStatus struct {
	Status int32  `json:"status,omitempty"`
	Info   string `json:"string,omitempty"`
	CbDef  string `json:"cbDef,omitempty"`
	Time   int64  `json:"time,omitempty"`
}

func (s *DevicesService) GetMessages(ctx context.Context, deviceID string, opt *DeviceMessagesOptions) (*DeviceMessages, error) {
	spath := fmt.Sprintf("/devices/%s/messages", deviceID)
	spath, err := addOptions(spath, opt)
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

	var messages DeviceMessages
	if err := decodeBody(res, &messages); err != nil {
		return nil, err
	}

	return &messages, nil
}

type DeviceMetric struct {
	LastDay   int32 `json:"lastDay"`
	LastWeek  int32 `json:"lastWeek"`
	LastMonth int32 `json:"lastMonth"`
}

func (s *DevicesService) GetMetric(ctx context.Context, deviceID string) (*DeviceMetric, error) {
	spath := fmt.Sprintf("/devices/%s/messages/metric", deviceID)

	req, err := s.client.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}
	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var deviceMetric DeviceMetric
	if err := decodeBody(res, &deviceMetric); err != nil {
		return nil, err
	}

	return &deviceMetric, nil
}

type CreateMultipleDevicesBody struct {
	DeviceTypeID          string                `json:"deviceTypeId"`
	Prefix                string                `json:"prefix,omitempty"`
	ProductCertificateKey string                `json:"productCertificateKey,omitempty"`
	Prototype             string                `json:"prototype,omitempty"`
	Activable             bool                  `json:"activable,omitempty"`
	Devices               []*DeviceCreationBulk `json:"devices"`
}

type DeviceCreationBulk struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	PAC              string  `json:"pac"`
	Lat              float64 `json:"lat,omitempty"`
	Lng              float64 `json:"lng,omitempty"`
	AutomaticRenewal bool    `json:"automaticRenewal,omitempty"`
}

type CreateMultipleDevicesOutput struct {
	Total int32  `json:"total,omitempty"`
	JobID string `json:"jobId,omitempty"`
}

func (s *DevicesService) CreateMultipleWithAsync(ctx context.Context, body *CreateMultipleDevicesBody) (*CreateMultipleDevicesOutput, error) {
	req, err := s.client.newRequest(ctx, "POST", "/devices/bulk", body)
	if err != nil {
		return nil, err
	}

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var output CreateMultipleDevicesOutput
	if err := decodeBody(res, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
