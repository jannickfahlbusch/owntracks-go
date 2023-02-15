//go:generate mockgen -destination=mocks/mock_client.go -package=mocks . Client
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"pkg.jf-projects.de/owntracks/types"
)

var _ Client = &clientInstance{}

const (
	DefaultTimeout    = 2 * time.Minute
	DefaultTimeFormat = time.RFC3339
)

const API_PATH = "/api/0"

type clientInstance struct {
	client *http.Client
	apiURL string
}

// Client is a client for the Owntracks Recorder API
type Client interface {
	Users(context.Context) ([]string, error)
	Devices(context.Context, string) ([]string, error)
	Locations(context.Context, string, string, time.Time, time.Time) (*types.LocationList, error)
	Publish(context.Context, string, string, *types.Location) error
	Version(context.Context) (*types.Version, error)

	// Exists is not present on the recorder API. This method will use Locations() internally
	Exists(context.Context, string, string, *types.Location) (bool, error)
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

	return http.NewRequestWithContext(ctx, http.MethodGet, endpointURL.String(), nil)
}

// Users returns a list of Users known to the Recorder
func (client *clientInstance) Users(ctx context.Context) ([]string, error) {
	request, err := client.newRequest(ctx, API_PATH+"/list", nil)
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
	request, err := client.newRequest(ctx, API_PATH+"/list", map[string]string{"user": user})
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
	request, err := client.newRequest(ctx, API_PATH+"/locations", parameters)
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
	request, err := client.newRequest(ctx, API_PATH+"/version", nil)
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

func (client *clientInstance) Publish(ctx context.Context, user, device string, location *types.Location) error {
	endpointURL, err := url.Parse(client.apiURL + "/pub")
	if err != nil {
		return err
	}

	jsonContent, err := json.Marshal(location)
	if err != nil {
		return err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL.String(), bytes.NewReader(jsonContent))
	if err != nil {
		return err
	}

	request.Header.Set("X-Limit-U", user)
	request.Header.Set("X-Limit-D", device)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return nil
}

func (client *clientInstance) Exists(ctx context.Context, user, device string, location *types.Location) (bool, error) {

	// From and To are exclusive and in UTC, so we need to substract/add a second to get the location in question in between
	fromExclusive := location.Timestamp.Add(-time.Second).UTC()
	toExclusive := location.Timestamp.Add(time.Second).UTC()

	locations, err := client.Locations(ctx, user, device, fromExclusive, toExclusive)
	if err != nil {
		return false, err
	}

	if locations.Count <= 0 {
		return false, nil
	}

	for _, availableLocation := range locations.Data {
		if availableLocation.IsSamePlaceAndTime(location) {
			return true, nil
		}
	}

	return false, nil
}
