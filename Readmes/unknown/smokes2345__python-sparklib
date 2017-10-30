
*** Implementation Example
Can be found in cloudbot.py
```python
#!/usr/bin/python
from cherryspark import CherrySpark
import cherryspark,sys,cherrypy
from pytz import timezone
from datetime import datetime
import time,pytz

class SparkBot(CherrySpark):

	def __init__(self,*args,**kwargs):
		CherrySpark.__init__(self,*args,**kwargs)

	@cherryspark.post
	def time(self,tz="UTC",t=None,fmt='%Y-%m-%d %H:%M:%S %Z%z'):
		try:
			if not t:
				t = datetime.now(timezone(tz))
			else:
				t = datetime.strptime(t,fmt)
			return t.strftime(fmt)
		except pytz.exceptions.UnknownTimeZoneError as e:
			return "I don't know timezone {}".format(str(e))


if __name__ == "__main__":

	rooms = ["CIS Hackathon","test room"]
	bots = {}
	hostname = "cherryspark.cisco.com"
	port = 5000
	key = sys.argv[1]

	cherrypy.log.screen = True
	cherrypy.config.update({'server.socket_host': hostname,'server.socket_port': port})

	man = cherryspark.SparkManager(key)
	man.mountpoint = "/"
	cherrypy.tree.mount(man,man.mountpoint)
	
	for room in rooms:
		mountpoint = "/" + room.replace(' ','_')
		mon = SparkBot(key,room,hostname,port,mountpoint)
		mon.setupHooks()
		config = {mountpoint: {'tools.trailing_slash.on': False}}
		cherrypy.tree.mount(mon,mountpoint,config)
		bots.update({room: mon})
	
	if hasattr(cherrypy.engine, 'block'):
		cherrypy.engine.start()
		cherrypy.engine.block()
	else:
		cherrypy.server.quickstart()
		cherrypy.engine.start()
```