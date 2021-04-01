package cli

type Data struct {
	Location string
	Unit     string
}

func ParseInput(string) Data {
	return Data{Location: "london", Unit: "metric"}
}
