// Code generated by zanzibar
// @generated

package bar

import (
	"context"
	"net/http"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients/bar"
	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"
)

// NormalHandler is the handler for "/bar/bar-path"
type NormalHandler struct {
	Clients *clients.Clients
}

// NewNormalEndpoint creates a handler
func NewNormalEndpoint(
	gateway *zanzibar.Gateway,
) *NormalHandler {
	return &NormalHandler{
		Clients: gateway.Clients.(*clients.Clients),
	}
}

// HandleRequest handles "/bar/bar-path".
func (handler *NormalHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {
	var requestBody NormalHTTPRequest
	if ok := req.ReadAndUnmarshalBody(&requestBody); !ok {
		return
	}

	workflow := NormalEndpoint{
		Clients: handler.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	response, respHeaders, err := workflow.Handle(ctx, req.Header, &requestBody)
	if err != nil {
		req.Logger.Warn("Workflow for endpoint returned error",
			zap.String("error", err.Error()),
		)
		res.SendErrorString(500, "Unexpected server error")
		return
	}

	res.WriteJSON(200, respHeaders, response)
}

// NormalEndpoint calls thrift client Bar.Normal
type NormalEndpoint struct {
	Clients *clients.Clients
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w NormalEndpoint) Handle(
	ctx context.Context,
	// TODO(sindelar): Switch to zanzibar.Headers when tchannel
	// generation is implemented.
	headers http.Header,
	r *NormalHTTPRequest,
) (*endpointsBarBar.BarResponse, map[string]string, error) {
	clientRequest := convertToNormalClientRequest(r)

	clientHeaders := map[string]string{}

	clientRespBody, _, err := w.Clients.Bar.Normal(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		w.Logger.Warn("Could not make client request",
			zap.String("error", err.Error()),
		)
		// TODO(sindelar): Consider returning partial headers in error case.
		return nil, nil, err
	}

	// Filter and map response headers from client to server response.
	endRespHead := map[string]string{}

	response := convertNormalClientResponse(clientRespBody)
	return response, endRespHead, nil
}

func convertToNormalClientRequest(body *NormalHTTPRequest) *barClient.NormalHTTPRequest {
	clientRequest := &barClient.NormalHTTPRequest{}

	clientRequest.Request = (*clientsBarBar.BarRequest)(body.Request)

	return clientRequest
}
func convertNormalClientResponse(body *clientsBarBar.BarResponse) *endpointsBarBar.BarResponse {
	// TODO: Add response fields mapping here.
	downstreamResponse := (*endpointsBarBar.BarResponse)(body)
	return downstreamResponse
}
