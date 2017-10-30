cpuAscii 
=====

Install
=======

- os requirements for snimpy

```
# Macos
#
# brew install libsmi
# easy_install snimpy
#
# /usr/local/Cellar/libsmi/0.4.8/share/mibs/iana/
#
# Debian
#
# apt-get install python-dev build-essential
# apt-get install libffi-dev
# easy_install snimpy
#
# vi /etc/apt/sources.list
# ... main non-free 
#
# apt-get update snmp-mibs-downloader
# dpkg --configure -a
# download-mibs
#
# ls /usr/share/mibs/ietf
#
# apt-get install smitools
# smilint -s -l1 CISCO-CDP-MIB.my 
```

Documentation
=============

```python
usage: cpuAscii.py [-h] -r DEVICENAME [-c DEVICECOMM] [-m] [-d]

cpu and memory realtime graph

optional arguments:
  -h, --help            show this help message and exit
  -r DEVICENAME, --routerName DEVICENAME
  -c DEVICECOMM, --community DEVICECOMM
  -m, --memory
  -d, --debug
  

xxxxx.xxxx.xxx

Cisco IOS Software, 7200 Software (C7200P-ADVENTERPRISEK9-M), Version 15.1(3)S3, RELEASE SOFTWARE (fc1)
Technical Support: http://www.cisco.com/techsupport
Copyright (c) 1986-2012 by Cisco Systems, Inc.
Compiled Fri 30-Mar-12 06:19 by prod_rel_team

CPU Graph (sh processes cpu sorted | i ^CPU)
###############################################################################
**************************                              30  cpmCPUTotal5sec    
*****************************                           33  cpmCPUTotal5sec    
***************************                             31  cpmCPUTotal5sec    
****************************                            32  cpmCPUTotal5sec    
**************************                              30  cpmCPUTotal5sec    
**************************                              30  cpmCPUTotal5sec    
**************************                              30  cpmCPUTotal5sec    
*****************************                           33  cpmCPUTotal5sec    
******************************************************  61  cpmCPUTotal5sec    
**************************                              30  cpmCPUTotal5sec <==

CPU Graph (sh processes cpu sorted | i ^CPU)
###############################################################################
██████████████████████████████████████████████████████████  34  cpmCPUTotal1min
██████████████████████████████████████████████████████████  34  cpmCPUTotal5min

MEM Graph (sh process memory sorted | i Processor Pool)
####################################################################################
++++++++++++++++++++++++++++++++++++++++++++++++++    1G  ciscoMemoryTotal          
+++++++++++++++++++++++++++++++++++++++++++++++       1G  ciscoMemoryPoolFree       
++                                                  108M  ciscoMemoryPoolUsed       
++++++++++++++++++++++++++++++++++++++++++++++        1G  ciscoMemoryPoolLargestFree


```


Thanks
======

License
=======