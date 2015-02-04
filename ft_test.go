package main

import (
	"testing"
	"strings"
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

var fakeAirportTests = []string{"HEEORF", "A"}

/*Test to make sure Airport code size works*/
func TestValidateAirportCode(t *testing.T) {
	for i, test := range airportTests {
		var actualICAO int32 = validateAirportCode(test.airportICAOCode)
		var actualIATA int32 = validateAirportCode(test.airportIATACode)

		if actualICAO != 1 {
			t.Error("Expected 1, got ", actualICAO, "Test %d", i)
		}
		if actualIATA != 0 {
			t.Error("Expected 0, got ", actualIATA, "Test %d", i)
		}

	}

	for i, ret := range fakeAirportTests {
		res := validateAirportCode(ret)
		if res != -2 {
			t.Error("Expected 2, got ", res, "Test %d", i)
		}
	}

}

/*Test to get airport indexes*/
func TestGetAirportIndex(t *testing.T) {
	for i, test := range airportTests {
		indexICAO := getAirportIndex(test.airportICAOCode)
		indexIATA := getAirportIndex(test.airportIATACode)

		if indexICAO == -2 || indexIATA == -2 {
			t.Error("Error! Got ", indexICAO, "Test %d", i)
		}
	}

	for i, ret := range fakeAirportTests {
		res := getAirportIndex(ret)
		if res != -2 {
			t.Error("Error! Got ", res, "Test %d", i)
		}
	}
}

/*Test to get airport data */
func TestGetAirportData(t *testing.T) {
	var airportInfo airportInformation
	airportInfo = getAirportData(getAirportIndex("KJFK"))
	if airportInfo.airportIndex != 3797 {
		t.Error("Error! Got %d but expected 3797", airportInfo.airportIndex)
	}
}

/*Test to get airport METAR*/
func TestGetAirportMETAR(t *testing.T){
	for _, test := range airportTests {
		result := getAirportMETAR(test.airportICAOCode)
		slices := strings.Split(result," ")
		firstCode := slices[0]
		if firstCode != test.airportICAOCode {
			t.Error("Error! Got %s but expected %s", firstCode, test.airportICAOCode)
		}
	}
}

func testGetFlightData(t *testing.T) {
	
}