package main

import "encoding/xml"

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

type Metar struct {
	airportICAOCode string
	raw_text string `xml "raw_text"`
	elevation float32 `xml "elevation"`
}

type Response struct {
	XMLName xml.Name `xml: "response"`
	Metars []Metar `xml: "METAR"`
}