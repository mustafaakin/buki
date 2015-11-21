package buki

import "gopkg.in/alexzorin/libvirt-go.v2"

func BuildConnection() libvirt.VirConnection {
	conn, err := libvirt.NewVirConnection("qemu:///system")
	if err != nil {
		panic(err)
	}
	return conn
}