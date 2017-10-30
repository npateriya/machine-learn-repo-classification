# Tropo-CallCenter-SparkLogger
Tropo dials a call center
Please askMessage untill an agent confirms or a timeout happens
Then logs the results to a Cisco Spark Room.

## Use case examples:
* Call Center Testing
* Call Center Load Testing
* Troubleshooing call center hold issues to capture timestamp info.


# Getting Started
## Tropo Setup
* Setup Tropo account
* Contact Support to enable outbound dialing support.
* Set your Spark Authorization token and Room ID in the script, and a few variables are clearly commented.

## Spark Setup
* Create a Cisco Spark Account
* Go to your Spark Room and add a new integration. Grab your Authorization token or create a new "Bot Integration"
* Grab the ID for the Bot and use that as the token
* Create a Spark Room for the Logger to post into.
* Go to web.ciscospark.com and open the logger Room you created. Grab the UID of the room like this https://web.ciscospark.com/rooms/1a758e30-xxxx-11e6-xxxx-696d2bd64b2e/chat
* That ID is your Room ID in the script.

## Scheduling it for free using IFTTT
* Signup at www.ifttt.com
* Create a new recipe using Date Time recurring and the Maker channel.
* Setup requests at 00,15,30,45 minute intervals.
* In the Maker channel add the web request for your Tropo API call like this
https://api.tropo.com/1.0/sessions?action=create&token=MyAuthToken&dialnumber=1231231234&dial247=true
