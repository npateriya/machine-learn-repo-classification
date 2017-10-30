# VMMaestro Remote Access Using Linux and PuTTY (KiTTY)

How to configure Cisco VIRL – VMMaestro to use external telnet and SSH Client on Linux with KiTTY.

KiTTY is a fork of PuTTY and offer a command options to exec a remote comand after login.

1 - First, download Wine and install the putty. Create the exec shurtcuts whereaver you want. The mine is on ~/bin

2 - Download [kitty](http://kitty.9bis.net/) and replace the putty.exe executable with kitty.exe

3 - Configure Maestro, on Prefences->Terminal->Cisco Terminal to use PuTTY, as bellow


```
Both Telnet and SSH command:
/home/<user>/bin/putty 

Telnet Arguments:
-telnet guest@%h %p %r

SSH arguments:
-ssh -pw guest -P %p guest@%k -cmd %r
```

That's all!
