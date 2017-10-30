# CCNA Class Notes
Some notes that might came in handy with the 'Cisco Certified Network Associate' or CCNA course.

WARNING: WORK IN PROGRESS, FORMATTING AND INFORMATION IS ALL VERY POOR!

## Network Basic's

**Topology**

**OSI Model**

**Binary**

**TCP vs. UDP**

**Packets**

All data sent to and from any computer or server is formatted and split into into packets. They have structure beginning with the header, this contains information like source and destination network addresses, error detection information and the data itself. There is a difference between TCP and UDP packets, UDP is 

**NAT or Network Address Translation**

## Devices

**Router**

**Switch**

**Server**

**End-Point**

## IPv4 Basic's

**IP Address**

This address is assigned to hosts. A host in computer network theory is simply a device that has an IP assigned to it, anything from a computer, a phone or even a router or router interfaces. All IPv4 addresses are split up into 4 sections called "octets" these contain 8 bits which can range from (0 - 255). This is sort of like an ID, so that packets can find their way to you.

*192.168.100.25* is an example of what one could look like

**Subnet Mask**

This mask is what we use to devide different networks. The 1's represent the network portion, while the 0's represent the host portion. If you take the binary from subnet mask and the IP and place them vertically, this help it make more sense. If a number 
```
11111111.11111111.11111111.11110000
255     .255     .255     .240
```

**Network Portion and Host Portion**

This is the most important thing you need to understand

```
11000000.10101000.01100100.00011001
192     .168     .100     .25
```

**IP Classes**

There are IP addresses reserved for private use by internal networks (such as your home and office). These address ranges CANNOT be used on the public 

**Network Address**

The first address in the network, not assigned to anything, just used to describe the subnetwork as a whole in topologies/network documentation.
```
11000000.10101000.01100100.00010000
192     .168     .100     .16
```

**Broadcast Address**

The final address in a subnetwork, It is reserved and cannot be assigned to a host/device.
```
11000000.10101000.01100100.00011001
00000000.00000000.00000000.00001111
11000000.10101000.01100100.00011111
192.168.100.31
```

**Default Gateway**

The IP address of the first router interface your computer goes to (never a switch). Basically if your computer goes to a switch and then into router interface g0/1, which the IP is set to 192.168.1.1. That IP is your default gateway, since it's the first router you reach.

**"CIDR notation" or "What does the slash mean?**

The network portion from the left in other words the amount of ones from left to right e.g.

192.168.0.0 /29 =

11111111.11111111.11111111.11111000



## IPv6 Basic's

...

## Subnet Mathematics
Host bits are zeros to the right of the subnet mask.
2^(number of host bits) -2 = Number of hosts.

## Understanding Cisco IOS
**Normal Mode**



**Enable Mode**



**Config Mode**




## Commands

Don't forget to take anything between the square brackets and replace it with actual information that applies to your situation.

**Show Commands**

Used to show certain information set

**Setting Hostname**

Give a name to your router or switch, so you can identify it.
```
(config)# hostname [NAME]
```

**Setting MOTD**

Display a messsage upon login.
```
(config)# banner motd
```


config > line console 0 > password ... > login > exit  -  add access to console line

config > enable password > exit  -  set/change a static enable password

config > enable secret > exit - add password to secure access

config > interface g0/0 or fe0/0 or vlan1 > ip address ... ... > no shut > exit  -  config IP to interface



service password-encryption

show run  /  show ip int brief

int [whatever] > ipv6 address fe80::2 link-local  -  set link local



---Default Route used when there is no known network to forward to---
(config)# ip route 0.0.0.0 0.0.0.0 ip add [IP address of where it is going or the exit interface]


---Static Route uses its own exit interface---
(config)# ip route [destination address] [destination subnet mask] [exit interface of current device]

# Resources
http://networkengineering.stackexchange.com/questions/7106/how-do-you-calculate-the-prefix-network-subnet-and-host-numbers

http://www.cisco.com/c/en/us/td/docs/ios/12_2/security/configuration/guide/fsecur_c/scfpass.pdf
