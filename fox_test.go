package fox_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/fox"
)

func TestEnergyMeterWithNoAPIKeySendsCurrentWorkingParameters(t *testing.T) {
	t.Parallel()

	wantPath := "/00/get_current_parameters"

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != wantPath {
			t.Errorf("invalid path:\n%s\n", cmp.Diff(wantPath, r.URL.Path))
		}
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(respBodyEnergy))
	}))
	defer ts.Close()

	em := fox.NewEnergyMeter(ts.URL)
	em.HTTPClient = ts.Client()

	got, err := em.CurrentParams()
	if err != nil {
		t.Fatal(err)
	}

	want := fox.EnergyMeterStats{
		Status:        "ok",
		Voltage:       "247.3",
		Current:       "0.00",
		PowerActive:   "",
		PowerReactive: "",
		Frequency:     "50.18",
		PowerFactor:   "",
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestEnergyMeterWithAPIKeySendsCurrentWorkingParameters(t *testing.T) {
	t.Parallel()

	energyMeterAPIKey := "DS123QWES12"

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wantPath := fmt.Sprintf("/%s/get_current_parameters", energyMeterAPIKey)
		if r.URL.Path != wantPath {
			t.Errorf("invalid path:\n%s\n", cmp.Diff(wantPath, r.URL.Path))
		}
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(respBodyEnergy))
	}))
	defer ts.Close()

	em := fox.NewEnergyMeter(ts.URL)
	em.HTTPClient = ts.Client()
	em.APIKey = energyMeterAPIKey

	got, err := em.CurrentParams()
	if err != nil {
		t.Fatal(err)
	}

	want := fox.EnergyMeterStats{
		Status:        "ok",
		Voltage:       "247.3",
		Current:       "0.00",
		PowerActive:   "",
		PowerReactive: "",
		Frequency:     "50.18",
		PowerFactor:   "",
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

var respBodyEnergy = `{
	"status":"ok",
	"voltage":"247.3",
	"current":"0.00",
	"power_active":"0.0",
	"power_reactive":"0.0",
	"frequency":"50.18",
	"power_factor":"1.00"
}`
