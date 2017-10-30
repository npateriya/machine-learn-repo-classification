serviceabilitytools
===================
Cisco serviceability tools
Service Control and Activation for Cisco Communications Manager and Instant Messaging and Presence
# Author: Patrick Nowicki - patrick.nowicki@cdw.com
# Version: 1.0
# Date: 9/1/2014
# support information https://discdungeon.cdw.com/vvtwiki/index.php/Serviceabilitytools


Usage:
This tools is run from a command line in this format
serviceabilitytools.exe <tool> <args>

Config File:
A configuration file with the name "serviceabilitytools.cfg" must be created in the same folder as the executable. The config file should look like this:

# Version
version=10.5
# User ID
user=ccmadministrator
# Password
password=password
# Primary Server IP - typically CUCM Publisher
primaryserver=192.168.1.1
# optional debug level
debug=0

Note: The password field must be populated before first use. After first run, it will be encrypted in the configuration file.

Tools:

activateservice - Use: Activate or Deactivate Services. This tool takes a single argument, the name of an input file. This is a tab separated file with the following columns:
Column 	Value 	Required?
Server 	IP/FQDN of server to be acted on 	Yes
Service 	Full Service name - e.g. "Cisco Tftp" not "Tftp" 	Yes
Deploy 	Deploy/Undeploy - to activate or deactivate service 	Yes


controlservice - Use: Start/Stop/Restart Services. This tool takes a single argument, the name of an input file. This is a tab separated file with the following columns:
Column 	Value 	Required?
Server 	IP/FQDN of server to be acted on 	Yes
Service 	Full Service name - e.g. "Cisco Tftp" not "Tftp" 	Yes
Action 	Action to preform on server - e.g. Start/Stop/Restart 	Yes


clusterserviceinfo - Use: To dump all service information for cluster, including CUCM and IM&P. This tool takes no arguments and keys off of the primary server IP in the config file.
Note: Due to Serviceability AXL Throttle typically being set to 15 requests per minute, per server, this tool can take a few minutes to run.
Note: The output of this tool will be dumped into a file called clusterinfo<timestamp>.txt and include, Server ip/fqdn, Service Name, Service Status, Uptime, Service Type, Dependencies, Service Group, and whether or not the service can be Activated or Deactivated.

clusterservicerestart - Use: To restart a particular service, cluster-wide. This tool takes a single argument, the full service name, and keys off of the primary server IP in the config file.
Note: If the service name contains a space, it should be surrounded in quotes.
Note: This tool will only restart the specified service on those servers where that service is activated and in the "Started" state.

clustertftprestart - Use: To restart the Cisco Tftp service, cluster-wide. This tool takes no arguments, and keys off of the primary server IP in the config file.
Note: This tool will only restart Cisco Tftp on those servers where that service is activated and in the "Started" state. 
