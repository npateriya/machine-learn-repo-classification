# myanyconnect
Cisco AnyConnect helper utility for OSX

While running `/opt/cisco/anyconnect/bin/vpn` is no big deal, I wanted something a little more tailored to how I would expect the VPN client to work.  This does nothing that the Cisco AnyConnect client doesn't already do.  Simply a work flow thing.

I use openconnect on Ubuntu, but I have issues with openconnect on OSX.  

# requirements
1. Cisco AnyConnect client must be installed. 
2. coreutils for timeout/gtimeout
    * `brew update && brew install coreutils`

# deploy
I copy to `/usr/local/bin`, but a link works or just keep it at the root of your home folder. 

# usage
```
Usage: myanyconnect OPTIONS HOST

VPN helper utility powered by Cicso AnyConnect client for OSX

Options:
  -d, --disconnect          Disconnect from the VPN
  -s, --status              Status of the VPN
  -u, --username            Username to use when connecting to the VPN
  -p, --password            Password to use when connecting to the VPN
  -h, --help                Display this help
```
# examples
## connecting
`myanyconnect vpn.mycompany.com`
```
username: jbfreels
password: 
Connected
```

## disconnecting
`myanyconnect -d`
```
Disconnected
```

## no user input connection
`myanyconnect -p MYSECUREPASSWORD vpn.mycompany.com`
```
Connected
```