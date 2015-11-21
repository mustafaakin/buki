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

	uuid, err := buki.CreateNATNetwork("mustafa", "172.16.0.1","255.255.248.0")
	println(uuid)
	println(err)

}
