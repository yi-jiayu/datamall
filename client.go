package datamall

import (
	"net/http"
)

// The current LTA DataMall endpoint and version.
const (
	DataMallEndpoint = "http://datamall2.mytransport.sg/ltaodataservice"
	DataMallVersion  = "3.7"
)

// APIClient contains the AccountKey and a http.Client to use to make requests to the LTA DataMall API.
type APIClient struct {
	Endpoint   string
	AccountKey string
	Client     *http.Client
}

// NewDefaultClient creates a new LTA DataMall APIClient using accountKey and http.DefaultClient.
// Not recommended in production due to the lack of timeouts in http.DefaultClient.
func NewDefaultClient(accountKey string) APIClient {
	return APIClient{
		Endpoint:   DataMallEndpoint,
		AccountKey: accountKey,
		Client:     http.DefaultClient,
	}
}

// NewClient creates a new LTA DataMall APIClient using accountKey and client.
func NewClient(accountKey string, client *http.Client) APIClient {
	return APIClient{
		Endpoint:   DataMallEndpoint,
		AccountKey: accountKey,
		Client:     client,
	}
}
