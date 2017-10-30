# Arduino-Alcohol-Watch
Cisco project - Y9
run in Arduino IDE

(click to edit code to get nicer formatting)

#include <Time.h>
#include <TimeLib.h>
#include <LiquidCrystal.h>
LiquidCrystallcd(12,11,5,4,3,2);

//time
time_t tnow; //declaring variable tnow as type time_t 
time_ttlastprinted=now();

//alcohol
const int AOUTpin=0; //the AOUT pin of alcohol sensor into analog pin A0 of arduino (initialized '0')
const int DOUTpin=8; //the DOUT pin of alcohol sensor into digital pin D8 of arduino (initialized '8')
const int ledPin=13; //the anode of LED connects the digital pin D13 of arduino

int limit; //threshold - changeable
int alcoholevel; //reading from sensor
int alclevelscaled; //divided by 100 to give value between 1 - 100

//emogis
byte smiley[8] = {
  B00000,
  B10001,
  B00000,
  B00000,
  B10001,
  B10001,
  B01110,
};
byte sad[8] = {
  B00000,
  B10001,
  B00000,
  B00000,
  B01110,
  B10001,
  B10001,
};
byte middle[8] = {
  B00000,
  B10001,
  B00000,
  B00000,
  B11111,
  B00000,
  B00000,
};

void setup() {
  //time
  setTime(6,50,50,12,3,2017); //manually set time
  lcd.begin(16,2); //sets size of LCD
  //lcd.clear();
  lcd.print("iAM Watch");

  //alcohol
  Serial.begin(115200);//9600
  pinMode(DOUTPin, INPUT); //sets pin as input to arduino - sensor is input for aduino to read and process value
  pinMode(ledPin, OUTPUT); //sets pin as output of arduino 

  //custom emogis
  lcd.createChar(0, smiley);
  lcd.createChar(1, middle);
  lcd.createChar(2, sad);
}

//time - LED
int myPrintTime(time_t t){
  //add '0' if HOUR less than 10 - nicer formatting
  if(hour(t)<10){
    lcd.print("0");
  }
  lcd.print(hour(t));
  lcd.print(":");

  if(minute(t)<10){
    lcd.print("0");
  }
  lcd.print(minute(t));
  lcd.print(":");

  if(second(t)<10){
    lcd.print("0");
  }
  lcd.print(second(t));
  return 1;
}

void loop() {
  //clock
  lcd.setCursor(0,1);//bottom line of LCD
  tnow=now();
  //if time has changed print it
  if(tnow!=tlastprinted){ 
    myPrintTime(tnow);
  }
  tlastprinted=tnow;//saves last time printed

  //alcohol (imogis)
  alcoholevel= anologRead(AOUTpin);//reads anolog value from alcohol AOUT pin
  //sets alcohol level - take out if have sensor
  //alcoholevel = 12;
  limit = digitalRead(DOUTpin);//reads digital value from alcohol DOUT pin
  //sets limit manually if no sensor - take out if have
  //limit = 1;

  //alcohol to serial moniter - can see on computer and put into graph - raw data input to app
  Serial.print("Alcohol value: ");
  Serial.print(alcoholevel);
  Serial.print(" Limit: ");
  Serial.println(limit);//prints limit reached as LOW or HIGH - alcohol above or below limit
  delay(100);

  //alcohol to LCD
  alclevelscaled = round(alcoholevel/100);// readings from sensor r between 0 and 1023
  lcd.setCursor(10,0);//sets right of LCD
  lcd.print("Intake");
  lcd.setCursor(11,1);//sets line below
  lcd.print(alclevelscaled); //value between 1 - 9
  lcd.setCursor(14,1);
  //emogis
  if(alclevelscaled < 4){
    lcd.write(byte(0));
  }
  else if(alclevelscaled < 7) {
    lcd.write(byte(1));
  }
  else {
    lcd.write(byte(2));
  }
  delay(100);

  //aclohol - LED
  if (limit == HIGH){
  digitalWrite(ledPin, HIGH); //if limit is reached, turn LED on as status indicator
  } 
  else {
  digitalWrite(ledPin, LOW); //if threshold not reached, LED remains off
  }
  delay(100); 
}


