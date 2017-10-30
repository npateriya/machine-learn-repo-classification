## AutoRun - Automated Emergency Protocols ##
## [Cisco DevNet Hackathon Winner](http://devnetevents.cisco.com/event/codingcamp-Calgary-2016) ##

#### **Authors**: [Christopher Billington](http://www.github.com/cbillington), Cal Nguyen, Craig MaCritchie ####
----------

### **About** ###
AutoRun is a service designed to automate emergency protocols for a fictional oil and gas scenario. This service leverages several technologies,

 1. An ASP.NET Web application for consuming/exposing REST API's written in C#.
 2. JavaScript scripts for working with [Cisco Tropo API](https://www.tropo.com/) to automate phone calls and SMS messaging.
 3.  Python scripts for leveraging a Raspberry PI to be an IoT device that streams temperature data to the cloud using [Relayr](https://relayr.io/).
 4. A [Cisco Spark Bot](https://www.ciscospark.com/) to act as the brain in the middle to create full automation by deciding what clusters to notify.

### **Hackathon Information** ###

#### **The Scenario** ####

Our customer Trollop Resources is a mid-size Oil and Gas producer. Head office is located in Calgary with 15 remote office locations across the province, some in very remote locations. Recently our customer has upgraded to the latest in Cisco networking, Wi-Fi, and collaboration solutions to ensure connectivity across alllocations. As a leader in the industry, safety is of the utmost importance to the organization and they would like to utilize their new technology to increase safety measures, improve field worker engagement and reduce costly management of field monitoring equipment.

Your team has been asked to come in and provide ideas and solutions to help solve some of the following business issues:

 - How can we improve safety notifications to field staff and provide improved reporting mechanisms?
 - How can we respond quicker to safety alerts?
 - What can we do to better engage with field employees and reduce their office time?
 - How can we utilize this new network environment to reduce the people cost related to monitoring field operations?

#### **Solution Documentation Submission** ####

**Assumptions**: Spoofed sensor data from Rasberri Pi to simulate fluctuating temperature changes. CMX location is available. Assumes location of sensors located in company database.

**How can we improve safety notifications to field staff and provide improved reporting mechanisms?**

 - To improve safety notifications, we’ve automated the process (using
   spark/tropo)
 - We allow employees to call and emergency number of the current
   situation and location (grab coordinates via CMX). The resulting call will notify all relevant personnel (via phone and text) of the emergency and what to do next.
 - Automatically detect warnings from instrumentation sensors (through
   Relayr). When we detect abnormal readings, we notify operators to troubleshoot the situation by giving them the location and a description of the problem.
 - Allowed administrators of the spark room to broadcast an emergency to
   all our personnel.
 - To provide improved reporting mechanisms Provide a monthly report of
   safety incidents and their respective locations (all saved into a
   database)
 - Demonstrate a reduction in safety incidents (with a monthly
   production report), as a direct result of operations responding to
   raised alarms before the issues become an incident/liability

**How can we respond quicker to safety alerts?**

 - Setting up command lines for BEEZBOT.
 - Used a third party application to notify personnel about hazardous
   weather conditions in their area.
 - Setup the BEEZBOT emergency command where everyone is notified via text message &amp; a phone call of the problem/situation.

**What can we do to better engage with field employees and reduce their office time?**

 - Set up data points using Relayr (and gateways to analyze the raw
   data). Say we look at compressor temperature for overheating, or the pressure if it is high.. And send an alarm to the operators to go to the location

 - Same goes with an oil/gas well.. We can monitor the pressure,
   temperature and torque and for high alarms we can send the operator to the location to fix the situation...Say if the area is an oil well, and the oil tanker is overflowing or at a high level, we can notify the operator to call truckers so we prevent spills. Operators no longer need to come to the office and have meetings.. As we have conference call on the fly for troubleshooting.

**How can we utilize this new network environment to reduce the people cost related to monitoring field operations?**

 - Reduced operator time and cost by routing field staff to necessary
   locations automatically, which reduces field costs and time.

 - With operators routed to proper locations automatically, we will
   reduce the down time for site operations (unnecessary pump changes
   due to catastrophic failures/equipment failure). This will in turn by
   fully automating safety emergency protocols and reporting, we allow
   significantly faster response times and allow us to gather a large
   collection of data about emergency responses and hotspots. This will
   greatly reduce costs by reducing overall incidents and standardizing
   small failures that become catastrophic.

### **Example** ###

#### **Features** ###

**Automated Call In**

This feature allows employees to call an emergency number with an emergency. The service will then transcribe the message and automatically report the issue to the proper personnel.

*Tropo Endpoint to Accept Call and Transcribe.*

```javascript
say("This is the Emergency Response Line!");

record("What is the emergency and your current location?", {
	beep:true,
	maxTime:80,
	terminator: "1",

//transcriptionOutURI:"http://requestb.in/1g331pt1"

			transcriptionOutURI:"https://3d27d5c4.ngrok.io/api/Tropo",
transcriptionID:callerID + "," + getRandomArbitrary(51.0442711,51.08) + ";," + getRandomArbitrary(-114.0685925,-114.09)
});
```
ASP.NET REST Endpoint will now capture the call from Tropo and notify all relevant personnel using Cisco Spark API (Cisco Spark is used as a "team" service which has a collection of all personnel)

```csharp
private static void TextMembers(List<Member> members, string message)
        {
            HttpClient textClient = Config.GetClient(Config.tropoBase);

            string phoneList = "[";
            foreach (var member in members)
            {
                phoneList += "\"" + member.phone + "\",";
            }

            phoneList = phoneList.Trim(',');
            phoneList += "]";

            var textClientContent = new FormUrlEncodedContent(new[]
            {
                    new KeyValuePair<string, string>("token","4c6176697958676e6f66746448455548685350554b455a456a4166756e644c75486b5649724e6e776a4b756a"),
                    new KeyValuePair<string, string>("msg", "Emergency Message: " + message),
                    new KeyValuePair<string, string>("networkToUse", "SMS"),
                    new KeyValuePair<string, string>("numbersToDial", phoneList),
                });
            textClient.PostAsync("", textClientContent);
        }
```

**Automated Notification based on IoT sensor data**

This feature uses a Rasberry PI as an IoT temperature sensor which streams data to Relayr cloud. The ASP.NET application can then pull sensor data and decide whether an employee needs to be routed to the location or if the temperature is past critical and notify relavent personell of an emergency by Calling/SMS with the emergency and location.


Rasberry PI is setup with the capability of posting sensor data to the cloud (acting as a gateway).


ASP.NET application decides whether an emergency is required and notifies personell by leveraging the Cisco Spark team and Tropo.

```csharp
private static void CheckSensor()
        {
            HttpClient getClient = new HttpClient();
            string getAddress = "https://api.relayr.io/devices/e84e2eb1-80bf-48e8-a5c1-c710c5310281/readings";
            getClient.BaseAddress = new Uri(getAddress);
            getClient.DefaultRequestHeaders.Accept.Add(
               new MediaTypeWithQualityHeaderValue("application/json"));

            getClient.DefaultRequestHeaders.Authorization = new AuthenticationHeaderValue("Bearer", "RNClQ25HqmxEtFquKuGRAv4b8OAnFbzvh5MpFgNtyqpJOOcngIQEGwf0Xw0vgF4n");

            HttpResponseMessage response = getClient.GetAsync("").Result;
            var readingWrapper = response.Content.ReadAsAsync<ReadingWrapper>().Result;

            decimal temperature = Convert.ToDecimal(readingWrapper.readings[0].value);

            Sensor sensor = new Sensor();
            string status = "Normal";

            string message = "";
            if (temperature > 30)
            {
                message = "Critical readings at: " + "\n" +
                          "LSD: 66-27-75-05W4 \n" +
                          "Temperature: " + temperature;

                SparkBot.NotifyText(message);
                status = "Alert";
            }
            else
            {
                message = "LSD: 66-27-75-05W4 \n" +
                          "Temperature: " + temperature;
            }
            SparkBot.PostMessage(message);

            sensor.LSD = "66-27-75-05W4";
            sensor.temperature = temperature.ToString();
            sensor.status = status;

            DBAccess.InsertSensor(sensor);
        }
