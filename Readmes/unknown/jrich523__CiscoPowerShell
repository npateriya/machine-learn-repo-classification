Cisco PowerShell Module
=======================

_Due to IOS differences it is not recommended you use this_

connects in to a cisco unit to gather and parse data.

IOS - 6.0(2)N1(2)

PowerShell - created/tested on V5, should be V3 compatible

# Install #

	iex (new-object System.Net.WebClient).DownloadString('https://raw.github.com/jrich523/CiscoPowerShell/master/Install.ps1')

# Example #
Currently very limited functionality

	#load module	
	Import-Module CiscoPowerShell
	
	#store creds
	$cred = Get-Credential
	
	#connect to unit
	Connect-CSUnit 10.11.12.1 $cred

	#gather interface data
	$interface = Get-CSInterfaceStatus

	#take a look at the data
	$interface | Sort Vlan | Format-Table
	
