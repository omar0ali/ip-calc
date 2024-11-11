# ip-calc
## IP Calculator
This command-line tool helps you analyze IP networks. It allows you to:
1. View detailed information about an IP address and its network (including subnet mask, network address, CIDR notation, broadcast address, and the range of usable hosts).
2. Divide a given IP network into smaller subnets, allowing you to create and visualize multiple subnets based on a specified number of divisions.


### Download the Pre-release:
You can download the pre-release version from the [GitHub Releases Page](https://github.com/omar0ali/ip-calc/releases/tag/v0.9).

>[!NOTE]
The subnet division feature is in an experimental stage (alpha), and the calculations may not always be accurate. Use at your own risk.

## Setup

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

### Usage

```
Usage:
        ip-calc <ip-address>/CIDR
          - Displays details of the IP address and network.
        ip-calc <ip-address>/CIDR <subnet-count>
          - Displays details of the IP address and network, and divides the subnet into the specified number of smaller subnets.

Example 1: ip-calc 192.168.0.0/24
        Displays details of the network: IP address, subnet mask, CIDR, broadcast address, range of hosts, total hosts, and usable hosts.

Example 2: ip-calc 192.168.0.0/24 4
        Divides the network into 4 subnets and shows the subnets' range, broadcast address, and CIDR.
```

```bash
go run . 192.168.0.0/24
```

**OUTPUT**

```
Internet Protocol: 192.168.0.0
Subnetmask: 255.255.255.0
Network Address: 192.168.0.0
CIDR: 24
Broadcast Address: 192.168.0.255
Range: 192.168.0.1 - 192.168.0.254
Total Hosts: 256
Usable Hosts: 254
```