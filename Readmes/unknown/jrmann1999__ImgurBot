# ImgurBot
This repository is an example using NodeJS to implement a listener for a Cisco Spark chat-bot.

It currently listens for mentions in rooms it is a member of, searching for Imgur animated gifs and returning them to the room. 
It can also be Private messaged and will subsequently reply.

As long as the script is running, it will remember the last 10 posts it has made, and can be deleted by telling it to delete (@bot delete).  Repeated operations will result in repeated deletions up to the last 10 posts.

You must have setup a Spark bot and obtained the Authorization token, additionally you will need to sign up for the Imgur API and have that token available.  Both will need to be specified as environment variables to run the bot successfully.

### Usage:

IMGURAPIKEY=KEY SPARKAPIKEY=KEY node server.js

No console output is expected with this.

### Debug:

I have utilized the debug module within NPM to setup custom debugging, the three named debug options are:

DEBUG=server<br>
This will give you debug information about the main server routine, generally giving you console logging indicating the steps the system is taking

DEBUG=imgur or spark<br>
This will highlight any errors or non HTTP status codes of 200, and what the data looks like when that occurs
