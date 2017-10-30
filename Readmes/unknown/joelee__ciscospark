# Cisco Spark API for Node.js

This is a super-simple lightweight JavaScript wrapper for [Cisco Spark API](https://developer.ciscospark.com/quick-reference.html). 

[![NPM](https://nodei.co/npm/node-ciscospark.png?downloads=true&downloadRank=true&stars=true)](https://nodei.co/npm/node-ciscospark/)

[![npm](https://img.shields.io/npm/v/node-ciscospark.svg)](https://www.npmjs.com/package/node-ciscospark) [![Build Status](https://travis-ci.org/joelee/ciscospark.svg?branch=master)](https://travis-ci.org/joelee/ciscospark) [![codecov](https://codecov.io/gh/joelee/ciscospark/branch/master/graph/badge.svg)](https://codecov.io/gh/joelee/ciscospark) [![Documentation](https://doc.esdoc.org/github.com/joelee/ciscospark/badge.svg)](https://doc.esdoc.org/github.com/joelee/ciscospark/) [![dependencies Status](https://david-dm.org/joelee/ciscospark/status.svg)](https://david-dm.org/joelee/ciscospark) [![Known Vulnerabilities](https://snyk.io/test/github/joelee/ciscospark/badge.svg)](https://snyk.io/test/github/joelee/ciscospark) [![](https://img.shields.io/github/issues-raw/joelee/ciscospark.svg)](https://github.com/joelee/ciscospark/issues) [![Stories in Ready](https://badge.waffle.io/joelee/ciscospark.png?label=ready&title=Ready)](https://waffle.io/joelee/ciscospark)  [![JavaScript Style Guide](https://img.shields.io/badge/code_style-standard-brightgreen.svg)](https://standardjs.com)

## Installation

```bash
$ npm install node-ciscospark
```

## Quick Start

### API

```javascript
var CiscoSpark = require('node-ciscospark');

var spark = new CiscoSpark(accessToken);
```

> `accessToken` can also be set in the `CISCOSPARK_ACCESS_TOKEN` environment variable.

### [Create a Message](https://developer.ciscospark.com/endpoint-messages-post.html)

```javascript
// spark.messages.create(parameters, callback)
spark.messages.create({
  roomId: sparkChatRoomId,
  text: 'This is a test'
}, function (err, result) {
  if (err) console.error(err);
  console.log(result);
});
```

### [List Rooms](https://developer.ciscospark.com/endpoint-rooms-get.html)

```javascript
// spark.rooms.list(parameters, callback)
spark.rooms.list({
  teamId: sparkTeamId
}, function (err, result, response) {
  if (err) console.error(err);
  console.log(result);
  console.log('Response Codes:', response.statusCode);
});
```

## Supported [Spark API](https://developer.ciscospark.com/quick-reference.html)

### [Messages](https://developer.ciscospark.com/resource-messages.html)

API methods | Usage
----------- | -----
[List Messages](https://developer.ciscospark.com/endpoint-messages-get.html) | `spark.messages.list(parameters, callback)`
[Create a Message](https://developer.ciscospark.com/endpoint-messages-post.html) | `spark.messages.create(parameters, callback)`
Create a Message to a Room | `spark.messages.createToRoom(roomId, markdownMessage, callback)`
Create a Message to a Person | `spark.messages.createToPersonId(personId, markdownMessage, callback)`
Create a Message to a Person's email | `spark.messages.createToPersonEmail(email, markdownMessage, callback)`
[Get Message Details](https://developer.ciscospark.com/endpoint-messages-messageId-get.html) | `spark.messages.get(messageId, callback)`
[Delete a Message](https://developer.ciscospark.com/endpoint-messages-messageId-delete.html) | `spark.messages.delete(messageId, callback)`

### [People](https://developer.ciscospark.com/resource-people.html)

API methods | Usage
----------- | -----
[List People](https://developer.ciscospark.com/endpoint-people-get.html) | `spark.people.list(parameters, callback)`
[Create a Person](https://developer.ciscospark.com/endpoint-people-post.html) | `spark.people.create(parameters, callback)`
[Get Person Details](https://developer.ciscospark.com/endpoint-people-personId-get.html) | `spark.people.get(personId, callback)`
[Update a Person](https://developer.ciscospark.com/endpoint-people-personId-put.html) | `spark.people.update(personId, parameters, callback)`
[Delete a Person](https://developer.ciscospark.com/endpoint-people-personId-delete.html) | `spark.people.delete(personId, callback)`
[Get my own details](https://developer.ciscospark.com/endpoint-people-me-get.html) | `spark.people.me(callback)`

### [Rooms](https://developer.ciscospark.com/resource-rooms.html)

API methods | Usage
----------- | -----
[List Rooms](https://developer.ciscospark.com/endpoint-rooms-get.html) | `spark.rooms.list(parameters, callback)`
[Create a Room](https://developer.ciscospark.com/endpoint-rooms-post.html) | `spark.rooms.create(parameters, callback)`
[Get Room Details](https://developer.ciscospark.com/endpoint-rooms-roomId-get.html) | `spark.rooms.get(roomId, callback)`
[Update a Room](https://developer.ciscospark.com/endpoint-rooms-roomId-put.html) | `spark.rooms.update(roomId, parameters, callback)`
[Delete a Room](https://developer.ciscospark.com/endpoint-rooms-roomId-delete.html) | `spark.rooms.delete(roomId, callback)`

### [Memberships](https://developer.ciscospark.com/resource-memberships.html)

API methods | Usage
----------- | -----
[List Memberships](https://developer.ciscospark.com/endpoint-memberships-get.html) | `spark.memberships.list(parameters, callback)`
[Create a Membership](https://developer.ciscospark.com/endpoint-memberships-post.html) | `spark.memberships.create(parameters, callback)`
[Get Membership Details](https://developer.ciscospark.com/endpoint-memberships-membershipId-get.html) | `spark.memberships.get(membershipId, callback)`
[Update a Membership](https://developer.ciscospark.com/endpoint-memberships-membershipId-put.html) | `spark.memberships.update(membershipId, parameters, callback)`
[Delete a Membership](https://developer.ciscospark.com/endpoint-memberships-membershipId-delete.html) | `spark.memberships.delete(membershipId, callback)`

### [Teams](https://developer.ciscospark.com/resource-teams.html)

API methods | Usage
----------- | -----
[List Teams](https://developer.ciscospark.com/endpoint-teams-get.html) | `spark.teams.list(parameters, callback)`
[Create a Team](https://developer.ciscospark.com/endpoint-teams-post.html) | `spark.teams.create(parameters, callback)`
[Get Team Details](https://developer.ciscospark.com/endpoint-teams-teamId-get.html) | `spark.teams.get(teamId, callback)`
[Update a Team](https://developer.ciscospark.com/endpoint-teams-teamId-put.html) | `spark.teams.update(teamId, parameters, callback)`
[Delete a Team](https://developer.ciscospark.com/endpoint-teams-teamId-delete.html) | `spark.teams.delete(teamId, callback)`

### [Team Memberships](https://developer.ciscospark.com/resource-team-memberships.html)

API methods | Usage
----------- | -----
[List Team Memberships](https://developer.ciscospark.com/endpoint-teammemberships-get.html) | `spark.teamMemberships.list(parameters, callback)`
[Create a Team Membership](https://developer.ciscospark.com/endpoint-teammemberships-post.html) | `spark.teamMemberships.create(parameters, callback)`
[Get Team Membership Details](https://developer.ciscospark.com/endpoint-teammemberships-membershipId-get.html) | `spark.teamMemberships.get(membershipId, callback)`
[Update a Team Membership](https://developer.ciscospark.com/endpoint-teammemberships-membershipId-put.html) | `spark.teamMemberships.update(membershipId, parameters, callback)`
[Delete a Team Membership](https://developer.ciscospark.com/endpoint-teammemberships-membershipId-delete.html) | `spark.teamMemberships.delete(membershipId, callback)`

### [Webhooks](https://developer.ciscospark.com/resource-webhooks.html)

API methods | Usage
----------- | -----
[List Webhooks](https://developer.ciscospark.com/endpoint-webhooks-get.html) | `spark.webhooks.list(parameters, callback)`
[Create a Webhook](https://developer.ciscospark.com/endpoint-webhooks-post.html) | `spark.webhooks.create(parameters, callback)`
[Get Webhook Details](https://developer.ciscospark.com/endpoint-webhooks-webhookId-get.html) | `spark.webhooks.get(webhookId, callback)`
[Update a Webhook](https://developer.ciscospark.com/endpoint-webhooks-webhookId-put.html) | `spark.webhooks.update(webhookId, parameters, callback)`
[Delete a Webhook](https://developer.ciscospark.com/endpoint-webhooks-webhookId-delete.html) | `spark.webhooks.delete(webhookId, callback)`

## Callbacks

The callbacks above return `error`, `resultBody` and `httpResponse`:

```javascript
  function (error, resultBody, httpResponse) { }
```

where:

* `error` is an Error object, or null
* `resultBody` is the JSON response
* `httpResponse` is a [http.ServerResponse](https://nodejs.org/api/http.html#http_class_http_serverresponse) object


## Code Documentation

See: https://doc.esdoc.org/github.com/joelee/ciscospark/

[![Documentation](https://doc.esdoc.org/github.com/joelee/ciscospark/badge.svg)](https://doc.esdoc.org/github.com/joelee/ciscospark/)


## License

The MIT License (MIT)

Copyright (C) 2017 Joseph Lee

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

---

## Contributor(s)

Joseph Lee [![Twitter Follow](https://img.shields.io/twitter/follow/joe_lee.svg?style=social&label=Follow)](https://twitter.com/joe_lee)
