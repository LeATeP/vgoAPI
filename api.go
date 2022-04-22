package main

func (s *queryStruct) fetchTable() error {
	rows, err := p.Query(s.query)
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
