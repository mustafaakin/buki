package main
import (
	"github.com/codegangsta/cli"
	"os"
	"github.com/mustafaakin/buki"
	"fmt"
)

func ListNetworks(c *cli.Context){
	nets := buki.GetNetworks()
	fmt.Printf("%10s %50s %10s %10s\n", "Name", "UUID", "Bridge", "Mode");
	fmt.Printf("%10s %50s %10s %10s\n", "====", "====", "======", "====");

	for _, net := range(nets) {
		fmt.Printf("%10s %50s %10s %10s\n", net.Name, net.UUID, net.Bridge.Name, net.Forward.Mode)
	}
}

func ListVMS(c* cli.Context){
	vms := buki.ListVM()
	fmt.Printf("%10s %50s %10s %10s %6s\n", "Name", "UUID", "CPUs", "Memory", "Active");
	fmt.Printf("%10s %50s %10s %10s %6s\n", "====", "====", "====", "======", "======");

	for _, vm := range(vms) {
		fmt.Printf("%10s %50s %10d %10d %6t\n", vm.Name, vm.UUID, vm.Cpus, (vm.Memory / 1024), vm.Active)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "buki"
	app.Version = "0.1"
	app.Usage = "creates vms"
	app.Commands = []cli.Command{{
		Name:      "network",
		Subcommands: []cli.Command{{
			Name: "list",
			Usage: "lists networks",
			Action: ListNetworks,
		}},
	},{
		Name:      "vm",
		Subcommands: []cli.Command{{
			Name: "list",
			Usage: "lists vms",
			Action: ListVMS,
		}},
	}}

	app.Run(os.Args)
}
