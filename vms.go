package buki
import (
	"encoding/xml"
	"strconv"

)

type VM struct {
	XMLName xml.Name       `xml:"domain" json:"-"`
	Name    string	       `xml:"name" json:"name"`
	UUID	string		   `xml:"uuid" json:"uuid"`
	Memory	int			   `xml:"memory" json:"memory"`
	Cpus	int			   `xml:"vcpu" json:"cpus"`
	Active 	bool		   `json:"active"`
	Disks	[]Disk		   `xml:"devices>disk" json:"disks"`
	Interface []Interface   `xml:"devices>interface" json:"interfaces"`
	Graphics Graphics	`xml:"devices>graphics" json:"graphics"`
}

type Disk struct{
	Type	string	`xml:"type,attr" json:"type"`
	Source	Source	`xml:"source" json:"source"`
	Target  Target `xml:"target" json:"target"`
}

type Source struct {
	// Add more for supporting others
	File  string    `xml:"file,attr" json:"file,omitempty"`
	Network string    `xml:"network,attr" json:"network,omitempty"`

}

type Target struct {
	Dev  string  `xml:"dev,attr" json:"dev"`
	Bus  string  `xml:"bus,attr" json:"bus,omitempty"`
}

type Interface struct {
	Type	string	`xml:"type,attr" json:"type"`
	Source	Source	`xml:"source" json:"source"`
	Target  Target `xml:"target" json:"target"`
	MAC		MAC	 `xml:"mac" json:"mac"`
	Model   Model  `xml:"model" json:"model"`

}

type MAC struct {
	Address string	`xml:"address,attr" json:"address"`
}

type Graphics struct {
	Type	string	`xml:"type,attr" json:"type"`
	Port    string	`xml:"port,attr" json:"port"`
	Listen	string	`xml:"listen,attr" json:"listen"`
}

type Model struct {
	Type  string	`xml:"type,attr" json:"type"`
}

func ListVM() []VM {
	conn := BuildConnection()
	defer conn.CloseConnection()

	vms, _ := conn.ListAllDomains(0)

	VMs := make([]VM, len(vms))

	for idx, vm := range vms {
		xmlResp, _ := vm.GetXMLDesc(0)
		// ignoring error right now
		xml.Unmarshal([]byte(xmlResp), &VMs[idx])
		isActive, _ := vm.IsActive()
		VMs[idx].Active = isActive
	}

	return VMs
}
// CreateBasicVM creates a basic VM with given parameters
func CreateBasicVM(image string, name string, cpus, ram int, diskSize, network, cloudConfig string) (*VM, error){
	// TODO: Escape all params
	conn := BuildConnection()
	defer conn.CloseConnection()

	// Copies the images
	CopyImage(image, name, diskSize)

	// Genereate Cloud Config
	cloudConfigFile, _ := CreateCloudConfig(name, cloudConfig)

	// TODO: Infer image type from given Disk Images with "qemu-img info --output=json disk.img"

	xmlString := `
		<domain type='kvm'>
			<name>`+ name +`</name>
			<memory>`+ strconv.Itoa(ram)  +`</memory>
			<vcpu>`+ strconv.Itoa(cpus) +`</vcpu>
			<os>
				<type>hvm</type>
				<boot dev="hd" />
			</os>
			<devices>
				<graphics type='vnc' port='-1'/>
				<disk type='file' device='disk'>
				    <driver name='qemu' type='qcow2'/>
					<source file='` + GetVMPrimaryDiskName(name) + `'/>
					<target dev='vda' bus='virtio'/>
				</disk>
				<disk type='file' device='disk'>
					<source file='` + cloudConfigFile + `'/>
					<target dev='vdb' bus="virtio"/>
				</disk>
				<interface type='network'>
					<source network='` + network + `'/>
					<mac address='` + GenerateMAC()+ `'/>
					<model type="virtio" />
				</interface>
		`

	xmlString += `
			</devices>
	</domain>`

	// Write XML file for reference purposes
	// ioutil.WriteFile(userDataFilePath , []byte(userData), 0777 )
	println(xmlString)

	dom, err := conn.DomainCreateXML(xmlString, 0)
	defer dom.Free()

	if err == nil {
		err = dom.Create()

		// Construct a new Network object to return it
		myVM := &VM{}
		respXml, _ := dom.GetXMLDesc(0);

		xml.Unmarshal([]byte(respXml), myVM)
		return myVM, nil
	} else {
		return nil, err
	}
}

func GetVM(name string) *VM {
	// TODO: Non existent check
	conn := BuildConnection()
	defer conn.CloseConnection()

	dom, _ := conn.LookupDomainByName(name)
	defer dom.Free()

	myVM := &VM{}
	respXml, _ := dom.GetXMLDesc(0);
	xml.Unmarshal([]byte(respXml), myVM)
	isActive, _ := dom.IsActive()
	myVM.Active = isActive

	return myVM
}

func StartVM(name string) {
	conn := BuildConnection()
	defer conn.CloseConnection()

	dom, _ := conn.LookupDomainByName(name)
	defer dom.Free()

	dom.Create()
}

func StopVM(name string) {
	conn := BuildConnection()
	defer conn.CloseConnection()

	dom, _ := conn.LookupDomainByName(name)
	defer dom.Free()

	dom.Shutdown()
}