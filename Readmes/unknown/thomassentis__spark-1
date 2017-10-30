# Cisco Spark in Pzartech
* [Purpose](#purpose)
* [Design Decisions](#design)
  * [Lack of Frameworks/Libraries](#frameworks)
  * [Architechture](#architecture)
* [Installation and Usage](#installation-and-usage)
  * [How to Host the App](#hosting)
  * [How to Use the App](#usage)
  * [How to Run Tests](#tests)
  * [Supported Browsers](#browsers)
  * [Spark Integration Process](#integration)


<a name="purpose"></a>
## Purpose
This demo app is designed to provide a clear understanding of the process of integrating Cisco's Spark JavaScript Source Developer Kit into your app for calling and receiving calls.

<a name="design"></a>
## Design Decisions

<a name="frameworks"></a>
### Lack of Frameworks/Libraries

We decided to minimize the use of any major JavaScript frameworks/libraries such as React or Angular so that the example is clear to all JavaScript developers. We are using JQuery to simplify some calls, Express for end points, and Browserify to bundle the javascript.

<a name="architecture"></a>
### Architecture

The demo app is broken into two parts. The login page begins the authentication process by authenticating with spark. The index page finishes the authentication process by registering our app with spark. Calls are made and answered from the index page via templates located in index.html and corresponding JavaScript files located in js/[templateName]Template.js). The Spark service is a light wrapper around the Spark SDK.

<a name="installation-and-usage"></a>
## Installation and Usage

<a name="hosting"></a>
### How to Host the App

* [Install NodeJS](https://nodejs.org/en/download/)
* Clone the repo: `$ git clone https://github.com/ciscospark/spark-js-sdk-example.git`
* `$ cd spark-js-sdk-example`
* `$ npm install`
* To launch the example app, you'll need Spark Integration credentials. These can be obtained by [creating a Spark account](https://web.ciscospark.com/signin) and following [this guide](https://developer.ciscospark.com/authentication.html).
    - Under Scope, enable spark:all
    - Under Redirect_URI, enter "http://localhost:8000/index.html"
    - Place your client id and your client secret in the .env.json file in the root directory of the project.
* You can also override the following environment variables by adding them to the .env.json file:
    - CISCOSPARK_REDIRECT_URI (needs to be changed if hosting anywhere aside from localhost:8000)
        - NOTE: You still must include /index.html at the end of your redirect URI in order for the app to work
    - CISCOSPARK_LOG_LEVEL (defaults to 'info'. Alternative log levels are: silent, error, warn, log, info, debug, and trace)
* `$ npm start`
* In your browser of choice, go to `http://localhost:8000` (Note: Internet Explorer and Safari are not supported)

<a name="usage"></a>
### How to Use the App

* You must have a [Spark account](https://web.ciscospark.com/signin) to use the app
* On the login screen, click the login button. This forwards you to Spark's OAuth login page if you aren't already authenticated. If you've logged in already, skip the next two steps.
* Enter your Spark email address and click 'Next'
* Enter your password for Spark and click 'Sign In'.
* You may be prompted to give the app permissions to use your Spark account. Click 'Accept'.
* You will be directed to the app's Index page.
* In the email field, type an email address of another existing Spark user and click one of the call buttons.
* Once a call is active you can toggle your outgoing audio & video with the buttons below your video feed
* You can then hangup with the hangup button.
* If you receive an incoming call at any time an overlay will be presented
    - Click one of the blue buttons to accept it with video or video and audio, or the red reject button to reject it.
* At any time, click the "sign out" button to log out of the app.

<a name="tests"></a>
### How to Run the Tests

#### Unit Tests
* From within the app's root directory: `$ npm run unit-test`

##### Integration Tests
* Java SDK is required
* `$ $(npm bin)/webdriver-manager update`

* You will need to add test user environment variables by adding them to the .env.json file:
``` .env.json
{
  "sparkUsername": "Insert Spark User E-mail inside these quotes",
  "sparkPassword": "Insert Spark User Password inside these quotes",
  "calleeUsername": "Insert Second Spark User E-mail inside these quotes",
  "calleePassword": "Insert Second Spark User Password inside these quotes",
  "fakeUser": "test"
}
```

* From within the app's root directory: `$ npm run integration-test`

<a name="browsers"></a>
### Supported Browsers

* The app is intended for Google Chrome or Mozilla Firefox. Internet Explorer & Safari are not supported by the Spark SDK.
* **Chrome**: Optional - When developing without a webcam, start Chrome with the ```--use-fake-device-for-media-stream``` and ```--use-fake-ui-for-media-stream``` flags to simulate a webcam.

* **Firefox**: Required - If the app is not secured with SSL, microphone and camera permissions will be denied. This can be overriden by changing the ```media.navigator.permission.disabled``` flag to true within the ```about:config``` of Firefox.

### Further Documentation
Also see [Cisco's Documentation](https://ciscospark.github.io/spark-js-sdk/) regarding the SDK.

<a name="integration"></a>
### Spark Integration Process

Before integrating Spark's JavaScript SDK into your app, there are a few major processes that you should familiarize yourself with.

#### Authorization
 * Before utilizing Spark's functionality, you must first authorize the user in order to have them receive a valid auth token. This can be done by calling ```spark.authorize()```.
   * If OAuth credentials have not been generated, this will redirect the user to a Cisco Spark OAuth login page. Upon completion, this will redirect to the provided CISCOSPARK_REDIRECT_URI.  
   * If OAuth credentials have already been generated, this will resolve a promise and continue code execution.

#### Registration
 * After being successfully authorized, the user's device needs to be registered with Spark. This needs to happen after Spark is completely authenticated, which can be done by assigning an event listener to Spark's authenticated flag:  ```spark.on('change:isAuthenticated', () => { spark.phone.register() });```
 * **In our demo app, we wrap the authentication function within a promise so that we can know precisely when Spark is finished authenticating and registering the device.**

#### Events
 There are multiple events that trigger upon the use of various functionality in the Spark SDK. Events occur on different objects within the Spark SDK (spark.phone, spark.call, spark, etc...). This demo application makes use of the following events:
 * ```spark.on('change:isAuthenticating', () => {})```
 * ```spark.on('change:isAuthenticated', () => {})```
 * ```spark.phone.on('call:incoming', () => {})```
 * ```spark.call.on('remoteMediaStream:change', () => {})```
 * ```spark.call.on('localMediaStream:change', () => {})```
 * ```spark.call.on('change:receivingVideo', () => {})```
 * ```spark.call.on('change:receivingAudio', () => {})```
 * ```spark.call.on('remoteVideoMuted:change', () => {})```
 * ```spark.call.on('remoteAudioMuted:change', () => {})```
 * ```spark.call.on('disconnected', () => {})```
 * ```spark.call.on('error', () => {})```

#### Calls
After the user is registered with Spark they can receive or create a call. You will also need to pass in certain parameters (like constraints) in order to dictate the nature of the call (audio, video, audio-video). When calling a user, that user's email must be specified in the ```spark.phone.dial``` function call. In order to make or receive a call, you will need to supply an ```offerOptions``` and a ```constraints``` object.
```
  return spark.phone.dial(user, {
    offerOptions: {
      offerToReceiveAudio: true,
      offerToReceiveVideo: true
    },
    constraints: {
      audio: true,
      video: true,
      fake: false
    }
  });
```
##### Listening
  * In order to listen for any incoming calls, you need to register an event for the Spark Phone's incoming call: ```spark.phone.on('call:incoming', (call) => { // Your code here });```. To let the caller know that their call to the user has become visible, you can use ```call.acknowledge()```.
  * You can answer a call using ```call.answer(offerOptions, constraints)``` on the call returned from the 'incoming' event.
##### Calling a User
  * In order to make a call you will need to supply:
    * the user's email
    * an ```offerOptions``` object
    * a ```constraints``` object
  ```
  return spark.phone.dial(userEmail, {
    offerOptions,
    constraints
  });
  ```
  * spark.phone.dial returns a Call object which you can then set up event listeners in order to handle the call functionality.
  ```
  Call.on('remoteMediaStream:change', () => {
   // Update video source, toggle displays, etc...
  });
  ```
  * Once a call has connected you can set the src attribute of your video tag to ```call.remoteMediaStreamUrl``` or ```call.localMediaStreamUrl``` corresponding to incoming stream and outgoing stream respectively.

##### Active-Call Functionality
Once in a call there are other various functions that you can execute such as toggling the receiving / sending audio and video. See ```activeCallTemplate.js``` for our usage of this functionality.

##### Call Feedback
In order to allow users to send feedback about their experience with a given call, you can supply them with a page or fields to capture that information and execute the following:
 ```
 call.sendFeedback({
  // String
  userComments: comments,
  // Boolean
  includeLogs: false,
  // Integer 1 - 5
  userRating: 3
 });
 ```

#### Useful snippets

* Log all events on a call: `call.on('all', (name) => console.log('currentCall event: ', name));`
