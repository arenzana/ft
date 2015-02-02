package main

import ( "os"
       "github.com/codegangsta/cli"
       "net/http"
       "io"
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
const userHome string = GetEnv("HOME")
const separator = os.PathSeparator
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
		      println("Airport Information: ", c.Args().First())
		    },
  		},
  		{
  			Name:	"track",
  			ShortName: "t",
  			Usage: "Track a Flight",
  			Action: func(c *cli.Context) {
		      println("Flight Information: ", c.Args().First())
		    },
  		},
  		{
  			Name:	"airlineinfo",
  			ShortName: "l",
  			Usage: "Airline Information",
  			Action: func(c *cli.Context) {
		      println("Airline Information: ", c.Args().First())
		    },
  		},
  	}
//	app.Action = func (c *cli.Context) {

	//Entry Point
//		println ("Boomshakalaka")
//	}

app.Run(os.Args)
}
/*
Validate airport code.
2 - Invalid code.
1 - ICAO code.
0 - IATA code.
*/
func getStaticData() int {
	var outputFileAirport string = userHome + "/."
}
func validateAirportCode (airportCode string) int {
	var airportEval int = 2

	return airportEval
}