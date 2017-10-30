# cisco-update

# Cisco bulk command script - web0 2014

This is kind of reinventing the wheel here and there are likely several better utilities for this. It started out a basic expect script to login to a device and eventually became more complicated.

 This can be used with a bash loop to apply updates to a list of devices
 See webster for lots of examples of how this can and has been used

 Commands and syntax vary depending on device and IOS version
 To overcome this I've typically used rancid or other config files to grep for 
 a specific target device and built out the commands in text files for each
 Don't forget to write memory

 Requires the Net::Telnet::Cisco cpan module which can be installed
 with 'sudo cpan Net::Telnet::Cisco' or with cpanminus

# Usage

Usage: ./cisco.pl -h hostname 

NOTE: Either -c or -f arguments are required

OPTIONS

 -c		Command to run on host, use ; for multiple lines

 -e		Set enable mode

 -f		Text file containing commands to run

 -h		Hostname or IP of device

 -n		Enable Mode Password - defaults to our standard

 -o		Log file to append output information to

 -p		Password - defaults to our standard pw

 -u		Username - defaults to spectrum

 -v		Verbose mode: display output on screen


