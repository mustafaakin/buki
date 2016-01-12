package main

import (
	"fmt"

	"github.com/mustafaakin/buki"
)

func main() {
	/*
		net, err := buki.CreateBridgedNetwork("br0-bridge", "br0")

		fmt.Printf("%+v \n", net)
		fmt.Printf("%+v \n", err)
	*/

	var cloudConfig = `#cloud-config
password: mustafa
chpasswd: { expire: False }
ssh_pwauth: True
hostname: dq
`
	image := "trusty-server-cloudimg-amd64-disk1"
	name := "deneme4"
	cpus := 2
	ram := 2048
	disk := "50G"
	network := "vmbr0"

	vm, err := buki.CreateBasicVM(image, name, cpus, ram*1024, disk, network, cloudConfig)

	fmt.Printf("%+v \n", vm)
	fmt.Printf("%+v \n", err)

	/*
		VM := buki.GetVM("myubuntu")
		fmt.Printf("%+v", VM.Active)

		buki.StartVM("myubuntu")

		VM = buki.GetVM("myubuntu")
		fmt.Printf("%+v", VM.Active)
	*/
	vm = buki.GetVM(name)
	fmt.Println(vm)

}
