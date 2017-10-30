# Pixels Camp Activities, Lisbon 2016

**If you're a participant to the PixelsCamp 2016 hackathon, 
take a look to the [Cisco API Challenge](https://github.com/ObjectIsAdvantag/hackathon-resources/tree/master/pixelscamp-lisbon) 
pass by the Cisco booth to discuss about APIs and innovation at Cisco,
and sign up at [DevNet - Cisco Developer Program](https://developer.cisco.com).**

> To give a try to the Pixels Camp Voice Machine
>
>    call **+351-308-801-061** 


## What's going on here ?

_Well, Cisco APIs at work: a REST API, a Tropo IVR and a [Cisco Spark Bot](https://github.com/ObjectIsAdvantag/sparkbot-webhook-samples/blob/master/examples/devnet/bot.js) playing all together_


This code takes a json event file and exposes it as a REST API with 3 ressources :
- [GET /](https://pixelscamp.herokuapp.com/) : returns all events, sorted by begin date
- [GET /next](https://pixelscamp.herokuapp.com/next?limit=10): shows upcoming activities
- [GET /now](https://pixelscamp.herokuapp.com/now) : returns events happening now

This API is consumed by a [Tropo Voice Machine](https://tropo.com) :
- Call +351-308-801-061 (a Tropo self-service IVR - Interactive Voice Response, featured: a Portuguish phone number) 
- have upcoming activities spoken to you (featured: Tropo Speech To Text),
- get details for an activity (featured: Tropo DTMF - touch tone),
- select an activity and receive details by SMS (featured: Tropo Outbound SMS)

Text your email to the +1-414-939-0591 (featured: Tropo bi-directional SMS support)
- have your [Cisco Spark](https://www.ciscospark.com) account added to the DevNet Support room (featured: Tropo Outbound requests to external Web APIs)
- note that [Tropo.eu](https://blog.tropo.eu/2016/09/20/tropo-in-europe-what-it-means/) has now launched, you can pick an SMS-enabled phone number in Europe,

Launch your favorite CiscoSpark client (Web, iOS, android, windows or mac)
- you're now part of a Spark Room, (featured: Cisco Spark Memverships API) 
- get real time info what's going on (featured: Cisco Spark Bot accounts & WebHook API)
- check if the session you're attending has begun: /next, /now


## Architecture

If you want to dig into the code:
- this project is leveraging SailsJS for the REST API aspects, the [ActivityController](api/controllers/ActivityController.js) holds main logic,
- the Voice Machine script custom logic starts here for [Voice interactions](tropo-IVR.js#L354) and there for [SMS interactions](tropo-IVR.js#L318).  
app.js

This API is deployed on [Heroku](https://pixelscamp.herokuapp.com/).


## Want to learn more

Check these webinars from Cisco Live Vegas 2016:
- DEVNET2002
- DEVNET3002


## MIT License

Feel free to use, reuse, extend, and contribute