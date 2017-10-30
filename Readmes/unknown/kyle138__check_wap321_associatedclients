# check_wap321_associatedclients
Nagios plugin to check associated clients in Cisco WAP321 devices

check_wap321_associatedclients.php
Check Associated Clients for Cisco WAP321
Manufacturer recommended levels are 20 for warning and 35 for critical.

USAGE: php check_wap321_associatedclients.php HOST COMMUNITY WARNING CRITICAL
HOST: IP or FQDN of the target Cisco WAP321 device.
COMMUNITY: SNMP Community Name
WARNING: Amount of Associated Client connections to trigger WARNING
CRITICAL: Amount of Associated Client connections to trigger CRITICAL
EXAMPLE: php check_wap321_associatedclients.php 192.168.1.1 public 20 35

Include in commands.cfg:
define command{
 command_name check_wap321_associatedclients
 command_line php /path/to/check_wap321_associatedclients.php $HOSTADDRESS $ARG1$ $ARG2$ $ARG3$
}

This plugin requires php5-snmp to be installed
