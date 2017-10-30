# sshconfigurator
Automates configuration of cisco devices using SSH<br />
<br />
Usage:<br />
Requires 3 files<br />
1) A file containing IP addresses of devices to configure in the following format<br />
1.2.3.4<br />
2.3.4.5<br />
3.4.5.6<br />
4.5.6.7<br />
<br />
2) A file containing the login credentials (telnet/ssh) in the following format<br />
user1,password1,enable1<br />
user2,password2<br />
<br />
NOTE: You can enter multiple credentials, it tries them top down<br />
<br />
3) A file containing the configuration to paste in the following format<br />
snmp-server community abcdef rw 55<br />
access-list 55 permit 1.2.3.4<br />
<br />
**ENSURE THE ABOVE DO NOT HAVE UNNECCESARY SPACES***<br />
<br />
<br />
1) The script asks for the file containing IP addresses, checks for validity<br />
2) The script asks for the file containing login credentials, checks for validity<br />
3) The script asks for the file containing the configuration, checks for file availability<br />
4) The script configures all devices and list successfully configured devices, failed devices and unreachable devices<br />
<br />
There is also the option of checking the IPs reachablility by  uncommenting the code from lines 56-73, this will make the program run longer<br />
<br />
Run with Python 2.7 and ensure you have the paramiko library installed.<br />
<br />
Have Fun!! :) 

