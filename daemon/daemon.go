package main

import (
//	"github.com/mustafaakin/buki"
//	"fmt"
	"github.com/mustafaakin/buki"
	"fmt"
	"encoding/json"
)

func main() {
/*
	nets := buki.GetNetworks()

	b, err := json.Marshal(nets)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
*/

//	newNet, err := buki.CreateNATNetwork("mustafa", "172.16.0.1","255.255.248.0", "172.16.1.10", "172.16.1.50")

//	fmt.Printf("%+v \n", newNet)

	/*
	err := buki.DownloadImage("https://cloud-images.ubuntu.com/trusty/current/trusty-server-cloudimg-amd64-disk1.img", "ubuntu_14-04")
	if err != nil  {		
		println("Error occured", err)
	} else {
		println("Image succesfully downloaded")
	}
	*/
	
//	fmt.Println(buki.GetAvailableImages())

/*
	for i := 0; i < 10; i++ {
		println(buki.GenerateMAC())
	}
*/

//	buki.CopyImage("ubuntu_14-04", "myubuntu", "10G")
/*
	userdata := `#cloud-config
password: mustafa
chpasswd: { expire: False }
ssh_pwauth: True`

	err := buki.CreateCloudConfig("myubuntu", userdata)
	if err != nil {
		println(err.Error())
	} else {
		println("Created cloud-config image")
	}

	buki.CreateVM("ubuntu")
	*/

	vms := buki.ListVM()
	b, err := json.Marshal(vms)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
