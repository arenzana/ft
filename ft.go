package main

import ( "os"
       "github.com/codegangsta/cli"
       "net/http"
       "io"
       "fmt"
       "path/filepath"
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

func main(){
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appDescription
	app.Email = email
	app.Author = author
	app.Version = version
	app.Commands = []cli.Command {
		{
	    	Name:      "airportinfo",
		    ShortName: "a",
		    Usage:     "Display Airport Information",
		    Action: func(c *cli.Context) {
		      if _, err := os.Stat(outputFileAirports); os.IsNotExist(err) {
				getStaticData()
				}
				validateAirportCode(c.Args()[0])

		    },
  		},
  		{
  			Name:	"track",
  			ShortName: "t",
  			Usage: "Track a Flight",
  			Action: func(c *cli.Context) {
		      if _, err := os.Stat(outputFileRoutes); os.IsNotExist(err) {
				getStaticData()
				}
			},
  		},
  		{
  			Name:	"airlineinfo",
  			ShortName: "l",
  			Usage: "Airline Information",
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
Download data from openflights.org to not make much use of the FlightAware API ($$)
*/
func getStaticData() int {
	os.MkdirAll(baseDir, 0777)
	downloadFromUrl("http://sourceforge.net/p/openflights/code/HEAD/tree/openflights/data/airports.dat",outputFileAirports)
	downloadFromUrl("http://sourceforge.net/p/openflights/code/HEAD/tree/openflights/data/airlines.dat",outputFileAirlines)
	downloadFromUrl("http://sourceforge.net/p/openflights/code/HEAD/tree/openflights/data/routes.dat",outputFileRoutes)
	return 0
}
/*
Validate airport code.
2 - Invalid code.
1 - ICAO code.
0 - IATA code.
*/
func validateAirportCode (airportCode string) int {

	var airportEval int = 2
	var airportCodeBytes=len([]rune(airportCode))

	if airportCodeBytes == 3 {
		airportEval = 0
	} else if airportCodeBytes == 4 {
		airportEval = 1
	} else {
		airportEval = 2
	}
	return airportEval
}

func downloadFromUrl(url string,fileName string) {

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
