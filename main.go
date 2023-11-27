package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/net-byte/go-gateway"
	"github.com/toqueteos/webbrowser"
)

func printMessageAndExit(msg string) {
	fmt.Println(msg)
	fmt.Print("Press enter key to exit")
	fmt.Scanln()
	os.Exit(0)
}

func main() {
	gateway, err := gateway.DiscoverGatewayIPv4()
	if err != nil {
		printMessageAndExit("Unable to discover Gateway IP")
	}

	fmt.Println("Gateway IP: ", gateway)

	var port int

	for {
		fmt.Print("Enter port number between 0 to 65535: (8080) ")
		var input string
		fmt.Scanln(&input)
		if input == "" {
			port = 8080
			break
		}
		port, err = strconv.Atoi(input)
		if err == nil && port >= 0 && port <= 65535 {
			break
		}
	}

	err = webbrowser.Open(fmt.Sprint("http://", gateway, ":", port, "/"))
	if err != nil {
		printMessageAndExit("Unable to open URL in Browser")
	}
}
