# cisco-spark-domoticz-integration

Receive Domoticz home automation system/Switch messages in Cisco Spark
paglanz@cisco.com


1. If you haven't done so, create a Cisco Spark account for the Domoticz system to use here: https://www.ciscospark.com
This account should be different from your Spark account you use for use daily business. 
You will require a separate email account to sign up.

2. Make sure that you have curl installed on your Domoticz raspberry or Debian flavoured server:
log on to your domoticz server via ssh. If you're sitting in front of a windows computer, you want to use Putty. Check here to dowload: http://www.putty.org/.
Raspbian uses the user „pi“ as the default user, with the default password „raspberry“. Your milage may vary, check your distribution documentation for default/root logon credentials. The script we're going to create will need the curl package to be installed, which it isn't per default. For Debian flavoured distributions,  let's do this real quick on the command line: sudo apt-get install curl

3. Gather your Spark Access token here: https://developer.ciscospark.com/getting-started.html (log in with your spark account created in step one)

4. Replace the string "your_personal_access_token_here" in the command line below with your actual access token you gathered in step3. Replace your_primary_spark_account@acme.com with the email you signed up with for your primary spark account.

curl https://api.ciscospark.com/v1/messages -X POST -H "Authorization:Bearer your-personal_acces_token_here” --data toPersonEmail=your_primary_spark_account@acme.com --data “text=Domoticz was here”

If everything works out okay, you should see a message pop up in your primary spark account, once you issue the command above.

5. Now you can create the script (or use sparkmessage.sh) on the domoticz box that sends notfications to your spark account:
Create a scripts directory in the home directory of the user pi:
mkdir scripts
Change directory into the scripts directory:
cd scripts
Now use nano (a fairly simple text editor) to create your script:
nano sparkmessage.sh
copy&paste text below into nano and edit your accesstoken+room ID (to discover room IDs available to you, use "listrooms.sh"

#!/bin/bash
accesstoken="your_spark_access_token_here"
roomId="the_room_id_of_the_room_you_want_domoticz_to_post_messages_to"

curl https://api.ciscospark.com/v1/messages -X POST -H "Authorization:Bearer $accesstoken" --data "roomId=$roomId" --data "text=$1"

Save your script by pressing Ctrl-O.
Exit the editor by pressing Ctrl-X
make the script executable: chmod 755 sparkmessage.sh

6. Now the domoticz part: Open the Domoticz Web UI and go to Setup/Settings/Notifications
Check the "Custom/HTTP/Action" checkbox and paste below line  into the URL/Action field:

script:///home/pi/scripts/sparkmessage.sh "#MESSAGE" 

Please note the triple slash: script:///

Once done, hit the Test button and you should see a message in Spark.

7. Now you can have domoticz send a notifications once a switch changes its status. Simply click on "Switches" in the main menu of the domoticz web ui, select a switch you want to receive notifications for, click on notifications and make sure that http is checked under "active systems". Edit your custom Message for on/off and you are done.


# cisco-spark-domoticz-integration
