package sigfox

import (
	"context"
	"fmt"
	"net/http"
)

type CoverageService service

type CoveragePredictionInput struct {
	Lat     float64 `url:"lat"`
	Lng     float64 `url:"lng"`
	Radius  int     `url:"radius,omitempty"`
	GroupID string  `url:"groupId,omitempty"`
}

type CoveragePredictionOutput struct {
	LocationCovered bool  `json:"locationCovered,omitempty"`
	Margins         []int `json:"margins,omitempty"`
}

// Predictions retrieve coverage predictions for any location.
func (s *CoverageService) Predictions(input *CoveragePredictionInput) (*CoveragePredictionOutput, *http.Response, error) {
	return s.PredictionsContext(context.Background(), input)
}

// PredictionsContext retrieve coverage predictions for any location with context.
func (s *CoverageService) PredictionsContext(ctx context.Context, input *CoveragePredictionInput) (*CoveragePredictionOutput, *http.Response, error) {
	spath := fmt.Sprintf("/coverages/global/predictions")
	spath, err := addOptions(spath, input)
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

	var out CoveragePredictionOutput
	if err := decodeBody(res, &out); err != nil {
		return nil, res, err
	}

	return &out, res, nil
}

type CoverageBatchPredictionInput struct {
	Locations []Location `json:"locations"`
	Radius    int        `json:"radius,omitempty"`
	GroupID   string     `json:"groupId,omitempty"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type CoverageBatchPredictionOutput struct {
	Data []struct {
		Location
		CoveragePredictionOutput
	} `json:"data,omitempty"`
}

// BatchPredictions retrieve coverage predictions for any batch of locations.
func (s *CoverageService) BatchPredictions(input *CoverageBatchPredictionInput) (*CoverageBatchPredictionOutput, *http.Response, error) {
	return s.BatchPredictionsContext(context.Background(), input)
}

// BatchPredictionsContext retrieve coverage predictions for any batch of locations with context.
func (s *CoverageService) BatchPredictionsContext(ctx context.Context, input *CoverageBatchPredictionInput) (*CoverageBatchPredictionOutput, *http.Response, error) {
	req, err := s.client.newRequest(ctx, "POST", "/coverages/global/predictions", input)
	if err != nil {
		return nil, nil, err
	}

	res, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, res, err
	}

	var out CoverageBatchPredictionOutput
	if err = decodeBody(res, &out); err != nil {
		return nil, res, err
	}

	return &out, res, nil
}

type CoverageRedundancyInput struct {
	Lat             float64 `url:"lat"`
	Lng             float64 `url:"lng"`
	OperatorID      string  `url:"operatorId,omitempty"`
	DeviceSituation string  `url:"deviceSituation,omitempty"`
	DeviceClassID   int     `url:"deviceClassId,omitempty"`
}

type CoverageRedundancyOutput struct {
	Redundancy int `json:"redundancy,omitempty"`
}

// Redundancy retrieve coverage redundancy for an operator.
func (s *CoverageService) Redundancy(input *CoverageRedundancyInput) (*CoverageRedundancyOutput, *http.Response, error) {
	return s.RedundancyContext(context.Background(), input)
}

// RedundancyContext retrieve coverage redundancy for an operator with context.
func (s *CoverageService) RedundancyContext(ctx context.Context, input *CoverageRedundancyInput) (*CoverageRedundancyOutput, *http.Response, error) {
	spath := fmt.Sprintf("/coverages/operators/redundancy")
	spath, err := addOptions(spath, input)
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

	var out CoverageRedundancyOutput
	if err := decodeBody(res, &out); err != nil {
		return nil, res, err
	}

	return &out, res, nil
}
