# cmxdemo
Sample application accessing a Cisco CMX Mobility Services server

## Server Installation

The server is a node.js application.  To run, install node.js and then invoke the
script bin/run. The default port is 3000.  Edit config.json to change it.

## Local Methods
/local/config - Get or update server configuration

The server supports one local method, /local/config, that can be used either
with GET or POST.  As a GET method, it returns JSON with the server's virtualize
and cmxServer settings.  As a POST method, it accepts a JSON argument specifying
values for any or all of virtualize, cmxServer, username, and password.  The
POST method returns the same JSON as the GET method, with the addition of an
authToken value if either username or password is specified.  The authToken can
be used for the CMX API pass-through methods.

## Pass-Through Methods
/api/xxxx?token=xxx - invoke CMX API

All methods beginning /api are passed to the CMX server, which provides the
response.  See CMX documentation for available methods and response.

All pass-through methods must include a token argument specifying an authorization
token obtained from a previous call to the /local/config POST method.

## Server Virtualization
The server passes all requests to a remote CMX server.  If a remote CMX
server is unavailable, or if you want predictable responses to a subset of
the server methods, activate server virtualization support.  This can be done in
the UI by setting the "Virtualize" checkbox on the login form, or with the API by
setting virtualize=true in a /local/config POST call.  Update config.json to
change the default for virtualize.

The following methods are available when virtualization is active:

* /api/location/v1/clients/count       - number of clients
* /api/location/v1/clients             - detailed info on all clients
* /api/location/v1/clients/xxx         - detailed info on the client with MAC address xxx
* /api/location/v1/history/clients/xxx - location history for the client with MAC address xxx
* /api/config/v1/maps/imagesource/xxx  - image of map xxx

The server obtains data for virtualized method response from JSON files in the
testData directory.

## UI

The cmxdemo UI is a jQuery application served on the same port as the API.  To
invoke it, use any browser (Chrome recommended) to navigate to http://localhost:3000.

After login (default credentials are learning/learning), the UI displays a map showing
all clients in the room where the first client returned by /api/location/v1/clients is
located.  Click on a client name to show the recent history of that client's movements.
Press the Show List of All Users button to show a table of all users known to CMX.  
Click on a user name in the list to show that user's recent location history.  Click 
on a room name in the list to show all users currently in the room.
