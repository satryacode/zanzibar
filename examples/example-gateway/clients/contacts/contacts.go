package contactsClient

import (
	"bytes"
	"net/http"
	"strconv"
)

// ContactsClient is the client itself
type ContactsClient struct {
	httpClient *http.Client
	baseURL    string
}

// SaveContactsResponse response object
type SaveContactsResponse struct {
	Res *http.Response
}

// SaveContacts will call POST /:uuid/contacts
func (contacts *ContactsClient) SaveContacts(
	save *SaveContactsRequest,
) (*SaveContactsResponse, error) {
	fullURL := contacts.baseURL + "/" + save.UserUUID + "/contacts"
	method := "POST"

	rawBody, err := save.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, fullURL, bytes.NewReader(rawBody))
	if err != nil {
		return nil, err
	}

	res, err := contacts.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return &SaveContactsResponse{
		Res: res,
	}, nil
}

// Options for creating Contacts client
type Options struct {
	IP   string
	Port int32
}

// Create makes a contacts client
func Create(opts *Options) *ContactsClient {
	baseURL := "http://" + opts.IP + ":" + strconv.Itoa(int(opts.Port))

	client := &ContactsClient{
		httpClient: &http.Client{
			Transport: &http.Transport{
				DisableKeepAlives:   false,
				MaxIdleConns:        500,
				MaxIdleConnsPerHost: 500,
			},
		},
		baseURL: baseURL,
	}
	return client
}