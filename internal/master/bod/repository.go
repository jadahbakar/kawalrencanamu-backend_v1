package bod

import (
	"database/sql"
	"fmt"

	"go.uber.org/zap"
)

type bodRepository struct {
	Conn *sql.DB
}

func NewBodRepository(Conn *sql.DB) BodRepository {
	return &bodRepository{Conn: Conn}
}

func (br *bodRepository) SearchAll() ([]Bod, error) {
	query := `SELECT bod_id, bod_call_sign, bod_nama_id, bod_nama_en
				FROM mst.bod ORDER BY bod_id`
	rows, err := br.Conn.Query(query)
	if err != nil {
		zap.Error(err)
		return nil, err
	}

	result := make([]Bod, 0)
	t := Bod{}
	for rows.Next() {
		err = rows.Scan(
			&t.BodId,
			&t.BodCallSign,
			&t.BodNamaId,
			&t.BodNamaEn,
		)
		if err != nil {
			zap.Error(err)
			return nil, err
		}

		result = append(result, t)
	}
	if rows.Err() != nil {
		// if any error occurred while reading rows.
		fmt.Println("Error will reading user table: ", err)
		zap.Error(err)
		return nil, rows.Err()
	}
	return result, nil
}

func (br *bodRepository) SearchById(Id int) (Bod, error) {
	var t Bod
	query := `SELECT bod_id, bod_call_sign, bod_nama_id, bod_nama_en
				FROM mst.bod WHERE BY bod_id=?`
	err := br.Conn.QueryRow(query, Id).Scan(
		&t.BodId,
		&t.BodCallSign,
		&t.BodNamaId,
		&t.BodNamaEn)
	if err != nil {
		fmt.Println("Error will reading user table: ", err)
		zap.Error(err)
	}
	return t, nil
}
