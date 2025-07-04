package main

import (
	"fmt"

	"github.com/qba73/fox"
)

func main() {
	// Create FOX Client

	fc := fox.NewClient()

	// Example: Read Energy Meter Stats

	status, err := fc.EnergyMeterCurrentStatus("http://192.168.50.122")
	if err != nil {
		// handle error
	}
	fmt.Printf("%+v\n", status)
	// {Status:ok Voltage:245.6 Current:0.00 PowerActive:0.0 PowerReactive:0.0 Frequency:50.04 PowerFactor:1.00}

	// Example: Read Energy Total

	total, err := fc.EnergyMeterTotal("http://192.168.50.122")
	if err != nil {
		// handle error
	}
	fmt.Printf("%+v\n", total)
	// {Status:ok ActiveEnergy:000 ReactiveEnergy:000 ActiveEnergyImport:000 ReactiveEnergyImport:000}

}
