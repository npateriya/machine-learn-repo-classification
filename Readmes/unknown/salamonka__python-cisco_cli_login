# python-cisco_cli_login
Python Script using Getpass to collect a username, password and enable password

This script is built for python2.7 and uses getpass

It prompts for a username to be collected to be used by other processes such as netmiko to send to a a cisco device for login
It prompts for a password and then asked to re-enter the password.  It compares the entries together to ensure the user is not making a typo.
It prompts for an enable password, asks to re-enter and compares to the initial password to ensure no typo

Using getpass, both password and enable password are masked from view but stored in memory in clear text as they need to be passed to the cisco device in clear text using netmiko
