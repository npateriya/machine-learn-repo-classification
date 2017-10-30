Cisco IOS Configuration Generator
=================================

This project consists of a collection of python tools that can be used for auto-generation of configuration files used on Cisco IOS devices, that could be incorporated in scripts for automating configuration of multiple devices.

Utilities include:

- `csinit.py`: device initialization tool
- `csint.py`: interface configuration tool

# csinit.py

Generator utility for Cisco IOS device initial configuration

```
usage: csinit.py [-h] [-d DOMAIN] [-l LOGIN] [-t] [-s] [-x6] [-K]
                 hostname console_pass enable_pass
```

Positional arguments:

```
  hostname              Hostname
  console_pass          Console Password
  enable_pass           Enable Secret
```

Optional arguments:

```
  -h, --help            show this help message and exit
  -d DOMAIN, --domain DOMAIN
                        Domain Name
  -l LOGIN, --login LOGIN
                        Login Credentials - username:password
  -t, --telnet          Enable Telnet; use -l username:password
  -s, --ssh             Enable SSH; use -l username:password
  -x6, --noipv6         Disable IPv6 unicast-routing
  -K, --packettracer    Packet Tracer syntax for key generation
```

# csint.py

Generator utility for Cisco IOS interface configuration

```
usage: csint.py [-h] [options] interface
```

Optional arguments:

```
  -h, --help            show this help message and exit
  -d DESC, --desc DESC  Description
  -e ENCAP, --encap ENCAP
                        Encapsulation mode; can use with -v
  -4 IPV4, --ipv4 IPV4  IPv4 address/CIDR
  -6 IPV6, --ipv6 IPV6  IPv6 address/CIDR
  -l LLOCAL, --llocal LLOCAL
                        IPv6 link-local address
  -m NETMASK, --netmask NETMASK
                        Specify IPv4 netmask instead of CIDR
  -sA, --access         Switchport mode access; use with -v to specify VLAN
  -sN, --native         Switchport mode trunk; use with -v to specify VLAN
  -sT, --trunk          Switchport mode trunk
  -v VLAN, --vlan VLAN  VLAN number; use with -a,-n,-e dot1q
  -X, --shutdown        Shutdown interface (default: no shutdown)
  -x, --exit            Add exit to end of interface configuration
```

### Examples:

**Interface g0/1 with IPv4 address using CIDR or dotted decimal notation:**

```
csint.py g0/1 -4 172.16.0.1/24
csint.py g0/1 -4 172.16.0.1 255.255.255.0
```

Will generate:

```
interface g0/1
 ip address 172.16.0.1 255.255.255.240
 no shutdown
```

**Interface lo1 with IPv6 address:**

```
csint.py lo1 -6 2001:db8:cafe:a::1/64 -l fe80::1
```

Will generate:

```
interface lo1
 ipv6 address 2001:db8:cafe:a::1/64
 ipv6 address fe80::1 link-local
 no shutdown
```

**Interface range configured with switchport access VLAN:**

```
csint.py "range g0/2-24" -sA -v 20
```
