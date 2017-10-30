# CiscoBackup
# Cisco MDS Fibre Channel Switch Configuration Backup Using SSH
#Author: Deepak Naidu

#Script: CiscoBackup.pl
#
#Info: Prerequisite config on switch for using remote SSH command to backup Cisco MDS switch config.

username CiscoBackup password 5 !  role network-operator

username CiscoBackup sshkey ssh-rsa “KEY KEY==” CiscoBackup@server

snmp-server user CiscoBackup network-operator 

no telnet server enable

ssh key rsa 1024 force

ssh server enable

