# Overview
This repo is for RecordPoint Take Home Test specifically my implementation. You can read the requirments [here](RecordPointTakeHomeTest.pdf). I documented my approach to each problem which can be found [here](My_Approach_to_RecordPoint_Take_Home_Test.md)

# Technologies

- [Docker](https://docker.com): As container platform
- [Golang](https://go.dev): As programing language for main executable

# Local Setup
to setup the environment you need to run `docker-compose up -d` to create the containers. To take down the environment run `docker-compose down` which will stop the containers

# Usage
Once the environment is setup you can go to localhost:80 to see the static content being served from the webserver. You can also run the `main` executable to query the employees table from the database. You could also run `go run main.go` but you would need go version 1.16 which is the version I developed it on.

# Warnings
I did not test this on a windows machine my entire development was on on linux specifically Ubunut. 
