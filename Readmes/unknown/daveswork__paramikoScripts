# paramikoScripts
Moving this over from my dead and dying other githubaccount
https://github.com/davescriptingspot
Attempted blog with further breakdown on this works:
https://davescriptingspot.wordpress.com/

OK, 

I'm creating this because it took me far too long to google the 
individual bits and pieces needed to get Paramiko to perform something
as simple as log in to priviliged mode and issue a "show" command. 

Most of the examples I've seen, while written very well and functional, are
a bit cluttered with error handling. While this is important, it was a little
challenging to filter out the basic mechanism that was the heart of the script
from the error handling portion. 

I'm taking the approach of writing small scripts with ample comments for future reference.
This is not intended to be a fully integrated suite of tools, but more of a dissection of 
various functions .

Hopefully this will help someone new or strugling. 

All of them will assume you have Python 2.7 and Paramiko installed. 


#Contents

#---------------------Cisco-----------------

cisco-01.py :  Login via ssh, enter priviliged mode, and issue a command , eg "Show vlan brief" 

cisco-02.py :  Login to a switch via ssh, enter priviliged mode, and grab the current 5 min average throughput on all the interfaces

cisco-03.py :  Essentially cisco-01.py with some error handling, "try..except" wrapper for the ssh connection, and created the "issue_command" function to cut back on repeating the same lines when sending a command to the device.

cisco-04.py :  Essentialy cisco-02.py with the error handling and "issue_command" function from cisco-03.py

cisco-05.py :  This parses a CSV file of login credentials and issues the "show run" command on a batch of devices. 

cisco-06.py :  This builds on cisco-05.py and parses the output of a "show archive config differences" per host. 

cisco-07.py :  This parses the saved configuration and writes it to a local file . 

#---------------------Juniper----------------

junos-01.py :  Similar to cisco-03.py but for junos, just to give you the jist of how to at least get started there. 

junos-02.py :  Sample Paramiko script using RSA key based authentication to log into the device and run "show version|no-more"

junos-03.py :  Sample Paramiko script which connects to a remote device through a bastion server, aka, 'jump' box.

junos-04.py :  Similar to junos-04, but with more prompts and error handling, prompt for password instead of being hardcoded.
