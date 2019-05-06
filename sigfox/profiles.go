package sigfox

import (
	"context"
	"net/http"
)

type ProfileService service

type ListProfilesInput struct {
	GroupID string `url:"groupId,omitempty"`
	Inherit bool   `url:"inherit,omitempty"`
	Fields  string `url:"fields,omitempty"`
	Limit   int32  `url:"limit,omitempty"`
	Offset  int32  `url:"offset,omitempty"`
}

type ListProfilesOutput struct {
	Data   []Profile  `json:"data,omitempty"`
	Paging Pagination `json:"paging,omitempty"`
}

type Profile struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Group MinimalGroup
	Roles []MinRole
}

type MinRole struct {
	ID   string        `json:"id,omitempty"`
	Name string        `json:"name,omitempty"`
	Path []MinMetaRole `json:"path,omitempty"`
}

type MinMetaRole struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// List retrieve a list of a Group's profiles according to visibility permissions and request filters.
func (s *ProfileService) List(input *ListProfilesInput) (*ListProfilesOutput, *http.Response, error) {
	return s.ListContext(context.Background(), input)
}

// ListContext retrieve a list of a Group's profiles according to visibility permissions and request filters with context.
func (s *ProfileService) ListContext(ctx context.Context, input *ListProfilesInput) (*ListProfilesOutput, *http.Response, error) {
	spath, err := addOptions("/profiles", input)
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

	var out ListProfilesOutput
	if err := decodeBody(res, &out); err != nil {
		return nil, res, err
	}

	return &out, res, nil
}
