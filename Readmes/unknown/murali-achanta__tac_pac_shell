# Cisco Nexus N7K tac-pac shell
Cisco Nexus N7K devices support tac-pac output collection for debugging customer issues, Also some times only part of the logs are collected using *show tech-support* commands. This application takes the uncompressed tac-pac text file and provide a CLI interface very similar to that of actual device. It prints out the command outputs from the supplied tac-pac text file.

This application is based on ****Cmd2**** package as it provides redirection of output to other linux utlities like ```less```, ```grep``` etc. 

Dependencies:
``` Cmd2 package, 'pip install cmd2'``` 

This CLI application provides

* TAB completion on available commands
* *listall* to list all available commands
* *attach module x* to navigate to module specific commands
* *list* to list context specific command
* it generates tmp cache pickle file to speed-up on BIG text file

Below is a sample output from mocked up tac-pac file

```
[my linux]$ ./shell_sim.py test.txt 
<< building cache file /tmp/89c793420806da3de364b42f33f670bd19829421.pcl >>
        
 ----
 Welcome to the NXOS type shell for tac-pac.         
 ----

test-switch# show module
`show module`
Mod  Ports  Module-Type                         Model              Status
---  -----  ----------------------------------- ------------------ ----------
1    100     100 Gbps Ethernet Module           N7K-MABCD-12      ok
2    400     1000 Mbps Ethernet Module          N7K-MABCD-11      ok
5    0      Supervisor module-90                N7K-SUP90         active *
6    0      Supervisor module-90                N7K-SUP90         ha-standby
test-switch# 
```
attach module uses recursion

```
test-switch# attach module 1
        
 ----
 Welcome to the NXOS type shell for tac-pac.         
 ----

test-switch# (module 1)# list
vdc 1
show clock
show version
test-switch# (module 1)# show clock
`show clock`
10:11:12.234 PST Fri Feb 13 2030
test-switch# (module 1)# exit
```

Please Note this is Python2 application, and uses nested dicts
