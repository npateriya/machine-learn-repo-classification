# Cisco Scripts

Authored Aug 20, 2015 by Ken Lemoine

### Summary

This repo contains a set of Python, Expect and BASH scripts designed to 'blackhole' an IP address on multiple Cisco routers.  This will effectively send all traffic from this IP to the 'bitbucket'.  The use case for us is to blackhole any IP that is involved in a DDos attack.  Everything is logged to ~/results.log. 

### Requirements

* Access to some Cisco routers, with privelege level high enough to run 'enable' and 'conf t'.
* The script assumes your username of the computer running the scripts matches the username you use to log into the routers.
* A file (not included) called 'device-list.txt' with a list of router hostnames or IP's, in the same directory as these scripts.
* Python, Expect and BASH obviously.  A Linux OS would be easiest, this was built and tested on Ubuntu 14.10, but any flavor should work fine.


### Installation Instructions

* Install Git
  * sudo apt-get install git
  * git config --global user.name "yourusername"
  * git config --global user.email "your@emailaddy"

* Checkout this repository
  * git clone https://github.com/KeyInfo/cisco_scripts.git 

* Install Expect
  * sudo apt-get install expect

* Create a device-list.txt file with a list of your routers, one per line.

* Create 'outfiles' directory
  * mkdir outfiles 

* Install (optional) Pip. Here's the Ubuntu steps -
  * sudo apt-get update
  * sudo apt-get -y install python-pip
  * pip -v  (verify's successful install)

* Install 'netaddr' Python library. 
  * sudo pip install netaddr

### Script Execution

* Run the blackhole.py file
  * python blackholeip.py

### Validation

* Run this command on the router to validate, if the process was successful you should see some output like below
  * sho run | inc <-ipaddresshere->
  * Output:
``` 
  network <-IP-Address-> mask 255.255.255.255
ip route <-IP-Address-> 255.255.255.255 Null0
ip prefix-list Blackhole seq xxxx permit <-IP-Address->/32
```

### Troubleshooting

* Check your home directory for the log file, ~/results.log
