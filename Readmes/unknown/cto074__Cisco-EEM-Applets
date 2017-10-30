# Cisco-EEM-Applets
Cisco Embedded Event Manager Applets
Embedded Event Manager Applet for Cisco routers that support EEM functionality.
This EEM applet calculate bandwidth of an interface subtracting the bytes of the interface with a 2 seconds interval. 

To allow SNMP polling at router we need to have "snmp-server manager" configured.
Value of .ifindex can be checked with command "show snmp mib ifmib ifindex"

Cisco routers refresh interface rate by default every few seconds but the value that is showed is an average of a 5 minutes ##interval (by default, it can be lowered to 30 seconds with command "load-interval 30". 
So in most cases, we can't observe traffic bursts, at least with this script we can quickly check the average rate in a 2 ##second interval.

Ideally we should use a SNMP polling server (example MRTG, Cacti,..) with a very small interval between SNMP requests, but ##not always is available. Also it should be checked the snmp counter update interval that varies depending on the IOS ##versions.

There are two commands that you can use to control how often the SNMP counters are updated.
The first command is the “service counters max age seconds” which configures the maximum "age" of cached counters, in seconds. 
The other one is a hidden command: “snmp-server hc poll centiseconds” this one configures how often the HC (high capacity) counters are polled.
You can use the following commands to update the counters every second:

service counters max age 1
snmp-server hc poll 100
