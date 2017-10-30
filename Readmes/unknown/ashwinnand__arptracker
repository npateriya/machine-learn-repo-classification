# arptracker 
Track and monitor arp cache changes in cisco network via routers.

(orginally imported from https://code.google.com/archive/p/arptracker/)

##Description
arptracker collects and tracks ARP data via SNMP from your (cisco) routers.

It assumes that your SNMP RO community string is uniform across your Cisco router network.

It tracks when a particular MAC address (per IP) was first and and last seen on a network,effectively giving a whole range of forensics and monitoring possibilities.

arptracker works by scraping arp tables from all routers once every 30minutes(configurable). The database has 2 tables, one is the current arp information and the other is historical, ie, the main arpdatabase itself, which gives us when a particular IP/mac combo was first and last network.

##Typical Goal

Many !!!!

a. How about tracking which machines are using which ip address over time intervals.

b. Tracking which machine(MAC) has what IP address. 

c. How many machines are 'active' on a LAN.

d. Detect Arp poisoning attacks and proxy arp in network . i.e which mac address has more than one IP adddress assigned to it.

e. Data repository for a quota and/or quarantine management.

