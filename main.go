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
	p *sql.DB
)

func main() {
	if err := psql_init(); err != nil {
		log.Fatal(err)
	}
	rou := gin.Default()
	rou.GET("/:table", GetTable)

	rou.Run(":8080")
}
func GetTable(c *gin.Context) {
	table := c.Param(`table`)
	va := []queryStruct{{table: "user_", query: "SELECT * FROM user_ order by id"},
						{table: "item", query: "SELECT * FROM item order by id"},
						{table: "unit", query: "SELECT * FROM unit order by id"}}
	for i, v := range va {
		if v.table == table {
			data, err := fetchUsers(&va[i])
			if err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			c.IndentedJSON(http.StatusOK, data)
			return
		}
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Table not found"})
}
// next implementation is to use generics, fetchUser[user_] for each, in separate fn, 
// use []any as a data gather in struct and pointers...
func fetchUsers(q *queryStruct) ([]any, error) {
	var data any
	var dataPool []any
	var pointers []any
	rows, err := p.Query(q.query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		switch q.table {
		case `user_`:
			u := user_{}
			pointers = []any{&u.Id, &u.Username, &u.Units, &u.Inventory}
			data = &u
		case `unit`:
			u := unit{}
			pointers = []any{&u.Id, &u.UserID, &u.Level, &u.Class, &u.Status, &u.Grade, &u.Stats.Health, &u.Stats.HealthFull, &u.Stats.Attack, &u.Stats.Defense, &u.Stats.Xp}
			data = &u
		case `item`:
			u := item{}
			pointers = []any{&u.Id, &u.UserID, &u.Name, &u.ItemLvl, &u.Category, &u.Rarity, &u.Tier, &u.Description}
			data = &u
		}
		err = rows.Scan(pointers...)
		if err != nil {
			return nil, err
		}
		dataPool = append(dataPool, data)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return dataPool, nil
}