# HomeCamSpark - home surveillance camera with Cisco Spark

This is a home surveillance camera integrated with Cisco Spark.

Features:

* Motion detection: when the camera detects a motion, it will take a picture and send to the Spark room.
* Speak: you may send text to the Spark room, and it will be spoken by the camera.

## Setup

1. Download CiscoSpark on your mobile device, sign in with your personal account
2. Install dependencies on Ubuntu 14.04: `apt-get install php5-cli motion festival`
3. Sign in to <https://web.ciscospark.com> with a separate camera account
4. Create a room with your personal account, and send an arbitrary message
5. Log in to <https://developer.ciscospark.com>, copy access token, and paste in `accesstoken.txt`
6. Execute `php5 choose-room.php` to choose a room that contains your personal account
7. Edit `motion.conf` if necessary: change `videodevice` to your webcam device name
8. Download [ngrok](https://ngrok.com), unzip and put binary in current directory
9. Login to <https://dashboard.ngrok.com>, execute the "ngrok authtoken" command line shown on the page

## Usage

1. Execute `./start-www.sh` to start a private web server to order to serve captured images
2. Press CTRL+C on `motion`, and execute `./stop-www.sh` to stop the service
