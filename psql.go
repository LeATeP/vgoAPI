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
func (s *queryStruct) fetchTable(args ...any) error {
	rows, err := p.Query(s.query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		s.makePointers()

		err = rows.Scan(s.pointers...)
		if err != nil {
			return err
		}
		s.dataPool = append(s.dataPool, s.data)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *queryStruct) makePointers() {
	switch s.table {
	case "user_":
		var d user_
		s.pointers = []any{&d.Id, &d.Username, &d.Units, &d.Inventory}
		s.data = &d
	case "unit":
		var d unit
		s.pointers = []any{&d.Id, &d.UserID, &d.Level, &d.Class, &d.Status, &d.Grade, &d.Stats.Health, &d.Stats.HealthFull, &d.Stats.Attack, &d.Stats.Defense, &d.Stats.Xp}
		s.data = &d
	case "item":
		var d item
		s.pointers = []any{&d.Id, &d.UserID, &d.Name, &d.ItemLvl, &d.Category, &d.Rarity, &d.Tier, &d.Description}
		s.data = &d
	}
}
