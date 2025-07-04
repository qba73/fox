package fox

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// EnergyMeterStats represents current data
// recorded by the Energy Meter
type EnergyMeterStats struct {
	Status        string `json:"status"`
	Voltage       string `json:"voltage"`
	Current       string `json:"current"`
	PowerActive   string `json:"power_active"`
	PowerReactive string `json:"power_reactive"`
	Frequency     string `json:"frequency"`
	PowerFactor   string `json:"power_factor"`
}

// EnergyTotal represents cumulative energy
// recorded by the meter.
type EnergyTotal struct {
	Status               string `json:"status"`
	ActiveEnergy         string `json:"active_energy"`
	ReactiveEnergy       string `json:"reactive_energy"`
	ActiveEnergyImport   string `json:"active_energy_import"`
	ReactiveEnergyImport string `json:"reactive_energy_import"`
}

// EnergyMeter represents a Client for Fox EnergyMeter.
type EnergyMeter struct {
	Name    string
	IP      string
	BaseURL string
	APIKey  string
}

// NewEnergyMeter creates a client for Fox Energy Meter.
// It takes a base URL where the IP represents Meter's IP address
// in the local network.
func NewEnergyMeter(baseURL string) *EnergyMeter {
	em := EnergyMeter{
		BaseURL: baseURL,
		APIKey:  "00",
	}
	return &em
}

type Client struct {
	HTTPClient *http.Client
}

func NewClient() *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
	}
}

func (c *Client) EnergyMeterCurrentStatus(baseURL string) (EnergyMeterStats, error) {
	url := fmt.Sprintf("%s/00/get_current_parameters", baseURL)
	var e EnergyMeterStats
	if err := c.get(context.Background(), url, &e); err != nil {
		return EnergyMeterStats{}, err
	}
	return e, nil
}

func (c *Client) EnergyMeterTotal(baseURL string) (EnergyTotal, error) {
	url := fmt.Sprintf("%s/00/get_total_energy", baseURL)
	var et EnergyTotal
	if err := c.get(context.Background(), url, &et); err != nil {
		return EnergyTotal{}, err
	}
	return et, nil
}

func (c *Client) get(ctx context.Context, url string, data any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("fox: creating HTTP request: %w", err)
	}
	res, err := c.HTTPClient.Do(req)
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

// GetEnergyMeterStatus takes a string representing Meter's IP address
// on a local network and returns current statistics or an error.
//
// It uses default Fox client and default HTTP Client.
func GetEnergyMeterStatus(ip string) (EnergyMeterStats, error) {
	return NewClient().EnergyMeterCurrentStatus("http://" + ip)
}

// GetEnergyTotal takes a string representing Meter's IP address
// on a local network and returns total energy reading or an error.
//
// It uses default Fox client and default HTTP Client.
func GetEnergyTotal(ip string) (EnergyTotal, error) {
	return NewClient().EnergyMeterTotal("http://" + ip)
}
