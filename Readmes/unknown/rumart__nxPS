# nxPS

### PowerShell module for administering Cisco Nexus 5k switches via the NX-API
##### NOTE! The NX-API on the 5k switches are different from th 7k and 9k's

This module is built in an existing production environment and it is created as a way to retrieve simple information from Nexus switches without having to know how to write NXOS commands. It may or may not work in other environments, but hopefully it can be an inspiration for others to write similar Powershell modules utilizing this API


**Cisco Programming guide**
http://www.cisco.com/c/en/us/td/docs/switches/datacenter/nexus5000/sw/programmability/guide/b_Cisco_Nexus_5K6K_Series_NX-OS_Programmability_Guide/nx_api.html


**Enabling the NX API**

_Enable without additional config_

```
configure terminal
feature nxapi
```

_Set specific ports (in conf)_
```
nxapi http port 8080
nxapi https port 8443
```

_Enable sandbox (in conf)_
```
nxapi sandbox
```

**Sandbox**

The NXAPI sandbox provides for an easy way to get to know the API through a website where you can write your commands and see both the output and how the API commands should be

------

**Usage**

Start with creating a connection to the Nexus switch with the New-NXConnection command. If you plan to connect to multiple switches place the output of the connection command in a variable which can be used in subsequent commands to the same switch. If only one connection is in use you can omit the connection parameter and use the one added as a global parameter.

To get information on the provided functions run the built in help function in Powershell

```powershell
Get-Help "function name" 
```
