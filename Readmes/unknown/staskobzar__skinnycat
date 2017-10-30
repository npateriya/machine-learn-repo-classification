## skinnycat: Small Cisco Skinny protocol command line utility

Small utility to emulate Skinny device communication. 
This utility is still under construction and for now provides only *register* and *keepalive* messages.

Requires Apache Portable Runtime library (apr-1.5). To install do:
```yum install apr-devel```
or
```apt-get install libapr1-dev```

To compile run: ```make```.
To generate documentation (requires doxygen): ```make doc```.
For testing (requires cmocka and dejagnu): ```make test```.

No install make target for now.

Example of use:
```
$ ./bin/skinnycat -M REGISTER -H 10.20.2.6 -m 5067AEE1D4D6 -v
[INFO]  Verbosity enabled.
[INFO]  Starting REGISTER to 10.20.2.6:2000 from device SEP5067AEE1D4D6
[INFO]  --> Register device 5067AEE1D4D6
[INFO]  --> IP port
[INFO]  <-- Register ACK
[INFO]  <-- Device buttons template. Total buttons: 2
[INFO]      [Button #1] : LINE
[INFO]      [Button #2] : SPEED DIAL
[INFO]  --> Date time template request
[INFO]  <-- DateTime response: 2016-10-18 14:21:39
[INFO]  <-- Clear prompt message.
[INFO]  <-- Display notify: Modulis Inc. - Skinny Protocol .
```

