package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func create_tables(db *sql.DB) {
	// CREATE table Country
	_, err := db.Exec("CREATE TABLE Country ( " +
						"name VARCHAR(64) NOT NULL PRIMARY KEY, " + 
						"team_initials CHAR(3) NOT NULL UNIQUE )")
	if err != nil {
        panic(err.Error())
    }

	// CREATE table Player
	_, err = db.Exec("CREATE TABLE Player ( " +
						"name VARCHAR(64) NOT NULL, " + 
						"shirt_number INT NOT NULL, " +
						"country VARCHAR(64) NOT NULL REFERENCES Country(name), " +
						"PRIMARY KEY(name, shirt_number, country) )")
	if err != nil {
        panic(err.Error())
    }

	// CREATE table WorldCup
	_, err = db.Exec("CREATE TABLE WorldCup ( " +
						"year YEAR NOT NULL PRIMARY KEY, " +
						"host_country VARCHAR(64) NOT NULL REFERENCES Country(name), " +
						"first_place_country VARCHAR(64) NOT NULL REFERENCES Country(name), " +
						"second_place_country VARCHAR(64) NOT NULL REFERENCES Country(name), " +
						"third_place_country VARCHAR(64) NOT NULL REFERENCES Country(name) )")
	if err != nil {
        panic(err.Error())
    }

	// CREATE table SoccerMatch (Match keyword unavailable)
	_, err = db.Exec("CREATE TABLE SoccerMatch ( " +
						"year YEAR NOT NULL REFERENCES WorldCup(year), " +
						"stage VARCHAR(64) NOT NULL, " +
						"home_team_country VARCHAR(64) NOT NULL REFERENCES Country(name), " +
						"away_team_country VARCHAR(64) NOT NULL REFERENCES Country(name), " +
						"datetime DATETIME NOT NULL, " +
						"city VARCHAR(64) NOT NULL, " +
						"stadium VARCHAR(64) NOT NULL, " +
						"attendence INT NOT NULL, " +
						"home_coach VARCHAR(64) NOT NULL, " +
						"away_coach VARCHAR(64) NOT NULL, " +
						"PRIMARY KEY(year, stage, home_team_country, away_team_country) )")
	if err != nil {
        panic(err.Error())
    }

	// CREATE table PlaysIn
	_, err = db.Exec("CREATE TABLE PlaysIn ( " +
						"player_name VARCHAR(64) NOT NULL, " +
						"shirt_number INT NOT NULL, " +
						"player_country VARCHAR(64) NOT NULL, " +
						"year YEAR NOT NULL, " +
						"stage VARCHAR(64) NOT NULL, " +
						"home_team_country VARCHAR(64) NOT NULL, " +
						"away_team_country VARCHAR(64) NOT NULL, " +
						"is_starter INT NOT NULL, " +
						"goals_scored INT NOT NULL, " +
						"PRIMARY KEY(player_name, shirt_number, player_country, " +
						"year, stage, home_team_country, away_team_country), " +
						"FOREIGN KEY(player_name, shirt_number, player_country) REFERENCES Player(name, shirt_number, country), " +
						"FOREIGN KEY(year, stage, home_team_country, away_team_country) REFERENCES " +
						"SoccerMatch(year, stage, home_team_country, away_team_country) )")
	if err != nil {
        panic(err.Error())
    }

}

func drop_tables(db *sql.DB){
	db.Exec("DROP TABLE Country")
	db.Exec("DROP TABLE Player")
	db.Exec("DROP TABLE WorldCup")
	db.Exec("DROP TABLE SoccerMatch")
	db.Exec("DROP TABLE PlaysIn")
}

func main() {
	// Open up our database connection.
    // The database is called testdb
    db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()

	drop_tables(db)
	create_tables(db)
}