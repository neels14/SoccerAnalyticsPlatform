# SoccerAnalyticsPlatform

Soccer Analytics Platform which provides users with a web interface to extract insights from past FIFA world cups data. For instance, users can query information related to matches, players, and teams.

## Technical Implementation
* Backend = Golang
* DB = MySQL (Local Environment) - containerized using Docker
* Frontend = Javascript + React

## Backend & DB Setup

1. Clone the repository, and ```cd``` into the `server` subdirectory

2. In a new terminal: ```docker-compose up``` - this step spins up a MySQL database in a docker container

Note: you must have Docker Desktop setup on your local machine in order to run the second step. If not, download it from the Docker website.

3. In another terminal window, ```cd``` into the `server` subdirectory again and run command: ```go run main.go```.
This will start the backend server, which begins by connecting to the MySQL database which you ran via docker, creates the necessary tables defined in our database schema, and then populates them with the sample test data.

### Frontend Setup

Ensure you have npm/Node.js installed before continuing.

#### Running the app in development mode

1. ```cd``` into the `client` subdirectory.

2. Run ```npm start```. Open [http://localhost:3000](http://localhost:3000) to view it in your browser.

#### Building the app for production

1. ```cd``` into the `client` subdirectory.

2. Run ```npm run build``` to build the app for production in the `build` folder. This correctly bundles React in production mode and optimizes the build for the best performance.

Check out [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.