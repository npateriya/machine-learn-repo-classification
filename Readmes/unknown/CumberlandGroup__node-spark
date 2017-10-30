# node-spark
**v2.2.0**

Cisco Spark API Library for Node JS based on a [Swagger](http://swagger.io/specification/) definition specification.

### Features

- [Rate limiting headers](https://developer.ciscospark.com/blog/blog-details-8193.html) inspected to adjust request rates based on Cisco Spark API. These are automatically re-queued and sent after the `retry-after` timer expires.
- [Pagination](https://developer.ciscospark.com/pagination.html) automatically invoked when requesting max results greater than the API max.
- Promises comply with [A+ standards.](https://promisesaplus.com/)
- Simple FIFO API queueing mechanism with adjustable delay.
- [Webhook](https://github.com/cumberlandgroup/node-spark-webhook) submodule.

### Project Setup/Install

```bash
mkdir myproject
cd myproject
npm init
npm install --save node-spark
touch index.js
```

**Example index.js:**
```js
var CiscoSpark = require('node-spark');

var options = {
  token:'OWQwOGEzMDgtZDYyTOKENjQwLWI2MTTOKEN5YmMzYTI1TOKENzVhNGNjTOKENDgx'
};

var spark = new CiscoSpark(options);

spark.connect()
  .then(client => client.rooms.getRooms())
  .then(res => console.log(res))
  .catch(err => console.log(err.message));
```

### Initialization / Config

The constructor accepts an `options` object. The only required object key property required is the `token`. Below shows the optional properties to override defaults for non-required options.

```js
var options = {
  token: 'OWQwOGEzMDgtZDYyTOKENjQwLWI2MTTOKEN5YmMzYTI1TOKENzVhNGNjTOKENDgx',
  swagger: 'https://raw.githubusercontent.com/CumberlandGroup/swagger-cisco-spark/master/cisco_spark_v1.json',
  delay: 600
};
```

**Options Object:**

- `token` : Spark API Token
- `swagger` : File path or URL to over-ride the internal swagger file definition
- `delay` : Delay in ms between outbound requests.

_**Note:** While this library will respect the Rate Limiting Headers, the outbound FIFO queue will help average out requests to minimize hitting the Rate Limiter. Once the API Rate Limiter is hit, the retry times returned are often in excess of 60 seconds which will cause significant delay in any real-time API interaction._

### Calling the Spark API

The `Spark.connect()` method returns a spark client object promise that includes the following methods and events.

`client.<resource>.<method>(<query>)`

The resource and method are defined in the Swagger definition. If not specified, an internal swagger definition file is used. For more information on the resource/method/query, reference this github [repository.](https://github.com/CumberlandGroup/swagger-cisco-spark)

### Events

`spark.on('<event>', function(<event params>) { // process event });`

Events Types:

- `request` - Emitted with each API request. The callback executed with the arguments:
  - `url` : requested URL
  - `headersObj` : object containing the headers of request
  - `bodyObj` :  object containing the contents body of the request
- `rate-limited` - Emitted when a response is returned that rate limit is hit. The callback executed with the arguments:
  - `retryAfter` : seconds that Spark API is requesting to wait before resending this request
  - `url` : requested URL
  - `headerObj` :  object containing the headers of request
  - `bodyObj` : object containing the contents body of the request
- `queued` - callback executed with the arguments:
  - `queueDepth` : size of queue
  - `url` : requested URL
  - `headersObj` :  object containing the headers of request
  - `bodyObj` : object containing the contents body of the request

### Webhook submodule

The webhook submodule can optionally be implemented directly from node-spark without a direct dependency on the node-spark-webhook package. For details on usage, refer to the submodule's [README.md](https://github.com/cumberlandgroup/node-spark-webhook).

**Project Setup:**
```bash
mkdir myproject
cd myproject
npm init
npm install --save node-spark
npm install --save express
npm install --save body-parser
touch index.js
```

**Example index.js: (embedded into an express.js app)**
```js
"use strict";

var Webhook = require('node-spark/webhook');
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
