# DevNet Events API

Public API for DevNet events featured at https://developer.cisco.com/site/devnet/events-contests/events/ 


## API Resources

Public ressources to search for events
- GET /events: all events ordered by date ASC
- GET /events/next: upcoming events (all events who begin after now, ie, do not comprise events happening now) 
- GET /events/current: events happening now


### POSTMAN

- collection: https://www.getpostman.com/collections/4a9ee2188317572e77a5
- documenter: https://documenter.getpostman.com/view/30210/devnet-events/2LNZ3k


## Architecture

The API is implemented on top of Sails.js, http://sailsjs.org/

An instance of the API is deployed on Heroku servers, and persists its data into PostgreSQL.
- https://devnet-events-api.herokuapp.com/api/v1/events/next


# Companions

## Cisco Spark bot

Launch Cisco Spark and invite CiscoDevNet@sparkbot.io to meet the bot.

The bot consumes the Events API to display current and upcoming events in Spark Rooms:
- bot source code: https://github.com/CiscoDevNet/node-sparkbot-samples/tree/master/examples/devnet 
- bot deployed at: https://devnet-events-sparkbot.herokuapp.com/


## Slack bot

Follow the [guidelines to run the bot](https://github.com/ObjectIsAdvantag/devnet-botkit-sample#how-to-run) on your local machine.

The bot consumes the Events API to display current and upcoming events in Slack channels
- bot source code: https://github.com/ObjectIsAdvantag/devnet-botkit-sample/blob/master/bot.js


# License

[MIT](LICENSE)

Feel free to use, reuse, extend, and contribute



