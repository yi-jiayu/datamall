package datamall

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// ArrivingBus contains information about an incoming bus at a bus stop.
type ArrivingBus struct {
	OriginCode       string
	DestinationCode  string
	EstimatedArrival time.Time
	Latitude         float64 `json:",string"`
	Longitude        float64 `json:",string"`
	VisitNumber      int     `json:",string"`
	Load             string
	Feature          string
	Type             string
}

func (b *ArrivingBus) UnmarshalJSON(data []byte) error {
	type arrivingBus struct {
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
	var ab arrivingBus
	err := json.Unmarshal(data, &ab)
	if err != nil {
		return err
	}
	b.OriginCode = ab.OriginCode
	b.DestinationCode = ab.DestinationCode
	if ab.EstimatedArrival != "" {
		t, err := time.Parse(time.RFC3339, ab.EstimatedArrival)
		if err != nil {
			return err
		}
		b.EstimatedArrival = t
	}
	if ab.Latitude != "" {
		lat, err := strconv.ParseFloat(ab.Latitude, 64)
		if err != nil {
			return err
		}
		b.Latitude = lat
	}
	if ab.Longitude != "" {
		lon, err := strconv.ParseFloat(ab.Longitude, 64)
		if err != nil {
			return err
		}
		b.Longitude = lon
	}
	if ab.VisitNumber != "" {
		n, err := strconv.Atoi(ab.VisitNumber)
		if err != nil {
			return err
		}
		b.VisitNumber = n
	}
	b.Load = ab.Load
	b.Feature = ab.Feature
	b.Type = ab.Type
	return nil
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
	ODataMetadata string `json:"odata.metadata"`

	// Deprecated: Use BusStopCode instead
	BusStopID string

	BusStopCode string
	Services    []Service
}

// GetBusArrival returns real-time Bus Arrival information of Bus Services at a queried Bus Stop,
// including Est. Arrival Time, Est. Current Location, Est. Current Load
func (c APIClient) GetBusArrival(busStopCode string, serviceNo string) (BusArrival, error) {
	req, err := http.NewRequest("GET", c.Endpoint+"/v3/BusArrival", nil)
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
		err := &Error{
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
