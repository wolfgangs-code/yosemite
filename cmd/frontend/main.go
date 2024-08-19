package main

import "fmt"

func main() {
	// ANSI color codes
	const (
		RESET = "\033[0m"
		GREEN = "\033[32m"
		BLUE  = "\033[34m"
		GREY  = "\033[90m"
	)
	fmt.Println(GREEN + "Yosemite" + RESET + " Frontend Server" + BLUE + " (Tunnel View)" + GREY + " v0.0.0" + RESET)
}
