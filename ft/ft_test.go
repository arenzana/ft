package ft

import (
	"testing"
	"strings"
	"os"
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

var FlightTests = []struct {
	FlightID string
	Destination string
}{
	{
		FlightID: "UA922",
		Destination: "EGLL",
	},
	{
		FlightID: "AA47",
		Destination: "KORD",		
	},
	{
		FlightID: "AF90",
		Destination: "KMIA",
	},
}

var fakeAirportTests = []string{"HEEORF", "A"}

func TestUnitAirportInfoEval(t *testing.T) {
	for _, test := range airportTests {
		AirportInfoEval(test.airportICAOCode)
	}
}

func TestUnitFlightTrackingEval(t *testing.T) {
	for _, test := range FlightTests {
		FlightTrackingEval(test.FlightID)
	}
}

/*Test to make sure Airport code size works*/
func TestUnitValidateAirportCode(t *testing.T) {
	for i, test := range airportTests {
		var actualICAO = ValidateAirportCode(test.airportICAOCode)
		var actualIATA = ValidateAirportCode(test.airportIATACode)

		if actualICAO != 1 {
			t.Error("Expected 1, got ", actualICAO, "Test %d", i)
		}
		if actualIATA != 0 {
			t.Error("Expected 0, got ", actualIATA, "Test %d", i)
		}

	}

	for i, ret := range fakeAirportTests {
		res := ValidateAirportCode(ret)
		if res != -2 {
			t.Error("Expected 2, got ", res, "Test %d", i)
		}
	}

}

/*Test to get airport indexes*/
func TestUnitGetAirportIndex(t *testing.T) {
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
func TestUnitGetAirportData(t *testing.T) {
	var airportInfo airportInformation
	airportInfo = getAirportData(getAirportIndex("KJFK"))
	if airportInfo.airportIndex != 3797 {
		t.Error("Error! Expected 3797 but got ", airportInfo.airportIndex)
	}
}

/*Test to get airport METAR*/
func TestUnitGetAirportMETAR(t *testing.T){
	for _, test := range airportTests {
		result := getAirportMETAR(test.airportICAOCode)
		slices := strings.Split(result," ")
		firstCode := slices[0]
		if firstCode != test.airportICAOCode {
			t.Error("Error! First code and ICAO Code: ", firstCode, test.airportICAOCode)
		}
	}
}

/*Test to get flight data*/
func TestUnitGetFlightData(t *testing.T) {
	var fi flightInformation

	for i, test := range FlightTests {
		fi = getFlightData(test.FlightID)
		if fi.FlightInfoExResult.Flights[0].Destination != test.Destination {
			t.Error("Expected ",test.Destination, "but got ",fi.FlightInfoExResult.Flights[i].Destination, "Test %d", i)
		}

	}
}

/*Test to make sure stats data works*/
func TestUnitGetStatisticData(t *testing.T) {
	getStaticData()
	if _, err := os.Stat(OutputFileRoutes); os.IsNotExist(err) {
		t.Error("Expected file ",OutputFileRoutes)
	}
}
/*Test to test download functionality*/
func TestUnitDownloadFromURL(t *testing.T) {
	downloadFromUrl("http://sourceforge.net/p/openflights/code/HEAD/tree/openflights/data/routes.dat?format=raw", OutputFileRoutes)
	if _, err := os.Stat(OutputFileRoutes); os.IsNotExist(err) {
		t.Error("Couldn't download file ",OutputFileRoutes)
	}
}
/*Test to make sure the env variables test func works*/
func TestUnitCheckEnvVariables( t *testing.T) {
}