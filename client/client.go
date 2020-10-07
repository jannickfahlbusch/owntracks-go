//go:generate mockgen -destination=mocks/mock_client.go -package=mocks . Client
package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/jannickfahlbusch/owntracks-go/types"
)

var _ Client = &clientInstance{}

const (
	DefaultTimeout    = 2 * time.Minute
	DefaultTimeFormat = time.RFC3339
)

type clientInstance struct {
	client *http.Client
	apiURL string
}

// Client is a client for the Owntracks Recorder API
type Client interface {
	Users(context.Context) ([]string, error)
	Devices(context.Context, string) ([]string, error)
	Locations(context.Context, string, string, time.Time, time.Time) (*types.LocationList, error)
	Version(ctx context.Context) (*types.Version, error)
}

// New creates a new Owntracks Client
func New(apiURL string) Client {
	httpClient := &http.Client{
		Timeout: DefaultTimeout,
	}

	return NewWithClient(apiURL, httpClient)
}

// NewWithClient creates a new Owntracks Client and uses the given HTTP client for communication
func NewWithClient(apiURL string, httpClient *http.Client) Client {
	return &clientInstance{
		client: httpClient,
		apiURL: apiURL,
	}
}

// do executes the given request and decodes the response into the 'into' parameter
func (client *clientInstance) do(request *http.Request, into interface{}) error {
	response, err := client.client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&into)

	return err
}

// newRequest creates a new request with the given path. Parameters is optional and will be encoded to query values if set
func (client *clientInstance) newRequest(ctx context.Context, path string, parameters map[string]string) (*http.Request, error) {
	endpointURL, err := url.Parse(client.apiURL)
	if err != nil {
		return nil, err
	}

	query := endpointURL.Query()
	for key, value := range parameters {
		query.Set(key, value)

	}
	endpointURL.RawQuery = query.Encode()

	endpointURL.Path += path

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL.String(), nil)

	return request, err
}

// Users returns a list of Users known to the Recorder
func (client *clientInstance) Users(ctx context.Context) ([]string, error) {
	request, err := client.newRequest(ctx, "/list", nil)
	if err != nil {
		return nil, err
	}

	listResponse := &types.ListResponse{}
	err = client.do(request, listResponse)

	if err != nil {
		return nil, err
	}

	return listResponse.Results, nil
}

// Devices returns the devices of the user 'user'
func (client *clientInstance) Devices(ctx context.Context, user string) ([]string, error) {
	request, err := client.newRequest(ctx, "/list", map[string]string{"user": user})
	if err != nil {
		return nil, err
	}

	listResponse := &types.ListResponse{}
	err = client.do(request, listResponse)

	if err != nil {
		return nil, err
	}

	return listResponse.Results, nil
}

// Locations returns the locations of the users device during the given timeframe
func (client *clientInstance) Locations(ctx context.Context, user, device string, from, to time.Time) (*types.LocationList, error) {
	parameters := map[string]string{
		"user":   user,
		"device": device,
		"from":   from.Format(DefaultTimeFormat),
		"to":     to.Format(DefaultTimeFormat),
	}
	request, err := client.newRequest(ctx, "/locations", parameters)
	if err != nil {
		return nil, err
	}

	locations := &types.LocationList{}
	err = client.do(request, locations)

	if err != nil {
		return nil, err
	}

	return locations, err
}

// Version returns the version of the recorder
func (client *clientInstance) Version(ctx context.Context) (*types.Version, error) {
	request, err := client.newRequest(ctx, "/version", nil)
	if err != nil {
		return nil, err
	}

	versionResponse := &types.Version{}
	err = client.do(request, versionResponse)

	if err != nil {
		return nil, err
	}

	return versionResponse, nil
}
