(Assumes all routers are cisco routers with telnet enabled.)

  A. This program connects to its first Cisco router(e.g. 192.168.2.101). 
  
  B. It then uses "sh ip int br" command to list all the local IPs, and add them into "visited IP" of the router.
  
  C. Router's routing table provides list of next available IP, and the list will be added to ip list database. If this contains new IP, it will try to telnet from current router to the next router. 
  
  D. Steps B and C will repeat until all the "new IP" equal to "visited IP" 
