# Spark SDK Tutorial - Making a call the easy way

## SEVT Track 1. Embedding Video into an iOS App in a few lines of code using a wrapper library *(Beginner)*
This project is based off a standard Swift Single View iOS App but has a configured dependency for the Spark SDK which has been tied to a ‘Start Call’ button. Using this tutorial you will be able to place a call using the Spark Video SDK for iOS with minimal setup.

## What is a Wrapper?
A wrapper library is a software library that is written primarily for the convenience of developers. 

Most APIs/SDKs that are published have a lot of capabilities but also a lot of Boilerplate code, Boilerplate code is what we call code which in most cases you will just copy and paste as a requirement of   using an SDK. For example most people that use the Cisco Spark Video SDK will need to implement functionality to handle what should happen when an end user presses the mute button on their video call. In most cases the code that is written to do that would be the same in most projects that are built on the iOS SDK. 

So the aim of a wrapper library is to provide a basic reference implementation of everything that the SDK needs to function but abstracted into a more simple from, so that instead of having to write a complete implementation using a lot of Boilerplate code you can just use a wrapper library which someone else has written for you so that in this case you can implement Video calling in just a few lines of code lines of code.

## Prerequisites
Building iOS Apps requires a number of prerequisites in order to work correctly. In this case please ensure that your system meets the following requirements.

* A Mac running the latest version of Mac OS X 10.12 Sierra
	- If you are unfamiliar with updating Mac OS X you can follow the instructions here [Update the software on your Mac - Apple Support](https://support.apple.com/en-us/HT201541)
* The latest version of XCode downloaded and **installed** via the Mac App Store (at time of writing 8.3.2 or greater)
* Cocoapods **installed** and **initial pod setup completed** (to do this, open your terminal and run the following commands):
	* `sudo gem install cocoapods`
	* `pod setup (can take 10-15 minutes sometimes)`


## Lab Steps:
1. Download or Clone this project to your local Mac, your desktop or documents folder for example.
2. Open the Terminal.

3. Navigate via the CLI to root directory of the project your just downloaded e.g. `cd Desktop/SparkSDKExample-master`.

4. Enter the command in the Terminal `pod install`.

5. Once the `pod install` command has finished, issue the command `open SparkSDKExample.xcworkspace/`, this will cause Xcode to open. If you have not used XCode before, you may be asked to activate Developer mode on your Mac, to do this just press accept at the prompt and provide your username and password if needed.

6. In XCode, click on `ViewController.swift` in the Xcode left navigation panel. (If this is your first time in the project, you may need to expand the 'SparkSDKExample' by clicking on the 'disclosure triangle' in the left panel).

7. Replace the text that says `“API_KEY”` with your API key from https://www.developer.ciscospark.com , ensuring that your key has “” around it. E.g. “AK54DGDG”. This API key is used to tell Spark which account is starting the call.

8. Replace the text that says `“RECIPIENT”` with a Spark account address or SIP URI that you wish to call. E.g. “jonfiel@cisco.com” or “sip:2123123@abcxyz.com”. The Recipient is the Spark account or SIP URI that will receive the call (the API key from step 8 **cannot** be from the same account as the recipient that you are calling. The reason for this is, **Spark does not support calling yourself**).

9. Press the “Play” button top left corner of Xcode to start the App.

![Alt text](/TutorialAssets/playButton.png "Play Button Example")


11. The app will start in the iOS Simulator, you should then press the “Start Call“ button and it will trigger a call to be made via the Cisco Spark SDK. The iOS Simulator does not support showing a self view at the moment, so the call will just show the remote party (recieve video but not display self view), unless you choose to run the app on your personal iOS device (ask your proctor).

![Alt text](/TutorialAssets/startCall.png "Start Call Example")

