Python program to tag ports on Cisco Nexus 5K switch with specified vlan.
Network was setup with Cisco 7K|5K|2K switches with servers L3 gateway on 7Ks and 5K switches only doing layer 2. 2Ks were dual homed to both parent 5Ks however that is optional for this script to work.

Script takes datacenter location to find nexus 7K ip to log into and vlan number want to tag. Also input as many different ports need to tag in the format of the 2K FEX serial number and port number. Uses regex to strip out those two pieces of information so all the following examples are valid:

FOX1739GGGG 45  
FOX1739GGGH;46  
AF04, U20, 3MMMMM; 1, <->, AF04, U44, FOX1739GGGJ;47

Script also uses cdp to branch out from starting point of 7Ks so no need to update if add more 5|2Ks.
