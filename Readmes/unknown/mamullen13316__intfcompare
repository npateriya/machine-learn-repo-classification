This script was written for a specific use case where I was upgrading a large number of Cisco Catalyst 3750
switches to Nexus 2248 Fabric Extenders.  The ports were configured identically on the new FEXs to match the 
existing 3750 port configurations.  However, due to the large number of switches it was not possible to capture
all of the day-to-day changes that would occur between the time the FEX configurations were generated and when 
the FEXs actually got installed. This script was written to identify any discrepancies between the interface 
configurations of the existing switch vs. the new switch that was to replace it. The intent was to reduce the burden
of having to manually SSH into each device, collect the configurations, and then load them into a program like
DiffMerge or Notepad++ to analyze the differences.

The script does the following:
1. SSH to each device, which can be specified at run time on the command line or in a .yml file.
2. Runs a set of commands which can be specified at run time on the command line or in a .yml file.
3. Compares the interface configurations and reports any differences
4. A copy of the command output is saved in the working directory

If using the .yml file, it must be specified on the command line using the -source option.  The format of the file
is as follows:

---
device1:

   ip: 192.168.1.1

   username: admin

   password: 'RQ269BYafin6M8X4SyVK3+1Bx+4WlN3+DlMKYfS/s6U='

   command: show running-config

   stack_number: '1'

device2:

   ip: 192.168.1.2

   username: admin

   password: 'yEo3W+wjD4JiA5an4B2iiXqKwUw6pteU1R0HUG2IttA='

   command: show running-config interface e100/1/1-48

Note the encrypted passwords in the file. The passwords are encrypted with AES using the included pwencrypt.py script.
To obtain the encrypted form of the password to put in the .yml file,  run the script and provide the clear-text password:

Enter encryption key:

Password:

Enter password:

Password:

MqwbtIlgzWsJdrbhJDQlOxoJ7sp/bZ5ygx+2TWs+gGA=


The encryption key is not stored anywhere,  it will need to be remembered and used when the intfcompare.py script is run.
The key will be used to decrypt the passwords in the .yml file at run time.  

Requirements:
The PyCrypto, Paramiko, and CiscoConfParse Python modules are required and must be installed using your favorite package manager. Below is an example using PIP:

sudo pip install pycrypto

sudo pip install paramiko

sudo pip install ciscoconfparse


