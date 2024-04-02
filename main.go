package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/julien040/go-ternary"
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
		fmt.Print("Enter port number between 0 to 65536: (Optional) ")
		var input string
		fmt.Scanln(&input)
		if input == "" {
			break
		}
		number, err := strconv.Atoi(input)
		if err == nil && number > 0 && number < 65536 {
			port = number
			break
		}
	}

	portOrEmptyString := ternary.If(port == 0, "", fmt.Sprint(":", port))
	err = webbrowser.Open(fmt.Sprint("http://", gateway, portOrEmptyString, "/"))
	if err != nil {
		printMessageAndExit("Unable to open URL in Browser")
	}
}
