package main

import ("testing"
)

var airportTests = []struct {
	airportICAOCode string
	airportIATACode string

	}{
	 {
		airportICAOCode: "LEMD",
		airportIATACode: "MAD",
	},
	{
		airportICAOCode: "KJFK",
		airportIATACode: "JFK",
	},
	{
		airportICAOCode: "WSSS",
		airportIATACode: "SIN",
	},
	{
		airportICAOCode: "YSSY",
		airportIATACode: "SYD",
	},
	{
		airportICAOCode: "FACT",
		airportIATACode: "CPT",
	},
}

/*Test to make sure Airport code validation works*/
func TestValidateAirportCode (t *testing.T) {
	for i,test := range airportTests {
		var actualICAO int = validateAirportCode(test.airportICAOCode)
		var actualIATA int = validateAirportCode(test.airportIATACode)

		if actualICAO != 1 {
			t.Error("Expected 1, got ", actualICAO, "Test %d", i)
		}
		if actualIATA != 0 {
			t.Error("Expected 0, got ", actualIATA, "Test %d", i)
		}

	}
}

func TestGetAirportData (t *testing.T){
	for i,test := range airportTests {
		_,resultICAO := getAirportData(test.airportICAOCode,1)
		_,resultIATA := getAirportData(test.airportIATACode,0)
	}
}