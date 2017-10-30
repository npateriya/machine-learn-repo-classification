snmpEngine 
=====

WORK in PROGRESS... 

Install
=======

To install, execute:

```
pip install snmpEngine
```

- os requirements for snimpy

```
# Macos
#
# brew install libsmi
# easy_install snimpy
#
# /usr/local/Cellar/libsmi/0.4.8/share/mibs/iana/
#
# Debian
#
# apt-get install python-dev build-essential
# apt-get install libffi-dev
# easy_install snimpy
#
# vi /etc/apt/sources.list
# ... main non-free 
#
# apt-get update snmp-mibs-downloader
# dpkg --configure -a
# download-mibs
#
# ls /usr/share/mibs/ietf
#
# apt-get install smitools
# smilint -s -l1 CISCO-CDP-MIB.my 
```

Documentation
=============

```python
from snmpEngine import MftSNMP

initialize the object
r1 = MftSNMP("x.x.x.x", community="public")
```

since now, you can recall a variable by walk the json tree

```python
>>> r1['XRv1']["interfaces"]["GigabitEthernet0/0/0/0.10"]["ipAdEntAddr"]
10.10.10.1
```

### json

return the info's as json

```python
print  r1
{
    "XRv1": {
        "BGP4-MIB": {
            "bgpIdentifier": "1.1.1.1"
        },
        "CISCO-CDP-MIB": {
            "cdpGlobalRun": "true(1)"
        },
        "OSPF-MIB": {
            "ospfRouterId": "1.1.1.1"
        },
        "global": {
            "sysDescr": "Cisco IOS XR Software (Cisco IOS XRv Series),  Version 5.2.0[Default] Copyright (c) 2014 by Cisco Systems, Inc.",
            "sysLocation": "virtualbox",
            "sysName": "XRv1",
            "sysObjectID": "1.3.6.1.4.1.9.1.613",
            "sysUpTime": "0:13:09.730000"
        },
        "interfaces": {
            "GigabitEthernet0/0/0/0": {
                "CISCO-CDP-MIB": {
                    "cdpInterfaceEnable": "true"
                },
                "ifAdminStatus": "up",
                "ifAlias": "",
                "ifDescr": "GigabitEthernet0/0/0/0",
                "ifIndex": 3,
                "ifMtu": 1514,
                "ifOperStatus": "up",
                "ifPhysAddress": "8:0:27:d6:4b:99",
                "ifType": ethernetCsmacd(6)
            },
            "GigabitEthernet0/0/0/0.10": {
                "MPLS-L3VPN-STD-MIB": {
                    "mplsL3VpnVrfName": "BLACK"
                },
                "ifAdminStatus": "up",
                "ifAlias": "",
                "ifDescr": "GigabitEthernet0/0/0/0.10",
                "ifIndex": 9,
                "ifMtu": 1518,
                "ifOperStatus": "up",
                "ifPhysAddress": "8:0:27:d6:4b:99",
                "ifType": l2vlan(135),
                "ipAdEntAddr": "10.10.10.1",
                "ipAdEntNetMask": "255.255.255.0"
            },
...
```

### show interfaces in table

```python
r1.show_interfaces()
+------------+------------+------------+------------+------------+------------+
|  ifIndex   |  ifDescr   |  ifAlias   |   IPadd    |     AS     |     OS     |
+============+============+============+============+============+============+
|     12     | Gi0/0/0/0. | GRPX       | 10.1.3.1   | up         | up         |
|            | 13         |            |            |            |            |
+------------+------------+------------+------------+------------+------------+
|     11     | Gi0/0/0/0. | GRPX       | 10.1.2.1   | up         | up         |
|            | 12         |            |            |            |            |
+------------+------------+------------+------------+------------+------------+
|     5      | Mgt0/0/CPU | Management | 192.168.56 | up         | up         |
|            | 0/0        | interface  | .109       |            |            |
+------------+------------+------------+------------+------------+------------+
|     9      | Gi0/0/0/0. |            | 10.10.10.1 | up         | up         |
|            | 10         |            |            |            |            |
+------------+------------+------------+------------+------------+------------+
|     13     | Gi0/0/0/0. |            | 10.14.14.1 | up         | up         |
|            | 14         |            |            |            |            |
+------------+------------+------------+------------+------------+------------+
|     2      | Null0      |            | --         | up         | up         |
+------------+------------+------------+------------+------------+------------+
|     10     | Gi0/0/0/0. |            | 10.11.11.1 | up         | up         |
|            | 11         |            |            |            |            |
+------------+------------+------------+------------+------------+------------+
|     8      | Lo0        | PROVA      | 1.1.1.1    | up         | up         |
+------------+------------+------------+------------+------------+------------+
|     4      | Gi0/0/0/1  |            | 10.0.3.11  | down       | down       |
+------------+------------+------------+------------+------------+------------+
|     3      | Gi0/0/0/0  |            | --         | up         | up         |
+------------+------------+------------+------------+------------+------------+
```

### ifStatEngine(choice, speedUnit="kb", secDelta=5)

return current the speed/sec on choosen interface (Out/in)

```python
>>> r1.ifStatEngine("5",speedUnit="kb",secDelta=1)
(0.0, 0.0)
```

### ifstat(choice="None")

show realtime speed of that interface

```python
>>> r1.ifstat("5")

5  -  MgmtEth0/0/CPU0/0
'Management interface'

Out     In
341.0   251.0
489.0   378.0
181.0   783.0
862.0   536.0
221.0   296.0
1120.0  371.0
1420.0  425.0
...
```

if you not pass nothing, show_interfaces() will be call to let you choose the ifIndex

### summary()

this shows main 

```python
>>> r1.summary()
hostname XRv1
OSPF RID 1.1.1.1
BGP RID 1.1.1.1
Interface no. 10
OSPF Adjacency 0
LDP Adjacency 0
```

### walkInterfaces()

```python
>>> r1.walkInterfaces()
interface  GigabitEthernet0/0/0/0.13
  description GRPX
interface  GigabitEthernet0/0/0/0.12
  description GRPX
interface  MgmtEth0/0/CPU0/0
  description Management interface
  ======= arp =======
  pint vrf MGT 192.168.56.102
  telnet 192.168.56.102 /vrf MGT
interface  GigabitEthernet0/0/0/0.10
interface  GigabitEthernet0/0/0/0.14
interface  Null0
interface  GigabitEthernet0/0/0/0.11
interface  Loopback0
  description PROVA
interface  GigabitEthernet0/0/0/1
interface  GigabitEthernet0/0/0/0
```

Thanks
======

License
=======
