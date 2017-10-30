#vpncwebctl

Control your Cisco VPN client remotely

##Description

Assuming that you have the Cisco VPN client already installed in your VPN gateway server, this tool provides you a web interface allowing you to control your client remotely.

Basically with this tool you'll be able to: start and stop the Cisco client and see the VPN connection status.

##Installation

This utility is written using Bottle, a Python web framework.

First of all, install the required dependencies.

Using the Ubuntu/Debian package manager:

`apt-get install python-bottle`

Or using the alternative python package installer

`pip install bottle`

Then clone this repo in the `/opt` folder and fill up the `server.conf` file to fit your needs.

```
cd /opt
git clone git@github.com:frommelmak/vpncwebctl.git
cd vpncwebctl
python server.py
```

##Optional

Finally you can install [supervisor](http://supervisord.org/) in order to manage the server process.
