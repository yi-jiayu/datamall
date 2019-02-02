package datamall

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// BusStop contains information about a bus stop
type BusStop struct {
	BusStopCode string
	RoadName    string
	Description string
	Latitude    float64
	Longitude   float64
}

// BusStops represents a response from the DataMall BusStops endpoint. A list of bus stops is contained in the Value
// property.
type BusStops struct {
	ODataMetadata string    `json:"odata.metadata"`
	Value         []BusStop `json:"value"`
}

// GetBusStops returns detailed information for all bus stops currently being serviced by buses.
func (c APIClient) GetBusStops(offset int) (BusStops, error) {
	req, err := http.NewRequest(http.MethodGet, c.Endpoint+"/BusStops", nil)
	if err != nil {
		return BusStops{}, err
	}

	req.Header.Set("AccountKey", c.AccountKey)

	q := req.URL.Query()
	q.Add("$skip", fmt.Sprintf("%d", offset))
	req.URL.RawQuery = q.Encode()

	res, err := c.Client.Do(req)
	if err != nil {
		return BusStops{}, err
	}
	defer res.Body.Close()

	var busStops BusStops
	err = json.NewDecoder(res.Body).Decode(&busStops)
	if err != nil {
		return BusStops{}, err
	}

	return busStops, nil
}
