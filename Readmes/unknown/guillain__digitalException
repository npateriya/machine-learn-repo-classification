# digitalException
Cisco Spark integration with SigFox button

## What is it?
* Bt.tn / Sigfox button as trigger
* <img src="doc/button.png" height="300px">
* Cisco Spark as Space for human and data
* ![](doc/chatbot.png)
* Local DB for historical and data post processing treatment... analytics ready :)

When the Button is pushed Cisco Spark space is created and people are added.

If **escalation** word is enter an escalation process is launch and additionnal member is added.

![](doc/workflow.png)

## Based on
* [Flask](http://flask.pocoo.org/) (Python)
* [bt.tn](https://my.bt.tn) (can provide simulator or physical button)
* [Cisco Spark](https://web.ciscospark.com) (for free)

## Features
* **Escalation:** add new member when keyword 'escalation' or 'escalade' is enter
* **Close:** close spacer when keyword 'close' is enter
* **Search:** search in the room history the previous message and print the list with web url to click on and jump in the original space.
* **Display:** print welcome and tips messages (to configure in the [settings.py](settings.py) file)
* **Tips:** random tips when people post new message

## PreRequisites
Configuration is provided for Apache and WSGI server.
But you can also get only the python with another web server, container...
* Apache2: to provide HTTP server
* MySQL: to record locally event as connector availability...
* python (2.7): development language
* * Flask
* * MySQLdb

## Doc
* [Install](doc/install.md)

## Report bug
[issues](issues)


Have fun ;)
