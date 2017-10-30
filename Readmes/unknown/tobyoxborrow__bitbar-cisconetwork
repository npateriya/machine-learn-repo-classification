# Cisco Corporate Network BitBar Plugin

## Introduction

BitBar plugin to display combined connection information indicating:
* Corporate network status
* Internet connectivity status
* VPN status (up or down)
* Country we appear to be from - if connected via VPN may be different from
  your local connection.

The output is intended to be minimal, only displaying information when
relevant. As such, when you have a basic working connection is when it displays
the least.

## Examples

Vanilla connection from HK, no issues.

    HK

Vanilla connection from HK, disrupted Internet connectivity.

    HK www:xx dns:ok

Connected to corporate network via VPN, appearing to be from Singapore.

    .:I:.:I:. <vpn> SG

No connectivity.

    Unknown connection www:xx dns:xx

## BitBar

BitBar lets you put the output from any script/program in your Mac OS X Menu
Bar.

Project home: https://github.com/matryer/bitbar

## Installation

Install BitBar following the project's own instructions.

To install the plugin:

    # Go to your BitBar plugins directory and clone the project
    cd /path/to/plugins/directory
    git clone https://github.com/tobyoxborrow/bitbar-cisconetwork CiscoNetwork

    # Ensure the script is executable
    chmod +x ./CiscoNetwork/cisconetwork.sh

    # Create Enabled directory if it doesn't already exist
    mkdir -p Enabled
    cd Enabled

    # Create symlink in the Enabled directory to the script and encode the time
    # (e.g. 10s = ten seconds) into the name so BitBar knows how frequently
    # to execute it.
    ln -s ../CiscoNetwork/cisconetwork.sh ./cisconetwork.10s.sh

Choose `Refresh` from the BitBar menus.
