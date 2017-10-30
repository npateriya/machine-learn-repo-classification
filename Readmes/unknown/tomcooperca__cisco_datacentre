# Cisco Datacentre Module

#### Table of Contents

1. [Description](#description)
2. [Setup - The basics of getting started with cisco_datacentre](#setup)
    * [What cisco_datacentre affects](#what-cisco_datacentre-affects)
    * [Setup requirements](#setup-requirements)
    * [Beginning with cisco_datacentre](#beginning-with-cisco_datacentre)
3. [Usage - Configuration options and additional functionality](#usage)
4. [Reference - An under-the-hood peek at what the module is doing and how](#reference)
5. [Limitations - OS compatibility, etc.](#limitations)
6. [Development - Guide for contributing to the module](#development)

## Description

This module will configure Cisco Nexus 3K/9K switches, based on generic Cisco EVPN design standards. Some key assumptions include:

* EVPN using BGP control plane for leaf-spine networks
* OSPF is used as the primary IGP and/or underlay routing protocol
* This module supports the `guestshell` only; not supported for
  NX-OS `bash-shell`

## Setup

### What cisco_datacentre affects

The base classes of this module will setup the following configuration:

* install prerequisite Ruby gems for the ciscopuppet module
* NTP (NX-OS style, not Linux style)
* SNMP community strings and access lists (SNMPv2 only)
* Security/compliance related configuration items such as logging parameters, login attempts,
  system clock and more

The EVPN classes will setup the following items:

* OSPF underlay routing & interfaces
* BGP EVPN overlay peering (spines are assumed to act as route reflectors)
* Virtual Port-Channel (vPC) between leaf pairs
* User VLANs (as VXLAN VNI's within the fabric) and downstream interfaces

### Setup Requirements

In order for this module to function, you must include
`cisco_datacentre::base` class in your Puppet code. It is recommended to use the roles & profiles
pattern in a control repository to consume these manifests.

### Beginning with cisco_datacentre

To start using the module, add the following classes to your code:

```
node default {
  include cisco_datacentre::base
}
```

## Usage

Most of the classes in this module are designed to work with large amounts of
Hiera lookup data. This can either be brought in via automatic parameter lookup
in this module's namespace or using class declarations from another class (such
as a profile).

*Example Leaf classes with APL*
```leaf.pp
  include cisco_datacentre::evpn::leafglobal
  include cisco_datacentre::evpn::underlay
  include cisco_datacentre::evpn::leafoverlay
```

*Example leaf.yaml Hiera file with various parameters*
```
---
cisco_datacentre::evpn::underlay::loopback0_ip: 10.10.255.1
cisco_datacentre::evpn::underlay::ospf_area: 0.0.0.1
cisco_datacentre::evpn::underlay::ptp_interfaces:
  Ethernet1/49:
    description: 'PTP'
    ipaddress: 10.10.254.0
<etc>
```

## Reference

### Base Classes

* `cisco_datacentre::base` - Base class to install the prerequisite Ruby gems
  for the ciscopuppet module; required for all other classes. Also sets an initial
  `root` password in the guestshell environment to `Password123!`
* `cisco_datacentre::ntp` - Configure NTP servers. Required parameter of `ntpservers`
  is an array of IP addresses. An optional `ntp_acl_name` can be passed to override
  the default ACL name of `'NTP_ACL'`
* `cisco_datacentre::sec` - Apply security-related global configuration settings.
  Optionally, you can specify a different management VRF for the syslog source
  interface. Otherwise, `management` will be the assumed default.
* `cisco_datacentre::snmp` - Configure SNMPv2 community strings & ACL's. Must pass
  in two hashes, `read_only` and `read_write`. Optionally, you can also pass in the
  SNMP sysLocation & sysContact strings. This class assumes that both R/O and R/W
  strings are required as well as having an access-list with an array of permitted
  CIDR networks to be attached to each string. The hash must be formatted as such:

_Hiera_
```
---
read_only:
  'readonly_string':
    'readonly_acl_name':
      - '10.0.0.0/24'
      - '10.0.1.0/24'
read_write:
  'readwrite_string':
    'readwrite_acl_name':
      - '10.0.2.0/24'
      - '10.0.3.0/24'
```
_Puppet DSL_
```
$read_only = {
  'readonly_string' => {
    'readonly_acl_name' => [
      '10.0.0.0/24',
      '10.0.1.0/24',
    ]
  }
}
$read_write = {
  'readwrite_string' => {
    'readwrite_acl_name' => [
      '10.0.2.0/24',
      '10.0.3.0/24',
    ]
  }
}
```

## Limitations

This module is compatible with NX-OS version 7.0(3)I4(2) and later. It is only
supported to work within the NX-OS guestshell environment. The required ciscopuppet
module is version 1.6.0.

## Development

For any pull requests to this module which are creating new classes, please ensure
to create corresponding RSpec tests to ensure catalog compilation is tested and validated.
