package main

import (
	"database/sql"
	"fmt"
	"os"
)

func psql_init() (err error) {
	config := init_config()
	p, err = sql.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		config.host, config.port, config.user, config.password, config.dbname))
	if err != nil {
		return
	}
	err = p.Ping()
	if err != nil {
		return err
	}
	fmt.Println("Connected to PostgreSQL!")
	return

}

func init_config() *con_config {
	return &con_config{
		hostname: os.Getenv("HOSTHAME"),
		host:     os.Getenv("PSQL_HOST"),
		dbname:   "postgres",
		user:     os.Getenv("PSQL_USER"),
		password: os.Getenv("PGPASSWORD"),
		port:     5432,
	}
}

// Type con_config is connection config to database
type con_config struct {
	hostname string // name of the machine / service / containerID
	host     string // IP address of db or name of the network
	dbname   string // name of db to connect to
	user     string // db user
	password string // db password
	port     int    // port of db
}
