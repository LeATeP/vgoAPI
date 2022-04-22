package main

type queryStruct struct {
	id int
	name string
	table    string
	query    string
	data     any
	pointers []any
	dataPool []any
}
type user_ struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Units     string `json:"units"`
	Inventory string `json:"inventory"`
}

type unit struct {
	Id     int64  `json:"id"`
	UserID int    `json:"userID"`
	Level  int    `json:"level"`
	Class  string `json:"class"`
	Status string `json:"status"`
	Grade  string `json:"grade"`
	Stats  Stats  `json:"stats"`
}
type Stats struct {
	Health     int   `json:"health"`
	HealthFull int   `json:"healthFull"`
	Attack     int   `json:"attack"`
	Defense    int   `json:"defense"`
	Xp         int64 `json:"xp"`
}
type item struct {
	Id          int    `json:"id"`
	UserID      int    `json:"userID"`
	Name        string `json:"name"`
	ItemLvl     int    `json:"itemLvl"`
	Category    string `json:"category"`
	Rarity      string `json:"rarity"`
	Tier        int    `json:"tier"`
	Description string `json:"description"`
}
