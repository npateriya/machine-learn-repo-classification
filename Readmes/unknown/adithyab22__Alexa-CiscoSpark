
[![Build Status](https://travis-ci.org/adithyab22/Alexa-CiscoSpark.svg?branch=master)](https://travis-ci.org/adithyab22/Alexa-CiscoSpark)

# Amazon Alexa-Cisco Spark
This is an amazon alexa skill which leverages voice experience of Amazon Alexa to perform actions on Cisco Spark.
Using this skill, you can register users, add contacts, create teams, send messages to contacts/teams and so on - all these - just using voice. 

# Usage  

*Once installed, this skill requires linking account with Cisco Spark.

## Sample utterances
The invocation name for this skill is "Sparky"
Just say: "Alexa, Ask Sparky to send message to John Doe"

These are a few things you can ask Sparky to do:

* Create team [Teamname]
* Create room [Roomname]
* Send message to [Person|Roomname|Teamname]
* Add member to [Teamname|Roomname]

## How it works?
1) Alexa uses invocation name "Sparky" to invoke this skill. 
2) Skill calls a web service running on AWS Lambda.
3) Lambda web service requests Cisco Spark API to perform the command.
4) Based on API response, Lambda returns the speech output to user.
