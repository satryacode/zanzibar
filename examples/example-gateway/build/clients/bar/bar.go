// Code generated by zanzibar
// @generated

package barClient

import (
	"context"
	"strconv"

	"github.com/pkg/errors"
	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	"github.com/uber/zanzibar/runtime"
)

// BarClient is the http client for service Bar.
type BarClient struct {
	ClientID   string
	HTTPClient *zanzibar.HTTPClient
}

// NewClient returns a new http client for service Bar.
func NewClient(
	gateway *zanzibar.Gateway,
) *BarClient {
	ip := gateway.Config.MustGetString("clients.bar.ip")
	port := gateway.Config.MustGetInt("clients.bar.port")

	baseURL := "http://" + ip + ":" + strconv.Itoa(int(port))
	return &BarClient{
		ClientID:   "bar",
		HTTPClient: zanzibar.NewHTTPClient(gateway, baseURL),
	}
}

// ArgNotStruct calls "/arg-not-struct-path" endpoint.
func (c *BarClient) ArgNotStruct(
	ctx context.Context,
	headers map[string]string,
	r *ArgNotStructHTTPRequest,
) (map[string]string, error) {

	req := zanzibar.NewClientHTTPRequest(
		c.ClientID, "argNotStruct", c.HTTPClient,
	)

	// Generate full URL.
	fullURL := c.HTTPClient.BaseURL + "/arg-not-struct-path"

	err := req.WriteJSON("POST", fullURL, headers, r)
	if err != nil {
		return nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{200, 403})

	switch res.StatusCode {
	case 200:
		// TODO: log about unexpected body bytes?
		_, err = res.ReadAll()
		if err != nil {
			return respHeaders, err
		}

		return respHeaders, nil

	case 403:
		var exception clientsBarBar.BarException
		err = res.ReadAndUnmarshalBody(&exception)
		if err != nil {
			return respHeaders, err
		}
		return respHeaders, &exception

	default:
		// TODO: log about unexpected body bytes?
		_, err = res.ReadAll()
		if err != nil {
			return respHeaders, err
		}
	}

	return respHeaders, errors.Errorf(
		"Unexpected http client response (%d)", res.StatusCode,
	)
}

// MissingArg calls "/missing-arg-path" endpoint.
func (c *BarClient) MissingArg(
	ctx context.Context,
	headers map[string]string,
) (*clientsBarBar.BarResponse, map[string]string, error) {

	req := zanzibar.NewClientHTTPRequest(
		c.ClientID, "missingArg", c.HTTPClient,
	)

	// Generate full URL.
	fullURL := c.HTTPClient.BaseURL + "/missing-arg-path"

	err := req.WriteJSON("GET", fullURL, headers, nil)
	if err != nil {
		return nil, nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return nil, nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{200, 403})

	switch res.StatusCode {
	case 200:
		var responseBody clientsBarBar.BarResponse
		err = res.ReadAndUnmarshalBody(&responseBody)
		if err != nil {
			return nil, respHeaders, err
		}

		return &responseBody, respHeaders, nil

	case 403:
		var exception clientsBarBar.BarException
		err = res.ReadAndUnmarshalBody(&exception)
		if err != nil {
			return nil, respHeaders, err
		}
		return nil, respHeaders, &exception

	default:
		// TODO: log about unexpected body bytes?
		_, err = res.ReadAll()
		if err != nil {
			return nil, respHeaders, err
		}
	}

	return nil, respHeaders, errors.Errorf(
		"Unexpected http client response (%d)", res.StatusCode,
	)
}

// NoRequest calls "/no-request-path" endpoint.
func (c *BarClient) NoRequest(
	ctx context.Context,
	headers map[string]string,
) (*clientsBarBar.BarResponse, map[string]string, error) {

	req := zanzibar.NewClientHTTPRequest(
		c.ClientID, "noRequest", c.HTTPClient,
	)

	// Generate full URL.
	fullURL := c.HTTPClient.BaseURL + "/no-request-path"

	err := req.WriteJSON("GET", fullURL, headers, nil)
	if err != nil {
		return nil, nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return nil, nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{200, 403})

	switch res.StatusCode {
	case 200:
		var responseBody clientsBarBar.BarResponse
		err = res.ReadAndUnmarshalBody(&responseBody)
		if err != nil {
			return nil, respHeaders, err
		}

		return &responseBody, respHeaders, nil

	case 403:
		var exception clientsBarBar.BarException
		err = res.ReadAndUnmarshalBody(&exception)
		if err != nil {
			return nil, respHeaders, err
		}
		return nil, respHeaders, &exception

	default:
		// TODO: log about unexpected body bytes?
		_, err = res.ReadAll()
		if err != nil {
			return nil, respHeaders, err
		}
	}

	return nil, respHeaders, errors.Errorf(
		"Unexpected http client response (%d)", res.StatusCode,
	)
}

// Normal calls "/bar-path" endpoint.
func (c *BarClient) Normal(
	ctx context.Context,
	headers map[string]string,
	r *NormalHTTPRequest,
) (*clientsBarBar.BarResponse, map[string]string, error) {

	req := zanzibar.NewClientHTTPRequest(
		c.ClientID, "normal", c.HTTPClient,
	)

	// Generate full URL.
	fullURL := c.HTTPClient.BaseURL + "/bar-path"

	err := req.WriteJSON("POST", fullURL, headers, r)
	if err != nil {
		return nil, nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return nil, nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{200, 403})

	switch res.StatusCode {
	case 200:
		var responseBody clientsBarBar.BarResponse
		err = res.ReadAndUnmarshalBody(&responseBody)
		if err != nil {
			return nil, respHeaders, err
		}

		return &responseBody, respHeaders, nil

	case 403:
		var exception clientsBarBar.BarException
		err = res.ReadAndUnmarshalBody(&exception)
		if err != nil {
			return nil, respHeaders, err
		}
		return nil, respHeaders, &exception

	default:
		// TODO: log about unexpected body bytes?
		_, err = res.ReadAll()
		if err != nil {
			return nil, respHeaders, err
		}
	}

	return nil, respHeaders, errors.Errorf(
		"Unexpected http client response (%d)", res.StatusCode,
	)
}

// TooManyArgs calls "/too-many-args-path" endpoint.
func (c *BarClient) TooManyArgs(
	ctx context.Context,
	headers map[string]string,
	r *TooManyArgsHTTPRequest,
) (*clientsBarBar.BarResponse, map[string]string, error) {

	req := zanzibar.NewClientHTTPRequest(
		c.ClientID, "tooManyArgs", c.HTTPClient,
	)

	// Generate full URL.
	fullURL := c.HTTPClient.BaseURL + "/too-many-args-path"

	err := req.WriteJSON("POST", fullURL, headers, r)
	if err != nil {
		return nil, nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return nil, nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{200, 403})

	switch res.StatusCode {
	case 200:
		var responseBody clientsBarBar.BarResponse
		err = res.ReadAndUnmarshalBody(&responseBody)
		if err != nil {
			return nil, respHeaders, err
		}

		return &responseBody, respHeaders, nil

	case 403:
		var exception clientsBarBar.BarException
		err = res.ReadAndUnmarshalBody(&exception)
		if err != nil {
			return nil, respHeaders, err
		}
		return nil, respHeaders, &exception

	default:
		// TODO: log about unexpected body bytes?
		_, err = res.ReadAll()
		if err != nil {
			return nil, respHeaders, err
		}
	}

	return nil, respHeaders, errors.Errorf(
		"Unexpected http client response (%d)", res.StatusCode,
	)
}
