# IOT_Project_Fullstack
## About

A Fullstack for MQTT controlled pressure system. Software has an intuitive user interface that sends JSON packages to LPC1549 though MQTT. 

![1111](https://user-images.githubusercontent.com/61495413/139274217-54b428eb-96f3-4019-abfa-593880833b60.png)

* LPC1549 Reads the pressure data using a pressure sensor, formats it into JSON and pushes the message to the MQTT server.
* The server then receives the message and parses the information, the data then is added to the ArangoDB.
* The frontend poll the server with a POST methoid to query the latest MQTT information.
* User sends the chosen pressure level or fan speed though the user interface using a POST methoid. 
* The JSON gets then pushed back to the LPC1549.


## 1. Authentication page
<img width="815" alt="截屏2021-10-27 23 38 26" src="https://user-images.githubusercontent.com/56063237/139145216-226c763d-54ad-4d14-8cdd-32b39913c0e8.png">

This page shows that first time opening the website, the user will be asked to enter an username and a password. For testing purposes current usernames are: “x” and “v”, and the password is the same as the username. All Logins attemps are logged, including wrong passwords.


## 2. Home Page
<img width="966" alt="截屏2021-10-28 00 21 05" src="https://user-images.githubusercontent.com/56063237/139148961-24117f59-5eb2-4fe5-8a2f-3c66b54ca667.png">

#### 1. Home page:
Home page shows the main functions of the project which is to visualize the Fan speed and Pressure of the Embedded device.
#### 2. Graph page: 
Graph page contains a graph of the collected fan and pressure data.
#### 3. Logins: 
Login page shows when the users have logged in and how many times.
#### 4. Clear Fan Data
This button can clean all the history of the Fan speed and Pressure.
#### 5. Mode switch
Mode switch can decide between the automatic or manual mode.
#### 6. Automatic mode (A)

#### 7. Manual mode (M)



## 3. Automatic mode (A)


<img width="819" alt="截屏2021-10-27 23 53 59" src="https://user-images.githubusercontent.com/56063237/139146390-1a02f7f0-4ad8-427b-b461-b8bb563abebe.png">

This is manual mode.\
(The screenshot shows a example)\
The users can set the Fan speed. Then MQTT reports back the Pressure in *Pa*.
#### Pressure:
Green( When the pressure is less than 80)\
Yellow( When the pressure is over 80 or less than 100)\
Pink( When the pressure is over 100)

#### Fan Speed:
Green( When the Fan speed is less than 60)\
Yellow( When the Fan speed is over 60 or less than 80)\
Pink( When the Fan speed is over 80)


## 4. Manual mode (M)
<img width="831" alt="截屏2021-10-27 23 53 50" src="https://user-images.githubusercontent.com/56063237/139146378-bb4df640-0dbe-4f71-8171-2fcbbf286d17.png">

This is automatic mode.\
(The screenshot shows a example)\
In this mode, users set the chosen pressure level.
#### Pressure:
Green( When the pressure is less than 80)\
Yellow( When the pressure is over 80 or less than 100)\
Pink( When the pressure is over 100)

#### Fan Speed:
Green( When the Fan speed is less than 60)\
Yellow( When the Fan speed is over 60 or less than 80)\
Pink( When the Fan speed is over 80)



## 5. Graphs page
<img width="827" alt="截屏2021-10-27 23 54 15" src="https://user-images.githubusercontent.com/56063237/139146411-557647ca-5839-4c35-ae1d-b94146bf8f87.png">

This is graph page.
A graph that shows the time table for each sample received from the MQTT. The y-axis is represent the number of fan speed and pressure, and those two colors  separate them.

## 6. Logins page
<img width="826" alt="截屏2021-10-27 23 54 24" src="https://user-images.githubusercontent.com/56063237/139146414-c12b43f0-1648-4817-92a8-3e215df47963.png">

Login logs are stored here.

