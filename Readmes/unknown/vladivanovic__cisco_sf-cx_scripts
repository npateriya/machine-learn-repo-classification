# cisco_sf-cx_scripts
These are Python scripts I built for working with Syslog data from Cisco CX and Sourcefire devices.

These devices through no fault of their own output quite a lot of data in Syslog format (between 40-60GB of day for an organisation I worked at) which is great news but hard to parse so much data.

This is where Splunk came in quite handy but limitations with the amount of logging I could input into the Splunk system I had access too meant that these big chunky files needing a lot of rewriting to compact into a format I could upload into Splunk but still obtain the data I needed.

<b>SF_Rewrite.py - Rewrites SourceFire Syslogs (takes about 1hour per 10GB of log) into CSV Format</b><br />
I used regex to pull the following variables:
Time and Date, Connection Type, Client, Application Protocol, Web Application detected, Initiator Bytes, Responder Bytes, and then lastly, connection type UDP/TCP, source/destination ports and source/destination IP Addresses.

<b>CX_Metered_Rewrite.py - Rewrites SourceFire Syslogs and Outputs bytes totalled (takes about 1hour per 10GB of log) into CSV Format</b><br />
I used regex to pull the following variables:
Bytes Recieved, Source and Destination IP, AVC App Name (application i.e. chrome as detected by NBAR), AVC App Type (application type, i.e. web, bitorrent, etc as detected by NBAR), URL Category (as categorised by Cisco Web Sec), Event Gen Time, Bytes Recieved, Bytes Sent and Bytes Total

<b>Convert_CX_to_DB.py</b><br />
My attempt at converting essentially CSV into SQLite3 format (attempting to find alternative methods than Splunk)

<b>CX-Met-OFFNET_Rewrite.py</b><br /> If your current internet connection is based on a peak/non-peak period of usage, this script would read through UTC Values provided by the CX Event Generation Time and pull out all lines and write to a new file for the given time period, i.e. if you want to analyse peak time only.

I also have some Netflow rewriters but they're specific to Netflow provided by the ISP at the time so I have chosen not to upload these here.
