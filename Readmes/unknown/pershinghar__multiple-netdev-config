# multiple-netdev-config
Remotely run commands on multiple cisco/dell/force10/juniper/... devices through SSH

usage: run-commands.py [-h] -d DEVICES -c CMDS -u USER [-p PASS] [-k KEY] [-e]
                       [--debug]
    example:
        
	    ./run-commands.py -d "123.234.56.7,10.2.3.4" -c "./cmds.txt" -u admin -p password

### arguments:
```  
  -h, --help            show this help message and exit
  -d DEVICES, --devs DEVICES
                        Devices which we send commands to, divided by commas, example: "10.1.2.3,10.2.3.4"
  -c CMDS, --cmds CMDS  Commands to send (file)
  -u USER, --user USER  SSH UserName
  -p PASS, --pass PASS  Password to device (plaintext)
  -k KEY, --key KEY     Key used to ssh login (when no pass specified)
  -e, --enb             Use enable command <not implemented yet>
  --debug               debug mode
```

This script is still in development, I'm willing to hear feedback - any requests, tips, questions.
