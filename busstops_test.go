package datamall

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestAPIClient_GetBusStops(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.Write([]byte(`{"odata.metadata":"http://datamall2.mytransport.sg/ltaodataservice/$metadata#BusStops","value":[{"BusStopCode":"01012","RoadName":"Victoria St","Description":"Hotel Grand Pacific","Latitude":1.29684825487647,"Longitude":103.85253591654006},{"BusStopCode":"01013","RoadName":"Victoria St","Description":"St. Joseph's Ch","Latitude":1.29770970610083,"Longitude":103.8532247463225},{"BusStopCode":"01019","RoadName":"Victoria St","Description":"Bras Basah Cplx","Latitude":1.29698951191332,"Longitude":103.85302201172507}]}`))
	}))
	defer ts.Close()

	client := APIClient{
		Endpoint: ts.URL,
		Client:   http.DefaultClient,
	}

	busStops, err := client.GetBusStops(0)
	if err != nil {
		t.Fatal(err)
	}

	actual := busStops
	expected := BusStops{
		OdataMetadata: "http://datamall2.mytransport.sg/ltaodataservice/$metadata#BusStops",
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
