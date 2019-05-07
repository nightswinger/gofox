package sigfox

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"reflect"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

const (
	defaultBaseURL = "https://api.sigfox.com/v2"
	userAgent      = "gofox"
)

type Client struct {
	baseURL    *url.URL
	HTTPClient *http.Client

	Login, Password string
	UserAgent       string

	common service

	ApiUser    *ApiUserService
	Coverage   *CoverageService
	Device     *DeviceService
	DeviceType *DeviceTypeService
	Group      *GroupService
	Profile    *ProfileService
	Tile       *TileService
}

type service struct {
	client *Client
}

type Pagination struct {
	Next string `json:"next"`
}

func NewClient(login, password string) (*Client, error) {
	if len(login) == 0 {
		return nil, errors.New("missing login key")
	}

	if len(password) == 0 {
		return nil, errors.New("missing login password")
	}

	parsedURL, err := url.ParseRequestURI(defaultBaseURL)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", defaultBaseURL)
	}

	c := &Client{HTTPClient: &http.Client{}, baseURL: parsedURL, UserAgent: userAgent, Login: login, Password: password}
	c.common.client = c
	c.ApiUser = (*ApiUserService)(&c.common)
	c.Coverage = (*CoverageService)(&c.common)
	c.Device = (*DeviceService)(&c.common)
	c.DeviceType = (*DeviceTypeService)(&c.common)
	c.Group = (*GroupService)(&c.common)
	c.Profile = (*ProfileService)(&c.common)
	c.Tile = (*TileService)(&c.common)

	return c, nil
}

func (c *Client) newRequest(ctx context.Context, method, spath string, body interface{}) (*http.Request, error) {
	u := *c.baseURL
	u.Path = path.Join(c.baseURL.Path, spath)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.SetBasicAuth(c.Login, c.Password)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

type ErrorResponse struct {
	Response *http.Response
	Message  string `json:"message"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}

// Do sends an API request and returns the API response.
func (c *Client) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	err = checkResponse(resp)
	if err != nil {
		return resp, err
	}

	return resp, err
}

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	return errorResponse
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()

	if w, ok := out.(io.Writer); ok {
		io.Copy(w, resp.Body)
	}
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

type QueryParams struct {
	Fields    string   `url:"fields,omitempty"`
	Since     int64    `url:"since,omitempty"`
	Before    int64    `url:"before,omitempty"`
	Limit     int      `url:"limit,omitempty"`
	Offset    int32    `url:"offset,omitempty"`
	ProfileID string   `url:"profileId,omitempty"`
	GroupIds  []string `url:"groupIds,omitempty"`
}

type QueryParam func(*QueryParams)

func Fields(s string) QueryParam {
	return func(q *QueryParams) { q.Fields = s }
}

func Since(i int64) QueryParam {
	return func(q *QueryParams) { q.Since = i }
}

func Before(i int64) QueryParam {
	return func(q *QueryParams) { q.Before = i }
}

func Limit(i int) QueryParam {
	return func(q *QueryParams) { q.Limit = i }
}

func Offset(i int32) QueryParam {
	return func(q *QueryParams) { q.Offset = i }
}

func ProfileID(id string) QueryParam {
	return func(q *QueryParams) { q.ProfileID = id }
}

func GroupIds(s []string) QueryParam {
	return func(q *QueryParams) { q.GroupIds = s }
}

func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	vs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = vs.Encode()
	return u.String(), nil
}
