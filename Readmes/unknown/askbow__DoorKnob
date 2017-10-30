# DoorKnob
DoorKnob is an Android app which has one function: it opens a door in Cisco Physical Access Manager (CPAM) 1.4.1 (and probably later versions as well) environment.

I designed it in MIT App Inventor, and thus the source code requires MITai.

# Intended usage
DoorKnob was designed with the following usage scenario in mind:

1. An employee goes outside his or her office, leaving personal access card at work desk.

2. Good thing, the company uses CPAM for building access control=)

3. Each employee has a personal login in CPAM (via AD integration, for example) with enough rights to open one door via web request.

4. Corporate WiFi is accessible in front of the door, which the employee in question has rights to open remotely. For obvious reasons, you must protect WiFi by means outside the scope of DoorKnob. I presuppose that enterprise network access is sufficiently secure.

5. Company CISO should think of protecting BYOD equipment via some means outside the scope of DoorKnob before allowing employees to use an app, which grants physical access to company premises.

# Structure, operation and Protocols
DoorKnob connects to CPAM via HTTP, sending GET requests to CPAM scripts designed [per Cisco documentation] to interact with Cisco IP-phones.

Basically:

        DoorKnob====HTTP-GET=====>CPAM

                <===HTTP-REPLY====
		
I use two CPAM web methods in DoorKnob's operation:

1. To request access to open the door specified in settings:

     http://cpam.example.com/QuickUnlock.jsp?UserName=USERNAME&pass=PASSWORD&method=authenticate&resourceId=DOORID&resourceType=door
	 
2. To discover the DOORID, which the specified user has rights to open via web on the CPAM, server specified:

     http://cpam.example.com/IPPhoneManager.jsp?UserName=USERNAME&pass=PASSWORD&method=authenticate

DoorKnob stores its settings internally via the MITai component TinyDB.

# Usage
Make sure, that you have created a separate user in CPAM to use with DoorKnob. The rights required to open one door remotely are quite limited, no reason to grant admin rights.
Install the app from the DoorKnob.apk file directly or use the DoorKnob.aia file to import the project to MITai and compile it yourself.

Run the app after installation.

In the main screen, click/tap Settings to open the Settings screen.
Specify CPAM server address, username/password and then click/tap Discover.
The Discovery screen might open briefly (even so briefly that you won't notice it, if there are no errors), connect to CPAM server specified and copy the resourseId of the first door listed in the server response to the Door ID field of the Settings screen. 
Click/tap the little arrow (back button) to return to the main screen.

Now you can click/tap OPEN THE DOOR. DoorKnob sends a request to grant access and briefly displays server's reply.

If everything is OK, the specified door should go to GranTAccess or Open state in CPAM.

# Disclaimer
This is a simple app for Android. It works (I have tested it), but it might still contain bugs, errors, mistakes. All the things you can typically find in work-in-progress state of any project.

Moreover, it is definitely insecure or at least you must consider it insecure. Use at your own risk.

Though I am the original author, exclusive rights for commercial use, as well as any and all others, are reserved by my employer --- Netcube LLC (Moscow, Russia; www.netcube.ru ).
For any use outside curiosity, please contact Netcube. Contact information is available on Netcube's website.

