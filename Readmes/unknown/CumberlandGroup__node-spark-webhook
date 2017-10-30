# node-spark-webhook

Cisco Spark API Web Hook Utility Library for Node JS.

This library parses an incoming Cisco Spark API web hook and emits events based on the content. This library does not interface with the Spark API and is intended to simplify functions around web hook parsing and authenticating [HMAC ](https://developer.ciscospark.com/webhooks-explained.html#auth) secured web hooks only.

**Example: (embedded into an express.js app)**
```js
"use strict";

var Webhook = require('node-spark-webhook');
var express = require('express');
var bodyParser = require('body-parser');
var path = require('path');

var webhook = new Webhook();

// add events
webhook.on('request', function(hook) {
  console.log('%s.%s web hook received', hook.resource, hook.event)
});

var app = express();
app.use(bodyParser.json());

// add route for path that which is listening for web hooks
app.post('/spark', webhook.listen());

// start express server
var server = app.listen('3000', function () {
  console.log('Listening on port %s', '3000');
});
```

#### Install:

```bash
npm install --save node-spark-webhook
```

### Initialization / Config

The constructor accepts an optional `options` object.

```js
var Webhook = require('node-spark-webhook');

var options = {
  secret: 'My HMAC Secret',
  reqObjProp: 'params'
};

var webhook = new Webhook(options);
```

#### Options Object

- `secret` : If specified, the incoming web hook will be [authenticated](https://developer.ciscospark.com/webhooks-explained.html#auth) before being processed through the `webhook.listen()` function.
- `reqObjProp` : Property of request object where the string or JSON object of the incoming body is parsed to. If not specified, defaults to "body". Some Connect based libraries (Express/Restify/etc) alter this location based on their defaults and which extensions you have enabled. Typically, if not "body", "params" is the request property where the body is passed.

### Reference

`webhook.listen()`

Returns a function with properties `req` and `res`. When embedded in a connect based web app such as express.js, allows easy processing and validation (including HMAC if secret is specified) of an incoming web hook. Once validated, emits events based on the web hook's `resource` value.

`webhook.auth(payload, signature, secret [, callback]))`

Should you simply need to authenticate a web hook, you can access the `auth` function directly.

- `payload` : request body in JSON or string format
- `signature` : HMAC signature retrieved from the 'x-spark-signature' header property
- `secret` : The secret used when creating the web hook from the Spark API
- `callback` : (optional) If specified executes a callback with properties `(err, body)`. If not specified, returns a promise fulfilled with the `body`, or rejected with the `error`.

### Events

The library will emit events based on the value of the `resource` property of the incoming webhook. For example, to capture a new room membership:

```js
webhook.on('memberships', function(event, membershipObj) {
  if(event == 'created') {
    console.log('%s added to room', membershipObj.personDisplayName);
  }
});
```

*Note: The web hooks you receive depend on how you originally created the web hook from the Spark API.*

Additionally, a more generic event named "request" is triggered on every validated web hook. This event returns a single property that contains the request `body` as an object. For example:

```js
webhook.on('request', function(hook) {
  console.log('%s.%s web hook received', hook.resource, hook.event)
});
```

### License

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
