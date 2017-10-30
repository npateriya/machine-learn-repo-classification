pyCiscoSparkAPI
================
A simple Python SDK interface for the Spark for Developers APIs.

Overview
---------
This SDK exposes the Spark for Developers APIs in simple Python class wrappers.

For more information on Sprark for Developers see http://developer.ciscospark.com

Installation
---------

#### Manual

```
git clone https://github.com/shenning00/pyCiscoSparkAPI.git
cd pyCiscoSparkAPI
python ./setup.py build
python ./setup.py install
```

Usage
----
Here are a few simple examples.

__Note__: This SDK uses the simple Personal Access Token. See here (see https://developer.ciscospark.com/getting-started.html) for more info on how to get the token.

###Getting Started####

Simple - import the SDK, create a client using a personal access token.
Get all of my rooms. Iterate through the rooms and print out the titles.

```python
from pyCiscoSparkAPI import SparkClient

token="abcdef....456790"
client = SparkClient(token)
rooms = client.rooms.getRooms()
for room in rooms:
    print(room.title())
```

###People###

Get all of the people with displayName that starts with some text.

```python
client = SparkClient(token)
# get all of the cool peopled named 'Scott'
people = client.people.getPeople(displayName="Scott")
for person in people:
    print(person.displayName())
```

###Memberships###

Get all of the memberships for a particular room by roomID.

```python
client = SparkClient(token)
memberships = client.memberships.getMembershipsByRoom(roomId)
    print(member.personEmail())
```

###Rooms###

Get a room by roomId.

```python
client = SparkClient(token)
rooms = client.rooms.getMessagesByRoom(roomId)
for room in rooms:
    print(room.title())
```

Create a new room.

```python
client = SparkClient(token)
room  = client.rooms.newRoom("3rd Quarter Projections")
print(room.title())
print(room.id())
```


###Messages###

Get all of the message in a room by roomId.

```python
client = SparkClient(token)
messages = client.messages.getMessagesByRoom(roomId)
for message in messages:
    print(message.text())
```

Send a message to a room.
```python
client = SparkClient(token)
message = client.messages.sendMessageToRoom(roomId,"Quarter projects are on target.")
print(message.id())
```



Seperate Class Documentation
----
