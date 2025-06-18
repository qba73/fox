package fox

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// EnergyMeterStats represents current data recorded
// by the Energy Meter
type EnergyMeterStats struct {
	Status        string
	Voltage       string
	Current       string
	PowerActive   string
	PowerReactive string
	Frequency     string
	PowerFactor   string
}

// EnergyMeter represents a Client for Fox EnergyMeter.
type EnergyMeter struct {
	Name       string
	IP         string
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

// NewEnergyMeter creates a client for Fox Energy Meter.
// It takes a base URL where the IP represents Meter's IP address
// in the local network.
func NewEnergyMeter(baseURL string) *EnergyMeter {
	em := EnergyMeter{
		BaseURL:    baseURL,
		APIKey:     "00",
		HTTPClient: http.DefaultClient,
	}
	return &em
}

// CurrentParams returns current Meter Status or error.
func (em *EnergyMeter) CurrentParams() (EnergyMeterStats, error) {
	url := fmt.Sprintf("%s/%s/get_current_parameters", em.BaseURL, em.APIKey)
	var e EnergyMeterStats
	if err := em.get(context.Background(), url, &e); err != nil {
		return EnergyMeterStats{}, err
	}
	return e, nil
}

func (em *EnergyMeter) get(ctx context.Context, url string, data any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("fox: creating HTTP request: %w", err)
	}
	res, err := em.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("fox: sending GET request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("fox: got response code: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("fox: reading response body: %w", err)
	}

	if err := json.Unmarshal(body, data); err != nil {
		return fmt.Errorf("fox: unmarshaling response body: %w", err)
	}
	return nil
}

// EnergyStats takes a string representing Meter's IP address
// on a local network and returns current statistics or error.
//
// It uses default Energy Meter client and default HTTP Client.
func EnergyStats(IP string) (EnergyMeterStats, error) {
	return NewEnergyMeter("http://" + IP).CurrentParams()
}
