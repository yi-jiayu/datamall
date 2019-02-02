package datamall

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func mustParseRFC3339(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}

func TestAPIClient_GetBusArrival(t *testing.T) {
	client := APIClient{
		Client: &MockHTTPClient{
			Body: `{
  "odata.metadata": "http://datamall2.mytransport.sg/ltaodataservice/$metadata#BusArrivalv2/@Element",
  "BusStopCode": "96049",
  "Services": [
    {
      "ServiceNo": "2",
      "Operator": "GAS",
      "NextBus": {
        "OriginCode": "99009",
        "DestinationCode": "10589",
        "EstimatedArrival": "2017-08-02T15:38:49+08:00",
        "Latitude": "1.3438965",
        "Longitude": "103.96221783333333",
        "VisitNumber": "1",
        "Load": "SEA",
        "Feature": "WAB",
        "Type": "DD"
      },
      "NextBus2": {
        "OriginCode": "99009",
        "DestinationCode": "10589",
        "EstimatedArrival": "2017-08-02T15:50:58+08:00",
        "Latitude": "1.3739711666666667",
        "Longitude": "103.97514966666667",
        "VisitNumber": "1",
        "Load": "SEA",
        "Feature": "WAB",
        "Type": "DD"
      },
      "NextBus3": {
        "OriginCode": "99009",
        "DestinationCode": "10589",
        "EstimatedArrival": "2017-08-02T15:57:50+08:00",
        "Latitude": "0",
        "Longitude": "0",
        "VisitNumber": "1",
        "Load": "SEA",
        "Feature": "WAB",
        "Type": "DD"
      }
    },
    {
      "ServiceNo": "24",
      "Operator": "SBST",
      "NextBus": {
        "OriginCode": "54009",
        "DestinationCode": "54009",
        "EstimatedArrival": "2017-08-02T15:40:56+08:00",
        "Latitude": "1.3445355",
        "Longitude": "103.968438",
        "VisitNumber": "1",
        "Load": "SDA",
        "Feature": "WAB",
        "Type": "SD"
      },
      "NextBus2": {
        "OriginCode": "54009",
        "DestinationCode": "54009",
        "EstimatedArrival": "2017-08-02T15:53:59+08:00",
        "Latitude": "1.359649",
        "Longitude": "103.99118333333334",
        "VisitNumber": "1",
        "Load": "SEA",
        "Feature": "WAB",
        "Type": "SD"
      },
      "NextBus3": {
        "OriginCode": "54009",
        "DestinationCode": "54009",
        "EstimatedArrival": "2017-08-02T16:06:00+08:00",
        "Latitude": "1.341015",
        "Longitude": "103.9712825",
        "VisitNumber": "1",
        "Load": "SEA",
        "Feature": "WAB",
        "Type": "SD"
      }
    }
  ]
}`,
		},
	}

	etas, err := client.GetBusArrival("", "")
	if err != nil {
		t.Fatalf("%+v", err)
	}

	expected := BusArrival{
		ODataMetadata: "http://datamall2.mytransport.sg/ltaodataservice/$metadata#BusArrivalv2/@Element",
		BusStopCode:   "96049",
		Services: []Service{
			{
				ServiceNo: "2",
				Operator:  "GAS",
				NextBus: ArrivingBus{
					OriginCode:       "99009",
					DestinationCode:  "10589",
					EstimatedArrival: mustParseRFC3339("2017-08-02T15:38:49+08:00"),
					Latitude:         1.3438965,
					Longitude:        103.96221783333333,
					VisitNumber:      1,
					Load:             "SEA",
					Feature:          "WAB",
					Type:             "DD",
				},
				NextBus2: ArrivingBus{
					OriginCode:       "99009",
					DestinationCode:  "10589",
					EstimatedArrival: mustParseRFC3339("2017-08-02T15:50:58+08:00"),
					Latitude:         1.3739711666666667,
					Longitude:        103.97514966666667,
					VisitNumber:      1,
					Load:             "SEA",
					Feature:          "WAB",
					Type:             "DD",
				},
				NextBus3: ArrivingBus{
					OriginCode:       "99009",
					DestinationCode:  "10589",
					EstimatedArrival: mustParseRFC3339("2017-08-02T15:57:50+08:00"),
					Latitude:         0,
					Longitude:        0,
					VisitNumber:      1,
					Load:             "SEA",
					Feature:          "WAB",
					Type:             "DD",
				},
			}, {
				ServiceNo: "24",
				Operator:  "SBST",
				NextBus: ArrivingBus{
					OriginCode:       "54009",
					DestinationCode:  "54009",
					EstimatedArrival: mustParseRFC3339("2017-08-02T15:40:56+08:00"),
					Latitude:         1.3445355,
					Longitude:        103.968438,
					VisitNumber:      1,
					Load:             "SDA",
					Feature:          "WAB",
					Type:             "SD",
				},
				NextBus2: ArrivingBus{
					OriginCode:       "54009",
					DestinationCode:  "54009",
					EstimatedArrival: mustParseRFC3339("2017-08-02T15:53:59+08:00"),
					Latitude:         1.359649,
					Longitude:        103.99118333333334,
					VisitNumber:      1,
					Load:             "SEA",
					Feature:          "WAB",
					Type:             "SD",
				},
				NextBus3: ArrivingBus{
					OriginCode:       "54009",
					DestinationCode:  "54009",
					EstimatedArrival: mustParseRFC3339("2017-08-02T16:06:00+08:00"),
					Latitude:         1.341015,
					Longitude:        103.9712825,
					VisitNumber:      1,
					Load:             "SEA",
					Feature:          "WAB",
					Type:             "SD",
				},
			},
		},
	}
	if !reflect.DeepEqual(expected, etas) {
		fmt.Printf("Expected:\n%#v\nActual:\n%#v", expected, etas)
		t.Fail()
	}

	t.Run("when datamall returns an error", func(t *testing.T) {
		client := APIClient{
			Client: &MockHTTPClient{
				StatusCode: 503,
			},
		}

		_, err := client.GetBusArrival("", "")
		assert.Error(t, err)

		expected := Error{
			StatusCode: 503,
		}
		assert.Equal(t, expected, err)
	})
}
