# CiscoTerminalServer
What will evolve into a full REST API for Cisco IOS and other devices

This is the initial codebase for development of a RESTful API for Cisco IOS. So far I have implemented a thread-safe job oriented system for pushing requests to the command line of the device. I haven't gone too deep into testing as I'm working mostly on parsers for the results of show commands.

Unlike other efforts of this sort which tend to float around the open source, I'm trying to be as complete as possible and am currently testing against a C3750X running IOS 15.0(2)SE2, a C3850 running a modern variant of IOS 16, and a 2911 with UCSE installed running IOS 
15.6(1)T0a... which now that I think of it is quite old.

Initially, the codebase is a crumby gui I can use for rapid development. I don't want to publish a REST API until I've considered all the issues I'm hoping to address. For now, I'm entirely focussed on creating the GET commands for all popular requests. I've tested my code (though not included in this version) for show run and show start. I haven't included these as it's too tempted to try and data mine from the config.

I have included pretty complete parsers for :
 - show inventory
 - show cdp entry *
 - show ip interface brief
 - show vlan brief
 - show ip route
 - show interface
 - show ip interface
 - show ip arp
 - show interface status
 
My near-term todo includes
 - show vrf
 - show ipv6 interface brief
 - show ipv6 interface
 - show ipv6 route
 - show ip ospf neighbor detail
 - show ip ospf interface
 - show ip ospf database router
 - show ip ospf database network
 - show ip ospf database summary
 - show ip ospf database asbr-summary
 - show ip ospf database external
 - show ip ospf database nssa-external
 - show ip protocol
 - show ip device tracking all
 - show ip access-list
 - show ipv6 access-list
 - show route-map
 
These each take a little time to do since instead of just plucking a simple expression out using a simple regular expression or fixed field input, I'm properly parsing them and trying to make the parsers dynamic enough to cope with all IOS versions I encounter. This means stuff like 'show ip route' can be very difficult'

I am not using other tools since I have been either extremely disappointed by the licenses (closed or highly restrictive), the code quality (chef, puppet, Cisco's yaml extensions), the resource consumption (Cisco Prime and APIC-EM), etc...

What I'm writing now should work on any system with the .NET Core framework available and possibly with ASP.NET... which means pretty much everything these days. In addition to a REST API, I'm also intending to massage the library into a self-contained Powershell module that can be auto-compiled using CSC which is included with pretty much everything these days. No, I'm not writing this in Python. There are dozens or maybe even hundreds of crap Python and Ansible modules for Cisco these days. Heck, Cisco makes 10 new ones every week and abandons them a week or two later. 

I'm focused more on quality, consistancy and compatibility than on simply pumping out another useless API. Cisco can continue to do whatever they want with their tools. I have tried most all of them now and while many of them are kinda cool, they're all basically designed by Cisco... for Cisco and they work with only the IOS versions Cisco wants to sell. I have yet to see any reason why a toolkit of this sort couldn't just work on pretty much anything.

I'm coding this as part of my enterprise/data center automation toolkit. This means it's primarily meant for automating through Powershell and while this part of the toolkit is focused on Cisco, the primary focus is on a fully Active Directory and Group Policy managed infrastructure. Cisco will simply provide equipment (servers, switches, routers), but it will all be configured through Windows Server, System Center and Powershell. 

Expectations :

I am perfectly happy to accept feature requests, unit tests, bug reports, code contributions etc... in the next commit or two, I'll include the licenses in the headings of the source files which is an MIT license. I might take money for feature requests people need "NOW!!!" or I'm not particularly interested in.

Goal :

I think I can implement policy based infrastructure without the need for new equipment. I believe it should be possible to get full policy networking on nearly all equipment Cisco has sold for 10 years. I hope to make it easier with ISE and TrustSec, but I can't see any reason why 802.1x combined with Active Directory and downloadable ACLs can't do almost all of it. Of course, it will require an SDK that can change the ACLs often. I think I'll have that covered soon.

How will I make money from this?

I don't intend to make money from this. At least not directly. I'm a Cisco instructor teaching in 6 different tracks. I am almost 100% focused on becoming a Microsoft guy now. The reason is, I don't think I can be good at networking if I can't integrate it properly with the rest of the infratructure. So, I'm automating everything and using it as a major learning experience. 

Notes :
I'm making extensive use of Renci.SshNet (also known as SSH.net) for communication. I'm also using Pegasus PEG parser generator for nearly all parsing. I'm doing my best to code readable recursive decent parsers for all input. I'm not overly excited about the lack of reuse of code as I'm not aware of how to implement an "include" so things like parsing IP addresses, prefixes, etc... don't end up spread out. I think I'll have to write some sort of templating code to make it easier to centrally manage highly repetitive code. For now, I'll do it the hardware and try to be consistant.

More notes :
'show ip route' is implemented but really really needs a lot of test data. I would like to support all 'show ip route' formats for all IOS versions going back to early in the 12.x series or even 11.x for good luck. I would appreciate help building a directory of test data for all versions of the file for the unit tests. I would also like some complex tables, that means bgp routing tables, DMVPN with NHRP, route replication, etc... Please shoot me as many different tables as you can. You can obsure the public IP addresses if you'd like.

Even more notes :
 'show interfaces' has been implemented true to the word. I have chosed to make it fail when it encounters unknown conditions as opposed to pushing past them since I believe it's extremely important to get full support for this show command. As of now, it seems to handle my 3750X, 3850, 2911, an old 803 with IOS 12.2 and BRI. However there is still very much to do. I will have to build a unit test framework very soon to work on this in more detail.
