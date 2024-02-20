package utils

import (
	"os"
	"strconv"
)

func GetPort() int {
	// Try to get the port number from the command line
	// If it's not there, use the default port number
	port := 15000
	err := error(nil)
	if len(os.Args) > 1 {
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			port = 15000
		}
		// check if the port number is valid
		// if it's not, use the default port number
		if port < 1024 || port > 65535 {
			port = 15000
		}
	}
	return port
}
