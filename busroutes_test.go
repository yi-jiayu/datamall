package datamall

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestAPIClient_GetBusRoutes(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.Write([]byte(`{"odata.metadata":"http://datamall2.mytransport.sg/ltaodataservice/$metadataBusRoutes","value":[{"ServiceNo":"10","Operator":"SBST","Direction":1,"StopSequence":51,"BusStopCode":"14081","Distance":22.5,"WD_FirstBus":"0606","WD_LastBus":"0003","SAT_FirstBus":"0602","SAT_LastBus":"0007","SUN_FirstBus":"0601","SUN_LastBus":"0004"},{"ServiceNo":"10","Operator":"SBST","Direction":1,"StopSequence":52,"BusStopCode":"14101","Distance":22.9,"WD_FirstBus":"0607","WD_LastBus":"0004","SAT_FirstBus":"0603","SAT_LastBus":"0008","SUN_FirstBus":"0602","SUN_LastBus":"0004"}]}`))
	}))
	defer ts.Close()

	client := APIClient{
		Endpoint: ts.URL,
		Client:   http.DefaultClient,
	}

	busRoutes, err := client.GetBusRoutes(0)
	if err != nil {
		t.Fatal(err)
	}

	actual := busRoutes
	expected := BusRoutes{
		OdataMetadata: "http://datamall2.mytransport.sg/ltaodataservice/$metadataBusRoutes",
		Value: []BusRoute{
			{
				ServiceNo:       "10",
				Operator:        "SBST",
				Direction:       1,
				StopSequence:    51,
				BusStopCode:     "14081",
				Distance:        22.5,
				WeekdayFirstBus: "0606",
				WeekdayLastBus:  "0003",
				SatFirstBus:     "0602",
				SatLastBus:      "0007",
				SunFirstBus:     "0601",
				SunLastBun:      "0004"},
			{
				ServiceNo:       "10",
				Operator:        "SBST",
				Direction:       1,
				StopSequence:    52,
				BusStopCode:     "14101",
				Distance:        22.9,
				WeekdayFirstBus: "0607",
				WeekdayLastBus:  "0004",
				SatFirstBus:     "0603",
				SatLastBus:      "0008",
				SunFirstBus:     "0602",
				SunLastBun:      "0004"},
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		fmt.Printf("Expected:\n%#v\nActual:\n%#v\n", expected, actual)
		t.Fail()
	}
}
