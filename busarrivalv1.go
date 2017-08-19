package datamall

import (
	"encoding/json"
	"net/http"
)

// ArrivingBusV1 contains information about an incoming bus at a bus stop.
type ArrivingBusV1 struct {
	EstimatedArrival string
	Latitude         string
	Longitude        string
	VisitNumber      string
	Load             string
	Feature          string
}

// ServiceV1 contains information about a bus service at a bus stop.
type ServiceV1 struct {
	ServiceNo      string
	Status         string
	Operator       string
	OriginatingID  string
	TerminatingID  string
	NextBus        ArrivingBusV1
	SubsequentBus  ArrivingBusV1
	SubsequentBus3 ArrivingBusV1
}

// BusArrivalV1 contains information about incoming buses at a bus stop.
type BusArrivalV1 struct {
	OdataMetadata string `json:"odata.metadata"`
	BusStopID     string
	Services      []ServiceV1
}

// GetBusArrivalV1Options contains optional parameters for GetBusArrivalV1.
type GetBusArrivalV1Options struct {
	ServiceNo string
	SST       bool
}

// GetBusArrivalV1 returns real-time bus arrival information at a bus stop from the LTA DataMall API. Returns information
// for a specific bus service or all services if options.serviceNo is empty. Returns timestamps in SST format if
// options.SST is true.
func (c APIClient) GetBusArrivalV1(busStopID string, options *GetBusArrivalV1Options) (BusArrivalV1, error) {
	req, err := http.NewRequest("GET", c.Endpoint+"/BusArrivalV2", nil)
	if err != nil {
		return BusArrivalV1{}, err
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
		return BusArrivalV1{}, err
	}
	defer res.Body.Close()

	var busArrival BusArrivalV1
	err = json.NewDecoder(res.Body).Decode(&busArrival)
	if err != nil {
		return BusArrivalV1{}, err
	}

	return busArrival, nil
}
