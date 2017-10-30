# nagios-bgp-ipv6-cisco
Check status of IPv6 BGP peers on a Cisco IOS device

This script will login to a Cisco IOS router return the status of IPv6 BGP peers.

## Requirements

* [netmiko](https://github.com/ktbyers/netmiko)
* [pynag](http://pynag.org/)

If you are using pip, the following command will install all requirements:

    pip install netmiko pynag

## Command line options

```
    -H HOST               Host to connect to (default: localhost)
    -l USERNAME           Username to login with
    -p PASSWORD           Password to login with
    --html                Colorize Nagios summary output
```

This check runs the command `show bgp ipv6 unicast neighbors` on the router.

## Icinga or Nagios configuration

This check needs login credentials to login to a router. It is recommended that you store this info in the [resource.cfg file](http://docs.icinga.org/latest/en/sample-resource.html).

Define a check command:

```
define command {
                command_name                          check_bgp_ipv6_cisco
                command_line                          $USER10$/check_bgp_ipv6_cisco.py -H $ARG1$ -l $USER21$ -p $USER22$ --html
}
```

Use the command in a service or host definition:

```
define service {
                host_name               myrouter
                service_description     bgp peers ipv6
                check_command           check_bgp_ipv6_cisco!$HOSTADDRESS$
}
```
