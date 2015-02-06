package main

import (
//	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
//	"github.com/jmoiron/jsonq"
	"github.com/moovweb/gokogiri"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
	//	"strings"
)

/*
Constant definitions
*/

const version string = "0.1"
const author string = "Ismael Arenzana"
const email string = "iarenzana@gmail.com"
const appName string = "ft"
const appDescription string = "Command-line flight tracker"
const flightAwareBase string = "https://flightxml.flightaware.com/json/FlightXML2/"

var userHome string = os.Getenv("HOME")
var flightAwareAPIKey string = os.Getenv("FLIGHTAWARE_API_KEY")
var flightAwareAPIUser string = os.Getenv("FLIGHTAWARE_API_USER")
var baseDir string = filepath.Dir(userHome + "/.ft/")
var outputFileAirports string = baseDir + "/airports.dat"
var outputFileAirlines string = baseDir + "/airlines.dat"
var outputFileRoutes string = baseDir + "/routes.dat"

/*
Main function
*/

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appDescription
	app.Email = email
	app.Author = author
	app.Version = version
	app.Commands = []cli.Command{
		{
			Name:      "airportinfo",
			ShortName: "a",
			Usage:     "Display Airport Information",
			Action: func(c *cli.Context) {
				var ai airportInformation
				var inputAirport string = c.Args()[0]
				if _, err := os.Stat(outputFileAirports); os.IsNotExist(err) {
					getStaticData()
				}
				airportVal := validateAirportCode(inputAirport)
				if airportVal == -2 {
					fmt.Println("Error, could not validate airport code.")
					os.Exit(1)
				}
				airportIndex := getAirportIndex(inputAirport)
				ai = getAirportData(airportIndex)
				fmt.Println("Airport Name:", ai.airportName)
				fmt.Print("Location    : ", ai.airportCity, ", ", ai.airportCountry, "\n")
				fmt.Println("ICAO        :", ai.airportICAOCode, " IATA: ", ai.airportIATACode)
				metar := getAirportMETAR(ai.airportICAOCode)
				fmt.Println("METAR       :", metar)
			},
		},
		{
			Name:      "track",
			ShortName: "t",
			Usage:     "Track a Flight",
			Action: func(c *cli.Context) {
				if _, err := os.Stat(outputFileRoutes); os.IsNotExist(err) {
					getStaticData()
				}
				var inputFlightToTrack string = c.Args()[0]
				flightData := getFlightData(inputFlightToTrack)
				for i := range flightData.FlightInfoExResult.Flights {
					fmt.Println("Origin City      : ", flightData.FlightInfoExResult.Flights[i].OriginCity)
					fmt.Println("Destination City : ", flightData.FlightInfoExResult.Flights[i].DestinationCity)
					fmt.Println("Aircraft Type    : ", flightData.FlightInfoExResult.Flights[i].Aircrafttype)
//					t := int64(flightData.FlightInfoExResult.Flights[i].FiledTime)
//					filedETA := time.Unix(t,0)
					t2 := int64(flightData.FlightInfoExResult.Flights[i].Estimatedarrivaltime)
					actualETA := time.Unix(t2,0)

					fmt.Println("Filed Arrival    : ", flightData.FlightInfoExResult.Flights[i].FiledEte)
					fmt.Println("Scheduled Arrival: ", actualETA)
					fmt.Println("Route            : ", flightData.FlightInfoExResult.Flights[i].Route)
				}

			},
		},
		{
			Name:      "airlineinfo",
			ShortName: "l",
			Usage:     "Airline Information",
			Action: func(c *cli.Context) {
				if _, err := os.Stat(outputFileAirlines); os.IsNotExist(err) {
					getStaticData()
				}

			},
		},
	}

	app.Run(os.Args)
}

/*
Validate airport code.
-2 - Invalid code.
1 - ICAO code.
0 - IATA code.
*/
func validateAirportCode(airportCode string) int32 {

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

	csvfile, err := os.Open(outputFileAirports)
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
	csvfile, err := os.Open(outputFileAirports)
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
	var METARURL string = "https://aviationweather.gov/adds/dataserver_current/httpparam?dataSource=metars&requestType=retrieve&format=xml&stationString=" + airportICAOCode + "&hoursBeforeNow=1"
	var metar string

	resp, err := http.Get(METARURL)
	if err != nil {
		fmt.Println("error %v", err)
		os.Exit(-1)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	doc, err := gokogiri.ParseXml(data)
	if err != nil {
		os.Exit(-1)
	}
	defer doc.Free()

	metars, err := doc.Root().Search("data/METAR/raw_text")
	if err != nil {
		fmt.Println("Error parsing XML!")
		os.Exit(-1)
	}
	for _, resultset := range metars {
		metar = resultset.Content()
		break
	}
	return metar
}

func getFlightData(flightNumber string) flightInformation {
	flightInfo := flightInformation{}
	var flightInfoURL string = flightAwareBase + "FlightInfoEx?ident=" + flightNumber + "&howMany=1&offset=2"
	client := &http.Client{}
	req, err := http.NewRequest("GET", flightInfoURL, nil)
	req.SetBasicAuth(flightAwareAPIUser, flightAwareAPIKey)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error %v", err)
		os.Exit(-1)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
/*	jsonfile, err := os.Open("/Users/iaren/tmp/flightinfo.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer jsonfile.Close()

	scanner := bufio.NewScanner(jsonfile)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var result string
	for _, l := range lines {
		result = string(l)
	}
*/
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
	os.MkdirAll(baseDir, 0777)
	downloadFromUrl("http://sourceforge.net/p/openflights/code/HEAD/tree/openflights/data/airports.dat?format=raw", outputFileAirports)
	downloadFromUrl("http://sourceforge.net/p/openflights/code/HEAD/tree/openflights/data/airlines.dat?format=raw", outputFileAirlines)
	downloadFromUrl("http://sourceforge.net/p/openflights/code/HEAD/tree/openflights/data/routes.dat?format=raw", outputFileRoutes)
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
