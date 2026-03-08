package domain


type Order struct {
	ID           int
	CustomerName string
	Total        int
	Payment      string
	Status       string
	Created_at   string
}
