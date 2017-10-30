Cisco-Cert-Unofficial-API
===========

Verify Cisco Certs from within your app. This utility will post the Cisco Cert verification code and name in order to verify a Cisco Cert for that person. CCIE and CCDE validation not included for now, as that uses a different tool and I am working on that as well.

Dependencies
===========
Node.js, request, cheerio, express, underscore

Usage
===========
Send an HTTP GET request with query string that includes code to get back the Cisco Certification name and the user's name.
E.g.

http://yourhosthere/verify?code=XXXXXXXXXXXXXXXXXXXXXX

Returned JSON object will look like:
    {"cert":"CCNP","name":"John Smith", "description":"Brief Description of Cisco Cert"}

If the Cert Verification number is not found, an HTTP 404 will be returned.

If the Cisco server times out or has another error, an HTTP 500 will be returned.

If the Cisco cert name doesn't match the pre-defined names, a description of "No Description Exists" will be returned.

Installation
===========
###With NPM:


    $ npm install cisco-cert-api-server
https://npmjs.org/package/cisco-cert-api-server

###With Heroku:

This can easily be deloyed into a Heroku App via the Heroku toolkit.  Install via npm, then create and push the app via Git. Instructions here: https://devcenter.heroku.com/articles/nodejs

