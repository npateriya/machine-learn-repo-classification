# TalosSays
Command-line JS (uses [PhantomJS](https://phantomjs.org)) to check email Reputation on Cisco Talos.  It uses the Talos reputation search URL in the form:

[https://www.talosintelligence.com/reputation_center/lookup?search=myhost.example.com]

Edit TalosSays.js and change *"var hostname = myhost.example.com"* to your web/email hostname.  Running is as simple as:

`phantomjs TalosSays.js`

This utility outputs just four lines of text, extracted from the Talos reputation result web page:

```
Hostname:              myhost.example.com
Forward/Reverse DNS:   Yes
Blacklisted:           No
Email Reputation:      Neutral
```
