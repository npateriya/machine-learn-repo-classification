TCL script for Cisco routers to check connectivity to the specific hosts on the internet.

The script generates syslog message with the result.


Configuration for running script periodically:

```
event manager applet PING_TEST
 event timer watchdog time 300 maxrun 250
 action 0.5 cli command "enable"
 action 1.0 cli command "tclsh flash:ping.tcl"
 action 2.0 syslog msg "$syslog_msg_output$"
 action 3.0 cli command "configure terminal"
 action 4.0 cli command "no event manager environment syslog_msg_output"
 action 5.0 cli command "end"
 action 6.0 cli command "exit"
```
