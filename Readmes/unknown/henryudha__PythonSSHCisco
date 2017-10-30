## SSH Cisco Device with Python

This is a Class to ssh Cisco device through Python Paramiko

## How to Use

1. put Cisco.py on your current directory.
2. at the same directory, open python console
3. or if you want to create your own script with this class, please put the Cisco.py on the same directory with your script

## Code Sample

### Open your console

    $ python
    # import class
    >>> from Cisco import Cisco
    >>>
    # define object 
    >>> ip = "192.168.1.1"
    >>> username = "yourusername"
    >>> password = "yourpassword"
    >>> port = 22
    >>> command1 = "terminal length 0"
    >>> command2 = "show clock"
 
    # connect to device. we use connect function on Cisco.py file
    >>> ab = Cisco(ip,port,username,password)
    >>> ab.connect()
    u'\r\nRouter-1 #'
    
    # execute cisco command. we use execute_command function on Cisco.py file
    >>> ac = ab.execute_command(command1)
    >>> ad = ab.execute_command(command2)

    # print output
    >>> print ac
    terminal length 0
    Router-1 #
    >>> print ad
    show clock
    12:47:22.553 GMT+7 Mon Oct 24 2016
    Router-1 #

    # close connection
    >>> ab.close()

