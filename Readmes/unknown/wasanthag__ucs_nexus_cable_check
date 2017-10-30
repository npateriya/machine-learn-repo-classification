# cisco_ucs_cable_patching
Script to verify patch connections between cisco ucs C series servers and nexus switches

This script verifies the cabling between a Cisco UCS server and upstream Nexus switches.
This is helpful when you want to confirm what your patch plan says in paper is really how 
it is cabled. This script assumes each UCS C series server is connected to two upstream Nexus 
switches(leaf1 and leaf2) for redundancy. Common server deployment has MLOM and PCIE VIC cards for high availability.
The script will log in to the Upstream Nexus switches and use LLDP neighbors to map out the MAC
addresses seen by the switches with what it extract from the UCS CIMC. Then it compares them and maps
out which server port is connected to which switch port.
Example:
Following server has MLOM (UCS VIC 1225) and PCIE(UCS VIC 1227) Cisco VICs. For redundancy bonding is configured
across first port of MLOM (E0) going to leaf 1 and first port of PCIE (E2) going to leaf2. Similar cabling is done for
E1 and E3. The script confirms the actual physical cabling with LLDP.
```
Logging in to server server3-cimc.local.
Successfully Logged in to the CIMC
Getting MAC Address details for MLOM NICS
MAC for MLOM1 is a8:9d:21:6b:04:9a
MAC for MLOM2 is a8:9d:21:6b:04:9e
E0 with MAC a8:9d:21:6b:04:9a is connected to Eth1/11 on leaf 1
E1 with MAC a8:9d:21:6b:04:9e is connected to Eth1/35 on leaf 2
Getting MAC Address details for PCIE NICS
MAC for E2 is 74:a2:e6:7a:a5:32
MAC for E3 is 74:a2:e6:7a:a5:36
E2 with MAC a8:9d:21:6b:04:9a is connected to Eth1/11 on leaf 2 
E3 with MAC 74:a2:e6:7a:a5:36 is connected to Eth1/35 on leaf 1 
MAC Address details collected sucessfully
```
The script expects the CIMC ip addresses in a file and writes the patch details to the screen and to a file 
specified with -o option. A log containing all outputs from Nexus switches and server CIMC is captured to a
file specified with -l.

To run simply issue, python check_ucs_conns.py -c cimc.txt -u1 <cimc user> -p1 <cimc password> -u2 <switch user> -p2 <switch password> -ip1 <switch1 ip> -ip2 <switch2 ip> -o patch.txt -l log.txt -v 1

```
$ python check_ucs_conns.py --h
usage: check_ucs_conns.py [-h] -c CIMC -u1 USERNAME1 -u2 USERNAME2 -p1 PASSWD1
                          -p2 PASSWD2 -o PATCH -ip1 IP1 -ip2 IP2 -l LOG -v
                          VERBOSE

CLI arg parser

optional arguments:
  -h, --help            show this help message and exit
  -c CIMC, --cimc CIMC  CIMC IP address file name
  -u1 USERNAME1, --username1 USERNAME1
                        Username to log into the CIMC.
  -u2 USERNAME2, --username2 USERNAME2
                        Username to log into the Leaf.
  -p1 PASSWD1, --passwd1 PASSWD1
                        Password to log in to CIMC.
  -p2 PASSWD2, --passwd2 PASSWD2
                        Password to log in to Leaf.
  -o PATCH, --patch PATCH
                        patch details file name to be created
  -ip1 IP1, --ip1 IP1   ip of leaf switch 1
  -ip2 IP2, --ip2 IP2   ip of leaf switch 2
  -l LOG, --log LOG     Log file name
  -v VERBOSE, --verbose VERBOSE
                        Verbose to see all CIMC and leaf output on stdout
```
