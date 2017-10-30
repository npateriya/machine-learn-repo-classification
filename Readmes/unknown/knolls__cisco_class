# NX API Scripts


### Cdp_Crawler
this will crawl all the device build objects that contain info about the crawled device then returns a list of these objects. Use this as a base to gather all the devices and do something with them. Both scripts below use Cdp_Crawler.

![Crawl output](http://i.imgur.com/Fll5ZJE.png "Crawl output")

### AutoShark Multi
AutoShark Multiwill crawl everything get the devices list,then monitor each device in its own thread. When it sees an interface with over 1000 packets (this is a place holder it would monitor errors) then it creates an acl to punt everything to the cpu and starts an etheranylizer. Its not very useful Multi spanner does a better job.

![shark output](http://i.imgur.com/eb9Rb5o.png "shark output")

### Multi Spanner

Multi Spanner will crawl everything get the devices list, then monitor each device in its own thread. When it sees an interface with over 1000 packets (this is a place holder it would monitor errors) then it will create an erspan session on that port for x amount of time. when the timer ends it will clean up the erspan session and continue to monitor the device.

![span output](http://i.imgur.com/o07GepF.png "span output")

# APIC Scripts

### APIC_Faults

APIC_Faults finds all the faults on the APIC, and scans all the interfaces for port resets over 3, and egress packet drops. It then prints them all on the screen. Screen Shot is below of the output:

![APIC_Faults output](http://i.imgur.com/JyDjJGH.png "APIC_Faults output")
*There are no egress packet drops in the fabric currently so there are none displayed in the above screen shot.*

