# MoodyBot
https://devpost.com/software/moodybot

## INCOMPLETE: A large part of this project was created on Built.io Flow's platform, however there is no current solution to sharing the project itself or extracting the code from it. The contents of this repository only include a local web-server used to work around the bugs found in Built.io's platform

## Inspiration
Have you ever wanted to listen to music based on how you’re feeling? Now, all you need to do is message MoodyBot a picture of yourself or text your mood, and you can listen to the Spotify playlist MoodyBot provides. Whether you’re feeling sad, happy, or frustrated, MoodyBot can help you find music that suits your mood!

## What it does
MoodyBot is a Cisco Spark Bot linked with Microsoft’s Emotion API and Spotify’s Web API that can detect your mood from a picture or a text. All you have to do is click the Spotify playlist link that MoodyBot sends back.

## How we built it
Using Cisco Spark, we created a chatbot that takes in portraits and gives the user an optimal playlist based on his or her mood. The chatbot itself was implemented on built.io which controls feeding image data through Microsoft’s Emotion API. Microsoft’s API outputs into a small Node.js server in order to compensate for the limited features of built.io. like it’s limitations when importing modules. From the external server we use moods classified by Microsoft’s API to select a Spotify playlist using Spotify’s Web API which is then sent back to the user on Cisco Spark.

## Challenges we ran into
Spotify’s Web API requires a new access token every hour. In the end, we were not able to find a solution to this problem. Our inexperience with Node.js also led to problems dealing with concurrency. However our biggest obstacle was probably handling Built.io's inability to import node modules in order to work around bugs in their software.

## Accomplishments that we're proud of
We were able to code around the fact that built.io would not encoding our images correctly. Built.io also was not able to implement other solutions to this problem that we tried to use.

## What we learned
Sometimes, the short cut is more work, or it won't work at all. Writing the code ourselves solved all the problems we were having with built.io.

## What's next for MoodyBot
MoodyBot has the potential to have its own app and automatically open the Spotify playlist it suggests. It could also connect over bluetooth to a speaker.

## Built With
javascript, node.js, cisco-spark, microsoft-emotion-api, spotify

## Credits
-[Ken Kang](https://github.com/kenkangg)<br>
-[Amin Mohtashami](https://github.com/Amin-Mohtashami)<br>
-[Krista Hayakawa](https://github.com/kehayakawa)<br>
-[Stephanie Stickel](https://github.com/stephstickel)<br>
