package datamall

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestAPIClient_GetBusArrival(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.Write([]byte(`{"odata.metadata":"http://datamall2.mytransport.sg/ltaodataservice/$metadata#BusArrival/@Element","BusStopID":"96049","Services":[{"ServiceNo":"2","Status":"In Operation","Operator":"GAS","OriginatingID":"99009","TerminatingID":"10589","NextBus":{"EstimatedArrival":"2017-05-20T06:56:49+00:00","Latitude":"1.350827","Longitude":"103.96423366666667","VisitNumber":"1","Load":"Seats Available","Feature":"WAB"},"SubsequentBus":{"EstimatedArrival":"2017-05-20T07:04:02+00:00","Latitude":"1.3684493333333334","Longitude":"103.97785566666667","VisitNumber":"1","Load":"Seats Available","Feature":"WAB"},"SubsequentBus3":{"EstimatedArrival":"2017-05-20T07:14:44+00:00","Latitude":"0","Longitude":"0","VisitNumber":"1","Load":"Seats Available","Feature":"WAB"}},{"ServiceNo":"24","Status":"In Operation","Operator":"SBST","OriginatingID":"54009","TerminatingID":"54009","NextBus":{"EstimatedArrival":"2017-05-20T06:58:04+00:00","Latitude":"1.3394471666666667","Longitude":"103.97383066666667","VisitNumber":"1","Load":"Standing Available","Feature":"WAB"},"SubsequentBus":{"EstimatedArrival":"2017-05-20T07:01:14+00:00","Latitude":"1.3445005","Longitude":"103.98324616666666","VisitNumber":"1","Load":"Seats Available","Feature":"WAB"},"SubsequentBus3":{"EstimatedArrival":"2017-05-20T07:07:23+00:00","Latitude":"1.3560058333333334","Longitude":"103.98930616666667","VisitNumber":"1","Load":"Seats Available","Feature":"WAB"}},{"ServiceNo":"5","Status":"In Operation","Operator":"SBST","OriginatingID":"77009","TerminatingID":"10009","NextBus":{"EstimatedArrival":"2017-05-20T06:53:53+00:00","Latitude":"1.3416796666666666","Longitude":"103.96159316666666","VisitNumber":"1","Load":"Seats Available","Feature":""},"SubsequentBus":{"EstimatedArrival":"2017-05-20T07:07:07+00:00","Latitude":"1.3731878333333334","Longitude":"103.96395","VisitNumber":"1","Load":"Seats Available","Feature":"WAB"},"SubsequentBus3":{"EstimatedArrival":"2017-05-20T07:18:43+00:00","Latitude":"0","Longitude":"0","VisitNumber":"1","Load":"Seats Available","Feature":"WAB"}}]}`))
	}))
	defer ts.Close()

	client := APIClient{
		Endpoint: ts.URL,
		Client:   http.DefaultClient,
	}

	t.Run("BusStopID only", func(t *testing.T) {
		etas, err := client.GetBusArrival("", nil)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := BusArrival{
			OdataMetadata: "http://datamall2.mytransport.sg/ltaodataservice/$metadata#BusArrival/@Element",
			BusStopID:     "96049",
			Services: []Service{
				{
					ServiceNo:     "2",
					Status:        "In Operation",
					Operator:      "GAS",
					OriginatingID: "99009",
					TerminatingID: "10589",
					NextBus: ArrivingBus{
						EstimatedArrival: "2017-05-20T06:56:49+00:00",
						Latitude:         "1.350827",
						Longitude:        "103.96423366666667",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus: ArrivingBus{
						EstimatedArrival: "2017-05-20T07:04:02+00:00",
						Latitude:         "1.3684493333333334",
						Longitude:        "103.97785566666667",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus3: ArrivingBus{
						EstimatedArrival: "2017-05-20T07:14:44+00:00",
						Latitude:         "0",
						Longitude:        "0",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
				},
				{
					ServiceNo:     "24",
					Status:        "In Operation",
					Operator:      "SBST",
					OriginatingID: "54009",
					TerminatingID: "54009",
					NextBus: ArrivingBus{
						EstimatedArrival: "2017-05-20T06:58:04+00:00",
						Latitude:         "1.3394471666666667",
						Longitude:        "103.97383066666667",
						VisitNumber:      "1",
						Load:             "Standing Available",
						Feature:          "WAB",
					},
					SubsequentBus: ArrivingBus{
						EstimatedArrival: "2017-05-20T07:01:14+00:00",
						Latitude:         "1.3445005",
						Longitude:        "103.98324616666666",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus3: ArrivingBus{
						EstimatedArrival: "2017-05-20T07:07:23+00:00",
						Latitude:         "1.3560058333333334",
						Longitude:        "103.98930616666667",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
				},
				{
					ServiceNo:     "5",
					Status:        "In Operation",
					Operator:      "SBST",
					OriginatingID: "77009",
					TerminatingID: "10009",
					NextBus: ArrivingBus{
						EstimatedArrival: "2017-05-20T06:53:53+00:00",
						Latitude:         "1.3416796666666666",
						Longitude:        "103.96159316666666",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "",
					},
					SubsequentBus: ArrivingBus{
						EstimatedArrival: "2017-05-20T07:07:07+00:00",
						Latitude:         "1.3731878333333334",
						Longitude:        "103.96395",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus3: ArrivingBus{
						EstimatedArrival: "2017-05-20T07:18:43+00:00",
						Latitude:         "0",
						Longitude:        "0",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
				},
			},
		}

		if !reflect.DeepEqual(expected, etas) {
			t.Fail()
		}
	})
	t.Run("BusStopID and ServiceNo", func(t *testing.T) {
		etas, err := client.GetBusArrival("", &GetBusArrivalOptions{ServiceNo: "24"})
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := BusArrival{
			OdataMetadata: "http://datamall2.mytransport.sg/ltaodataservice/$metadata#BusArrival/@Element",
			BusStopID:     "96049",
			Services: []Service{
				{
					ServiceNo:     "2",
					Status:        "In Operation",
					Operator:      "GAS",
					OriginatingID: "99009",
					TerminatingID: "10589",
					NextBus: ArrivingBus{
						EstimatedArrival: "2017-05-20T06:56:49+00:00",
						Latitude:         "1.350827",
						Longitude:        "103.96423366666667",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus: ArrivingBus{
						EstimatedArrival: "2017-05-20T07:04:02+00:00",
						Latitude:         "1.3684493333333334",
						Longitude:        "103.97785566666667",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus3: ArrivingBus{
						EstimatedArrival: "2017-05-20T07:14:44+00:00",
						Latitude:         "0",
						Longitude:        "0",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
				},
				{
					ServiceNo:     "24",
					Status:        "In Operation",
					Operator:      "SBST",
					OriginatingID: "54009",
					TerminatingID: "54009",
					NextBus: ArrivingBus{
						EstimatedArrival: "2017-05-20T06:58:04+00:00",
						Latitude:         "1.3394471666666667",
						Longitude:        "103.97383066666667",
						VisitNumber:      "1",
						Load:             "Standing Available",
						Feature:          "WAB",
					},
					SubsequentBus: ArrivingBus{
						EstimatedArrival: "2017-05-20T07:01:14+00:00",
						Latitude:         "1.3445005",
						Longitude:        "103.98324616666666",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus3: ArrivingBus{
						EstimatedArrival: "2017-05-20T07:07:23+00:00",
						Latitude:         "1.3560058333333334",
						Longitude:        "103.98930616666667",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
				},
				{
					ServiceNo:     "5",
					Status:        "In Operation",
					Operator:      "SBST",
					OriginatingID: "77009",
					TerminatingID: "10009",
					NextBus: ArrivingBus{
						EstimatedArrival: "2017-05-20T06:53:53+00:00",
						Latitude:         "1.3416796666666666",
						Longitude:        "103.96159316666666",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "",
					},
					SubsequentBus: ArrivingBus{
						EstimatedArrival: "2017-05-20T07:07:07+00:00",
						Latitude:         "1.3731878333333334",
						Longitude:        "103.96395",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus3: ArrivingBus{
						EstimatedArrival: "2017-05-20T07:18:43+00:00",
						Latitude:         "0",
						Longitude:        "0",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
				},
			},
		}

		if !reflect.DeepEqual(expected, etas) {
			t.Fail()
		}
	})
	t.Run("BusStopID, ServiceNo and SST", func(t *testing.T) {
		etas, err := client.GetBusArrival("", &GetBusArrivalOptions{ServiceNo: "24", SST: true})
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := BusArrival{
			OdataMetadata: "http://datamall2.mytransport.sg/ltaodataservice/$metadata#BusArrival/@Element",
			BusStopID:     "96049",
			Services: []Service{
				{
					ServiceNo:     "2",
					Status:        "In Operation",
					Operator:      "GAS",
					OriginatingID: "99009",
					TerminatingID: "10589",
					NextBus: ArrivingBus{
						EstimatedArrival: "2017-05-20T06:56:49+00:00",
						Latitude:         "1.350827",
						Longitude:        "103.96423366666667",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus: ArrivingBus{
						EstimatedArrival: "2017-05-20T07:04:02+00:00",
						Latitude:         "1.3684493333333334",
						Longitude:        "103.97785566666667",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus3: ArrivingBus{
						EstimatedArrival: "2017-05-20T07:14:44+00:00",
						Latitude:         "0",
						Longitude:        "0",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
				},
				{
					ServiceNo:     "24",
					Status:        "In Operation",
					Operator:      "SBST",
					OriginatingID: "54009",
					TerminatingID: "54009",
					NextBus: ArrivingBus{
						EstimatedArrival: "2017-05-20T06:58:04+00:00",
						Latitude:         "1.3394471666666667",
						Longitude:        "103.97383066666667",
						VisitNumber:      "1",
						Load:             "Standing Available",
						Feature:          "WAB",
					},
					SubsequentBus: ArrivingBus{
						EstimatedArrival: "2017-05-20T07:01:14+00:00",
						Latitude:         "1.3445005",
						Longitude:        "103.98324616666666",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus3: ArrivingBus{
						EstimatedArrival: "2017-05-20T07:07:23+00:00",
						Latitude:         "1.3560058333333334",
						Longitude:        "103.98930616666667",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
				},
				{
					ServiceNo:     "5",
					Status:        "In Operation",
					Operator:      "SBST",
					OriginatingID: "77009",
					TerminatingID: "10009",
					NextBus: ArrivingBus{
						EstimatedArrival: "2017-05-20T06:53:53+00:00",
						Latitude:         "1.3416796666666666",
						Longitude:        "103.96159316666666",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "",
					},
					SubsequentBus: ArrivingBus{
						EstimatedArrival: "2017-05-20T07:07:07+00:00",
						Latitude:         "1.3731878333333334",
						Longitude:        "103.96395",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus3: ArrivingBus{
						EstimatedArrival: "2017-05-20T07:18:43+00:00",
						Latitude:         "0",
						Longitude:        "0",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
				},
			},
		}

		if !reflect.DeepEqual(expected, etas) {
			t.Fail()
		}
	})
}

func TestAPIClient_GetBusArrivalV2(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.Write([]byte(`{"odata.metadata":"http://datamall2.mytransport.sg/ltaodataservice/$metadata#BusArrivalv2/@Element","BusStopCode":"96049","Services":[{"ServiceNo":"2","Operator":"GAS","NextBus":{"OriginCode":"99009","DestinationCode":"10589","EstimatedArrival":"2017-08-02T15:38:49+08:00","Latitude":"1.3438965","Longitude":"103.96221783333333","VisitNumber":"1","Load":"SEA","Feature":"WAB","Type":"DD"},"NextBus2":{"OriginCode":"99009","DestinationCode":"10589","EstimatedArrival":"2017-08-02T15:50:58+08:00","Latitude":"1.3739711666666667","Longitude":"103.97514966666667","VisitNumber":"1","Load":"SEA","Feature":"WAB","Type":"DD"},"NextBus3":{"OriginCode":"99009","DestinationCode":"10589","EstimatedArrival":"2017-08-02T15:57:50+08:00","Latitude":"0","Longitude":"0","VisitNumber":"1","Load":"SEA","Feature":"WAB","Type":"DD"}},{"ServiceNo":"24","Operator":"SBST","NextBus":{"OriginCode":"54009","DestinationCode":"54009","EstimatedArrival":"2017-08-02T15:40:56+08:00","Latitude":"1.3445355","Longitude":"103.968438","VisitNumber":"1","Load":"SDA","Feature":"WAB","Type":"SD"},"NextBus2":{"OriginCode":"54009","DestinationCode":"54009","EstimatedArrival":"2017-08-02T15:53:59+08:00","Latitude":"1.359649","Longitude":"103.99118333333334","VisitNumber":"1","Load":"SEA","Feature":"WAB","Type":"SD"},"NextBus3":{"OriginCode":"54009","DestinationCode":"54009","EstimatedArrival":"2017-08-02T16:06:00+08:00","Latitude":"1.341015","Longitude":"103.9712825","VisitNumber":"1","Load":"SEA","Feature":"WAB","Type":"SD"}}]}`))
	}))
	defer ts.Close()

	client := APIClient{
		Endpoint: ts.URL,
		Client:   http.DefaultClient,
	}

	etas, err := client.GetBusArrivalV2("", "")
	if err != nil {
		t.Fatalf("%+v", err)
	}

	expected := BusArrival{
		OdataMetadata: "http://datamall2.mytransport.sg/ltaodataservice/$metadata#BusArrivalv2/@Element",
		BusStopCode:   "96049",
		Services: []Service{
			{
				ServiceNo:     "2",
				Operator:      "GAS",
				NextBus: ArrivingBus{
					OriginCode:       "99009",
					DestinationCode:  "10589",
					EstimatedArrival: "2017-08-02T15:38:49+08:00",
					Latitude:         "1.3438965",
					Longitude:        "103.96221783333333",
					VisitNumber:      "1",
					Load:             "SEA",
					Feature:          "WAB",
					Type:             "DD",
				},
				NextBus2: ArrivingBus{
					OriginCode:       "99009",
					DestinationCode:  "10589",
					EstimatedArrival: "2017-08-02T15:50:58+08:00",
					Latitude:         "1.3739711666666667",
					Longitude:        "103.97514966666667",
					VisitNumber:      "1",
					Load:             "SEA",
					Feature:          "WAB",
					Type:             "DD",
				},
				NextBus3: ArrivingBus{
					OriginCode:       "99009",
					DestinationCode:  "10589",
					EstimatedArrival: "2017-08-02T15:57:50+08:00",
					Latitude:         "0",
					Longitude:        "0",
					VisitNumber:      "1",
					Load:             "SEA",
					Feature:          "WAB",
					Type:             "DD",
				},
			}, {
				ServiceNo:     "24",
				Operator:      "SBST",
				NextBus: ArrivingBus{
					OriginCode:       "54009",
					DestinationCode:  "54009",
					EstimatedArrival: "2017-08-02T15:40:56+08:00",
					Latitude:         "1.3445355",
					Longitude:        "103.968438",
					VisitNumber:      "1",
					Load:             "SDA",
					Feature:          "WAB",
					Type:             "SD",
				},
				NextBus2: ArrivingBus{
					OriginCode:       "54009",
					DestinationCode:  "54009",
					EstimatedArrival: "2017-08-02T15:53:59+08:00",
					Latitude:         "1.359649",
					Longitude:        "103.99118333333334",
					VisitNumber:      "1",
					Load:             "SEA",
					Feature:          "WAB",
					Type:             "SD",
				},
				NextBus3: ArrivingBus{
					OriginCode:       "54009",
					DestinationCode:  "54009",
					EstimatedArrival: "2017-08-02T16:06:00+08:00",
					Latitude:         "1.341015",
					Longitude:        "103.9712825",
					VisitNumber:      "1",
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
}
