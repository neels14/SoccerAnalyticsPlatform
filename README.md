# SoccerAnalyticsPlatform

Soccer Analytics Platform which provides users with a web interface to extract insights from past FIFA world cups data. For instance, users can query information related to matches, players, and teams.

## Technical Implementation
* Backend = Golang
* DB = MySQL (Local Environment) - containerized using Docker
* Frontend = Javascript + React (pending)

## Setup

1. Clone the repository, and ```cd``` into the cloned directory

2. In a new terminal: ```docker-compose up``` - this step spins up a MySQL database in a docker container

Note: you must have Docker Desktop setup on your local machine in order to run the second step. If not, download it from the Docker website.

3. In another terminal window, change your current directory into /cmd by running ```cd cmd/```. Then, run command: ```go run main.go```.
This will start the program, which begins by connecting to the MySQL database which you ran via docker, creates the necessary tables defined in our database schema, and then populates them with the sample test data.

### Program CLI

From here, the program will prompt you (the user) to enter certain parameters which our sql query features will use to run and provide you with information you might be interested in. Some features which are outputted require a user input (parameter) while others run without a user input.

<img width="838" alt="Screen Shot 2022-10-20 at 7 54 13 AM" src="https://user-images.githubusercontent.com/113378391/196941486-797a97c3-16f4-4382-8a2f-3f0115cb4711.png">



