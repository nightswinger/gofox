package sigfox

import (
	"context"
	"net/http"
)

type TileService service

type TileMonarchOutput struct {
	BaseImgURL     string `json:"baseImgUrl,omitempty"`
	TmsTemplateURL string `json:"tmsTemplateUrl,omitempty"`
	Bounds         Bounds `json:"bounds,omitempty"`
}

type Bounds struct {
	Sw LatLng `json:"sw,omitempty"`
	Ne LatLng `json:"ne,omitempty"`
}

type LatLng struct {
	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"lng,omitempty"`
}

// Monarch retrieve the information needed to display Sigfox Monarch service coverage.
func (s *TileService) Monarch() (*TileMonarchOutput, *http.Response, error) {
	return s.MonarchContext(context.Background())
}

// MonarchContext retrieve the information needed to display Sigfox Monarch service coverage with context.
func (s *TileService) MonarchContext(ctx context.Context) (*TileMonarchOutput, *http.Response, error) {
	req, err := s.client.newRequest(ctx, "GET", "/tiles/monarch", nil)
	if err != nil {
		return nil, nil, err
	}

	res, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, res, err
	}

	var out TileMonarchOutput
	if err := decodeBody(res, &out); err != nil {
		return nil, res, err
	}

	return &out, res, nil
}
