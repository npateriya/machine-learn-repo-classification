# Cluster - Multithreaded Cisco IOS Cluster Deployment Tool
![Cluster Demo](./clusterdemo.png)
*(screenshot taken from [Network Lab](https://github.com/shibusa/networklab) deploy)*

### Requirements:
- Python 2.7.10
- napalm_ios (https://napalm.readthedocs.io/en/latest/#)
- privileged, SSH access to respective Cisco IOS devices

### Installation
1. Install system dependencies
```
https://napalm.readthedocs.io/en/latest/installation/ios.html
```

2. Install "napalm_ios" library.
```
pip install napalm-ios==0.6.1
```

3. Make cluster executable
```
chmod +x cluster
```

### Usage
1. Template file(s) - Create a file(s) with the respective clustername in the same directory as 'cluster'.  The file should contain standard Cisco IOS formatting with "version #" on the top and "end" at the bottom.  Commands can be written as normal any.  Reusable parts of the template can be denoted in brackets.

Example:
Filename: basiccluster
```
version {versionnum}

hostname {hostname}

end
```

2. Create a json file with a dictionary of clusternames in the same directory as 'cluster'.  Within each cluster name should be a list of hosts belonging to the cluster and dictionary of options matching cluster file.  Option keys must match the brackets defined in the template file.

```javascript
{
  "bassiccluster":{
    "hosts": ["HOST1","HOST2","HOST3"],
    "options":{
      "versionnum":"12.4",
      "hostname":"router"
    }
  }
  "CLUSTERNAME2":{
    "hosts": ["HOST4","HOST5"],
    "options":{
      "versionnum":"12.4",
      "hostname":"switch"
    }
  }
}
```

2a. (OPTIONAL) - With multiple hosts, numerical fields can be incremented/decremented by an additional "adjustments" field alongside with "hosts" and "options".  The "adjustments" field needs to be in a dictionary format and match the same key found in the "options" dictionary.

Cluster File
```
version {versionnum}

int range fa1/{intnum}
ip address {network}{host} {netmask}

end
```

JSON File
```javascript
{
  "example":{
    "hosts": ["192.168.1.2","192.168.1.3"],
    "options":{
      "versionnum":"12.4",
      "intnum":0,
      "network":"10.0.0."
      "host":240
      "netmask:":"255.0.0.0"
    },
    "adjustments":{
      "intnum":1
      "host":-4
    }
  }
}
```

Host 192.168.1.2 Output
```
version 12.4

int range fa1/0
ip address 10.0.0.240 255.0.0.0

end
```

Host 192.168.1.3 Output
```
version 12.4

int range fa1/1
ip address 10.0.0.236 255.0.0.0

end
```

3. Run cluster
```
./cluster FILE.json
```
