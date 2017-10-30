# cisco-config-snips
Configuration snippets for Cisco devices

ASA-outbound-DNS-NAT.txt
see blog post at http://billyc5022.blogspot.com/2016/12/using-cisco-asa-nat-to-translate.html

So you have decided to use Cisco Umbrella or OpenDNS as your recursive DNS. Good choice! You update your internal DNS servers to point to 208.67.222.222.

All done, right?

Then you check your firewall logs and notice there are devices sending DNS queries directly to public DNS servers. How can you force those devices to use 208.67.222.222?

With NAT!

First you need to identify which external DNS servers are being used. Then you need to NAT DNS requests to those external DNS servers to the OpenDNS server.
