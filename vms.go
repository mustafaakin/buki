package buki

func ListVM(){

}

// CreateVM creates VM with given parameters 
func CreateVM(image string, name string, cpus, ram int, disks []string, networks []string){
	// TODO: Escape all params
	// CDROM with meta-data and user-data, just mount folder

	conn := BuildConnection()
	ram = ram * 1024;

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

