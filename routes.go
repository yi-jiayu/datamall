package datamall

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// BusRoute contains information about a single bus service at a particular bus stop.
type BusRoute struct {
	ServiceNo       string
	Operator        string
	Direction       int
	StopSequence    int
	BusStopCode     string
	Distance        float32
	WeekdayFirstBus string `json:"WD_FirstBus"`
	WeekdayLastBus  string `json:"WD_LastBus"`
	SatFirstBus     string `json:"SAT_FirstBus"`
	SatLastBus      string `json:"SAT_LastBus"`
	SunFirstBus     string `json:"SUN_FirstBus"`
	SunLastBun      string `json:"SUN_LastBus"`
}

// BusRoutes represents a response from the DataMall BusRoutes endpoint. An array of BusRoute is contained in the Value
// property.
type BusRoutes struct {
	// Deprecated: Use ODataMetadata instead
	OdataMetadata string
	ODataMetadata string     `json:"odata.metadata"`
	Value         []BusRoute `json:"value"`
}

// GetBusRoutes returns detailed route information for all services currently in operation, including: all bus stops
// along each route, first/last bus timings for each stop.
func (c APIClient) GetBusRoutes(offset int) (BusRoutes, error) {
	req, err := http.NewRequest(http.MethodGet, c.Endpoint+"/BusRoutes", nil)
	if err != nil {
		return BusRoutes{}, err
	}

	req.Header.Set("AccountKey", c.AccountKey)

	q := req.URL.Query()
	q.Add("$skip", fmt.Sprintf("%d", offset))
	req.URL.RawQuery = q.Encode()

	res, err := c.Client.Do(req)
	if err != nil {
		return BusRoutes{}, err
	}
	defer res.Body.Close()

	var busRoutes BusRoutes
	err = json.NewDecoder(res.Body).Decode(&busRoutes)
	if err != nil {
		return BusRoutes{}, err
	}

	return busRoutes, nil
}
