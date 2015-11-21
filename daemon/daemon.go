package main

import (
	"github.com/mustafaakin/buki"
	"fmt"
	"encoding/json"
)

func main() {
	nets := buki.GetNetworks()

	b, err := json.Marshal(nets)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

//	newNet, err := buki.CreateNATNetwork("mustafa", "172.16.0.1","255.255.248.0", "172.16.1.10", "172.16.1.50")

//	fmt.Printf("%+v \n", newNet)
}
