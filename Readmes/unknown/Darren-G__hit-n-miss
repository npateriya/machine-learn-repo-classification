hit-n-miss
==========

Letterkenny Coder Dojo CISCO Cisco Challenge entry


About
-----
Hit'n'Miss is a game, designed, coded and packaged in less than a week as an entry for the Cisco Transatlantic CoderDojo Challenge 2013. It uses Python and Pygame. The game is a simple battleship-type clone, in which you hunt down a hidden archery target with your bow and arrow. If you are lucky enough to hit one of the outer rings, the appropriatley coloured arrow head serves as a guide to help in your quest to get that bulls-eye. The game can be played in 1-Player practise mode where you hunt down a randomly placed target, or in 2-Player mode against an opponent over LAN or Internet, where each player places their own target, and tries to shoot the other in the least amount of goes.

The game is designed to be easily ported to a phone or handheld device.


Installation
------------
An executable installer is provided to allow any windows user to run/play the game without the need to install Python or Pygame on their PC. Just follow the installer instructions.

For Mac/Linux, the installed Python (2.5 or greater) will need the pygame module to allow the game to run.


Help
----
The single player game can be run standalone, just keep trying to find those targets!

The 2-Player experience requires a player start up the game in 'Server' mode. The 'Client' player will need to know the Server's IP addess and update their settings.ini file accordingly. The port number may need to be updated to by-pass any firewall rules and/or router configurations.


Credits
-------
A big thank you to Cara House Family Resource Center in Letterkenny for providing the premisis for our CoderDojo.

To the team that worked on this Game in such a short time:

 - Grace Doherty (Game design and coding)
 - Laura Young (The Name! and screen designs)
 - Ciara Phelan (Backgrounds and other art)
 - Patrick Dorrian (Networking and Game Logo)
 - Mitchell Goudie (Sound/music and coding)


Dependencies
------------
Python and Pygame are required to run the game from source.


Todo/Ideas
----------
Due to the shortage of time from finding out about the competition till the closing date, there are quite a few things that we would like to update:

 - The network code - if Pygame ever gets a decent network module!
 - Standalone server, to allow for multiple connections / simulteneous games
 - An android / iPhone port
 - A javascript / web client
 

Diaries
-------
Entry from Grace Doherty

Day 1
Discussed possible project ideas iincluding chat room and X's and O's games
Decided on a battleship type game
discussed different themes- spaceships, planets, balloons etc.
decided on Robin Hood theme - Arrows and targets
We talked about different methods of scoring and designs for the game including points based games and multiplayer scenarios
I started working on the target, what it looked like and how it was placed.
I managed to also work out a few bugs straight away such as not letting the player click outside the light green area.

Day 2
I worked on how to make the target transparent after placed, decided that using an image would be easier than drawing my own circles.
downloaded a suitable image and edited it using gimp so that the background colour was completly different to any other colours used in the program.
I loaded it to the game code.
I then changed how it was loaded as at first it was loading from the corner instead of the center - makes it look much better.
I then worked on changing the background to an image and blending the dark green area to a more suitable colour.

Day 3
Diecided to use a different layout for the code, keeping in mind the different game states and screens needed- Start, Place, Play and End
Compiled what everyone was working on to make a "real" game.
Discussed possiple win/lose screens and different ways of incuraging future play- stars, percentage etc.


Entry from Mitchell Goudie

Day 1:
Today there were only two members of the team as some were missing. We decided to work on a game rather than a chat server as they're more fun and grab peoples attention. I worked on the arrow placement by starting to place circles where the user clicked rather than images to make sure it worked. I came up with the idea of having it robin-hood esque. I was assigned to work on the sounds for the game. 

Day 2:
The entire team is in today, thank goodness.Today I worked on getting the screen to display an arrow where I clicked. This took some time and I'm making the noises in audacity. We've decided to go with a target game where users place their target and then click around the screen to try and find their opponents bullseye. The target has several rings which are in a colour sequence so the users has some idea of which way the centre is. I now have to somehow tweak the image so that it loads a different colour of arrow depending on the ring it hits. My teacher told me that it would be ideal if they were in the same image so we can load it as a sprite sheet, so I have now incorporated that but it draws the image in the corner, which is bad. My teacher suggested that I try and centre it by bring it over and down the correct amount of pixels, which worked brilliantly.

