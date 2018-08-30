# My Postback
## A Kochava Postback Delivery Mini-Project

### Description
The purpose of this project was to create a program that was capable of ingesting RAW POST DATA and delivering an response 
to an http endpoint. A php application (ingestion-agent) is used to accept the RAW POST DATA and format it to a "postback" object that is sent to
 a Redis queue. Next, a Golang application (delivery-agent) pulls these postback objects from the Redis queue, further formats them, and then 
 sends a request via an http endpoint. The responses from these requests are then stored in a log file. 
 
### Installation and Usage
Docker must be installed on your local machine (https://www.docker.com/get-started). Pull this github repo into your desired 
directory. </br> </br>
Run in your terminal: 
</br>
`docker-compose up --build`
</br>
OR (to run in background)
</br>
`docker-compose up --build -d` 
</br>
Shut down
</br>
`docker-compose down`
</br> </br>

#### Accessing log file
The log file of responses (e.g. delivery time, response code, response time, and response body) is stored in log.txt and is 
created when the program is run and sent RAW POST DATA. Alternatively, vim is installed in the docker image for the delivery-agent
and can be accessed by go into the CLI of the image.
</br> </br>
After inititializing docker-compose in background: 
</br>
Find the CONTAINER_ID using
</br>
`docker container ls`
</br>
Using that CONTAINER_ID enter the docker image CLI
</br>
`docker exec -it <CONTAINER_ID> bash`)
</br>
Use vim to view the log.txt file
</br>
`vim log.txt` 

#### Sending RAW POST DATA
The most conventient way to do this is to use the Postman app (https://www.getpostman.com/apps)
to send the POST requests.
