package main

import (
	"fmt"
	"os"
	"weather_cli/cli"
)

func main() {
	// cmd := []string{"cmd", "--city=\"London\"", "--unit=\"metric\""}
	s, err := cli.Parse(os.Args)
	fmt.Println(s, err)
}
