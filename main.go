package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	_ = fmt.Print
	_ = sql.Open
	_ = json.NewEncoder
	_ = http.ListenAndServe
	_ = os.Getenv
)

var (
	p      *sql.DB
	tables = []queryStruct{
		{id: 1, table: "user_", query: "SELECT * FROM user_ order by id"},
		{id: 2, table: "item", query: "SELECT * FROM item order by id"},
		{id: 3, table: "unit", query: "SELECT * FROM unit order by id"},
		{id: 4, table: "unit", name: "userById", query: "SELECT * FROM unit WHERE userid = $1 order by id"},
		{id: 5, table: "item", name: "userById", query: "SELECT * FROM item WHERE userid = $1 order by id"},
	}
)

func main() {
	if err := psql_init(); err != nil {
		log.Fatal(err)
	}
	rou := gin.Default()
	rou.GET("/all=:view", GetTable)
	rou.GET("/user=:user/:view", GetUserData)

	_ = rou.Run(":8080")
}