package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/karagog/polygon-io-client-golang/client"
)

// Set the environment variable like so:
// $ export POLYGON_IO_API_KEY=<abcd>
const apiKeyEnvVar = "POLYGON_IO_API_KEY"

var ctx = context.Background()

func main() {
	tickersCmd := flag.NewFlagSet("tickers", flag.ExitOnError)
	tickersSymbols := tickersCmd.String("symbols", "", "Comma-separated symbols to query")
	tickersSearch := tickersCmd.String("search", "", "Search terms")
	tickersCmd.Usage = func() {
		fmt.Printf("Usage of the tickers command:\n\n")
		fmt.Println("\ttestcli tickers --symbols={comma-separated-list}")
	}

	// Get the API key from the environment.
	apiKey := os.Getenv(apiKeyEnvVar)
	if apiKey == "" {
		panic(fmt.Sprintf("API key not specified in environment variable: %s", apiKeyEnvVar))
	}
	c := client.New(apiKey)

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n\n", os.Args[0])
		fmt.Println("\ttestcli {command} [-h]")
		fmt.Println()
		fmt.Println("Run with -h to get usage details for each command")
	}
	flag.Parse()
	flag.Set("alsologtostderr", "true")

	if len(os.Args) < 2 {
		panic("You must specify a command")
	}

	var response []byte
	var err error
	cmd := os.Args[1]
	switch cmd {
	case "tickers":
		tickersCmd.Parse(os.Args[2:])
		response, err = c.Tickers(ctx, &client.TickersOptions{
			Symbols: strings.Split(*tickersSymbols, ","),
			Search:  *tickersSearch,
		})
	default:
		panic(fmt.Sprintf("Unknown command: %s", cmd))
	}
	if err != nil {
		panic(err)
	}

	// Show the raw response, with indentations to make it readable.
	var indented bytes.Buffer
	if err := json.Indent(&indented, response, "", "\t"); err != nil {
		panic(err)
	}
	fmt.Println("Raw json response: \n", indented.String())
}
