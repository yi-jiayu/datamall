package datamall

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type BusStop struct {
	BusStopCode string
	RoadName string
	Description string
	Latitude float64
	Longitude float64
}

type BusStops struct {
	OdataMetadata string `json:"odata.metadata"`
	Value []BusStop `json:"value"`
}

func (c APIClient) GetBusStops(offset int) (BusStops, error) {
	req, err := http.NewRequest(http.MethodGet, c.Endpoint + "/BusStops", nil)
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
