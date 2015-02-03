package main

import (
	"testing"
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

/*Test to get airport data
func TestGetAirportData(t *testing.T) {
	for i, test := range airportTests {
		_, resultICAO := getAirportData(test.airportICAOCode, validateAirportCode(test.airportICAOCode))
		_, resultIATA := getAirportData(test.airportIATACode, validateAirportCode(test.airportIATACode))
	}
}
*/
