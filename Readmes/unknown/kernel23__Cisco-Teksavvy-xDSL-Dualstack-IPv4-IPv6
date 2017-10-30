Cisco and Teksavvy DSL dualstack IPv4 & IPv6
============================================

## Setup for my configuration

Sample configuration for Cisco IOS to run IPv4 and IPv6 over Teksavvy xDSL ISP.
My setup is as follow. Please adjust to your needs :

* VLAN 1 is LAN
* Dialer 1 is IPv4
* Dialer 2 is IPv6

I am using a Cisco series 800.

## Got IPv6 ?

To get IPv6 credentials, you need to contact Teksavvy at "ipv6@teksavvy.com".
NOC will assign you two IPv6 subnet.

* One is /64 (2607:....:....:....::/64)
* Second is /56 (2607:....:....:....::/56)

The subnet "/64" will be for routing. Don't use it in your local network to avoid overlap.
It should not appear in the configuration since i use "ipv6 address autoconfig".

First, split your /56 to use a smaller subnet. For example one "/64" is a goog choice and should be enough
for your local network).

In the sample configuration, i split my "/56" subnet into 256x "/64" and choose to use the first one.
It is identified as "2607:....:....:....::/64".

You can use online tool to help with this :

* http://www.gestioip.net/cgi-bin/subnet_calculator.cgi

Teksavvy IPv6 resolvers at this time are :

* 2607:f2c0::1
* 2607:f2c0::2

## Security / ACL ?

Basic ACL are included to prevent local network to be exposed as each hosts will NOW have public IPv6.
You should also disable IPv6 on all hosts you don't want to reachable.

> echo 1 > /proc/sys/net/ipv6/conf/all/disable_ipv6

## Testing ?

Quick test to check (ping Google Public DNS over IPv6) :

> [root@host ~]# ping6 2001:4860:4860::8888

> PING 2001:4860:4860::8888(2001:4860:4860::8888) 56 data bytes

> 64 bytes from 2001:4860:4860::8888: icmp_seq=1 ttl=53 time=52.1 ms

> `64 bytes from 2001:4860:4860::8888: icmp_seq=2 ttl=53 time=52.0 ms

> ^C

Also you can browse to https://ipv6.google.com !

Thanks Teksavvy, you rock !
