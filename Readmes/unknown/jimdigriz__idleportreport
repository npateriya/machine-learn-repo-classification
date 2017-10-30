This script hopefully can help any Network Sysadmin who finds tedious querying a switch that has no spare ports for patching, on which ports are 'available' as they are idle. The script should work on any Layer2 aware switch that SNMP::Info can talk to; of course I have been only able to test it against Cisco kit. It uses 'standard' MIB data and none of the Cisco specific bits so do let me know how you get along with it.

The following shows a trial running of what the script does on a switch with a mixture of idle and adminstratively 'shutdown' ports:

    ac56@node0:~$ idleportreport 172.16.1.172
    querying 172.16.1.172...
    uptime: 2w1d1h
      name: LibReadingRm-3548-2
     class: SNMP::Info::Layer2::C2900

    int    link admin time   description
    Fa0/1  down up    2w1d1h 049
    Fa0/2  down up    2w1d1h
    Fa0/3  down up    2w1d1h 051
    Fa0/4  down up    2w1d1h 052
    Fa0/15 down up    2w1d1h 063
    Fa0/18 down down  n/a    066
    Fa0/19 down down  n/a    067
    Fa0/25 down up    2w1d1h 073
    Fa0/29 down up    2w1d1h 077
    Fa0/33 down up    2w1d1h 081
    Fa0/39 down down  n/a    087
    Fa0/41 down down  n/a    089
    Fa0/43 down down  n/a    091
    Fa0/45 down down  n/a    093
    Fa0/47 down down  n/a    095
    Gi0/1  down up    2w1d1h

The script should work on any Perl enabled machine that has SNMP::Info installed. Simply download idleportreport and run it as in the above example. To use, remember to configure your SNMP communities/credentials at the top of the script to match your environment. And no, 'megasecret' is not our SNMPv2c community secret...
