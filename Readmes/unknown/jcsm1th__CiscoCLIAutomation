=============================SSH and Telnet Automation====================================
This script can be used to access devices remotely either via Telnet or SSH.

=============================Usage========================================================

=========Input Parameters==========

-h : help
-u <userid> : Userid
-n <username> : User Name - use ' '
-p <password> : Password - use ' '
-e <enable password> : Enable Password - use ' '
-d <text file> : Device List
-t <ssh/telnet> : Transport
-c <text file> : Command List
-r <Repository> : Repository Directory
-s <sub directory> : Sub Directory

Usage: ssh_telnet_automation.py -t telnet -u jasons2 -d device-list.txt -c command-list.txt

All of the arguments above are optional.  If an argument is not provided, the script will
prompt for the value during run time.

=============================Command and Device File Format===============================
Naming: The name of these files does not matter, but they should be text files.

Format: The files should have device names or commands listed on separate lines.

Example Device File (note that DNS or IP addresses are allowed):

somerouter.cisco.com
someotherrouter.cisco.com
10.10.5.9

Example Command File:

show version
show hardware
show ip interface brief

=============================Output=======================================================
The script provides output back to the screen as well as writes output to files.

The file output is held in a repository and subfolder.  Both of these can be provided as
input parameters.

The expected structure is as follows:

<repository>
    |
    <sub directory>
    <sub directory>
          |
          README       - file containing specifics regarding user, commands and devices
          OUTPUT_CMD_1 - file containing the output of command 1
          OUTPUT_CMD_2 - file containing the output of command 2
          etc...
          
