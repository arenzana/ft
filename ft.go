package main

import (
	"encoding/csv"
	"fmt"
	"github.com/codegangsta/cli"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

/*
Constant definitions
*/

const version string = "0.1"
const author string = "Ismael Arenzana"
const email string = "iarenzana@gmail.com"
const appName string = "ft"
const appDescription string = "Command-line flight tracker"

var userHome string = os.Getenv("HOME")
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
				var inputAirport string = c.Args()[0]
				if _, err := os.Stat(outputFileAirports); os.IsNotExist(err) {
					getStaticData()
				}
				airportVal := validateAirportCode(inputAirport)
				if airportVal == -2 {
					fmt.Println("Error, could not validate airport code.")
					os.Exit(1)
				}
				fmt.Println("Validated!")
				airportIndex := getAirportIndex(inputAirport)
				fmt.Println("Index is ", airportIndex)
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
func getAirportData(airportCode string, codeType int) airportInformation {
	var airportDataResult airportInformation
	return airportDataResult
}

/*
Some standard trivial functions to avoid repetition.
*/

/*
Download data from openflights.org to not make much use of the FlightAware API ($$)
*/
func getStaticData() int {
	os.MkdirAll(baseDir, 0777)
	downloadFromUrl("http://sourceforge.net/p/openflights/code/HEAD/tree/openflights/data/airports.dat", outputFileAirports)
	downloadFromUrl("http://sourceforge.net/p/openflights/code/HEAD/tree/openflights/data/airlines.dat", outputFileAirlines)
	downloadFromUrl("http://sourceforge.net/p/openflights/code/HEAD/tree/openflights/data/routes.dat", outputFileRoutes)
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

func getAirportIndex(airportCode string) int {
	csvfile, err := os.Open(outputFileAirports)
	if err != nil {
		fmt.Println(err)
		return -2
	}
	fmt.Println("Open file...")
	defer csvfile.Close()
	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1
	fmt.Println("Read...")

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		//		os.Exit(1)
	}
	fmt.Println("Begin Iteration...")
	for _, each := range rawCSVdata {
		fmt.Printf("Airport: %s IATA: %s, ICAO: %s\n, Index: %d", each[1], each[4], each[5], each[0])
		if each[4] == airportCode || each[5] == airportCode {
			fmt.Printf("Airport: %s IATA: %s, ICAO: %s\n, Index: %d", each[1], each[4], each[5], each[0])
			index, _ := strconv.Atoi(each[0])
			return index
		}
	}
	return -2
}
