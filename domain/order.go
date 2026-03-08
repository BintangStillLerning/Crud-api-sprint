package domain

import "database/sql"

type Order struct {
	ID           int
	CustomerName string
	Total        int
	Payment      string
	Status       string
	Created_at   sql.NullString
}
