this code was only tested with catalyst switches (2960s, 2960g, 2960).

# get vlans
snmpwalk -On -v1 -c public 192.168.32.35 .1.3.6.1.4.1.9.9.46.1.3.1.1.2.1
[...]
.1.3.6.1.4.1.9.9.46.1.3.1.1.2.1.33 = INTEGER: 1

# get macs learnt on this vlan (33)
snmpwalk -On -v1 -c public@33 192.168.32.35 .1.3.6.1.2.1.17.4.3.1.1 <= mind the vlan in community
[...]
.1.3.6.1.2.1.17.4.3.1.1.104.163.196.239.201.18 = Hex-STRING: 68 A3 C4 EF C9 12
.1.3.6.1.2.1.17.4.3.1.1.108.98.109.57.250.165 = Hex-STRING: 6C 62 6D 39 FA A5  <= let's search this one
.1.3.6.1.2.1.17.4.3.1.1.208.87.76.163.240.18 = Hex-STRING: D0 57 4C A3 F0 12

# bridge number
snmpwalk -On -v1 -c public@33 192.168.32.35 .1.3.6.1.2.1.17.4.3.1.2.   108.98.109.57.250.165 <= mind spaces !
.1.3.6.1.2.1.17.4.3.1.2.108.98.109.57.250.165 = INTEGER: 19 <= bridge number 19

# port number
snmpwalk -On -v1 -c public@33 2960g-1 .1.3.6.1.2.1.17.1.4.1.2.  19 <= mind spaces !
.1.3.6.1.2.1.17.1.4.1.2.19 = INTEGER: 10119 <= port number 10119

# port name
snmpwalk -On -v1 -c public 192.168.32.35 .1.3.6.1.2.1.31.1.1.1.1.  10119 <= mind spaces !
.1.3.6.1.2.1.31.1.1.1.1.10119 = STRING: "Gi0/19" <= port name
