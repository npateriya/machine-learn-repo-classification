# ASA-monitoring
It's a (Zabbix) custom template for an old Cisco ASA 5520 with some scripting

Some macros are needed:
{$SNMPHOST} = I know you can use {HOST.CONN} but I preferred to create my own "variable"
{$SNMP_COMMUNITY} = As usual
{$PEERIP} = Peer IP of one IPSec Tunnel (if you have more than one just clone the monitor)

The monitors for interfaces use "Discovery". There is simple RegEx which works for me just taking the interfaces containing: TMC or Untrust or Fail on their names. Kindly modify that RegEx to make it fits with you.
