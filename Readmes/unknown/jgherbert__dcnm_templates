# dcnm_templates
A repo of configuration deployment templates for Cisco DCNM.

Author: John Herbert (Twitter: @mrtugs, Web: http://movingpackets.net/)

# About

When I tried to develop some configuration deployment templates in Cisco's Data Center Network Manager (DCNM) tool, I found that there wasn't very much on the web by way of a reference, perhaps because it's not the most popular automation tool out there. Nonetheless, since I've written a few fairly basic configuration templates, I figured I'd share them. If you have suggestions/improvements, however unlikely that may be, feel free to make a pull request.

# Scripts

## ADD_A_NEW_VLAN

When run, the user is prompted for a VLAN ID (a number) and a Description. Currently a VLAN ID from 100-4096 is acceptable (but that can of course be changed to suit needs), and the description is checked for spaces, which are not allowed. The script will then run through the devices presented to it by DCNM and create a VLAN with the provided ID and description, and set it to "mode fabricpath".
If the hostname of a device includes the text "fabricedge", the script will additionally configure port-channel 1 and 2 (which it assumes are switchport trunks with a pruned VLAN list) to allow the new vlan on those trunks. This is useful for connectivity to edge routers, and again can be tweaked as necessary.
There isn't any way to interact with the user in DCNM, so when errors occur they are presented in the form of NXOS comments (i.e. lines starting with an exclamation point). 
