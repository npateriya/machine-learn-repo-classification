routerdefense
=============

Router Defense deep dives into Cisco routers and switches configuration and do security recommandations.
It gives the opportunity to audit network devices in a quick, efficient way and actionable practices.

The tool has been released at the BRUCON 2010 conference and includes around 140 tests.
It requires Python 2.7.

Project status: Abandoned

## Linux installation

Install reportlab to render PDF reports.

## Windows installation

Download & install http://www.python.org/ftp/python/2.7.5/python-2.7.5.msi
Downoload https://bitbucket.org/pypa/setuptools/raw/bootstrap/ez_setup.py
Execute c:\python27\python.exe ez_setup.py
Execute c:\python27\Scripts\easy_install.exe reportlab

## Help

```
./run_routerdefense -h
usage: run_routerdefense [-h] --iosconfig IOSCONFIG
                         [--iosplatform IOSPLATFORM] [--ip4inbound IP4INBOUND]
                         [--ip4outbound IP4OUTBOUND]
                         [--reportformat REPORTFORMAT]
                         [--reportfilename REPORTFILENAME]

Cisco IOS configuration security assessment tool.

optional arguments:
  -h, --help            show this help message and exit
  --iosconfig IOSCONFIG
                        Cisco IOS configuration file.
  --iosplatform IOSPLATFORM
                        Cisco IOS platform: router, switch or both (multiple
                        layers like Cat4k or 6k).
  --ip4inbound IP4INBOUND
                        Authorized nodes to do VTY inbound sessions (Telnet,
                        SSH). Possibles values examples (comma-separated) :
                        192.168.100.0/24, 192.168.1.1
  --ip4outbound IP4OUTBOUND
                        Authorized nodes to communicate with the device in the
                        outbound direction: SNMP, SYSLOG. Possibles values
                        examples (comma-separated) : 192.168.100.0/24,
                        192.168.1.1
  --reportformat REPORTFORMAT
                        Report format support: stdout, csv, html, pdf.
  --reportfilename REPORTFILENAME
                        Report filename.
```

##Demo

./run_routerdefense --iosconfig /tmp/show_run.txt 

______            _             ______      __
| ___ \          | |            |  _  \    / _|
| |_/ /___  _   _| |_ ___ _ __  | | | |___| |_ ___ _ __  ___  ___
|    // _ \| | | | __/ _ \ '__| | | | / _ \  _/ _ \ '_ \/ __|/ _ \
| |\ \ (_) | |_| | ||  __/ |    | |/ /  __/ ||  __/ | | \__ \  __/
\_| \_\___/ \__,_|\__\___|_|    |___/ \___|_| \___|_| |_|___/\___|

=[ Cisco IOS security assessment tool
=[ https://github.com/pello/routerdefense 
=[ version 2013.12


=[ Generic information

    => Hostname: LAB-A
    => IOS version: 11.1
    => Switching: Fast switching
    => Multicast: Disabled
    => QoS: Disabled
    => IPv6: Disabled
    => IPSEC VPN: Disabled
=[Management plane]=
The management plane consists of functions that achieve the management goals of the network. This includes interactive management sessions using SSH, as well as statistics-gathering with SNMP or NetFlow. When you consider the security of a network device, it is critical that the management plane be protected. If a security incident is able to undermine the functions of the management plane, it can be impossible for you to recover or stabilize the network.


=[ CDP

    => What: Leak network topology
    => Threat: Someone could leak information network topology
    => Patch impact: The NMS could stop gathering information from this device
    => Score: 1.9/10
    => How to fix: 
    Run the command "no cdp run" to globally disable CDP.
    Or, run the "no cdp enable" command on the following interfaces if they are not needed:
    Ethernet0, Ethernet1, Serial0, Serial1, Ethernet0, Serial0, Serial1, BRI0, Ethernet0, Serial0, Serial1, BRI0, Ethernet0, Serial0, Serial1, Ethernet0, Serial0, Serial1.

    CDP protocol is already disabled on those interfaces:
    .

Number of threatInfo(s) to fix: 1/1
...
=[ summary ]=

Management Plane

CDP: 1/1
LLDP: 1/1
Console port: 2/2
Aux port: 4/4
MOTD banner: 1/2
LOGIN banner: 1/2
EXEC banner: 1/2
IOS TCP/UDP services: 15/15
CPU/Memory: 9/9
Exceptions/crashes: 1/1
Passwords and authentication management: 1/4
Management protection: 5/7
Tacacs+ servers redundancy: 1/1
Tacacs+ authentication: 3/3
Tacacs+ authorization: 4/5
Tacacs+ accounting: 4/5
SNMP: 3/9
Syslog: 8/9
Configuration Replace/Rollback: 3/5
Vty lines: 4/5

Control Plane

ICMPv4 unreachable: 1/1
ARP proxy: 1/1
NTP: 1/1
TCP: 1/1
BGP: 0/6
EIGRP: 0/4
RIP: 0/1
OSPF: 0/5
GLBP: 0/1
HSRP: 0/1
VRRP: 0/1
TCLSH shell scripting: 1/1

Data Plane

ICMPv4 redirects: 1/1
IPv4 Options: 1/1
IPv4 source route: 1/1
ICMP deny any any: 1/1
IPv4 fragments: 1/1
Unicast Reverse Path Forwarding (IPv4): 1/1
Netflow: 1/1
Port Security: 3/5
Level 2: 3/9

