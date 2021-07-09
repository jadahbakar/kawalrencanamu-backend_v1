package propinsi

import "database/sql"

type propinsiRepository struct {
	Conn *sql.DB
}

// func NewRepository(Conn *sql.DB) PropinsiRepository {
// 	// return &propinsiRepository{Conn: Conn}
// }

// func (pr *propinsiRepository) FindAll() ([]*Propinsi, error) {

// }