Day 3:
Everyone is in today, so the workload will be even. Today I worked on the sounds and converting them to a format that pregame would recognise by default, which is .wav . We've decided that my background music wasn't very happy in tone, so I've replaced it with a loop of birds chirping as the background is a forest. My teacher has been helping Patrick with the network side of things, so we can make it multiplayer.

Day 4:
Everyone is in today again, great. Today I worked on creating an olden age sign to act as the background to a button which will be styled as a scroll. I've created multiple lengths of the sign using gimp to edit the images. I've also finalised the sounds. The name of the game has been decided as Hit'N'Miss by Laura.

The thought that all this was accomplished in roughly a week is astounding.


Entry from Darren Gallagher (Mentor)

Monday update:
We've decided against a chat server - I had a look at the code required and it would've been too much to try and complete within a week.
We are going to work on a simple 2-Player game that will run over a network (or internet, if we have time).
The game is a variant of Battleships, where each player chooses where to place their archery target. Each player in turn clicks on their screen in an attempt find to 'shoot an arrow' into their opponents target. Any hits will be recorded on their screen with the colour corresponding to which ring of the target they hit. A bulls-eye hit wins the game for you...
Grace has the target placement working; Mitchell edited some sound effects and started on the hit recorder and drawing.
We will need to think about screen designs for the various game states (splash screen, username entry, game running and game over). Maybe some background music - we can 'borrow' this from the internet but it would be much better if it was an original piece...
I have found a nice bit of client/server code that we can re-use the code from.

Tuesday Update:
Grace got the target (now an image file) to move on the playing screen and fade-out when placed.
Mitchell worked with the coloured arrow heads as markers for any shots / hits.
Ciara gathered up backgrounds and image files, cropping and resizing.
Laura started looking at the screens for the various states of the game.
Patrick started looking at the client / server network code, but there is something funny with the router setup in Cara House - it is preventing us from proper testing.
I've combined some of these updates into a 'working' demo... The files can be downloaded from:
https://github.com/Darren-G/archery-game
Here, I have generated a random player starting position, which you can click around to try and find. I've also embedded the sound files - these may not work on some installations.
I'll take some network equipment with me tomorrow, to see if we can push on with the network testing - we really need to get this working to give us a valid entry.

Wednesday Update:
Code base re-write... This makes it easier to work on each of the game states (start, placement, playing, end) independently. We now have a main game engine that triggers mini-engines within each of the game states.
Still struggling a bit getting network code to work... Maybe tomorrow!
Laura & Ciara are working together on art and layout for the various screens - check the Android development sites for inspiration...
Mitchell has recorded custom sound effects and maybe a backing track (mutable if possible!) - suggestion to 'happy' it up a bit.
Patrick, while being very patient waiting on network code, has started on the game instructions and may have a logo for us tomorrow. Everyone should be thinking about a name for the game.
Grace is going to compile a write-up to be used as development documentation for the project - this always impresses people - it shows that a development process was followed.
I have sent the code on to a games developer (works for CodeMasters) to get feedback from someone in the industry. He's impressed with the initial view. One suggestion was to include plenty of comments into the code.

Thursday Update:
The network code is gonna take too long - I wish we had more time! Currently the Server and Client will swap target positions and both can then play their own game.
Today, we agreed on the name for the project: Hit'n'Miss (thank you, Laura!)
We have our logo and background images (including a nice wooden sign and some scrolls for the buttons)
A few new screens have been added.
All that left now, is to add the final music and sound effects, compile the README file and package it up.
Check it out: https://github.com/Darren-G/hit-n-miss

Friday Update:
Pulled it all together nicely. Today we used cx_freeze to to convert the python script into a standalone executable and Inno to package it all up into a handy installer.
Some additional testing on the 'Network' part proved that it does indeed run on separate machines in a LAN... Nice!
I'll push the code up to Github later and include the README and excerpts from the diaries that we have been keeping.

Well done guys, you did in less than a week, what some development teams would have taken months to complete. Unbelievable!
