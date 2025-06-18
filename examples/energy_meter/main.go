package main

import (
	"fmt"

	"github.com/qba73/fox"
)

func main() {
	// =========================================================
	// Example for devices with configured API Key
	// =========================================================

	energyMeter := fox.NewEnergyMeter("http://192.168.50.122")
	energyMeter.APIKey = "f01333d9c779eaec58be22c6a6"

	cp, err := energyMeter.CurrentParameters()
	if err != nil {
		// handle error
	}
	fmt.Printf("%+v\n", cp)
	// {Status:ok Voltage:245.6 Current:0.00 PowerActive:0.0 PowerReactive:0.0 Frequency:50.04 PowerFactor:1.00}

	// =========================================================
	// Example for devices without configured API Key
	// =========================================================

	energy, err := fox.GetEnergyStats("192.168.50.122")
	if err != nil {
		// handle error
	}
	fmt.Printf("%+v\n", energy)
	// {Status:ok Voltage:240.8 Current:0.00 PowerActive:0.0 PowerReactive:0.0 Frequency:50.12 PowerFactor:1.00}

	total, err := fox.GetEnergyTotal("192.168.50.122")
	if err != nil {
		// handle error
	}
	fmt.Printf("%+v\n", total)
	// {Status:ok ActiveEnergy:000 ReactiveEnergy:000 ActiveEnergyImport:000 ReactiveEnergyImport:000}

}
