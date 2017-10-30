# Cisco_Syslog_Parser
Powershell script to parse Cisco Syslog messages to monitor infilteration

This script extracts the login events from the time filtered entries from syslog file, and it is up to you to further process using the data. This file is created to monitor network infilteration in cisco devices. The script runs from N-able scheduled task and collect timely information to get notifications on malicious activities. This might not be the best way to do the monitoring, but one of the ways to do the monitoring.

Example syslog entries

	Wed Apr 20 12:56:36 2016: <188>327: Apr 20 12:56:35.697 NZST: %SEC_LOGIN-4-LOGIN_FAILED: Login failed [user: theitteam] [Source: 192.168.56.112] [localport: 22] [Reason: Login Authentication Failed] at 12:56:35 NZST Wed Apr 20 2016
	Wed Apr 20 12:56:40 2016: <189>328: Apr 20 12:56:39.754 NZST: %SEC_LOGIN-5-LOGIN_SUCCESS: Login Success [user: theitteam] [Source: 192.168.56.112] [localport: 22] at 12:56:39 NZST Wed Apr 20

This same script could be used to filter syslog entries from other devices too. The filtering regex and the logic for filtering need to be changed accordingly.
