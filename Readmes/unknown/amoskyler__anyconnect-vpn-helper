anyconnect-vpn-helper
=====================
Simple Cisco AnyConnect VPN CLI tool.


## Setup
Step one: Get the VPN bash script from the github repo (curl command below)

    curl -o ./vpn https://raw.githubusercontent.com/amoskyler/anyconnect-vpn-helper/master/bin/vpn && chmod 700 ./vpn && mv ./vpn /usr/local/bin

This will do three things:
  1. Download the vpn CLI script from the github repository and store it in a file at your current directory named `vpn`
  2. Give the vpn script read, write, and execute permissions for only the current user
  3. and move the file `vpn` to the path `usr/local/bin`

Step two: Substitute in your own VPN credentials
  1. `subl /usr/local/bin/vpn`
  2. Now modify the following values to your VPN credentials
    - `__VPN_ADDRESS` // URL to vpn
    - `__CLIENT` // path to the Cisco AnyConnect vpn binary
    - `__GROUP` // VPN group (0 for two factor)
    - `__USERNAME`
    - `__PASSWORD`
    - `__SECOND_PASSWORD` // DUO method (e.g. PUSH)

Step three:
  If you're using bash change the first line to read #!/bin/bash (or to whatever shell you're using)


## Usage


### `vpn connect`
  - Type `vpn connect` to start the connection process...Complete the process by completing the two factor step

### `vpn disconnect`
  - Type `vpn disconnect` to disconnect from the vpn
  - This is sometimes necessary if the VPN has disconnected automatically

### `vpn status`
  - Type `vpn status` to check if the vpn is currently connected

### `-v` (verbose)
  - type `-v` as the second command argument to enable verbose output (e.g. `vpn connect -v`)
  - __notice:__ This will output your vpn authentication information.

## Known Issues
  - Disconnects unexpectedly due to 'inactivity'
  - The VPN will not connect (but will also not throw an obvious error due to output suppression) if another vpn client is open
  - Verbose flag only support `-v` but not alternate forms (i.e. `--verbose`, `-verbose`)
