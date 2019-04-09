package sigfox

import (
	"context"
	"fmt"
)

type GroupService service

type ListGroupsOptions struct {
	ParentID []string `url:"parentId,omitempty"`
	Deep     bool     `url:"deep,omitempty"`
	Name     string   `url:"name,omitempty"`
	Types    []int32  `url:"types,omitempty"`
	Fields   []string `url:"fields,omitempty"`
	Sort     string   `url:"sort,omitempty"`
	Limit    int32    `url:"limit,omitempty"`
	Offset   int32    `url:"offset,omitempty"`
	PageID   string   `url:"pageId,omitempty"`
}

type ListGroupsOutput struct {
	Data   []Group    `json:"data"`
	Paging Pagination `json:"paging"`
}

type Group struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type        int32  `json:"type,omitempty"`
	Timezone    string `json:"timezone,omitempty"`
	ID          string `json:"id,omitempty"`
	NameCl      string `json:"nameCl,omitempty"`
	//Path
	NetworkOperatorID string `json:"networkOperatorId,omitempty"`
}

func (s *GroupService) List(ctx context.Context, opt *ListGroupsOptions) (*ListGroupsOutput, error) {
	spath := fmt.Sprintf("/groups")
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

	var out ListGroupsOutput
	if err := decodeBody(res, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
