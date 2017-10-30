# ciscoCLI4groovy

Cisco CLI generic command executors

#About 

This is a groovy expect wrapper of Cisco IOS CLI interface. It allows you to automate in groovy, java 
and jvm kind of programing languages the communication with Cisco IOS routers and switches.

The library allows your application to send CLI commands in privilege and config terminal modes and also to apply
configuration templates over telnet or ssh without the need of specific command parsing. 

As a difference to many other such expect attempts (either open or close source) our attempt is generic so you can
send any command in any device mode without taking too much care on what the output will be, which symbols should be 
expected or how exactly to send the actual command or commands template. 

# Usage


## Login

In this example we will demonstrate how you can login to the router in privilege mode (privilege level 15). 
If you want to do a telnet session just change the protocol and the port. 


```java
        Hashtable<String, String> config = new Hashtable<String, String>();
        config.put("StrictHostKeyChecking", "no");

        Map<String, Object> params = new HashMap<String, Object>();
        params.put("protocol", "ssh");
        params.put("username", "user");
        params.put("password", "pass!");
        params.put("enable-password", "pass321!");
        params.put("address", "1.1.1.1");
        params.put("port", 22);
        params.put("timeout",30000);
        params.put("config",config);
        params.put("LOGGING_LEVEL","DEBUG");
        params.put("prompt",">");
        params.put("defaultTerminator","\r");
        params.put("powerUserPrompt","#");
        params.put("retries",2);

        Expect4GroovyScriptLauncher launcher = new Expect4GroovyScriptLauncher();
        //Try to login 

        Map<String, Object> loginResult = launcher.open(new String[]{"conf/groovy/cisco/ios" + File.separator}, "cisco_login.groovy", params);
        //Check login result
        if (loginResult.get("status").equals(2)) {
        //Can't login.
            logger.debug(loginResult);
        } else {
        //Login sucessful send further commands. 
        }
```


## Send privilege mode command over an already logged in session
In this example we will demonstrate how you can send a privilege mode (15) command over ssh. 
If you want to send it over telnet just change the protocol and the port. 

```java
        
            //Assuming that the login has already happen and we want to send a single privilege mode command. In this case this will be "show running-config"
            
            Map<String, Object> cmdParams = new LinkedHashMap<String, Object>();            
           
            cmdParams.put("command", "show running-config");
            cmdParams.put("mode", loginResult.get("mode"));

            result = launcher.sendCommand("cisco_sendCommand.groovy",cmdParams);
            //Check the command execution result 
            if(result.get("status").equals(1)){
            //Success check the response of the command. For examle in this case the "data" will be the output of\"show run" command"
            //Or the error produced during the execution of it.
                System.out.println(result.get("data"));
                
            }else{
            
                System.out.println("Command has not been executed sucessfully!!!: "+result.get("data"));
                System.out.println("ReportResult"+result.get("reportResult"));
                System.out.println("EvalResult"+result.get("evalResult"));

            }
            
```
            

## Send a single configuration command over an already logged in session

```java
            cmdParams.put("evalScript", null);
            cmdParams.put("configCommand","ip route 10.200.1.0 255.255.255.0 192.0.2.1");
            cmdParams.put("mode", loginResult.get("mode"));
            cmdParams.put("hostname", loginResult.get("hostname"));
            cmdParams.putAll(params);
            Map<String, Object> result = null;

            result = launcher.sendCommand("cisco_sendConfigCommand.groovy", cmdParams);
```

## Send multiple configuration commands (configuration template) over an already logged in session

```java
            cmdParams.put("mode", result.get("mode"));
            cmdParams.put("hostname", loginResult.get("hostname"));

            String configTemplate = "interface loopback 101\n" +
                    "ip address 127.0.0.101 255.255.255.255\n" +
                    "description useless loopback\n"+
                    "no shutdown\n";


            cmdParams.put("configTemplate", configTemplate);
            cmdParams.remove("configCommand");

            result = launcher.sendCommand("cisco_sendConfigCommand.groovy",cmdParams);

```