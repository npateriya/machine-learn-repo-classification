check_asa_sessions.py
=====================
v0.2

usage: check_asa_sessions.py [-h] [--version] [--debug] [-w WARNING]
                             [-c CRITICAL] [-wl WL] [-cl CL]
                             SNMP_COMMUNITY HOST

Icinga (Nagios) plugin that checks the total amount of current, concurrent
sessions on a Cisco ASA and evaluates them against 'warning' and 'critical'
value thresholds.

positional arguments:
  SNMP_COMMUNITY        the SNMP community string of the remote device
  HOST                  the IP of the remote host you want to check

optional arguments:
  -h, --help            show this help message and exit
  --version             show program's version number and exit
  --debug               debug output
  -w WARNING, --warning WARNING
                        set high warning threshold
  -c CRITICAL, --critical CRITICAL
                        set high critical threshold
  -wl WL                set low warning threshold
  -cl CL                set low critical threshold

Written in Python 3.