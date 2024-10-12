package main

import (
	"fmt"

	lib "github.com/omar0ali/ip-calc/lib"
)

func main() {
	address := lib.CreateAddress(10, 0, 0, 0).SetCIDR(16)
	// address.SetSubnet(255, 255, 255, 0)
	fmt.Printf("Internet Protocal: %v\n", address.GetIPAddress())
	fmt.Printf("Subnetmask: %v\n", address.GetSubnet())
	fmt.Printf("Network Address: %v\n", address.GetNetworkAddres())
	fmt.Printf("CIDR: %v\n", address.GetCIDR())
	fmt.Printf("Broadcast Address: %v\n", address.GetBroadCastAddres())
	fmt.Printf("Range: %v\n", address.GetRangeOfAvailableHosts())
	fmt.Printf("Total Hosts: %v\n", address.GetTotalHosts())
	fmt.Printf("Usable Hosts: %v\n", address.GetUsableHosts())
	listOfaddresses := address.DivideEvenlyBy(4) //Still work in progress
	fmt.Printf("---- Division Evenly by %v ----\n", len(listOfaddresses))
	for _, i := range listOfaddresses {
		fmt.Println(i.GetNetworkAddres(), "/", i.GetCIDR(), "<-Range->", i.GetRangeOfAvailableHosts())
	}
}
