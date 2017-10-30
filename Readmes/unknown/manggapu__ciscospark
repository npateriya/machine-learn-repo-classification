# Cisco Spark Python SDK
Welcome to the Python SDK for Cisco Spark API

Build notes:
- Tested on Python 2.7
- Requires pip install requests
- Requires pip install requests-toolbelt

Features:
- Maps to all functions and parameters of the 7 main APIs 
- These include people, rooms, messages, teams, memberships, team memberships, webhooks

sample.py shows you how to start using ciscospark today!

------------------------------------------------------------

### People
- people.get_details (token, personId)  
- people.get_me (token)  
- people.list (token, email=optional, displayName=optional, max=optional)  
- see https://developer.ciscospark.com/resource-people.html 

### Rooms
- rooms.create (token, title, teamId=optional)  
- rooms.delete (token, roomId)  
- rooms.get_details (token, roomId)  
- rooms.list (token, teamId=optional, max=optional, type=optional)  
- rooms.update (token, roomId, title)  
- see https://developer.ciscospark.com/resource-rooms.html  

### Memberships
- memberships.create (token, roomId, personId=optional, personEmail=optional, isModerator=optional)  
- memberships.delete (token, membershipId)  
- memberships.get_details (token, membershipId)  
- memberships.list (token, roomId=optional, personId=optional, personEmail=optional, max=optional)  
- memberships.update (token, membershipId, isModerator)  
- see https://developer.ciscospark.com/resource-memberships.html  

### Messages
- messages.create (token, roomId=optional, toPersonId=optional, toPersonEmail=optional, text=optional, markdown=optional, files=optional)  
- messages.delete (token, messageId)  
- messages.get_details (token, messageId)  
- messages.list (token, roomId, mentionedPeople=optional, before=optional, beforeMessage=optional, max=optional)  
- see https://developer.ciscospark.com/resource-messages.html  

### Teams
- teams.create (token, name)  
- teams.delete (token, teamId)  
- teams.get_details (token, teamId)  
- teams.list (token, max=optional)  
- teams.update (token, teamId, name)  
- see https://developer.ciscospark.com/resource-teams.html  

### Team Memberships
- team_memberships.create (token, teamId, personId=optional, personEmail=optional, isModerator=optional)  
- team_memberships.delete (token, membershipId)  
- team_memberships.get_details (token, membershipId)  
- team_memberships.list (token, teamId, max=optional)  
- team_memberships.update (token, membershipId, isModerator)  
- see https://developer.ciscospark.com/resource-team-memberships.html  

### Webhooks
- webhooks.create (token, name, targetUrl, resource, event, secret, filter=optional)  
- webhooks.delete (token, webhookId)  
- webhooks.get_details (token, webhookId)  
- webhooks.list (token, max=optional)  
- webhooks.update (token, webhookId, name, targetUrl)  
- see https://developer.ciscospark.com/resource-webhooks.html 
