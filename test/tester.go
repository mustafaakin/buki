package main
import (
	"github.com/mustafaakin/buki"
	"fmt"
)

func main() {
	VM := buki.GetVM("myubuntu")
	fmt.Printf("%+v", VM.Active)

	buki.StartVM("myubuntu")

	VM = buki.GetVM("myubuntu")
	fmt.Printf("%+v", VM.Active)

}
