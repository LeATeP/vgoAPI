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
		{table: "user_", query: "SELECT * FROM user_ order by id"},
		{table: "item", query: "SELECT * FROM item order by id"},
		{table: "unit", query: "SELECT * FROM unit order by id"},
		{id: 10, name: "query units by user", table: "unit", query: "SELECT * FROM unit WHERE user_id = $1 order by id"},
	}
)

func main() {
	if err := psql_init(); err != nil {
		log.Fatal(err)
	}
	rou := gin.Default()
	rou.GET("/:view", GetTable)

	_ = rou.Run(":8080")
}

func GetTable(c *gin.Context) {
	table := c.Param(`view`)
	for _, v := range tables {
		if v.table == table {
			if err := v.fetchTable(); err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			c.IndentedJSON(http.StatusOK, v.dataPool)
			return
		}
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Table not found"})
}
