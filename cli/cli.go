package cli

import "fmt"

type Data struct {
	Location string
	Unit     string
}

func Run(args []string) Data {
	fmt.Println(args[1])
	return Data{}
}
