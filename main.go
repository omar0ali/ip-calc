package main

import (
	"fmt"

	"os"
	"strconv"
	"strings"

	lib "github.com/omar0ali/ip-calc/lib"
)

func main() {
	// read input from user
	args := os.Args
	if len(args) < 2 || len(args) > 3 {
		fmt.Println("Usage:")
		fmt.Println("\tip-calc <ip-address>/CIDR\n\t  - Displays details of the IP address and network.")
		fmt.Println("\tip-calc <ip-address>/CIDR <subnet-count>\n\t  - Displays details of the IP address and network, and divides the subnet into the specified number of smaller subnets.")
		fmt.Println("\nExample 1: ip-calc 192.168.0.0/24")
		fmt.Println("\tDisplays details of the network: IP address, subnet mask, CIDR, broadcast address, range of hosts, total hosts, and usable hosts.")
		fmt.Println("\nExample 2: ip-calc 192.168.0.0/24 4")
		fmt.Println("\tDivides the network into 4 subnets and shows the subnets' range, broadcast address, and CIDR.")
		return
	}

	var (
		division     int
		err          error
		cidr_block   int
		octetsValues []uint8
	)

	network := strings.Split(args[1], "/")
	cidr_block, err = strconv.Atoi(network[1])
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	octets := strings.Split(network[0], ".")
	for _, octet := range octets {
		value, err := strconv.ParseUint(octet, 10, 8)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}
		octetsValues = append(octetsValues, uint8(value))
	}

	address := lib.CreateAddress(octetsValues[0], octetsValues[1], octetsValues[2], octetsValues[3]).SetCIDR(uint8(cidr_block))
	fmt.Printf("Internet Protocal: %v\n", address.GetIPAddress())
	fmt.Printf("Subnetmask: %v\n", address.GetSubnet())
	fmt.Printf("Network Address: %v\n", address.GetNetworkAddres())
	fmt.Printf("CIDR: %v\n", address.GetCIDR())
	fmt.Printf("Broadcast Address: %v\n", address.GetBroadCastAddres())
	fmt.Printf("Range: %v\n", address.GetRangeOfAvailableHosts())
	fmt.Printf("Total Hosts: %v\n", address.GetTotalHosts())
	fmt.Printf("Usable Hosts: %v\n", address.GetUsableHosts())
	//Still work in progress
	if len(args) > 2 {
		division, err = strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("*Skipped divisoin.")
		} else {
			listOfaddresses := address.DivideEvenlyBy(uint8(division))
			fmt.Printf("---- Division Evenly by %v ----\n", len(listOfaddresses))
			for _, i := range listOfaddresses {
				fmt.Println(i.GetNetworkAddres(), "/", i.GetCIDR(), "<-Range->\t", i.GetRangeOfAvailableHosts(), "\tBroadCast:", i.GetBroadCastAddres())
			}
		}
	}
}
