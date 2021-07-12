# Account-Service
A microservice to handle accounting in the Sensedia Demo 

## Running the application
This Application is ready to run in containers, the image can be built using the command:
docker build -t account-service . 

This application contains is able to run with a .env file containing the 
necessary environment variables that it needs to run the environments are 
described below

It's also possible to change the variables using environment 
variables during the container start up, the variables available are the
same variables in the .env file to run the container the Command below
can be used:
docker run -i -t -p 5005:5005 -e APP_PORT=5005 -e DATABASE_PORT=1111  customer-service

## Exposed Variables
PROFILE: Gin gonic profile example "debug"
APP_PORT: The port that the application will run 
DATABASE_PORT: Cockroach db port
DATABASE_NAME: Database name
DATABASE_USER: Database user
DATABASE_PASSWORD: Database password
DATABASE_HOST: Database host


