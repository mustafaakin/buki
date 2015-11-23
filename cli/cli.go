package main
import (
	"github.com/codegangsta/cli"
	"os"
	"github.com/mustafaakin/buki"
	"fmt"
	"io/ioutil"
)

func ListImages(c *cli.Context){
	images := buki.GetAvailableImages()
	fmt.Printf("%20s\n", "Name");
	fmt.Printf("%20s\n", "====");

	for _, image:= range(images) {
		fmt.Printf("%20s\n", image)
	}
}

func ListNetworks(c *cli.Context){
	nets := buki.GetNetworks()
	fmt.Printf("%10s %10s %10s\n", "Name", "Bridge", "Mode");
	fmt.Printf("%10s %10s %10s\n", "====", "======", "====");

	for _, net := range(nets) {
		fmt.Printf("%10s %10s %10s\n", net.Name, net.UUID, net.Bridge.Name, net.Forward.Mode)
	}
}

func ListVMS(c* cli.Context){
	vms := buki.ListVM()
	fmt.Printf("%10s %10s %10s %10s %6s\n", "Name", "CPUs", "Memory", "Networks", "Active");
	fmt.Printf("%10s %10s %10s %10s %6s\n", "====", "====", "======", "========", "======");

	for _, vm := range(vms) {
		fmt.Printf("%10s %10d %10d %10s %6t\n", vm.Name, vm.Cpus, (vm.Memory / 1024), vm.Interface[0].Source.Network, vm.Active)
	}
}

func StopVM(c* cli.Context){
	name := c.String("name")
	buki.StopVM(name)
}

func StartVM(c* cli.Context){
	name := c.String("name")
	buki.StartVM(name)
}

func CreateVM(c* cli.Context){
	name := c.String("name")
	image := c.String("image")
	cpus := c.Int("CPUs")
	ram := c.Int("RAM") * 1024
	diskSize := c.String("diskSize")
	network := c.String("network")

	// TODO: If file does not exist, the default cloudConfigFile
	cloudConfigFile, _ := ioutil.ReadFile(c.String("cloudConfigFile"))
	cloudConfig := string(cloudConfigFile)

	VM, _ := buki.CreateBasicVM(image, name, cpus, ram, diskSize, network, cloudConfig)
	fmt.Printf("Created VM from image '%s' : %s \n", name, VM)
}

func main() {
	app := cli.NewApp()
	app.Name = "buki"
	app.Version = "0.1"
	app.Author = "Mustafa Akin <mustafa91@gmail.com>"
	app.Usage = "easy libvirt virtualization management tool"
	app.Commands = []cli.Command{{
		Name:      "image",
		Subcommands: []cli.Command{{
			Name: "list",
			Usage: "Lists available VM Images",
			Action: ListImages,
		}},
	}, {
		Name:      "network",
		Subcommands: []cli.Command{{
			Name: "list",
			Usage: "lists networks",
			Action: ListNetworks,
		}},
	},{
		Name:      "vm",
		Usage:	"Mange Virtual machines",
		Subcommands: []cli.Command{{
			Name: "list",
			Usage: "lists vms",
			Action: ListVMS,
		}, {
			Name: "create",
			Usage: "Create a VM",
			Action: CreateVM,
			Flags: []cli.Flag {
				cli.StringFlag{
					Name: "name",
					Usage: "VM name",
				},
				cli.StringFlag{
					Name: "image",
					Usage: "Base VM image to copy",
				},
				cli.IntFlag{
					Name: "CPUs",
					Value: 1,
					Usage: "Number of CPUs to allocate",
				},
				cli.IntFlag{
					Name: "RAM",
					Value: 1024,
					Usage: "RAM size in MBs",
				},
				cli.StringFlag{
					Name: "diskSize",
					Value: "10G",
					Usage: "Primary OS disk size i.e. 100M, 5G, 20G, 1TB etc.",
				},
				cli.StringFlag{
					Name: "cloudConfigFile",
					Usage: "Cloud config file to set some variables",
				},
				cli.StringFlag{
					Name: "network",
					Value: "default",
					Usage: "Network for the virtual machine",
				},
			},
		},{
			Name: "stop",
			Usage: "Stops a VM given by name",
			Action: StopVM,
			Flags: []cli.Flag{cli.StringFlag{
				Name: "name",
				Usage: "VM name",
			}},
		},{
			Name: "start",
			Usage: "Start a VM given by name",
			Action: StartVM,
			Flags: []cli.Flag{cli.StringFlag{
				Name: "name",
				Usage: "VM name",
			}},
		}},
	}}

	app.Run(os.Args)
}
