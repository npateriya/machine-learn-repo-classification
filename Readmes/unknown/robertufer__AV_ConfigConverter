# AV_ConfigConverter
Avaya switch configuration converter to Cisco IOS

The goal of this project is to grab the important information from Avaya Passprt switch configurations 
and translate the commands into a format that can be pasted into a Cisco switch, for converting from one device to another.
It is very alpha, and will likely not be very polished....


To use it:
    create folders [Inputfiles] and [Outputfiles] in the same dir as the script. 
    If you want to apply a base configuration for the common stuff, change the file open function target in def fileinout, 
    and place the config file in the same directory as the script. If not, comment it out (I don't really want to make it test
    for the file yet and react if it doesn't exist, because I don't care). If you don't want to comment out a bunch of stuff,
    just make a blank file.
    
    load the Inputfiles folder with the scripts to convert - this is built for Avaya Passport configs, 
    I would like to test it out with other types in the future.

Notes:
    the OSPF routes are commented out, to allow for putting them in the interface configs (the new cisco way) or as 
    network statements in the routing process (old style).
    The vlans to port / port channel mappings are a mess now, I am eventually going to redo that part (probably load them into 
    an array due to the way each brand of switch handles the vlans.
