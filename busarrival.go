package datamall

import (
	"encoding/json"
	"net/http"
)

// ArrivingBus contains information about an incoming bus at a bus stop.
type ArrivingBus struct {
	EstimatedArrival string
	Latitude         string
	Longitude        string
	VisitNumber      string
	Load             string
	Feature          string
}

// Service contains information about a bus service at a bus stop.
type Service struct {
	ServiceNo      string
	Status         string
	Operator       string
	OriginatingID  string
	TerminatingID  string
	NextBus        ArrivingBus
	SubsequentBus  ArrivingBus
	SubsequentBus3 ArrivingBus
}

// BusArrival contains information about incoming buses at a bus stop.
type BusArrival struct {
	OdataMetadata string `json:"odata.metadata"`
	BusStopID     string
	Services      []Service
}

// GetBusArrivalOptions contains optional parameters for GetBusArrival.
type GetBusArrivalOptions struct {
	ServiceNo string
	SST       bool
}

// GetBusArrival returns real-time bus arrival information at a bus stop from the LTA DataMall API.
func (c APIClient) GetBusArrival(busStopID string, options *GetBusArrivalOptions) (BusArrival, error) {
	req, err := http.NewRequest("GET", c.Endpoint+"/BusArrival", nil)
	if err != nil {
		return BusArrival{}, err
	}

	req.Header.Set("AccountKey", c.AccountKey)

	q := req.URL.Query()
	q.Add("BusStopID", busStopID)

	if options != nil {
		if serviceNo := options.ServiceNo; serviceNo != "" {
			q.Add("ServiceNo", serviceNo)
		}

		if options.SST {
			q.Add("SST", "True")
		}
	}

	req.URL.RawQuery = q.Encode()

	res, err := c.Client.Do(req)
	if err != nil {
		return BusArrival{}, err
	}
	defer res.Body.Close()

	var busArrival BusArrival
	err = json.NewDecoder(res.Body).Decode(&busArrival)
	if err != nil {
		return BusArrival{}, err
	}

	return busArrival, nil
}
