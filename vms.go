package buki
import "encoding/xml"

/*
<domain type='kvm' id='16'>
  <name>myubuntu</name>
  <uuid>dfbdd086-44d9-4dd4-8bf4-1e81b2364dde</uuid>
  <memory unit='KiB'>1048576</memory>
  <currentMemory unit='KiB'>1048576</currentMemory>
  <vcpu placement='static'>2</vcpu>
  <resource>
    <partition>/machine</partition>
  </resource>
  <os>
    <type arch='x86_64' machine='pc-i440fx-trusty'>hvm</type>
    <boot dev='hd'/>
  </os>
  <features>
    <acpi/>
  </features>
  <clock offset='utc'/>
  <on_poweroff>destroy</on_poweroff>
  <on_reboot>restart</on_reboot>
  <on_crash>destroy</on_crash>
  <devices>
    <emulator>/usr/bin/kvm-spice</emulator>
    <disk type='file' device='disk'>
      <driver name='qemu' type='qcow2' cache='none'/>
      <source file='/home/mustafa/buki/vms/myubuntu/disk0.img'/>
      <target dev='vda' bus='virtio'/>
      <alias name='virtio-disk0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x0'/>
    </disk>
    <disk type='file' device='disk'>
      <driver name='qemu' type='raw'/>
      <source file='/home/mustafa/buki/vms/myubuntu/cloud-init.img'/>
      <target dev='vdb' bus='virtio'/>
      <alias name='virtio-disk1'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x0'/>
    </disk>
    <controller type='usb' index='0'>
      <alias name='usb0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x01' function='0x2'/>
    </controller>
    <controller type='pci' index='0' model='pci-root'>
      <alias name='pci.0'/>
    </controller>
    <interface type='network'>
      <mac address='24:42:53:21:52:45'/>
      <source network='default'/>
      <target dev='vnet0'/>
      <model type='virtio'/>
      <alias name='net0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x03' function='0x0'/>
    </interface>
    <serial type='pty'>
      <source path='/dev/pts/0'/>
      <target port='0'/>
      <alias name='serial0'/>
    </serial>
    <console type='pty' tty='/dev/pts/0'>
      <source path='/dev/pts/0'/>
      <target type='serial' port='0'/>
      <alias name='serial0'/>
    </console>
    <input type='mouse' bus='ps2'/>
    <input type='keyboard' bus='ps2'/>
    <graphics type='vnc' port='5900' autoport='yes' listen='127.0.0.1'>
      <listen type='address' address='127.0.0.1'/>
    </graphics>
  </devices>
</domain>

 */

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
// CreateVM creates VM with given parameters 
func CreateVM(image string, name string, cpus, ram int, disks []string, networks []string, cloudConfig string){
	// TODO: Escape all params
	conn := BuildConnection()

	xmlString := `
		<domain type='kvm'>
			<name>`+ name +`</name>
			<memory>`+ string(ram)  +`</memory>
			<vcpu>`+ string(cpus) +`</vcpu>
			<os>
				<type>hvm</type>
				<boot dev="hd" />
			</os>
			<devices>
				<graphics type='vnc' port='-1'/>
		`

	// Add disks
	for idx, disk := range disks {
		dev := "vd" + string('a' + idx);
		xmlString += `
				<disk type='file' device='disk'>
					<source file='` + disk + `'/>
					<target dev='` + dev + `'/>
				</disk>
		`
	}

	for _, network := range networks {
		xmlString += `
				<interface type='network'>
					<source network='` + network + `'/>
					<mac address='` + GenerateMAC()+ `'/>
				</interface>
		`
	}

	// Add network interfaces
	xmlString += `
			</devices>
	</domain>`

	conn.DomainCreateXML(xmlString, 0)
}

