# digitalDelegation
Cisco Spark integration with SigFox button

## What is it?
* Bt.tn / Sigfox button as trigger
* Cisco Spark as Space for human and data
* => When the Button is pushed Cisco Spark space is created and people are added

## Based on
* [sparky](https://github.com/flint-bot/sparky) (nodejs + express)
* [bt.tn](https://my.bt.tn) (can also provide simulator)

## HowTo
* Clone localy

> git clone https://github.com/guillain/digitalDelegation.git

* Go into the folder

> cd digitalDelegation

* Install dependencies

> npm install

* Config your [bt.tn](https://my.bt.tn/home) with:
* * action: HTTP
* * HTTP URL: your url
* * HTTP Method: POST
* * Arguments: application/json
```
{"ID":"<ID>","EID":"<EID>","DEVICEID":"<DEVICEID>","counter":"<COUNTER>","date":"<DATE>", "time": "<TIME>","name":"<NAME>","user":"<USER>","location":"<LOCATION>","emailaddress":"<EMAILADDRESS>"}
```
* Config your app with your [spark bot](https://developer.ciscospark.com/apps.html)

> cp config.js.default config.js
> vi config.js

* Run the application, two configuration availables

* 1/ For the dev, node is used

> ./app manual

* 2/ For the prod, pm2 is used (install also this dependency)

> ./app [start|stop|restart|show|staus|log]

* Push your button!!!
* Check what happen on your Cisco Spark app


Have fun
