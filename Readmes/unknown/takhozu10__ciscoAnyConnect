# Cisco AnyConnect for Mac 4.4  
Customize installation for enterpriase organization.  

__CiscoChoices.XML__ allows Cisco AnyConnect 4.4 package to be installed using the vendor provided pacakge but customize which modules to install.  
__anyConnectConfig.sh__ will install and customize the login by retrieving and replacing logged in user's username and vpn domain name to organization specific settings.  
__anyconnect__ file is a configuration file pre-configure to organization specific information.  

## CiscoChoices.xml file
This file allows custom installation method to be applied to the package.  
Currently, this file only enables VPN module during the installation. In order to enable the other modules during the installation, change the integer tag to the following  
```<integer>1</integer>```

## Edit anyconnect file
Enter the vpn server name here  
```<DefaultHostName>vpn.company.com</DefaultHostName>```

## Install Instructions  
1. Copy vendor provided AnyConnect.pkg file and CiscoChoices.xml file to the computer, in the /tmp directory.
2. Copy anyconnect file to /Users/```<username>```/.
3. Run anyConnectConfig.sh script to customize using user specific configuration.

## Install using JAMF Pro
1. Create DMG file that contains anyConnect.pkg, CiscoChoices.xml, and anyconnect file. 
    a. Place both anyConnect.pkg and CiscoChoices.xml in /tmp/ directory.
    b. Place anyconnect /Users/admin/ directory for now. We will use FEU (Fill existing user home direcotries) feature in JSS to place the .anyconnect file to existing user's directory.
2. Run anyConnectConfig.sh to install and configure Cisco AnyConnect application.