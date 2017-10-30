Description about "cisco_assignment_neha.sh"
********************************************
********************************************

Script "cisco_assignment_neha.sh" is a menu driven program to do the following : 
1. Change Password 2. See the disk space 3. Login to other box using ssh 4. Show all Service running 5. Show all ports opened 6. Show all java apps running 7. Facility to kill a app 8. Exit

FILES:-
=======
1. cisco_assignment_neha.sh:- Main script which performs all the mentioned tasks.
2. login.sh script is an expect script which uses shh to login to another machine



EXECUTION FLOW:-
==============

1) Copy the script at the desired file system location
2) Own the script by running "chown username cisco_assignment_neha.sh"
3) Give execute permission to the script by running "chmod +x cisco_assignment_neha.sh"
4) Run dos2unix utility as the script was written in textpad on Windows.
4) Run the script "./cisco_assignment_neha.sh"
5) Give the choice which is needed

Now lets evaluate how each and every option in the script is executed:-
=====================================================================

1) User is prompted to enter old passwd followed by new passwd and then confirmation of new passwd.
2) User gets to see a list of the various disk drives and the free/used space. Default is checking disk space for entire filesystem.
3) User enters remote machine credentials (hostname, username and password). Once the user is prompted to enter "yes" for getting the remote machine credential in the handshaked machines' certificate s tore on the local box.
After that, whenever this script is run for ssh, make sure that the remote machine keys using ssh-keygen utility. This way user will not be prompted everytime for saying "yes/no" in ssh.
4) User gets to view all running services on his terminal.
5) User can view all the ports which are open(in listening mode)
6) User is able to get a list of all running java applications
7) User specifies the process name which is to be killed
8) User exits from the script

LOGFILE:-
=========
neha_task.log file is created in the directory from where the main script is run.
