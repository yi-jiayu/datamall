package datamall

import (
	"encoding/json"
	"net/http"
)

// ArrivingBus contains information about an incoming bus at a bus stop.
type ArrivingBus struct {
	OriginCode       string
	DestinationCode  string
	EstimatedArrival string
	Latitude         string
	Longitude        string
	VisitNumber      string
	Load             string
	Feature          string
	Type             string
}

// Service contains information about a bus service at a bus stop.
type Service struct {
	ServiceNo string
	Operator  string
	NextBus   ArrivingBus
	NextBus2  ArrivingBus
	NextBus3  ArrivingBus
}

// BusArrival contains information about incoming buses at a bus stop.
type BusArrival struct {
	// Deprecated: Use ODataMetadata instead
	OdataMetadata string
	ODataMetadata string `json:"odata.metadata"`
	BusStopID     string
	BusStopCode   string
	Services      []Service
}

// GetBusArrival returns real-time Bus Arrival information of Bus Services at a queried Bus Stop,
// including Est. Arrival Time, Est. Current Location, Est. Current Load
func (c APIClient) GetBusArrival(busStopCode string, serviceNo string) (BusArrival, error) {
	req, err := http.NewRequest("GET", c.Endpoint+"/BusArrivalv2", nil)
	if err != nil {
		return BusArrival{}, err
	}

	req.Header.Set("AccountKey", c.AccountKey)

	q := req.URL.Query()
	q.Add("BusStopCode", busStopCode)

	if serviceNo != "" {
		q.Add("ServiceNo", serviceNo)
	}

	req.URL.RawQuery = q.Encode()

	res, err := c.Client.Do(req)
	if err != nil {
		return BusArrival{}, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		err := Error{
			StatusCode: res.StatusCode,
		}
		return BusArrival{}, err
	}

	var busArrival BusArrival
	err = json.NewDecoder(res.Body).Decode(&busArrival)
	if err != nil {
		return BusArrival{}, err
	}

	return busArrival, nil
}
