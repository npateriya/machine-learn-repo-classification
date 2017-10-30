# asa-dump-to-pcap
This project converts the "show capture xyz dump" output from a Cisco ASA and converts it into a PCAP file.

This script is designed to receive the ASCII text output of the "show capture xyz dump" command on a Cisco ASA firewall and turn it into a PCAP file.  This script is written in PHP, which means that it must run on a webserver, not as a command line script (i.e. like Perl does).  The benefit of this is that it can be hosted publicly on a internet webserver (by you) without having to authenticate to a shell.

This script works by receiving an html form submission that contains a bunch of text output that you paste in from the ASA firewall.  It digests that text and creates a file on the webserver with the PCAP output.

## Use case
You are troubleshooting some traffic through a Cisco ASA firewall.  You set up a packet capture on the firewall using the command "capture xyz interface outside" command.  After the capture has collected the packets, you want to view the data in Wireshark.  However, this is a multi-context firewall, so you don't have access to the flash memory to copy the capture there.  You also do not have ASDM access to the firewall, so you can't run the capture from there, instead.  All you can do is view the captured packets on the command line.  That is where this script comes in.  You can run the "show capture xyz dump" command, which gives you hex output of all the packets.  Now you can copy-and-paste that ASCII output into this script, which will generate a PCAP file out of it. Problem solved.
