# IOT_Project_Fullstack
## About

A Fullstack for MQTT controlled pressure system. Software has an intuitive user interface that sends JSON packages to LPC1549 though MQTT. 

![1111](https://user-images.githubusercontent.com/61495413/139274217-54b428eb-96f3-4019-abfa-593880833b60.png)


## 1. Authentication page
<img width="815" alt="截屏2021-10-27 23 38 26" src="https://user-images.githubusercontent.com/56063237/139145216-226c763d-54ad-4d14-8cdd-32b39913c0e8.png">

First time opening the website the user will be asked to enter an username and a password. For testing purposes current usernames are: “x” and “v”, and the password is the same as the username. All Logins attemps are logged, including wrong passwords.


## 2. Home Page
<img width="966" alt="截屏2021-10-28 00 21 05" src="https://user-images.githubusercontent.com/56063237/139148961-24117f59-5eb2-4fe5-8a2f-3c66b54ca667.png">

#### 1. Home page:
home page shows the main functions of the project which is to visualize the fan speed and pressure of the Embedded device.
#### 2. Graph page: 
Graph page is the graph that shows the history. 
#### 3. Logins: 
Login page shows the time and users which used the project website.
#### 4. Clear Fan Data
This buttom can clean all the history of the fan speed and pressure
#### 5. Mode swtich
it can decide the mode whether automatic or manu mode
#### 6. Automatic mode (A)

#### 7. Manual mode (M)



## 3. Automatic mode (A)
<img width="831" alt="截屏2021-10-27 23 53 50" src="https://user-images.githubusercontent.com/56063237/139146378-bb4df640-0dbe-4f71-8171-2fcbbf286d17.png">

This is manual mode.
And users can set the Fan speed. Then MQTT reports back the pressure in Pa.
#### Pressure:
green(<80)\
yellow(>80 <100)\
pink(>100)



## 4. Manual mode (M)
<img width="819" alt="截屏2021-10-27 23 53 59" src="https://user-images.githubusercontent.com/56063237/139146390-1a02f7f0-4ad8-427b-b461-b8bb563abebe.png">

This is automatic mode.
In this mode,users set the chosen pressure level.
#### Fan Speed:
green(<60)\
yellow(>60 <80)\
pink(>80)



## 5. Graphs page
<img width="827" alt="截屏2021-10-27 23 54 15" src="https://user-images.githubusercontent.com/56063237/139146411-557647ca-5839-4c35-ae1d-b94146bf8f87.png">

This is graph page.
A graph that shows the time table for each sample received from the MQTT.

## 6. Logins page
<img width="826" alt="截屏2021-10-27 23 54 24" src="https://user-images.githubusercontent.com/56063237/139146414-c12b43f0-1648-4817-92a8-3e215df47963.png">

Login logs are stored here.
