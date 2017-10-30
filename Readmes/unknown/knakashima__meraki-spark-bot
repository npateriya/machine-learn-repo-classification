# Meraki Spark Bot

A sample Cisco Spark Bot using Spark, Meraki Dashboard, and Tropo API's.

## Parts list
  * Amazon Web Services (AWS) API Gateway & Lambda
    * See Cory Guynn's nice writeup here: https://github.com/dexterlabora/cmxreceiver-lambda-inline
  * Cisco Spark API
    * developer.ciscospark.com
  * Meraki Dashboard API
    * developers.meraki.com
  * Tropo API
    * www.tropo.com/docs

## Description
The Spark Bot is tied to an AWS Lambda function (Python) that performs the fetching of information triggered by Spark messages.

The Spark Bot posts responses back to the originating Spark room.

When triggered, the Webhook will send an HTTP POST to a configured targetUrl.

An AWS API Gateway endpoint is used in this example

The Lambda function extracts the Spark message id from the Webhook triggered POST.

Next, the Lambda function sends an API call back to Spark to get the text of the message (only the message id itself is received initially). The text of the message is matched to "commands" mapped to different functions. The functions execute one or more API calls to the Meraki Dashboard for information.

Each function gathers the requested information, assembles the response into a table, and sends an API call to Spark to post the response in the room that originated the request.


*Sample event json passed by the Webhook*

```
event = {
    "status": "active",
    "resource": "messages",
    "name": “my webhook",
    "created": "2016-11-11T03:24:14.246Z",
    "appId": “<app id>“,
    "id": “<event id>“,
    "filter": "mentionedPeople=<bot id used in this example>”,
    "orgId": "",
    "createdBy": "",
    "targetUrl": “<URL such as an AWS API Gateway…>“,
    "ownedBy": "creator",
    "actorId": “<person id that triggered the webhook>“,
    "data": {
        "roomType": "group",
        "created": "2016-12-19T20:16:53.470Z",
        "personId": “<person id that triggered the webhook>“,
        "personEmail": "<user>@cisco.com",
        "mentionedPeople": [
            “<bot id used in this example>“
        ],
        "roomId": “<room id>“,
        "id": “<message id>“
    },
    "event": "created"
}
```

