# traceip
Find a device on a cisco network/LAN, given it's IP

    Usage:
     traceip [OPTIONS] <IP_ADDRESS>
    OPTIONS:
     --help -? -h
     --debug         ;# show debugging info
     --router        ;# specify a router to query
     --switch        ;# specify a switch to query
     --vlan          ;# specify the vlan number of the IP address
     --mac           ;# specify the mac address for the IP address
     --noping        ;# do not ping IP address first

-------------------

    The "traceip" command queries routers+switches via SNMP.
    Unlike other tools, traceip is designed to do minimal SNMP queries in order to find a single device,
    not dump *all* arp/mac information from every router + switch
    It works in my (cisco catalyst) environment.
    It does *not* currently work for nexus routers.
    It has builtin assumptions that work for my network.
     It queries a single router (arp table) but will query many switches (mac address tables) in search of a host/device.
     This is consistent with (works well in) an end-to-end VLAN network.
     This does *not* work well with daisy-chained routers.
     (changing this would require dumping the entire route table for each traversed router
      and doing some IP/SUBNET "math" to find the best/matching route)
     It assumes that the 3rd octet of an IP/SUBNET matches the VLAN that SUBNET resides in.
     It does some strange community_string/VLAN "magic", particular to querying MAC tables in certain older cisco switches.
     It will discover neighboring switches via CDP and store that information to accelerate future searches.
     (if the switch topology changes, this information may become out of date and need to be cleared)
  
