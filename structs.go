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
	flightNumber 			string
	actualarrivaltime		int
	actualdeparturetime		int
	aircrafttype			string
	destination				string
	destinationCity			string
	destinationName			string
	diverted				string
	estimatedarrivaltime	int
	faFlightID				string
	filed_airspeed_kts		int
	filed_airspeed_mach		string
	filed_altitude			int
	filed_departuretime		int
	filed_ete				string
	filed_time 				int
	ident					string
	origin					string
	originCity				string
	originName				string
	route					string
}