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
}

type Bridge struct {
	Name  string         `xml:"name,attr" json:"name"`
}

func GetNetworks() []Network{
	conn := BuildConnection()
	// Do not forget to close connection
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