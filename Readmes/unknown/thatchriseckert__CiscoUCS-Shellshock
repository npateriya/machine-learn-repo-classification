**Cisco UCS Manager 2.1(1b) Shellshock Exploit**
---
- CVE-2014-6278
- Confirmed on version 2.1(1b), but more are likely vulnerable.
- Cisco's advisory: https://tools.cisco.com/security/center/content/CiscoSecurityAdvisory/cisco-sa-20140926-bash
- Exploit generates a reverse shell to a nc listener.
- Exploit Author: @thatchriseckert

Exploit goes after a specific cgi script in Cisco UCS manager to create a reverse shell back to the ip/port specified.  Also posted on edb:  https://www.exploit-db.com/exploits/39568/

Usage -
```
[*]Usage: <Victim IP> <Attacking Host> <Reverse Shell Port>
[*]Example: shellshock.py 127.0.0.1 127.0.0.1 4444
[*]Listener: nc -lvp <port>
```



