//go:generate go run gen-accessors.go

package servicenow

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"

	"github.com/google/go-querystring/query"
)

const (
	userAgent = "go-servicenow"
	jsonv2Opt = "JSONv2"
)

// A Client manages communication with the ServiceNow API.
type Client struct {
	clientMu sync.Mutex   // clientMu protects the client during calls that modify the CheckRedirect func.
	client   *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. BaseURL should always be specified with a trailing slash.
	BaseURL *url.URL

	// User agent used when communicating with the ServiceNow API.
	UserAgent string

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the ServiceNow API.
	Incidents               *IncidentsService
	ChangeRequests          *ChangeRequestsService
	StandardChangeTemplates *StandardChangeTemplatesService
}

type service struct {
	client *Client
}

type internalFields struct {
	SysparmSysID  *string           `url:"sysparm_sys_id,omitempty"`
	SysparmAction SysparmActionType `url:"sysparm_action"`
	SysparmQuery  string            `url:"sysparm_query,omitempty"`
}

type SysparmActionType string

const (
	SysparmActionInsert         = "insert"
	SysparmActionInsertMultiple = "insertMultiple"
	SysparmActionUpdate         = "update"
	SysparmActionDelete         = "deleteRecord"
	SysparmActionDeleteMultiple = "deleteMultiple"
)

type DisplayValueType string

const (
	DisplayValueTrue  DisplayValueType = "true"
	DisplayValueFalse DisplayValueType = "false"
	DisplayValueAll   DisplayValueType = "all"
)

const (
	Eq         OperandType = "="
	Ne         OperandType = "!="
	AND        OperandType = "^"
	LOR        OperandType = "^OR"
	LIKE       OperandType = "LIKE"
	STARTSWITH OperandType = "STARTSWITH"
	ENDSWITH   OperandType = "ENDSWITH"
)

type OperandType string

type QueryOpts struct {
	Key string
	Op  OperandType
	Val string
}

type ListOptions struct {
	Limit        string           `url:"sysparm_record_count,omitempty"`
	DisplayValue DisplayValueType `url:"displayvalue,omitempty"`

	QueryOpts []QueryOpts `url:"-"`

	internalFields
}

type GetOptions struct {
	DisplayValue DisplayValueType `url:"displayvalue,omitempty"`

	internalFields
}

type CreateOptions struct {
	DisplayValue DisplayValueType `url:"displayvalue,omitempty"`

	internalFields
}

type UpdateOptions struct {
	DisplayValue DisplayValueType `url:"displayvalue,omitempty"`

	internalFields
}

// addOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	qs.Add(jsonv2Opt, "")

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// NewClient returns a new ServiceNow API client. If a nil httpClient is
// provided, a new http.Client will be used. To use API methods which require
// authentication, provide a http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(baseURL string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	baseEndpoint, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	if !strings.HasSuffix(baseEndpoint.Path, "/") {
		baseEndpoint.Path += "/"
	}

	c := &Client{client: httpClient, BaseURL: baseEndpoint, UserAgent: userAgent}
	c.common.client = c
	c.Incidents = (*IncidentsService)(&c.common)
	c.ChangeRequests = (*ChangeRequestsService)(&c.common)
	c.StandardChangeTemplates = (*StandardChangeTemplatesService)(&c.common)
	return c, nil
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Response is a ServiceNow API response. This wraps the standard http.Response
// returned from ServiceNow and provides convenient access to things like
// pagination links.
type Response struct {
	*http.Response
}

// newResponse creates a new Response for the provided http.Response.
// r must not be nil.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	if ctx == nil {
		return nil, errors.New("context must be non-nil")
	}
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// If the error type is *url.Error, sanitize its URL before returning.
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = sanitizeURL(url).String()
				return nil, e
			}
		}

		return nil, err
	}

	defer resp.Body.Close()

	response := newResponse(resp)

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return response, err
}

// sanitizeURL redacts the client_secret parameter from the URL which may be
// exposed to the user.
func sanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("client_secret")) > 0 {
		params.Set("client_secret", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}

func setCredentialsAsHeaders(req *http.Request, id, secret string) *http.Request {
	// To set extra headers, we must make a copy of the Request so
	// that we don't modify the Request we were given. This is required by the
	// specification of http.RoundTripper.
	//
	// Since we are going to modify only req.Header here, we only need a deep copy
	// of req.Header.
	convertedRequest := new(http.Request)
	*convertedRequest = *req
	convertedRequest.Header = make(http.Header, len(req.Header))

	for k, s := range req.Header {
		convertedRequest.Header[k] = append([]string(nil), s...)
	}
	convertedRequest.SetBasicAuth(id, secret)
	return convertedRequest
}

// BasicAuthTransport is an http.RoundTripper that authenticates all requests
// using HTTP Basic Authentication with the provided username and password.
type BasicAuthTransport struct {
	Username string // ServiceNow username
	Password string // ServiceNow password

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

// RoundTrip implements the RoundTripper interface.
func (t *BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := setCredentialsAsHeaders(req, t.Username, t.Password)
	return t.transport().RoundTrip(req2)
}

// Client returns an *http.Client that makes requests that are authenticated
// using HTTP Basic Authentication.
func (t *BasicAuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *BasicAuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}
