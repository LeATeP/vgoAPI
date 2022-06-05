package main

import (
	"net/http"
	"github.com/gin-gonic/gin"

)

func GetUserData(c *gin.Context) {
	user := c.Param(`user`)
	table := c.Param(`view`)

	for _, v := range tables {
		if v.name == `userById` && v.table == table {
			if err := v.fetchTable(user); err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.IndentedJSON(http.StatusOK, v.dataPool)
			return
		}
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Table not found"})
}

func GetTable(c *gin.Context) {
	table := c.Param(`view`)

	for _, v := range tables {
		if v.table == table {
			if err := v.fetchTable(); err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.IndentedJSON(http.StatusOK, v.dataPool)
			return
		}
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Table not found"})
}