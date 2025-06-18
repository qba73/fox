package main

import (
	"fmt"

	"github.com/qba73/fox"
)

func main() {
	energyMeter := fox.NewEnergyMeter("http://192.168.50.120")
	energyMeter.APIKey = "f01333d9c779eaec58be22c6a6"

	cp, err := energyMeter.CurrentParameters()
	if err != nil {
		// handle error
	}

	fmt.Printf("%+v\n", cp)
	// {Status:ok Voltage:245.6 Current:0.00 PowerActive:0.0 PowerReactive:0.0 Frequency:50.04 PowerFactor:1.00}
}
