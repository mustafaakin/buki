package buki

import (
	"gopkg.in/alexzorin/libvirt-go.v2"
	"math/rand"
	"fmt"
	"time"
)

func BuildConnection() libvirt.VirConnection {
	conn, err := libvirt.NewVirConnection("qemu:///system")
	if err != nil {
		panic(err)
	}
	return conn
}

func GenerateMAC() string {
	// Reference: https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/5/html/Virtualization/sect-Virtualization-Tips_and_tricks-Generating_a_new_unique_MAC_address.html
	mac := [6]byte{};
	mac[0] = 0x00
	mac[1] = 0x16
	mac[2] = 0x3e
	
	// Make sure we have a good seed right now, not sure it is the best way for that
	rand.Seed( time.Now().UTC().UnixNano())
	mac[3] = byte(rand.Intn(255))
	mac[4] = byte(rand.Intn(255))
	mac[5] = byte(rand.Intn(255))
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", mac[0], mac[1], mac[2], mac[3], mac[4], mac[5])
}