package main

import ("testing"
)

var airportTests = []struct {
	airportICAOCode string
	airportIATACode string

	{
		airportICAOCode="LEMD",
		airportIATACode = "MAD"
	},
	{
		airportICAOCode="KJFK"
		airportIATACode = "JFK"
	},
	{
		airportICAOCode="WSSS"
		airportIATACode = "SIN"
	},
	{
		airportICAOCode="YSSY"
		airportIATACode = "SYD"
	},
	{
		airportICAOCode="FACT"
		airportIATACode = "CPT"
	},
}

func TestValidateAirportCode (t *testing.T) {
	for i,test := range airportTests {
		test.airportIATACode.Len = 4;
	}
}
