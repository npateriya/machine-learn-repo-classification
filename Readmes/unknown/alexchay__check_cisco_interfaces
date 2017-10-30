check_cisco_interfaces
================

This plugin uses the bulk-get mode of SNMPv2 by default. Support for SNMPv3 with and without privacy is included.


### Installation

In order to compile this plugin you will need the `NET SNMP Development` package
(libsnmp-dev under Debian) as well as `cmake` and the standard compilation tools.

Clone or download the project from GitHub

    mkdir build && cd build
    cmake ..

Running "make" should successfully compile the plugin, and "make install" will install them under
the configured path.

    make
    make install


### Usage


(see also check_cisco_interface --help)

    check_cisco_interface -c public -h 192.168.0.1 -r 'FastEth' -p '$SERVICEPERFDATA$' -t $LASTSERVICECHECK$ -a

    Options;
     -h                 address of device

     -c|--community     community (default public)
     -r|--regex         interface list regexp
     -R|--exclude-regex interface list negative regexp
                            Regex to match interfaces (important, this is a Regular Expression
                            not a simple wildcard string, see below)
     -e|--errors        number of in errors to consider a warning (default 50)
                            Only warn if errors increase by more than this amount between checks
     -f|--out-errors    number of out errors to consider a warning
                            Defaults to the same value as for errors
     -p|--perfdata      last check perfdata
                            Performance data from previous check (used to calculate traffic)
     -t|--lastcheck     last checktime (unixtime)
                            Last service check time in unixtime (also used to calculate traffic)
     -b|--bandwidth     bandwidth warn level in %
     -s|--speed         override speed detection with this value (bits per sec)
     -x|--trim          cut this number of characters from the start of interface descriptions
     -j|--auth-proto    SNMPv3 Auth Protocol (SHA|MD5)
     -J|--auth-phrase   SNMPv3 Auth Phrase
     -k|--priv-proto    SNMPv3 Privacy Protocol (AES|DES) (optional)
     -K|--priv-phrase   SNMPv3 Privacy Phrase
     -u|--user          SNMPv3 User
     -d|--down-is-ok    disables critical alerts for down interfaces
                            i.e do not consider a down interface to be critical
     -D|--regex_down    interface list regexp for down interfaces to be critical
     -a|--aliases       retrieves the interface description
     -A|--match-aliases also test the Alias against the Regexes
     -N|--if-names      use ifName instead of ifDescr
        --timeout       sets the SNMP timeout (in ms)
        --sleep         sleep between every SNMP query (in ms)


### Large Plugin Output


Be aware that this plugin may generate large outputs.  Your version of Nagios / Icinga may cut off the output and cause you problems with various graphing tools; for best results restrict the list of interfaces using the -r option

### Regular Expressions

The following patterns can be used to match strings

     .          anything
     ^          beginning of string
     $          end of string (WARNING: you need to use $$ in a Nagios configuration file!)
     (abc|def)  either abc or def
     [0-9a-z]   a range
     *          the previous pattern multiple times


Examples;

     Eth        match any strings containing "Eth"
     ^FastEth   match any strings beginning with "FastEth"
     Eth(0|2)$  match Eth0 or Eth2
     Eth(0|2)   as above but would also match Eth20, Eth21, Eth22 etc



