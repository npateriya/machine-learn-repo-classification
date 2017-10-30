# PythonPlusCisco
# Automating boring tasks using Python to interface with putty -> cisco switches
# So far I've got the def for portsecurity made, and am trying to find ways to launch putty with a python command.

def portsecurity(password, ip_add, Port_num):
	print "%r" % ip_add #this is a placeholder for when i find out how to use python for opening ssh
	print "user.name" #This is for using TACACS+ to log in to an SSH session
	print "%s" % password
	raw_input("CTRL-C to cancel, or hit enter") #if your login fails, it doesn't push the rest of the code meaninglessly
	print "conf t"
	print "int %s" % (Port_num)
	print "no switchport mac-add stick"
	print "switchport mac-add stick"
	print "shut"
	print "no shut"

ip_add = raw_input("IP Add? ")
Port_num = raw_input("Port number? (include F/G etc.) ")
password = raw_input("Password? ")	
portsecurity(password, ip_add, Port_num)

	
#The next few lines just read the shit you just put in	
#filename = "test2.txt"	
#target2 = open(filename, 'r') 
#message = target2.read()
#print(message)
#target2.close()
