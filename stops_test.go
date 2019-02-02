package datamall

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAPIClient_GetBusStops(t *testing.T) {
	client := APIClient{
		Client: &MockHTTPClient{
			Body: `{
  "odata.metadata": "http://datamall2.mytransport.sg/ltaodataservice/$metadata#BusStops",
  "value": [
    {
      "BusStopCode": "01012",
      "RoadName": "Victoria St",
      "Description": "Hotel Grand Pacific",
      "Latitude": 1.29684825487647,
      "Longitude": 103.85253591654006
    },
    {
      "BusStopCode": "01013",
      "RoadName": "Victoria St",
      "Description": "St. Joseph's Ch",
      "Latitude": 1.29770970610083,
      "Longitude": 103.8532247463225
    },
    {
      "BusStopCode": "01019",
      "RoadName": "Victoria St",
      "Description": "Bras Basah Cplx",
      "Latitude": 1.29698951191332,
      "Longitude": 103.85302201172507
    }
  ]
}`,
		},
	}

	busStops, err := client.GetBusStops(0)
	if err != nil {
		t.Fatal(err)
	}

	actual := busStops
	expected := BusStops{
		ODataMetadata: "http://datamall2.mytransport.sg/ltaodataservice/$metadata#BusStops",
		Value: []BusStop{
			{
				BusStopCode: "01012",
				RoadName:    "Victoria St",
				Description: "Hotel Grand Pacific",
				Latitude:    1.29684825487647,
				Longitude:   103.85253591654006,
			},
			{
				BusStopCode: "01013",
				RoadName:    "Victoria St",
				Description: "St. Joseph's Ch",
				Latitude:    1.29770970610083,
				Longitude:   103.8532247463225,
			},
			{
				BusStopCode: "01019",
				RoadName:    "Victoria St",
				Description: "Bras Basah Cplx",
				Latitude:    1.29698951191332,
				Longitude:   103.85302201172507,
			},
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		fmt.Printf("Expected:\n%#v\nActual:\n%#v\n", expected, actual)
		t.Fail()
	}
}
