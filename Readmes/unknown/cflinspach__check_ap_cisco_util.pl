# check_ap_cisco_util.pl
Check plugin for nagios/icinga to record utilization of Cisco AP's.

Currently does not have warning or critical thresholds and is used primarily to gather data for graphs.

Usage:
check_ap_cisco_util.pl -H [host] -C [community] -O [ap name oid] -r [ap receive oid] -t [ap transmit oid]

Meru example:
check_ap_cisco_util.pl -H 192.168.0.6 -C public -O .1.3.6.1.4.1.14179.2.2.1.1.3 -r .1.3.6.1.4.1.14179.2.2.13.1.2 -t .1.3.6.1.4.1.14179.2.2.13.1.1
