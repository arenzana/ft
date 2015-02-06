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

type flightInformation struct {
	FlightInfoExResult struct {
		Flights []struct {
			Actualarrivaltime    int    `json:"actualarrivaltime"`
			Actualdeparturetime  int    `json:"actualdeparturetime"`
			Aircrafttype         string `json:"aircrafttype"`
			Destination          string `json:"destination"`
			DestinationCity      string `json:"destinationCity"`
			DestinationName      string `json:"destinationName"`
			Diverted             string `json:"diverted"`
			Estimatedarrivaltime int    `json:"estimatedarrivaltime"`
			FaFlightID           string `json:"faFlightID"`
			FiledAirspeedKts     int    `json:"filed_airspeed_kts"`
			FiledAirspeedMach    string `json:"filed_airspeed_mach"`
			FiledAltitude        int    `json:"filed_altitude"`
			FiledDeparturetime   int    `json:"filed_departuretime"`
			FiledEte             string `json:"filed_ete"`
			FiledTime            int    `json:"filed_time"`
			Ident                string `json:"ident"`
			Origin               string `json:"origin"`
			OriginCity           string `json:"originCity"`
			OriginName           string `json:"originName"`
			Route                string `json:"route"`
		} `json:"flights"`
		NextOffset int `json:"next_offset"`
	} `json:"FlightInfoExResult"`
}
