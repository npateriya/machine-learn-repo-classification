Cisco_ACI
=========

Scripts for Cisco ACI

Various scripts written in python for use with Cisco ACI (Application Centric Infastructure).

create_full_tenant.py

''' create_full_tenant.py
    The goal of this application is to create a new Tenant and associated Security Domain.
    It then creates a local user with full admin access to that Tenant.
    I have also put some pauses in the script so that you could use it to show a customer
    that the items are being built in the APIC before moving to the next item.
      
    This application does the following:
    1.  Collect a current valid admin username and password
    2.  Login to the system
    3.  Collect a new Tenants name
    4.  Create the new security group
    5.  Create the new Tenant
    6.  Collect the new admin username and password
    7.  Create the new user and place them in the proper security group
    *.  Display information back to the user (some day)
    
    Application created by Tige Phillips.  
    Please note that very little error checking is being done.  It's entirely possible 
    that my code could crash your system (not likely).  You have been warned!
'''
