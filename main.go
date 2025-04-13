package main

import (
	"fmt"
	"sneakytools/helpers"
	"sneakytools/scanners"
	"strings"
)

func main() {
	helpers.CliInit()

	keepGoing := "y"
	var choice uint

	for {

		helpers.CallClear()
		helpers.PrintBanner()
		helpers.PrintOptions()

		fmt.Print("choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var scanner = scanners.TcpScanner{}
			scanner.GatherContext()
			scanner.TcpScan()
		default:
			panic("invalid choice")
		}

		fmt.Print("Would you like to run another operation? (y/n)")
		fmt.Scan(&keepGoing)
		if strings.ToUpper(keepGoing) != "Y" {
			break
		}
	}
}
