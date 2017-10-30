# Mobile LED Message Detection

###Useful Links
	- configuring opencv in android studio: http://blog.hig.no/gtl/2014/08/28/opencv-and-android-studio/

	- communication standard: http://en.wikipedia.org/wiki/Asynchronous_serial_communication
	
	- coercing android to set the camera fps: http://answers.opencv.org/question/21613/how-can-i-change-fps-of-javacamera-on-android/

	- managing screen orientations: http://stackoverflow.com/questions/1512045/how-to-disable-orientation-change-in-android

	- integration of native code using jni: http://code.tutsplus.com/tutorials/advanced-android-getting-started-with-the-ndk--mobile-2152
	
	- thresholding in opencv: http://docs.opencv.org/2.4.2/doc/tutorials/imgproc/threshold/threshold.html

##Clarifying Questions
	
	1. 
		Q.) How are the LEDs read -- is the application expected to read binary messages conveyed through flashing lights, 
		is it supposed to read static lights, is it supposed to be able to distinguish color?

		A.) Application is expected to get a string of bytes by reading flashing LEDs. If LEDs can flash at frequency unseen by the eye 
		– it would be an advantage, but is low priority at this stage. Using color might be interesting to convey more information, but 
		is low priority at this stage.

	2. 
		Q.) Is the application expected to be able to distinguish between various LED arrangements on different computer types? i.e. the 
		arrangement of LEDs on the front of one computer might be different than that of another computer

		A.) This applies when >1 LED is used for signalling. It is likely this will be needed as amount of information transferred via single 
		LED might be very small. >1 LED is lower priority now, but it would be wise to make provisions to simplify testing this path should it 
		become necessary. For >1 LED we need a heuristic (for example LEDs are numbered left to right, top to bottom etc) for decoding information.

	3. 
		Q.) For the data being transmitted, does the transmission follow a specific standard? If so, what is it? Is the application expected to work 
		with LED information which is already in place, or is part of the project deciding how the computers should convey information through LEDs?

		A.) Yes, latter is part of the project and is, in fact, main priority now. Essentially, proof of concept should answer the following questions fast: 
 
		Achievable data rate / LED (and per camera FPS: 15/30/45/60/90/120) Error rate (or data rate with error correction if it is used) How 
		disruptive to conventional use of LEDs (i.e. can this be done completely transparently with LED appearing on/off is its main function would dictate)  
 
		Same as above but when using handheld camera (think of imperfect orientation, moving glare, some shaking, time limit – it’d be awkward 
		to hold camera steady for longer than probably ~5-10 sec – can find this limit experimentally as well)

	4. 
		Q.) What should the app identify? e.g. machine specifications, machine status, error codes, etc.

		A.) This is TBD, but initial idea is something along the lines: what device it is (name / type / status – i.e. if something is abnormal – 
		what it is, can be derived from syslog or other sources of device state information – depends on how much 

###Timing Visualizations
Original protocol:
![Timing Diagram](timingDOne.png)
Updated protocol:
![Updated Timing Diagram](updatedTimingD.png)

###State Diagrams
App Level FSM
![App Level FSM](appLevelFSM.png)
Bit Level FSM
![Bit Level FSM](bitLevelFSM.png)

##Preliminary Limitations

	1.	Cannot push camera FPS beyond 30, and difficult to hold FPS at anything above 25

	2.	Best FPS quality results from grayscale and 480p resolution, thus color coding information works against speed of data transfer

	3.	Will most likely need to code processing functionality in native language to eliminate overhead

###Stats
	1.	Start bit: 310ms
	2.	"1" 205
	3.	"0" 85
	4.	Stop bit: 400ms
	5.	LED on for 84ms
