package main

type airportInformation struct {
	airportName     string
	airportICAOCode string
	airportIATACode string
	airportIndex    int
	airportCountry  string
	airportCity     string
	airportLat      float64
	airportLong     float64
}

/*type airportMETAR struct {
	airportICAOCode string
	XMLName xml.Name `xml:"respsonse"`
	raw_text string `xml: "response>data>METAR"`
	elevation float32 `xml: "response>data>METAR"`
}*/