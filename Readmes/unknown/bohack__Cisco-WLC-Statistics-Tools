# Cisco-WLC-Statistics-Tools

This set of scripts pulls statistics from a WLC for a list of MAC addresses. You can then parsing the data
and replay connection type, speed, SNR and a number of diagnostic statistics for each MAC address.

The parser requires the SQLite3 ODBC driver installed. You will need to install the 32-bit version.
http://www.ch-werner.de/sqliteodbc/

You will also need to enter in the command wscript //H:CScript to allow for console scripting of VBS programs.

The script works in two parts:

Collection
The GetDebugsFromWLC.vbs script collects information from a list of MAC addresses.
The MAC addresses should be formated in Cisco formatting i.e 94e9.7922.167d.
The MAC address list file should have a MAC address per each line.
It does this by using the send keys VBS command to virtually type into a Putty session the pertainant commands.
Before you begin collection of data, make sure that you set the Putty logging to log out to a file.

Parsing
The ParserDebugsFromLogs.vbs script will parse the information logged to the putty file during the collection process. The parser will output to a CSV.
Once it is output to a CSV, the infromation can be opened in Excel and the data can be filtered and analysed further.
