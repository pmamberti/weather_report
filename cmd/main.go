package main

import (
	"os"
	"weather"
)

func main() {
	weather.RunCLI(os.Args[1:], os.Stdout)
}
