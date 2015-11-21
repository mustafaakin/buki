package buki

func ListVM(){

}

func CreateVM(image string, name string, cpus, ram int, disks []string, networks []string){
	// TODO: Escape all params

	conn := BuildConnection()
	ram = ram * 1024;

	xmlString := `
		<domain type='kvm'>
			<name>`+ name +`</name>
			<memory>`+ string(ram)  +`</memory>
			<vcpu>`+ string(cpus) +`</vcpu>
			<os>
				<type arch="i686">hvm</type>
				<boot dev="hd" />
			</os>
			<devices>
				<graphics type='vnc' port='-1'/>
		`

	// Add disks
	for _, elem := range disks {
		diskName := "hd";
		xmlString += `
			<disk type='file' device='disk'>
				<source file='` + elem + `'/>
				<target dev='` + diskName + `'/>
			</disk>
		`
	}


 `
			</devices>
	</domain>`

	conn.DomainCreateXML(xmlString, 0)

}

