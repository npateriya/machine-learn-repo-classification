# ConfigExtractor
Used to emulate some cisco comand pipe outputs, currently:
* exclude
* include
* section

## Details
Reads configurations in from named files (additional sources planned) and converts to a ConfigList object

ConfigList objects have a number of functions emulating Cisco behaviour:
* ConfigList.exclude(<regular expression>)
* ConfigList.include(<regular expression>)
* ConfigList.section(<regular expression>)

All above functions return ConfigList objects so functions are able to be chained as show in the examples below.

## Examples
From the included cisco.config show config lines that begin with interface, but not serial interfaces
```
from ConfExtractor.ConfExtractor import  ConfExtractor

c = ConfExtractor()
c.fromFile('cisco.config')
print (c.section(r'^iinterface').exclude(r'[Cc]ellular'))
```

Based on the below bgp congfiguration, return the non static redistributions from vrf cust1
```
from ConfExtractor.ConfExtractor import  ConfExtractor

    config = '''
router bgp 65000
 bgp log-neighbor-changes
 bgp listen range 0.0.0.0/0 peer-group InternetL3VPN
 no bgp default ipv4-unicast
 timers bgp 7 21
 neighbor InternetL3VPN peer-group
 neighbor InternetL3VPN remote-as 65535
 neighbor InternetL3VPN ebgp-multihop 255
 !
 address-family ipv4
  redistribute connected
 exit-address-family
 !
 address-family vpnv4
  neighbor InternetL3VPN activate
  neighbor InternetL3VPN send-community both
  neighbor InternetL3VPN route-map L3VPN_Internet in
 exit-address-family
 !
 address-family ipv4 vrf cust1
  redistribute connected
  redistribute static
  default-information originate
 exit-address-family
'''
    c = ConfExtractor()
    c.fromString(config)
    print (c.section(r'router bgp').section(r'vrf cust1').include(r'redist').exclude(r'static'))

```
