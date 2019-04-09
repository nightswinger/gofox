package sigfox

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"reflect"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

const (
	defaultBaseURL = "https://api.sigfox.com/v2"
)

type Client struct {
	baseURL    *url.URL
	HTTPClient *http.Client

	Login, Password string
	UserAgent       string

	common service

	Devices    *DevicesService
	DeviceType *DeviceTypeService
	Group      *GroupService
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

	c := &Client{HTTPClient: &http.Client{}, baseURL: parsedURL, Login: login, Password: password}
	c.common.client = c
	c.Devices = (*DevicesService)(&c.common)
	c.DeviceType = (*DeviceTypeService)(&c.common)
	c.Group = (*GroupService)(&c.common)

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

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()

	if w, ok := out.(io.Writer); ok {
		io.Copy(w, resp.Body)
	}
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
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
