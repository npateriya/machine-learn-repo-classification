
Cisco Spark - List spaces Python

## Step 1. List Your Spark Spaces
Copy/paste the following code to your script. It imports the required modules

```python
# coding=utf-8
import json
import requests
```

> **NOTE**: The “```# coding=utf-8```” handles special characters in space names


## Step 2. set the Personal Access Token variable
and update your token.
```python
myToken="YOUR_PERSONAL_ACCESS_TOKEN_HERE"
myToken="Bearer "+myToken
```

The token send to Spark has the format "Bearer XjY3ZMD…", added in the 2nd line above


## Step 3. function to get your space data.
Add this function (a collection of code that can re-used) that lists all of your spaces when called with 1 parameter:  your Token. 
```python
def get_spaces(mytoken):
	# The header is send to authenticate
	header = {'Authorization':mytoken, 'content-type':'application/json'}
	# Send GET request with above header, put result in 'result'
	result=requests.get(url='https://api.ciscospark.com/v1/rooms',headers=header)
	# Encode 'result' as JSON
	JSONresponse = result.json()
	# Create an Array for all spaces: ('space1','space2','space3')
	spacelist_array = []
	# For each item in the JSON data
	for EachSpace in JSONresponse['items']:
		# add the 'title' + 'space ID' to the spacelist_array
		spacelist_array.append(EachSpace.get('title') +
						' ** ' + EachSpace.get('id'))
	# Return the list of members
	return spacelist_array
```


Sample JSON response after GET request:
```json
{
"items": [{
    "lastActivity": "2016-03-23T09:25:08.217Z",
    "created": "2016-03-22T16:22:50.745Z",
    "isLocked": True,
    "id": "Y2lzY29zcGFyazovL30YS0xMWU1LTlhY2ItZWQwMThkMzRlMDU2",
    "title": "EMEAR Cloud Collab SE"
    }]
}
```

## Step 4. Printing the results
Next add this code that executes the get_spaces function and print a list of my spaces

```python
# Execute get_spaces and put result in 'SparkResult'
SparkResult = get_spaces(myToken)
# Loop through the SparkResult array
for space in SparkResult:
	# print space name
	print ("> Space:", space)
# Print the number of spaces (length of SparkResult array)
print ("----- TOTAL Space Count:", len(SparkResult))
```

Execute the code by clicking "RUN" and look at the output. It should look something like this: 

```
> Space: API doc team  ** YU21ZiMjYtZjN1ZigxYjDk2TAcGFyMDcx1ZiMWU21ZiMjYtZjN1ZigxYjBmNDk2
> Space: Customer ABC ** YAcGFyMDcx1ZiMWU21ZiMjYtZjN1ZigxYjBmNDk2Y2zcazovcGFyL1JPcGFyTlhO
> Space: Team planning ** Y2l1Zi29zcGFyazovcGFyL1JPcGFyTTcx1ZiMWU21ZiMjYtZjN1ZigxYjBmNDk2
```

> **NOTE** that this code will only return a maximum of 100 spaces. This is the default max of the Spark API unless you specify the Max parameter (see list spaces documentation) 


**NEXT LAB** [click here](http://gsx-clust-external-12qq4jfx2i278-442547460.us-west-1.elb.amazonaws.com:5555/tutorial?tut=https://github.com/JockDaRock/Time2Code/blob/master/QuickCodeTutorial.md?raw=True)

What happend to the hyperlink above?
