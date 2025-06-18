package fox_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/fox"
)

func TestEnergyMeterWithNoAPIKeyReadsCurrentWorkingParameters(t *testing.T) {
	t.Parallel()

	wantPath := "/00/get_current_parameters"

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != wantPath {
			t.Errorf("invalid path:\n%s\n", cmp.Diff(wantPath, r.URL.Path))
		}
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(respBodyCurrentParameters))
	}))
	defer ts.Close()

	em := fox.NewEnergyMeter(ts.URL)
	em.HTTPClient = ts.Client()

	got, err := em.CurrentParameters()
	if err != nil {
		t.Fatal(err)
	}

	want := fox.EnergyMeterStats{
		Status:        "ok",
		Voltage:       "247.3",
		Current:       "0.00",
		PowerActive:   "0.0",
		PowerReactive: "0.0",
		Frequency:     "50.18",
		PowerFactor:   "1.00",
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestEnergyMeterWithAPIKeyReadsCurrentWorkingParameters(t *testing.T) {
	t.Parallel()

	energyMeterAPIKey := "DS123QWES12"

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wantPath := fmt.Sprintf("/%s/get_current_parameters", energyMeterAPIKey)
		if r.URL.Path != wantPath {
			t.Errorf("invalid path:\n%s\n", cmp.Diff(wantPath, r.URL.Path))
		}
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(respBodyCurrentParameters))
	}))
	defer ts.Close()

	em := fox.NewEnergyMeter(ts.URL)
	em.HTTPClient = ts.Client()
	em.APIKey = energyMeterAPIKey

	got, err := em.CurrentParameters()
	if err != nil {
		t.Fatal(err)
	}

	want := fox.EnergyMeterStats{
		Status:        "ok",
		Voltage:       "247.3",
		Current:       "0.00",
		PowerActive:   "0.0",
		PowerReactive: "0.0",
		Frequency:     "50.18",
		PowerFactor:   "1.00",
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestEnergyMeterWithNoAPIKeyReadsTotalEnergy(t *testing.T) {
	t.Parallel()

	wantPath := "/00/get_total_energy"

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != wantPath {
			t.Errorf("invalid path:\n%s\n", cmp.Diff(wantPath, r.URL.Path))
		}
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(respBodyEnergyTotal))
	}))
	defer ts.Close()

	em := fox.NewEnergyMeter(ts.URL)
	em.HTTPClient = ts.Client()

	got, err := em.Total()
	if err != nil {
		t.Fatal(err)
	}

	want := fox.EnergyTotal{
		Status:               "ok",
		ActiveEnergy:         "000",
		ReactiveEnergy:       "000",
		ActiveEnergyImport:   "000",
		ReactiveEnergyImport: "000",
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestEnergyMeterWithAPIKeyReadsTotalEnergy(t *testing.T) {
	t.Parallel()

	energyMeterAPIKey := "DS123QWES12"

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wantPath := fmt.Sprintf("/%s/get_total_energy", energyMeterAPIKey)
		if r.URL.Path != wantPath {
			t.Errorf("invalid path:\n%s\n", cmp.Diff(wantPath, r.URL.Path))
		}
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(respBodyEnergyTotal))
	}))
	defer ts.Close()

	em := fox.NewEnergyMeter(ts.URL)
	em.HTTPClient = ts.Client()
	em.APIKey = energyMeterAPIKey

	got, err := em.Total()
	if err != nil {
		t.Fatal(err)
	}

	want := fox.EnergyTotal{
		Status:               "ok",
		ActiveEnergy:         "000",
		ReactiveEnergy:       "000",
		ActiveEnergyImport:   "000",
		ReactiveEnergyImport: "000",
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

var (
	respBodyCurrentParameters = `{
	"status":"ok",
	"voltage":"247.3",
	"current":"0.00",
	"power_active":"0.0",
	"power_reactive":"0.0",
	"frequency":"50.18",
	"power_factor":"1.00"
}`

	respBodyEnergyTotal = `{
	"status":"ok",
	"active_energy":"000",
	"reactive_energy":"000",
	"active_energy_import":"000",
	"reactive_energy_import":"000"
}`
)
