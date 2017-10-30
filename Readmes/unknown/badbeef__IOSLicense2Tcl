# IOSLicense2Tcl
New Cisco IOS routers require licenses to add new functionalities. But in order to write install license, you need to first copy license file to the router. This may not always be easy.

Here is a small utility, written in C, to generate Tcl script which generates the entire command sequence to install the license. As long as you have console access to the router, you can install license.

To use it, you need to first compile the program (using C compiler). The program takes the license file from standard input:

./a.out < my_license.lic

and then you can copy & paste the output to router.

Alternatively, you can redirect the output to a file and copy & paste the license file to the program:

./a.out > install.tcl

