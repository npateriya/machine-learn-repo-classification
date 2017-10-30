# CiscoSLA
OMD checks for IP SLA on Cisco devices.
Only supports ICMPECHO and UDPJITTER at this time.
show ip sla applications

Place all files in /omd/sites/<sitename>/share/check_mk/
Need to change a lot of things, check the issues for more info on future featrures

icmp-echo SLA:
ip sla 3
 icmp-echo 3.3.3.3
ip sla schedule 3 life forever start-time now

udp-jitter SLA(Remmenber IP SLA RESPONDER on the destination):
ip sla 4
 udp-jitter 10.0.101.250 5000 source-ip 192.168.30.1 codec g711alaw
 verify-data
ip sla schedule 4 life forever start-time now

