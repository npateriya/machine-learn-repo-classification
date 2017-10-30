
What does it do?
================

The Script pulls the Status Page of a my Modem (Cisco EPC3212), extracts the Signal Levels and signal to noise ratio (SNR) and logs them via rrdtool

It can also create a graph of the logged data, so one could see how the levels changed over the time.

It can also login to the admin section and pull the error log (The modem clears it's log on power loss).  But currently it does not store them in some Database.

Why?
====

Because my ISP claims to "constantly optimize" the network. 
Now, if they "optimize" the signal levels, so that their amplifiers don't use that much power, I can easily recognize that :-)


Usage
=====

set USERNAME and PASSWORD.
I did not include them, you never know if it is telling some big-company-secret. btw. you can easiliy google them :-)

Modemlogger.sh init
to dreate output and database dirs and init the database

Modemlogger.sh getStatus
to pull the status page and log

Modemlogger.sh getErrors
to pull the status page, log in admin section and log data

Modemlogger.sh graph
to generate PNGs



TODO
====

Evaluate the error log and store it somewhere safe.

Have the pngs formatted nice.

Think about having the db evaluate min and max values instead of avg




