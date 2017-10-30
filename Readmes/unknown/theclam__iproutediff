# iproutediff

Compares the output of two "show ip route" commands from Cisco IOS CLI and notes what
changes, additions and removals have taken place between the two.

This script requires Python 3.

Usage - Linux / Mac
===================

Set the iproutediff.py file as executable (chmod 755 iproutediff.py) then run it

Usage - Windows
===============

Run python 3 then execute:

exec(open("./iproutediff.py").read())

Example
=======
Harrys-MacBook-Air:iproutediff foeh$ ./iproutediff.py 
iproutediff - takes the output of two "show ip route" commands and notes the differences

v0.1 alpha, By Foeh Mannay, September 2017

Enter filename of first 'show ip route' (A): A.txt
Enter filename of second 'show ip route' (B): B.txt

Changed:

<<< B       192.168.115.0/27 [200/10] via 10.120.21.17  1d02h
>>> B       192.168.115.0/27 [200/15] via 10.120.21.17  1d11h


Changed:

<<< O       10.155.21.7/32 [110/1021] via 10.140.100.234 TenGigabitEthernet5/2 1d12h
>>> O       10.155.21.7/32 [110/1010] via 10.140.100.230 TenGigabitEthernet5/1 01:02:00


Removed:

<<< B       192.168.115.152/30 [200/10] via 10.132.210.19  2w2d


Added:

>>> B       192.168.115.80/32 [200/10] via 10.132.210.1  01:02:00



Bugs
====

I've barely tested the code so probably loooooooads!
