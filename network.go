package buki
import (
	"encoding/xml"
)
/*
<network>
  <name>default</name>
  <uuid>fb6e64fa-4d11-41f0-b623-378374fa6c47</uuid>
  <forward mode='nat'>
    <nat>
      <port start='1024' end='65535'/>
    </nat>
  </forward>
  <bridge name='virbr0' stp='on' delay='0'/>
  <ip address='192.168.122.1' netmask='255.255.255.0'>
    <dhcp>
      <range start='192.168.122.2' end='192.168.122.254'/>
    </dhcp>
  </ip>
</network>
 */

type Network struct {
	XMLName xml.Name       `xml:"network" json:"-"`
	Name    string	       `xml:"name" json:"name"`
	UUID    string	       `xml:"uuid" json:"uuid"`
	Forward Forward        `xml:"forward" json:"forward"`
	Bridge	Bridge         `xml:"bridge" json:"bridge"`
	IPs		[]IPAddress    `xml:"ip" json:"ip"`
	Active	bool           `json:"active"`
}

type Forward struct {
	Mode    string         `xml:"mode,attr" json:"mode"`
}

type IPAddress struct {
	Address string         `xml:"address,attr" json:"address"`
	Netmask string         `xml:"netmask,attr" json:"netmask"`
	IPRange DHCPRange		`xml:"dhcp>range" json:"range"`
}

type Bridge struct {
	Name  string         `xml:"name,attr" json:"name"`
}

type DHCPRange struct {
	Start  string   `xml:"start,attr" json:"start"`
	End    string   `xml:"end,attr" json:"end"`
}

// GetNetworks allows the tools 
func GetNetworks() []Network{
	conn := BuildConnection()
	defer conn.CloseConnection()

	nets, _ := conn.ListAllNetworks(0)

	Networks := make([]Network, len(nets))

	for idx, net := range nets {
		xmlResp, _ := net.GetXMLDesc(0)
		// ignoring error right now
		xml.Unmarshal([]byte(xmlResp), &Networks[idx])
		isActive, _ := net.IsActive()
		Networks[idx].Active = isActive
	}

	return Networks
}

// CreateNATNetwork creates a new NAT network
func CreateNATNetwork(name, address, netmask, start, end string) (*Network, error) {
	// TODO: Add error cheks
	conn := BuildConnection()

	xmlString :=
		`<network>
			<name>` + name + `</name>
			<bridge name="` + name + `"/>
			<forward/>
			<ip address="` + address + `" netmask="` + netmask +`">
			    <dhcp>
      				<range start='` + start + `' end='` + end + `'/>
    			</dhcp>
			</ip>
	    </network>`

	net, err := conn.NetworkDefineXML(xmlString)

	defer func(){
		conn.CloseConnection()
		net.Free()
	}()

	if err == nil {
		err = net.Create()
		net.SetAutostart(true)

		// Construct a new Network object to return it
		myNet := &Network{}
		respXml, _ := net.GetXMLDesc(0);

		xml.Unmarshal([]byte(respXml), myNet)
		return myNet, nil
	} else {
		return nil, err
	}
}

func CreateBridgedNetwork(){
	// TODO: Add this
}

func DeleteNetwork(uuid string){
	// Stops and deletes the network
}