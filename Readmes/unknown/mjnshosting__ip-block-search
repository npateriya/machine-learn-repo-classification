# ip-block-search
Collect route tables from Cisco devices then search the collected data for desired block ranges.

##Special Thanks to:
Andres R.

Jared J.

##Story/Purpose/Objective:
Tasked with tracing IP blocks across 100+ devices to note if they were free or used by someone. I looked at the task
and said, "You gotta be fucking kidding me! Im to dam PRETTY for UGLY work!" So It took me a week to develop and test the following 
scripts along with my coworkers who were doing the process manually. It grew from just 5 devices to scanning all devices across the
network which stretches from the east to the west coast. 

It is four files and while testing created 750MB+ of files that consisted of route tables. Its done in bash and only requires clogin with a proper .cloginrc file which is included in the rancid package and needs to be mofified to suit your environment. So if you have the non minimal version of your OS this should not require anything more than 
rancid. 

Split into four files: 

device-list - each line contains the name or ip of a network device you want clogin to access and dump the route table from.

collect-route.sh - collects the route table for all devices listed in device-list. There are plans for this to do more later. 

free-ip.sh - takes in your /16 block range (examples below) and uses the collected route data and extracts information pertaining 
			 to your specific block. 

.cloginrc - Not something I created but it is required. check the example that I included. Google for further documentation. 

#Requirements:

Install the "rancid" package from your OS's repo. This script was written on Fedora 22 system but was also tested on a Ubuntu system.

.cloginrc needs to be in directory of script.

All devices must be placed in the "device-list" on their own line. 

#Instructions:
nano free-ip.sh

                copy and paste the code contained the attached zip file "free-ip.txt". Then press ctrl+x then ‘y’ then ‘Enter’

nano collect-route.sh

                copy and paste the code contained the attached zip file "collect-route.txt". Then press ctrl+x then ‘y’ then ‘Enter’

chmod +x free-ip.sh collect-route.sh

                make them executable 

nano device-list

		copy and paste the code contained the attached zip file "host_net-devices.txt". Then press ctrl+x then ‘y’ then ‘Enter’

nohup ./collect-route.sh &

nohup ./free-ip.sh <value of first octet> <value of second octet> <value of start of range third octet> <value of start of range third octet> &
		
		Some files will be created but the important one is in the results folder named <oct1>-<oct2>-<oct3_start>-to-<oct3_end>-free-ip-report.txt use nano to look at it and 
		copy and paste the output to a notepad file for you to fill in the spreadsheet or documentation.         

nano results/<oct1>-<oct2>-<oct3_start>-to-<oct3_end>-free-ip-report.txt

#Run Example: 

nohup ./collect-route.sh & 

		I like this program to be ran in the background as it took an nearly two hours for 105 devices and created 753MB of total data that will be used later. 

tail -f nohup.out
			
		**When using nohup tail the nohup.out file for an update to see when the script is finished or just down use nohup with the & to run this in the foreground. Not suggested because this can take hours and will not finish where it left off. **

nohup ./free-ip.sh 192 168 0 26 & 
		
		This should be ran in the background as well. This will search through the collected *-route.txt files and finally create a sorted -report.txt file under the results folder.

tail -f nohup.out

cat result/192-168-0-to-26-free-ip-report.txt

		It should look something like this. 

On device FastEthernet0/7 route C 192.168.21.0/27 is directly connected belongs to customer Cust007:Bond-Networks-London:DIA

On 192.168.61.1 Vlan07 route C 192.168.21.192/28 is directly connected belongs to customer Cust007-T1

On r0.ga.mjnshosting.com GigabitEthernet2/1 route C 192.168.38.0/24 is directly connected belongs to customer Cust07-29-16:Jason-Bourne-Services

On edge-r0.colo1.ga.mjnshosting.com !!Please Investigate Manually!! route S 192.168.21.224/27 [1/0] via 192.168.32.20 belongs to customer !!Please Investigate Manually!!

On coredevice1000 Null0 route S 192.168.1.31/32 is directly connected belongs to customer

On 11.20.85.30 Null0 route S 192.168.26.4/32 is directly connected belongs to customer

#Explanation:

##Output Structure:

On <device name from device-list> <interface on device> route <route from device that contains our desired blick> belongs to customer <customer/interface description for interface on device>

Its pretty easy to read and use the given output. Some things do need to be pointed out. 

	1. Only Connected (C), Static (S), Local (L) routes are shown. We only need to see where blocks are routed so edge devices are good but sometimes all devices may be best for backbone devices that have 10GE customers. 

	2. If there is no interface as you can see with the Static (S) route the script will output "!!Please Investigate Manually!!" in place of the interface name since one is not given. 

		On edge-r0.colo1.ga.mjnshosting.com !!Please Investigate Manually!! route S 192.168.21.224/27 [1/0] via 192.168.32.20 belongs to customer !!Please Investigate Manually!!

	3. If there is no output for "sh int <interface> | i escription" on the deivce then nothing will be printed. You will see this on Null0 interfaces like the example shown below. 

		On 11.20.85.30 Null0 route S 192.168.26.4/32 is directly connected belongs to customer
