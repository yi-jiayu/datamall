package datamall

import (
	"encoding/json"
	"net/http"
)

// ArrivingBusV2 contains information about an incoming bus at a bus stop.
type ArrivingBusV2 struct {
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

// ServiceV2 contains information about a bus service at a bus stop.
type ServiceV2 struct {
	ServiceNo string
	Operator  string
	NextBus   ArrivingBusV2
	NextBus2  ArrivingBusV2
	NextBus3  ArrivingBusV2
}

// BusArrivalV2 contains information about incoming buses at a bus stop.
type BusArrivalV2 struct {
	OdataMetadata string `json:"odata.metadata"`
	BusStopID     string
	BusStopCode   string
	Services      []ServiceV2
}

// GetBusArrivalV2 returns real-time bus arrival information from the LTA DataMall API. Returns information for a
// specific bus service or all services if serviceNo is empty.
//
// Changes from GetBusArrivalV1:
//   - BusStopID attribute in BusArrivalV1 renamed to BusStopCode.
//   - New Type attribute in ArrivingBusV2.
//   - New OriginCode and DestinationCode attributes in ArrivingBusV2, replacing OriginatingID and TerminatingID in ServiceV1.
//   - Removal of Status attribute from ArrivingBusV1. Services which are not in operation will not be included in the
//   response.
//   - SubsequentBus and SubsequentBus3 attributes renamed to NextBus2 and NextBus3.
//   - Removal of SST parameter.
func (c APIClient) GetBusArrivalV2(busStopCode string, serviceNo string) (BusArrivalV2, error) {
	req, err := http.NewRequest("GET", c.Endpoint+"/BusArrivalv2", nil)
	if err != nil {
		return BusArrivalV2{}, err
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
		return BusArrivalV2{}, err
	}
	defer res.Body.Close()

	var busArrival BusArrivalV2
	err = json.NewDecoder(res.Body).Decode(&busArrival)
	if err != nil {
		return BusArrivalV2{}, err
	}

	return busArrival, nil
}
