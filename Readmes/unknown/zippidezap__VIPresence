# VIPresence
Cisco Spark location detection Bot with Meraki Location Analytics


Meraki CMX Spark Integration:
Retail:

This is an integration between Meraki Location Analytics, Spark. The idea here is for a retail store to be alerted via Spark when a high value customer enters the store. There are many use cases that also use this technology but high value/importance customers were the immediate ask for this project.

Our bot is fully interactive and allows you to monitor as many different MAC address' as you wish. Meraki Location Analytics posts data every time there is a change in the network to our application, we interpret the data sent to see if a device our bot is monitoring is in range. If so we then notify the end user as to the persons presence.

Below is a link to our overview video of the project:

https://cisco.box.com/s/scbq07uy7of2hzg9z9lkl609d37n6ary

Implementation:

The code is run using Node JS and BotKit. So these will both need installing first. We also used NGROK such that we can host the bot locally rather than using AWS or the like, however this is your choice whether you wish to use this.

https://github.com/howdyai/botkit-starter-ciscospark

https://nodejs.org/en/

https://ngrok.com

You will also need to change a few variables within the spark_bot.js file. In order to point the code to not just ngrok but for the Meraki Post reciever.

Change 'public_address' line 7: to the web address the bot will be hosted on, so your AWS deployment or your NGROK IP.

Change 'secret' line 434: to the secret used in your Meraki POST API settings.

Change 'validator' line 435: to the validator provided to you in the Meraki POST API settings.

Once this is done in the Meraki POST API settings you will also need to point your POST dumps to the URL being used by NGROK or AWS.

Once you have installed the relevant files and changed the variables, navigate to the VIPresence directory in the command line. Then execute the main application by doing

'node spark_bot.js' this will spin up the bot and should be running.
