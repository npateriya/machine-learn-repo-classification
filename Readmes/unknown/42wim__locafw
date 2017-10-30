# locafw

Build a location-aware firewall for your wireless users using linux, netfilter, ipset and (cisco) wireless lan controllers.  
Very much a WIP and not much documentation yet.

Use-case: only allowing specific classrooms (containing AP's) to specific servers (e.g. because of licensing issues).

Locafw connects to your wireless lan controllers, asks information about accesspoints and client.   
In the locafw config file you specify which destination addresses can be accessed by the clients on the specified accesspoints.  
Using this information locafw creates an [ipset](http://ipset.netfilter.org/) set, which you can use on your firewall.  

## binaries
not yet

## building
Go 1.6+ is required. Make sure you have [Go](https://golang.org/doc/install) properly installed, including setting up your [GOPATH] (https://golang.org/doc/code.html#GOPATH)

```
cd $GOPATH
go get github.com/42wim/locafw
```

You should now have locafw binary in the bin directory:

```
$ ls bin/
locafw
```

## running
1) Copy the locafw.yaml.sample to locafw.yaml in the same directory as the locafw binary.  
2) Edit locafw.yaml with the settings for your environment. See below for more config information.  
3) Now you can run locafw

### web
* http://localhost:8888/ip - returns a JSON array of all AP's with clients and foreach AP and array of IP addresses
* http://localhost:8888/ipset - returns a generated ipset set, can be used with ```ipset -! restore```
* http://localhost:8888/realod - reloads the configuration file


## config
### locafw
locafw looks for locafw.yaml in current directory. (use -conf to specify another file)

Look at locafw.yaml.sample for an example


```
web:
    listen: ":8888" #ip and port to listen on
#list of your cisco WLC's, need public commmunity and snmp v2c
controllers:
    - 192.168.1.1
    - 192.168.1.2
#ipset configuration, name of the set and TTL of the entries
ipset:
    name: locafw
    timeout: 1500
#ACLs, name is the name of your AP's (can be a regexp)
#dstip are the destination addresses that are allowed to be accessed from the specified AP
acl:
- name: lwap-buildingA-1
  dstip:
  - 10.0.0.5
  - 10.0.0.6
- name: "lwap-buildingB-*"
  dstip:
  - 1.4.5.7
  - 2.4.5.7
```

### iptables
```iptables -I FORWARD -m set --match-set locafw src,dst -j ACCEPT```

### ipset
```ipset create locafw hash:net,net family inet hashsize 2048 maxelem 65536 timeout 1800```

## TODO
Documentation
