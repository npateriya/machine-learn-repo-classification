Script to get end points from Cisco APIC.

usage: getACIEndpoints.py [-h] [-i IPS] [-m MAC] [--csv]
                          [--aci-user ADMINUSER] [-s SERVERNAME]
                          [--aci-pass PASSWORD]


This tool connects to ACI and pulls end point information. If no end point information is provided, all end points are return.

optional arguments:

    -h, --help            show this help message and exit

    -i IPS                Provide a list of IPv4 addresses to search for (one or many)

    -m MAC                Provide a list of MAC Addresses to search for (one or many)

    --csv                 Switch to enable exporting to file. File will be written to working directory.

    --aci-user ADMINUSER  Provide the user name for ACI access. Default is admin

    -s SERVERNAME         Provide APIC DNS name or IP address

    --aci-pass PASSWORD   Enter Password for APIC access. If none provided, you will be prompted

  
  > Examples
  

  __Get the entire end point list from the APIC__

    getACIEndpoints.py -s <ServerIP>

  __Get the entire end point list from the APIC written to file__

    getACIEndpoints.py -s <ServerIP> -csv

  __Get a specific end point from an APIC by IP address__

    getACIEndpoints.py -s <serverIP> -i <SearchIP>

  __Get a specific end point from an APIC by MAC Address__

    getACIEndpoints.py -s <serverIP> -m <SearchMAC>


***Note:*** By default csv files are written to the working directory.



