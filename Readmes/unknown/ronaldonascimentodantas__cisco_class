#THIS IS THE README FILE!

#N9K Scripts

Script: show-command.py is used to execute ANY show command using NX-API, without necessary to login on Sandbox web interface to check the show command availabity on NXAPI.

Syntax: show-command.py < switch-name > < username > < password > < show-command >

Example: show-command.py 172.16.1.1 admin cisco123 "interface description"

-------

Script: show-cdp-inventory-interface.py Return the neighbor name and neighbor interface for a local interface on Nexus 9000 using CDP protocol.

Syntax: show-cdp-inventory-interface.py < switch-name > < username > < password > < interface_name >

Example: show-cdp-inventory-interface.py 172.16.1.1 admin cisco123 Ethernet1/1

-------

Script: show-simple-inventory.py used to show a simple inventory interface from a device, all fieds are printed on CSV format, and could be imported on Excel.

Fields: Interface, Speed, Description, Neighbor Name, Neighbor Interface

Purpose: Help documentation process like an AS-Built

Syntax: Syntax: show-simple-inventory.py < switch-name > < username > < password >

Example: show-simple-inventory.py 172.16.1.1 admin cisco123

#APIC Scripts

Script: apic_show-mac-address.py is used to show all if a mac-address is inside the fabric
Syntax: python apic_show-mac-address.py < hostname > < username > < password > < mac-address >

Examples:

(1) Single mac address inside the fabric

python apic_show-mac-address.py apic admin XXXX 00:44:56:6F:85:AD

(2) Group mac address inside the fabric

python apic_show-mac-address.py apic admin XXXX 00:44:56

-------

Script: apic_find-epg.py is used to find a specific EPG inside the Fabric
or find a sub-string EPG inside the fabric

Syntax: python apic_find-epg.py < hostname > < username > < password > < epg > [-s]

[-s] specify if should be used sub-string on search or not.

Examples:

(1) Find a EPG string inside the fabric

python apic_find-epg_v2.py apic admin XXXX App

(2) Find a EPG sub-string inside the fabric

python apic_find-epg_v2.py apic admin XXXX App -s
