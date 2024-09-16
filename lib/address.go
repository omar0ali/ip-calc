package network

import (
	"fmt"
	"strings"
)

type Address struct {
	address [4]Octet
	subnet  [4]Octet
}

/*
Creates Address in Binary,
*/
func CreateAddressInBinary(a string, b string, c string, d string) Address {
	o1, err := CreateOctet(a)
	if err != nil {
		panic(err)
	}
	o2, err := CreateOctet(b)
	if err != nil {
		panic(err)
	}
	o3, err := CreateOctet(c)
	if err != nil {
		panic(err)
	}
	o4, err := CreateOctet(d)
	if err != nil {
		panic(err)
	}
	return Address{
		address: [4]Octet{o1, o2, o3, o4},
	}
}

/*
Creates Subnet in Binary
*/
func (adrs *Address) SetSubnetInBinary(a string, b string, c string, d string) {
	o1, err := CreateOctet(a)
	if err != nil {
		panic(err)
	}
	o2, err := CreateOctet(b)
	if err != nil {
		panic(err)
	}
	o3, err := CreateOctet(c)
	if err != nil {
		panic(err)
	}
	o4, err := CreateOctet(d)
	if err != nil {
		panic(err)
	}
	adrs.subnet = [4]Octet{o1, o2, o3, o4}
}

/*
Creates Address in Decimal,
*/
func CreateAddress(a uint8, b uint8, c uint8, d uint8) Address {
	return CreateAddressInBinary(toStr(a), toStr(b), toStr(c), toStr(d))
}

func (adrs *Address) SetSubnet(a uint8, b uint8, c uint8, d uint8) {
	adrs.SetSubnetInBinary(toStr(a), toStr(b), toStr(c), toStr(d))
}

/*
Converts binary octet to decimal
*/
func toStr(octet uint8) string {
	z := [8]uint8{128, 64, 32, 16, 8, 4, 2, 1}
	var builder strings.Builder
	for _, j := range z {
		if octet >= j {
			builder.WriteString("1")
			octet = octet - j
		} else {
			builder.WriteString("0")
		}
	}
	return builder.String()
}

func (a Address) GetIPAddress() string {
	return fmt.Sprintf("%v.%v.%v.%v", a.address[0].GetDecimal(), a.address[1].GetDecimal(), a.address[2].GetDecimal(), a.address[3].GetDecimal())
}

func (a Address) GetSubnet() string {
	if a.subnet[0].GetDecimal() == 0 {
		return "subnet hasn't been set yet"
	}
	return fmt.Sprintf("%v.%v.%v.%v", a.subnet[0].GetDecimal(), a.subnet[1].GetDecimal(), a.subnet[2].GetDecimal(), a.subnet[3].GetDecimal())
}

// This is the Classless Inter-Domain Routing
func (a Address) GetCIDR() uint8 {
	var count uint8 = 0
	for _, j := range a.subnet {
		for _, k := range j.octet {
			if k == 1 {
				count++
			}
		}
	}
	return count
}

// Get the broadcast address in Binary
func (a Address) GetBroadCastAddresInBinary() [4]Octet {
	var broadcastAddress [4]Octet
	var err error
	for i := 0; i < 4; i++ {
		var builder strings.Builder
		for j := 0; j < 8; j++ {
			if a.subnet[i].octet[j] == 1 {
				builder.WriteString(fmt.Sprintf("%v", a.address[i].octet[j]))
			} else {
				builder.WriteString(string("1"))
			}
		}
		broadcastAddress[i], err = CreateOctet(builder.String())
		if err != nil {
			panic(err)
		}
	}
	return broadcastAddress
}

// Get the broadcast address in Decimal

func (a Address) GetBroadCastAddres() string {
	broadcastAddress := a.GetBroadCastAddresInBinary()
	return fmt.Sprintf("%v.%v.%v.%v", broadcastAddress[0].GetDecimal(), broadcastAddress[1].GetDecimal(), broadcastAddress[2].GetDecimal(), broadcastAddress[3].GetDecimal())
}

// Getting the Network Address in Binary

func (a Address) GetNetworkAddresInBinary() [4]Octet {
	var networkAddress [4]Octet
	var err error
	for i := 0; i < 4; i++ {
		var builder strings.Builder
		for j := 0; j < 8; j++ {
			if a.subnet[i].octet[j] == 1 {
				builder.WriteString(fmt.Sprintf("%v", a.address[i].octet[j]))
			} else {
				builder.WriteString(string("0"))
			}
		}
		networkAddress[i], err = CreateOctet(builder.String())
		if err != nil {
			panic(err)
		}
	}
	return networkAddress
}

func (a Address) GetNetworkAddres() string {
	networkAddress := a.GetNetworkAddresInBinary()
	return fmt.Sprintf("%v.%v.%v.%v", networkAddress[0].GetDecimal(), networkAddress[1].GetDecimal(), networkAddress[2].GetDecimal(), networkAddress[3].GetDecimal())
}

// Get range of address
func (a Address) GetRangeOfAvailableHosts() string {
	networkAddress := a.GetNetworkAddresInBinary()
	broadcastAddress := a.GetBroadCastAddresInBinary()
	firstAddress := fmt.Sprintf("%v.%v.%v.%v", networkAddress[0].GetDecimal(), networkAddress[1].GetDecimal(), networkAddress[2].GetDecimal(), networkAddress[3].GetDecimal()+1)
	lastAddress := fmt.Sprintf("%v.%v.%v.%v", broadcastAddress[0].GetDecimal(), broadcastAddress[1].GetDecimal(), broadcastAddress[2].GetDecimal(), broadcastAddress[3].GetDecimal()-1)
	return fmt.Sprintf("%v - %v", firstAddress, lastAddress)
}

func (a Address) GetTotalHosts() uint {
	return uint(1 << (32 - (a.GetCIDR()))) // 1 << shifting bitwise to the left, is working as powering by two.
}

func (a Address) GetUsableHosts() uint {
	return a.GetTotalHosts() - 2
}
