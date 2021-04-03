package main

import (
	"fmt"
	"weather_cli/cli"
)

func main() {
	cmd := []string{"cmd", "--city=\"London\"", "--unit=\"metric\""}
	c, u, err := cli.Parse(cmd)
	fmt.Println(c, u, err)
}
