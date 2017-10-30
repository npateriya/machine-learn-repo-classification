


Cisco ACI APIC Model Viewer
===========================

*****************************************************************************************************
IF YOU ARE LOOKING FOR THE STATELESS APP TO RUN EMBEDDED IN THE APIC (YOU SHOULD BE :-) ) 
YOU NEED TO GO HERE >> https://haystacknetworks.com/cisco-apic-managed-object-browser
*****************************************************************************************************

The ACI Model Viewer visualises the APIC configuration through the use of graphs illustrating the configuration of the APIC. The model viewer has been used and added to over the last few years and hopefully will be of help to the community as it has been to me and my colleagues.  


Uses
====

   * Assists in learning how the model has been implemented
   * Enables troubelshooting where components have not been configured correctly
   * Assists in writing scripts against the APIC by drilling down into the model to find data or objects
   
   
Drilling Down On A L3 Out   
![alt text](/ss-1.PNG)
   
Drilling Through a Leaf Policy Group Bi-Directionally
![alt text](/ss-2.PNG)


Important Notes
===============
 
As this has been built over time for my use there are certain restrictions (or untested uses) of the code. 
 
   * The code uses cross site scripting to the APIC, so to avoid errors and the script being denied the browser web security must be disabled. In Chrome this is done by starting Chrome with the --disable-web-security switch. A Windows Chrome lnk file has been provided to start chrome in this mode.
   * Due to the operation of XHR in JS and that the APIC in HTTPS mode only uses a locally generated certificate, XHR will fail to      permit the operation connecting to the APIC due to the invalid certificate. This would not be an issue where the APIC has had      public signed certificate installed. To workaround this issuee, when the browser is first opened without web security as discussed, open a page to the APIC and accept the locally signed certificate. Once this is done, then open up the model viewer html page (index.html)
   * To store the APIC IP, username and password, edit the default credentials in apic-comms.js
   
   
Usage Steps
===========
1. Open apic-comms.js and change login details, app will prompt for them but wont save them between sessions, this will setup defautl fields.

2. Use the link "chrome(no sec)" to open Chrome with web security disabled, this is due to cross site scripting and HTTPS restrictions.

3. Open a tab in the chrome session and enter the APIC IP and accept the invalid certificate (only required if the APIC has the default self signed cert - if a public signed one has been installed then this does not need to be done)

4. Open a tab and browse to and open index.html

5. Validate login details and click submit, the app should login to the APIC.
   
Author
======
  Simon Birtles http://linkedin.com/in/simonbirtles



Licence
========
  * Copyright (c) 2015-2017, 
  * All rights reserved.
  * Dual licensed under the MIT and GPL licenses.

