package ft

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var FlightAwareBase string = "https://flightxml.flightaware.com/json/FlightXML2/"
var UserHome string = os.Getenv("HOME")
var FlightAwareAPIKey string = os.Getenv("FLIGHTAWARE_API_KEY")
var FlightAwareAPIUser string = os.Getenv("FLIGHTAWARE_API_USER")
var BaseDir string = filepath.Dir(UserHome + "/.ft/")
var OutputFileAirports string = BaseDir + "/airports.dat"
var OutputFileAirlines string = BaseDir + "/airlines.dat"
var OutputFileRoutes string = BaseDir + "/routes.dat"

/*
Main airportInfo function
*/

func AirportInfoEval(inputAirport string) {

	if _, err := os.Stat(OutputFileAirports); os.IsNotExist(err) {
		getStaticData()
	}


	var ai airportInformation

	airportIndex := getAirportIndex(inputAirport)
	ai = getAirportData(airportIndex)
	fmt.Println("Airport Name:", ai.airportName)
	fmt.Print("Location    : ", ai.airportCity, ", ", ai.airportCountry, "\n")
	fmt.Println("ICAO        :", ai.airportICAOCode, " IATA: ", ai.airportIATACode)
	metar := getAirportMETAR(ai.airportICAOCode)
	fmt.Println("METAR       :", metar)

}

func FlightTrackingEval(inputFlightToTrack string) {

	if _, err := os.Stat(OutputFileRoutes); os.IsNotExist(err) {
		getStaticData()
	}

	if checkEnvVariables() != 0 {
		fmt.Println("Please set the FLIGHTAWARE_API_KEY and FLIGHTAWARE_API_USER variables.")
		os.Exit(-1)
	}

	flightData := getFlightData(inputFlightToTrack)
	for i := range flightData.FlightInfoExResult.Flights {
		fmt.Println("Origin City      : ", flightData.FlightInfoExResult.Flights[i].OriginCity)
		fmt.Println("Destination City : ", flightData.FlightInfoExResult.Flights[i].DestinationCity)
		fmt.Println("Aircraft Type    : ", flightData.FlightInfoExResult.Flights[i].Aircrafttype)
		t2 := int64(flightData.FlightInfoExResult.Flights[i].Estimatedarrivaltime)
		actualETA := time.Unix(t2, 0)

		fmt.Println("Filed Arrival    : ", flightData.FlightInfoExResult.Flights[i].FiledEte)
		fmt.Println("Scheduled Arrival: ", actualETA)
		fmt.Println("Route            : ", flightData.FlightInfoExResult.Flights[i].Route)
	}

}

/*
Validate airport code.
-2 - Invalid code.
1 - ICAO code.
0 - IATA code.
*/
func ValidateAirportCode(airportCode string) int32 {

	var airportIndex int32 = -2
	var airportCodeBytes = len([]rune(airportCode))

	if airportCodeBytes == 3 {
		airportIndex = 0
	} else if airportCodeBytes == 4 {
		airportIndex = 1
	} else {
		airportIndex = -2
	}
	return airportIndex
}

/*
Get basic airport info
*/
func getAirportData(airportIndex int) airportInformation {
	var airportDataResult airportInformation

	csvfile, err := os.Open(OutputFileAirports)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer csvfile.Close()
	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, each := range rawCSVdata {
		index, _ := strconv.Atoi(each[0])
		if index == airportIndex {
			lat, _ := strconv.ParseFloat(each[6], 32)
			lon, _ := strconv.ParseFloat(each[7], 32)
			airportDataResult.airportName = each[1]
			airportDataResult.airportICAOCode = each[5]
			airportDataResult.airportIATACode = each[4]
			airportDataResult.airportIndex = airportIndex
			airportDataResult.airportCountry = each[3]
			airportDataResult.airportCity = each[2]
			airportDataResult.airportLat = lat
			airportDataResult.airportLong = lon
		}
	}

	return airportDataResult
}

/*
Get index of an airport code so we can gather all the data by index
*/
func getAirportIndex(airportCode string) int {
	csvfile, err := os.Open(OutputFileAirports)
	if err != nil {
		fmt.Println(err)
		return -2
	}
	defer csvfile.Close()
	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, each := range rawCSVdata {
		if each[4] == airportCode || each[5] == airportCode {
			index, _ := strconv.Atoi(each[0])
			return index
		}
	}
	return -2
}

/*
Function to get an airport METAR
*/
func getAirportMETAR(airportICAOCode string) string {
	//	var METARURL string = "https://aviationweather.gov/adds/dataserver_current/httpparam?dataSource=metars&requestType=retrieve&format=xml&stationString=" + airportICAOCode + "&hoursBeforeNow=1"
	var METARURL string = "http://dev.geonames.org/weatherIcaoJSON?ICAO=" + airportICAOCode
	met := Metar{}

	resp, err := http.Get(METARURL)
	if err != nil {
		fmt.Println("error %v", err)
		os.Exit(-1)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		os.Exit(-1)
	}

	err2 := json.Unmarshal([]byte(string(data)), &met)
	if err2 != nil {
		fmt.Println("Error parsing data! ", err2)
		os.Exit(-1)
	}
	return met.WeatherObservation.Observation
}

func getFlightData(flightNumber string) flightInformation {
	flightInfo := flightInformation{}
	var flightInfoURL string = FlightAwareBase + "FlightInfoEx?ident=" + flightNumber + "&howMany=1&offset=2"
	client := &http.Client{}
	req, err := http.NewRequest("GET", flightInfoURL, nil)
	req.SetBasicAuth(FlightAwareAPIUser, FlightAwareAPIKey)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error %v", err)
		os.Exit(-1)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(string(bodyText)), &flightInfo)
	if err != nil {
		fmt.Println("Ha petao! ", err)
		os.Exit(1)
	}

	for i := range flightInfo.FlightInfoExResult.Flights {
		//		fmt.Println(flightInfo.FlightInfoExResult.Flights[i].Ident)
		_ = i
	}

	return flightInfo
}

/*
================== Some standard trivial functions to avoid repetition. =============================
*/

/*
Download data from openflights.org to not make much use of the FlightAware API ($$)
*/
func getStaticData() int {
	os.MkdirAll(BaseDir, 0777)
	downloadFromUrl("http://sourceforge.net/p/openflights/code/HEAD/tree/openflights/data/airports.dat?format=raw", OutputFileAirports)
	downloadFromUrl("http://sourceforge.net/p/openflights/code/HEAD/tree/openflights/data/airlines.dat?format=raw", OutputFileAirlines)
	downloadFromUrl("http://sourceforge.net/p/openflights/code/HEAD/tree/openflights/data/routes.dat?format=raw", OutputFileRoutes)
	return 0
}

/*
Download a file from the URL
*/
func downloadFromUrl(url string, fileName string) {

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(filepath.FromSlash(fileName))

	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	bytesRead, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	_ = bytesRead
}
func checkEnvVariables() int {
	if FlightAwareAPIKey == "" || FlightAwareAPIUser == "" {
		return 1
	}
	return 0
}
