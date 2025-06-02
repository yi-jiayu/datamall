package main

import (
	"fmt"
	"github.com/yi-jiayu/datamall/v3"
	"os"
)

func main() {
	client := datamall.NewDefaultClient(os.Getenv("DATAMALL_ACCOUNT_KEY"))
	arrivals, err := client.GetBusArrival("96049", "")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", arrivals)
}
