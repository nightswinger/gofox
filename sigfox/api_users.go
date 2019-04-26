package sigfox

import (
	"context"
	"fmt"
	"net/http"
)

type ApiUserService service

type ApiUser struct {
	Name         string       `json:"name,omitempty"`
	Timezone     string       `json:"timezone,omitempty"`
	Group        MinimalGroup `json:"group,omitempty"`
	CreationTime int64        `json:"creationTime,omitempty"`
	ID           string       `json:"id,omitempty"`
	AccessToken  string       `json:"accessToken,omitempty"`
	Profiles     []Profile    `json:"profiles,omitempty"`
}

type MinimalGroup struct {
	Name  string `json:"name,omitempty"`
	Type  int32  `json:"type,omitempty"`
	ID    string `json:"id,omitempty"`
	Level int32  `json:"level,omitempty"`
}

type Profile struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ListApiUsersOutput struct {
	Data   []ApiUser  `json:"data"`
	Paging Pagination `json:"paging"`
}

func (s *ApiUserService) List(params ...QueryParam) (*ListApiUsersOutput, *http.Response, error) {
	return s.ListContext(context.Background(), params...)
}

func (s *ApiUserService) ListContext(ctx context.Context, params ...QueryParam) (*ListApiUsersOutput, *http.Response, error) {
	opt := &QueryParams{}
	for _, param := range params {
		param(opt)
	}

	spath := fmt.Sprintf("api-users/")
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

	var out ListApiUsersOutput
	if err := decodeBody(res, &out); err != nil {
		return nil, res, err
	}

	return &out, res, nil
}

func (s *ApiUserService) Info(ApiUserID string, params ...QueryParam) (*ApiUser, *http.Response, error) {
	return s.InfoContext(context.Background(), ApiUserID, params...)
}

func (s *ApiUserService) InfoContext(ctx context.Context, ApiUserID string, params ...QueryParam) (*ApiUser, *http.Response, error) {
	opt := &QueryParams{}
	for _, param := range params {
		param(opt)
	}

	spath := fmt.Sprintf("api-users/%s", ApiUserID)
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

	var out ApiUser
	if err := decodeBody(res, &out); err != nil {
		return nil, res, err
	}

	return &out, res, nil
}
