# ip-calc

### How to use

1. Make sure golang is installed.

    - You can view the instructions from here.
        - [https://go.dev/doc/install](https://go.dev/doc/install)

2. Clone the repo

```bash
git clone git@github.com:omar0ali/ip-calc.git
```

3. Get into the ip-calc directory

```bash
cd ip-calc
```

4. Before starting the app, you can view `main.go` get to know how everything is working, you can modify and set ip address to see the results

```golang
...
func main() {
	address := lib.CreateAddress(10, 0, 0, 0).SetCIDR(24)
	// address.SetSubnet(255, 255, 255, 0)
	fmt.Printf("Internet Protocal: %v\n", address.GetIPAddress())
	fmt.Printf("Subnetmask: %v\n", address.GetSubnet())
	fmt.Printf("Network Address: %v\n", address.GetNetworkAddres())
	fmt.Printf("CIDR: %v\n", address.GetCIDR())
	fmt.Printf("Broadcast Address: %v\n", address.GetBroadCastAddres())
	fmt.Printf("Range: %v\n", address.GetRangeOfAvailableHosts())
	fmt.Printf("Total Hosts: %v\n", address.GetTotalHosts())
	fmt.Printf("Usable Hosts: %v\n", address.GetUsableHosts())
    // The following block is not fully tested
    listOfaddresses := address.DivideEvenlyBy(4)
    fmt.Printf("---- Division Evenly by %v ----\n", len(listOfaddresses))
	for _, i := range listOfaddresses {
		fmt.Println(i.GetNetworkAddres(), "/", i.GetCIDR(), "<-Range->", i.GetRangeOfAvailableHosts())
	}
}
```

-   Example **Output**

```
Internet Protocal: 10.0.0.0
Subnetmask: 255.255.0.0
Network Address: 10.0.0.0
CIDR: 16
Broadcast Address: 10.0.255.255
Range: 10.0.0.1 - 10.0.255.254
Total Hosts: 65536
Usable Hosts: 65534
---- Division Evenly by 4 ----
10.0.0.0 / 18 <-Range-> 10.0.0.1 - 10.0.63.254
10.0.64.0 / 18 <-Range-> 10.0.64.1 - 10.0.127.254
10.0.128.0 / 18 <-Range-> 10.0.128.1 - 10.0.191.254
10.0.192.0 / 18 <-Range-> 10.0.192.1 - 10.0.255.254
```

5. To start the app

```bash
go run .
```

### Features to add

-   [ ] Use https://github.com/spf13/cobra for more CLI features i.e `go run . 10.0.0.0/16` instead of relaying on editing main.go file every time.
-   [x] Divide a subnet equally i.e `go run . 10.0.0.0 255.0.0.0 --divide 4` this is a CIDR 16
-   [ ] Writing test cases
