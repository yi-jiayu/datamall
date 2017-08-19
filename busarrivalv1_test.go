package datamall

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestAPIClient_GetBusArrivalV1(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.Write([]byte(`{"odata.metadata":"http://datamall2.mytransport.sg/ltaodataservice/$metadata#BusArrivalV2/@Element","BusStopID":"96049","Services":[{"ServiceNo":"2","Status":"In Operation","Operator":"GAS","OriginatingID":"99009","TerminatingID":"10589","NextBus":{"EstimatedArrival":"2017-05-20T06:56:49+00:00","Latitude":"1.350827","Longitude":"103.96423366666667","VisitNumber":"1","Load":"Seats Available","Feature":"WAB"},"SubsequentBus":{"EstimatedArrival":"2017-05-20T07:04:02+00:00","Latitude":"1.3684493333333334","Longitude":"103.97785566666667","VisitNumber":"1","Load":"Seats Available","Feature":"WAB"},"SubsequentBus3":{"EstimatedArrival":"2017-05-20T07:14:44+00:00","Latitude":"0","Longitude":"0","VisitNumber":"1","Load":"Seats Available","Feature":"WAB"}},{"ServiceNo":"24","Status":"In Operation","Operator":"SBST","OriginatingID":"54009","TerminatingID":"54009","NextBus":{"EstimatedArrival":"2017-05-20T06:58:04+00:00","Latitude":"1.3394471666666667","Longitude":"103.97383066666667","VisitNumber":"1","Load":"Standing Available","Feature":"WAB"},"SubsequentBus":{"EstimatedArrival":"2017-05-20T07:01:14+00:00","Latitude":"1.3445005","Longitude":"103.98324616666666","VisitNumber":"1","Load":"Seats Available","Feature":"WAB"},"SubsequentBus3":{"EstimatedArrival":"2017-05-20T07:07:23+00:00","Latitude":"1.3560058333333334","Longitude":"103.98930616666667","VisitNumber":"1","Load":"Seats Available","Feature":"WAB"}},{"ServiceNo":"5","Status":"In Operation","Operator":"SBST","OriginatingID":"77009","TerminatingID":"10009","NextBus":{"EstimatedArrival":"2017-05-20T06:53:53+00:00","Latitude":"1.3416796666666666","Longitude":"103.96159316666666","VisitNumber":"1","Load":"Seats Available","Feature":""},"SubsequentBus":{"EstimatedArrival":"2017-05-20T07:07:07+00:00","Latitude":"1.3731878333333334","Longitude":"103.96395","VisitNumber":"1","Load":"Seats Available","Feature":"WAB"},"SubsequentBus3":{"EstimatedArrival":"2017-05-20T07:18:43+00:00","Latitude":"0","Longitude":"0","VisitNumber":"1","Load":"Seats Available","Feature":"WAB"}}]}`))
	}))
	defer ts.Close()

	client := APIClient{
		Endpoint: ts.URL,
		Client:   http.DefaultClient,
	}

	t.Run("BusStopID only", func(t *testing.T) {
		etas, err := client.GetBusArrivalV1("", nil)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := BusArrivalV1{
			OdataMetadata: "http://datamall2.mytransport.sg/ltaodataservice/$metadata#BusArrivalV2/@Element",
			BusStopID:     "96049",
			Services: []ServiceV1{
				{
					ServiceNo:     "2",
					Status:        "In Operation",
					Operator:      "GAS",
					OriginatingID: "99009",
					TerminatingID: "10589",
					NextBus: ArrivingBusV1{
						EstimatedArrival: "2017-05-20T06:56:49+00:00",
						Latitude:         "1.350827",
						Longitude:        "103.96423366666667",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus: ArrivingBusV1{
						EstimatedArrival: "2017-05-20T07:04:02+00:00",
						Latitude:         "1.3684493333333334",
						Longitude:        "103.97785566666667",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus3: ArrivingBusV1{
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
					NextBus: ArrivingBusV1{
						EstimatedArrival: "2017-05-20T06:58:04+00:00",
						Latitude:         "1.3394471666666667",
						Longitude:        "103.97383066666667",
						VisitNumber:      "1",
						Load:             "Standing Available",
						Feature:          "WAB",
					},
					SubsequentBus: ArrivingBusV1{
						EstimatedArrival: "2017-05-20T07:01:14+00:00",
						Latitude:         "1.3445005",
						Longitude:        "103.98324616666666",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus3: ArrivingBusV1{
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
					NextBus: ArrivingBusV1{
						EstimatedArrival: "2017-05-20T06:53:53+00:00",
						Latitude:         "1.3416796666666666",
						Longitude:        "103.96159316666666",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "",
					},
					SubsequentBus: ArrivingBusV1{
						EstimatedArrival: "2017-05-20T07:07:07+00:00",
						Latitude:         "1.3731878333333334",
						Longitude:        "103.96395",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus3: ArrivingBusV1{
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
		etas, err := client.GetBusArrivalV1("", &GetBusArrivalV1Options{ServiceNo: "24"})
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := BusArrivalV1{
			OdataMetadata: "http://datamall2.mytransport.sg/ltaodataservice/$metadata#BusArrivalV2/@Element",
			BusStopID:     "96049",
			Services: []ServiceV1{
				{
					ServiceNo:     "2",
					Status:        "In Operation",
					Operator:      "GAS",
					OriginatingID: "99009",
					TerminatingID: "10589",
					NextBus: ArrivingBusV1{
						EstimatedArrival: "2017-05-20T06:56:49+00:00",
						Latitude:         "1.350827",
						Longitude:        "103.96423366666667",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus: ArrivingBusV1{
						EstimatedArrival: "2017-05-20T07:04:02+00:00",
						Latitude:         "1.3684493333333334",
						Longitude:        "103.97785566666667",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus3: ArrivingBusV1{
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
					NextBus: ArrivingBusV1{
						EstimatedArrival: "2017-05-20T06:58:04+00:00",
						Latitude:         "1.3394471666666667",
						Longitude:        "103.97383066666667",
						VisitNumber:      "1",
						Load:             "Standing Available",
						Feature:          "WAB",
					},
					SubsequentBus: ArrivingBusV1{
						EstimatedArrival: "2017-05-20T07:01:14+00:00",
						Latitude:         "1.3445005",
						Longitude:        "103.98324616666666",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus3: ArrivingBusV1{
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
					NextBus: ArrivingBusV1{
						EstimatedArrival: "2017-05-20T06:53:53+00:00",
						Latitude:         "1.3416796666666666",
						Longitude:        "103.96159316666666",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "",
					},
					SubsequentBus: ArrivingBusV1{
						EstimatedArrival: "2017-05-20T07:07:07+00:00",
						Latitude:         "1.3731878333333334",
						Longitude:        "103.96395",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus3: ArrivingBusV1{
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
		etas, err := client.GetBusArrivalV1("", &GetBusArrivalV1Options{ServiceNo: "24", SST: true})
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := BusArrivalV1{
			OdataMetadata: "http://datamall2.mytransport.sg/ltaodataservice/$metadata#BusArrivalV2/@Element",
			BusStopID:     "96049",
			Services: []ServiceV1{
				{
					ServiceNo:     "2",
					Status:        "In Operation",
					Operator:      "GAS",
					OriginatingID: "99009",
					TerminatingID: "10589",
					NextBus: ArrivingBusV1{
						EstimatedArrival: "2017-05-20T06:56:49+00:00",
						Latitude:         "1.350827",
						Longitude:        "103.96423366666667",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus: ArrivingBusV1{
						EstimatedArrival: "2017-05-20T07:04:02+00:00",
						Latitude:         "1.3684493333333334",
						Longitude:        "103.97785566666667",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus3: ArrivingBusV1{
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
					NextBus: ArrivingBusV1{
						EstimatedArrival: "2017-05-20T06:58:04+00:00",
						Latitude:         "1.3394471666666667",
						Longitude:        "103.97383066666667",
						VisitNumber:      "1",
						Load:             "Standing Available",
						Feature:          "WAB",
					},
					SubsequentBus: ArrivingBusV1{
						EstimatedArrival: "2017-05-20T07:01:14+00:00",
						Latitude:         "1.3445005",
						Longitude:        "103.98324616666666",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus3: ArrivingBusV1{
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
					NextBus: ArrivingBusV1{
						EstimatedArrival: "2017-05-20T06:53:53+00:00",
						Latitude:         "1.3416796666666666",
						Longitude:        "103.96159316666666",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "",
					},
					SubsequentBus: ArrivingBusV1{
						EstimatedArrival: "2017-05-20T07:07:07+00:00",
						Latitude:         "1.3731878333333334",
						Longitude:        "103.96395",
						VisitNumber:      "1",
						Load:             "Seats Available",
						Feature:          "WAB",
					},
					SubsequentBus3: ArrivingBusV1{
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
