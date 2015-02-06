# ft
Command line flight tracker

##Installation

You can clone the repo and build it yourself:

You will need three environment variables set to run the application:

GOPATH (only if you're building it.)

You will need a [Flightaware API key] (http://flightaware.com/commercial/flightxml/) to retrieve the data. Once you have the key, make sure you have FLIGHTAWARE_API_KEY and FLIGHTAWARE_API_USER populated.

```
git clone https://github.com/iarenzana/ft
git submodule init
cd ft
go build
```

This will generate a binary for your platform.

##Usage

```
NAME:
   ft - Command-line flight tracker

USAGE:
   ft [global options] command [command options] [arguments...]

VERSION:
   0.1

AUTHOR:
  Ismael Arenzana - <iarenzana@gmail.com>

COMMANDS:
   airportinfo, a	Display Airport Information
   track, t		Track a Flight
   airlineinfo, l	Airline Information
   help, h		Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```


##Other Libraries used

* [cli] (https://github.com/codegangsta/cli) - Allows me to create a command line application super fast.
* [pb] (https://github.com/cheggaaa/pb) - For progress bars.
* [jflect] (https://github.com/str1ngs/jflect) - Go struct generator from JSON. Life Saver.
* [gox] (https://github.com/mitchellh/gox) - To compile the code for several architectures.
