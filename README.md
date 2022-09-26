# SoccerAnalyticsPlatform

Soccer Analytics Platform which provides users with a web interface to extract insights from past FIFA world cups data. For instance, users can query information related to matches, players, and teams. It will include all of the core CRUD (create, read, update, and delete) operations.

## Technical Implementation
* Backend = Golang
* DB = MySQL (Local Environment) - containerized using Docker
* Frontend = Javascript + React (pending)

## Setup

1. Clone the repository, and cd into the cloned directory
2. Run command: "docker-compose up" - this step spins up a MySQL database in a docker container

Note: you must have Docker Desktop setup on your local machine in order to run the second step. If not, download it from the Docker website.

4. In another terminal, run command: "go run main.go" - runs the backend code to connect to the MySQL database and run sample queries.

### Sample Queries

CREATE TABLE ...

INSERT ....

SELECT * FROM ...

(insert ss)
