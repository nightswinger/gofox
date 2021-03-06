package sigfox

import (
	"context"
	"fmt"
)

type DeviceService service

type Device struct {
	ID                  string `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	Prototype           bool   `json:"prototype,omitempty"`
	PAC                 string `json:"pac,omitempty"`
	SequenceNumber      int32  `json:"sequenceNumber,omitempty"`
	TrashSequenceNumber int32  `json:"trashSequenceNumber,omitempty"`
	LastCom             int64  `json:"lastCom,omitempty"`
	Lqi                 int32  `json:"lqi,omitempty"`
	AverageSnr          string `json:"averageSnr,omitempty"`
	AverageRssi         string `json:"averageRssi,omitempty"`
	ActivationTime      int64  `json:"activationTime,omitempty"`
	CreationTime        int64  `json:"creationTime,omitempty"`
	State               int32  `json:"state,omitempty"`
	ComState            int32  `json:"comState,omitempty"`
	//Token
	UnsubscriptionTime     int64  `json:"unsubscriptionTime,omitempty"`
	CreatedBy              string `json:"createdBy,omitempty"`
	LastEditionTime        int64  `json:"lastEditionTime,omitempty"`
	AutomaticRenewal       bool   `json:"automaticRenewal,omitempty"`
	AutomaticRenewalStatus int32  `json:"automaticRenewalStatus,omitempty"`
	Activable              bool   `json:"activable,omitempty"`
}

type DeviceListOptions struct {
	ID           string   `url:"id,omitempty"`
	GroupIds     []string `url:"groupIds,omitempty"`
	Deep         bool     `url:"deep,omitempty"`
	DeviceTypeID string   `url:"deviceTypeId,omitempty"`
	OperatorID   string   `url:"operatorId,omitempty"`
	Sort         string   `url:"sort,omitempty"`
	MinID        string   `url:"minId,omitempty"`
	MaxID        string   `url:"maxId,omitempty"`
	Fields       []string `url:"fields,omitempty"`
	Limit        int32    `url:"limit,omitempty"`
	Offset       int32    `url:"offset,omitempty"`
	PageID       string   `url:"pageId,omitempty"`
}

type ListDevices struct {
	Data   []Device   `json:"data"`
	Paging Pagination `json:"paging"`
}

func (s *DeviceService) List(opt *DeviceListOptions) (*ListDevices, error) {
	return s.ListContext(context.Background(), opt)
}

func (s *DeviceService) ListContext(ctx context.Context, opt *DeviceListOptions) (*ListDevices, error) {
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

func (s *DeviceService) Create(body *CreateDeviceBody) (*CreateDeviceOutput, error) {
	return s.CreateContext(context.Background(), body)
}

func (s *DeviceService) CreateContext(ctx context.Context, body *CreateDeviceBody) (*CreateDeviceOutput, error) {
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

func (s *DeviceService) Info(deviceID string) (*Device, error) {
	return s.InfoContext(context.Background(), deviceID)
}

func (s *DeviceService) InfoContext(ctx context.Context, deviceID string) (*Device, error) {
	spath := fmt.Sprintf("/devices/%s", deviceID)
	req, err := s.client.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(ctx, req)
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

func (s *DeviceService) Update(deviceID string, body *UpdateDeviceBody) error {
	return s.UpdateContext(context.Background(), deviceID, body)
}

func (s *DeviceService) UpdateContext(ctx context.Context, deviceID string, body *UpdateDeviceBody) error {
	spath := fmt.Sprintf("/devices/%s", deviceID)

	req, err := s.client.newRequest(ctx, "PUT", spath, body)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req)
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

func (s *DeviceService) ListUndeliveredCallbacks(deviceID string, opt *UndeliveredCallbacksOptions) (*UndeliveredCallbacks, error) {
	return s.ListUndeliveredCallbacksContext(context.Background(), deviceID, opt)
}

func (s *DeviceService) ListUndeliveredCallbacksContext(ctx context.Context, deviceID string, opt *UndeliveredCallbacksOptions) (*UndeliveredCallbacks, error) {
	spath := fmt.Sprintf("/devices/%s/callbacks-not-delivered", deviceID)
	spath, err := addOptions(spath, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	var listUndelivered UndeliveredCallbacks
	if err := decodeBody(res, &listUndelivered); err != nil {
		return nil, err
	}

	return &listUndelivered, nil
}

func (s *DeviceService) DisengageSequenceNumber(deviceID string) error {
	return s.DisengageSequenceNumberContext(context.Background(), deviceID)
}

func (s *DeviceService) DisengageSequenceNumberContext(ctx context.Context, deviceID string) error {
	spath := fmt.Sprintf("/devices/%s/disengage", deviceID)

	req, err := s.client.newRequest(ctx, "POST", spath, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(ctx, req)
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

func (s *DeviceService) Messages(deviceID string, params ...QueryParam) (*DeviceMessages, error) {
	return s.MessagesContext(context.Background(), deviceID, params...)
}

func (s *DeviceService) MessagesContext(ctx context.Context, deviceID string, params ...QueryParam) (*DeviceMessages, error) {
	spath := fmt.Sprintf("/devices/%s/messages", deviceID)

	opt := &QueryParams{}
	for _, param := range params {
		param(opt)
	}
	spath, err := addOptions(spath, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(ctx, req)
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

// Perhaps metric is updated at 1:00AM UTC
func (s *DeviceService) Metric(deviceID string) (*DeviceMetric, error) {
	return s.MetricContext(context.Background(), deviceID)
}

func (s *DeviceService) MetricContext(ctx context.Context, deviceID string) (*DeviceMetric, error) {
	spath := fmt.Sprintf("/devices/%s/messages/metric", deviceID)

	req, err := s.client.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}
	res, err := s.client.Do(ctx, req)
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

func (s *DeviceService) CreateMultipleWithAsync(body *CreateMultipleDevicesBody) (*CreateMultipleDevicesOutput, error) {
	return s.CreateMultipleWithAsyncContext(context.Background(), body)
}

func (s *DeviceService) CreateMultipleWithAsyncContext(ctx context.Context, body *CreateMultipleDevicesBody) (*CreateMultipleDevicesOutput, error) {
	req, err := s.client.newRequest(ctx, "POST", "/devices/bulk", body)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	var output CreateMultipleDevicesOutput
	if err := decodeBody(res, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
