# mse
Python module to interface with Cisco MSE
##Functions
###WirelessClientLocation getClient(string server, string user, string password, string client)
 *Returns a single WirelessClientLocation based on the search 'client'.*
###[WirelessClientLocation] getClients(string server, string user, string password)
 *Returns a list of WirelessClientLocation from list of server.*
###[WirelessClientLocation] getAllClients(string user, string password)
 *Returns a list of WirelessClientLocation from list of servers.*
##Classes
*MapInformation*
  * imageName
  * floorRefId
  * offsetX
  * offsetY
  * height
  * width
  * length
  * unit
  * mapHierarchyString

*MapCoordinatePair*
  * y
  * x
  * unit

*WirelessClientStatistics*
  * currentServerTime
  * lastLocatedTime
  * firstLocatedTime

*WirelessClientLocation*
  * mseServer
  * userName
  * macAddress
  * isGuestUser
  * Statistics (*WirelessClientStatistics*)
  * currentlyTracked
  * ssId
  * dot11Status
  * band
  * MapCoordinate (*MapCoordinatePair*)
  * apMacAddress
  * confidenceFactor
  * ipAddress
  * MapInfo (*MapInformation*)
 
##TODO
  * *Combine getClients and getAllClients, let it accept either a single string or list of strings of server addresses.*
