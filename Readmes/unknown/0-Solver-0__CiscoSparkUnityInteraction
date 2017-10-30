"# CiscoSparkUnityInteraction" 
1.	Clone the repository:  https://github.com/0-Solver-0/CiscoSparkUnityInteraction
2.	Download ngrok from https://ngrok.com/download to get the IP adresse, which make a tunnel between the local network and the internet. Use the following instruction to create a tunnel on port 3000 : ngrok http 3000
 ![ngrok command](https://image.ibb.co/kWfOkQ/ngrok.png)
3.	Creating a CiscoSpark account, and get the User access token.
4.	Creating a Webhook https://developer.ciscospark.com/endpoint-webhooks-post.html and get the webhook ID.
 ![WebHook creation](https://image.ibb.co/kG1kzk/webhook.png)
5.	After setting all the variables In the ComSpace.js file, install nodeJS and run this file as : node ComSpace.js.
6.	The used ChatBot is a Ciscospark bot with the adresse probook@sparkbot.io which is located into cloud repository in heroku.com. 

process.env['public_address'] = 'https://fa4f672e.ngrok.io;
process.env['access_token'] = 'User access_token';
process.env['TCP_PORT'] = '8000';
process.env['HTTP_PORT'] = '3000';
process.env['BOT_ADRESS'] = 'probook@sparkbot.io';
process.env['webhock_id'] = '';
process.env['event_name'] = 'messages_created';
