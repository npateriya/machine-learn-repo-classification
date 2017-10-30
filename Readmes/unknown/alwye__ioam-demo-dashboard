# iOAM demo dashboard
Demo dashboard of iOAM to present at Cisco Live Las Vegas 2016

## Overview

![](images/main-screen.png)
Figure 1. Main screen

![](images/settings-panel.png)
Figure 2. Settings panel revealed

## Getting started
### For users
For those who simply want to run iOAM demo dashboard.

You need to clone this repo, navigate (```cd```) to the project's root folder and follow the steps below:

1. Install latest version of [Node.js](https://nodejs.org/en/download/) along with npm (*Node.js Package Manager*). The build was tested on Node.js v4.4.7 and npm 2.15.8.
2. Once Node.js is installed, run ```node start_server.js```. You will see the message ```Server running on 8080...```
3. Access web-GUI from a web-browser. If you are about to do it on the same computer you are running the GUI on, the address will be [http://localhost:8080](http://localhost:8080)
4. Hit *Settings* in the top toolbar, and fill out the Settings form. Once you hit *Save*, the new settings will apply and the page will refresh. You will see charts.
5. To stop web-server, press *Ctrl + C*.

**Note:** If you need to see static data, crack "start_server.js" open and change ```var dirName = __dirname + "/gui/build";``` to ```var dirName = __dirname + "/gui/build_static_07072016";```, then restart the web-server and access GUI.

### For developers
This section will help you to understand how change stuff in the iOAM dashboard.

This application is built with [AngularJS](https://angularjs.org), [Angular Material](https://material.angularjs.org/latest/) and [Chart.js](http://www.chartjs.org). It also employs a bunch of external libraries that the ones above implicitly depend on.

#### Prerequisites
Install [Bower](https://bower.io) and [Node.js](https://nodejs.org/en/download/).

1. Run ```npm install`` in project's root folder
2. Navigate (```cd```) to /gui/ and run ```npm install```, then ```bower install``` in there.

#### File structure

The file/folder structure goes like that:

```
/gui/ - root folder of all GUI sources and build code
	/build/ - destion for the built code
	/build_static_07072016/ - generated static API response will allow you to see the app without connection to the backend
	/node_modules/ - a set of vendor scripts for development
	/src/ - source code of frontend
		/app/ - Angular's pieces of code
			/controllers/ - each controller belongs here
			/modules/ - normally the only one module should be inside
			/services/ - home to Angular services
			/templates/ - HTML templates go here
		/assets/ - css/less/stuff
		/media/ - images/whatever
		/vendor/ - alias for "bower_components", contains a bunch vendor scripts for frontend
		index.html - must be clear enough
		main.js - list of JavaScript files to be concatenated and minified after compilation. Order matters.
	gulfile.js - the guy responsible for compilation of the whole thing
start_server.js - web server to serve static content from /gui/build/
```

#### Required commands
If you make changes to ```/gui/src/``` , make sure to run ```gulp clean && gulp build``` when you are ready to compile your frontend code. Once you run it, the output will end up in ```/gui/build/```

